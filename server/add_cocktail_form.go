package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/GrayMan124/ordering/internal/database"
)

func (cfg *ApiConfig) AddCocktailFromData(w http.ResponseWriter, r *http.Request) {
	cockName := r.FormValue("cocktailName")
	checkCocktails, err := cfg.Queries.CheckCocktail(r.Context(), cockName)
	if err != nil {
		log.Printf("Failed to execute query with error: %v", err)
		w.WriteHeader(500)
		return
	}
	if checkCocktails != 0 {
		log.Printf("Cocktail: %v \nalready exsits in the DB\n", cockName)
		w.WriteHeader(400)
		return
	}
	log.Printf("Got: %v cocktail\nFound: %v records", cockName, checkCocktails)
	baseSpirit := r.FormValue("BaseSpirit")
	cocktailImg := r.FormValue("CocktailIMG")
	cocktailType := r.FormValue("CocktailType")

	created_cocktail, err := cfg.Queries.AddCocktail(r.Context(), database.AddCocktailParams{
		BaseSpirit:   baseSpirit,
		CocktailType: "TBD",
		Name:         cockName,
		ImgName:      sql.NullString{String: cocktailImg, Valid: true},
		Type:         sql.NullString{String: cocktailType, Valid: true}})
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to insert cocktail into DB: %v", err)
		return
	}
	numIngr := r.FormValue("NumIngr")
	numIngrInt, _ := strconv.ParseInt(numIngr, 10, 32)
	var ingrList []Ingredient
	for idx := range numIngrInt {
		ingrName := r.FormValue(fmt.Sprintf("Ingredient_%v", idx))
		quantity := r.FormValue(fmt.Sprintf("Amount_%v", idx))
		abv := r.FormValue(fmt.Sprintf("ABV_%v", idx))
		unit := r.FormValue(fmt.Sprintf("Unit_%v", idx))
		abvF, _ := strconv.ParseFloat(abv, 32)
		quantI, _ := strconv.ParseInt(quantity, 10, 32)
		ingrList = append(ingrList, Ingredient{
			Name:   ingrName,
			ABV:    float32(abvF),
			Amount: int32(quantI),
			Unit:   unit,
		})
	}

	ingrMap, err := cfg.getIngredientsId(ingrList, r)
	if err != nil {
		log.Printf("Failed to retrieve the ingredient Ids: %v", err)
		w.WriteHeader(500)
		return
	}
	//Add all the ingredients into the db
	err = cfg.addRecipie(created_cocktail, ingrList, ingrMap, r)
	if err != nil {
		log.Printf("Failed to add parts of recipie into recipies: %v", err)
		w.WriteHeader(500)
		return
	}
	_ = ReturnCocktail{Name: created_cocktail.Name, ID: created_cocktail.ID.UUID}

	w.WriteHeader(201)
	w.Header().Set("Content-Type", "application/json")
	return
}
