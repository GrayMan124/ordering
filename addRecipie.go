package main

import (
	// "database/sql"
	// "github.com/GrayMan124/ordering/internal/ui"
	// "github.com/google/uuid"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/GrayMan124/ordering/internal/database"
)

func (cfg *apiConfig) AddCocktailFromData(w http.ResponseWriter, r *http.Request) {
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
	baseSpirit := r.FormValue("baseSpirit")
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
	cockId := created_cocktail.ID.UUID
	numIngr := r.FormValue("NumIngr")
	numIngrInt, _ := strconv.ParseInt(numIngr, 10, 32)
	for idx := range numIngrInt {
		ingrName := r.FormValue(fmt.Sprintf("Ingredient_%v", idx))
		quantity := r.FormValue(fmt.Sprintf("Amount_%v", idx))
		abv := r.FormValue(fmt.Sprintf("ABV_%v", idx))
		abvF, _ := strconv.ParseFloat(abv, 32)
		quantI, _ := strconv.ParseInt(quantity, 10, 32)
		abvVal := float32(abvF)
		quantVal := int32(quantI)
		addedRecipie, err := cfg.Queries.AddRecipie(r.Context(), database.AddRecipieParams{
			Name:       ingrName,
			Quantity:   quantVal,
			Abv:        abvVal,
			CocktailID: cockId})
		if err != nil {
			w.WriteHeader(500)
			log.Printf("Failed to add ingredient:\n%v", err)
		}
		log.Printf("Added ingredient: %s\n", addedRecipie.Name)

	}
	w.WriteHeader(201)
	return
}
