package main

import (
	// "database/sql"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"

	"github.com/GrayMan124/ordering/internal/database"
	"github.com/google/uuid"
)

type Ingredient struct {
	Name   string  `json:"name"`
	Amount int32   `json:"amount"`
	ABV    float32 `json:"abv"`
	Unit   string  `json:"unit"`
}

type SendCocktail struct {
	Base_spirit   string       `json:"base_spirit"`
	Cocktail_type string       `json:"cocktail_type"`
	Name          string       `json:"name"`
	ImgName       string       `json:"img_name"`
	Type          string       `json:"type"`
	Ingredients   []Ingredient `json:"ingredients"`
}

type ReturnCocktail struct {
	Name string    `json:"name"`
	ID   uuid.UUID `json:"uuid"`
}

func (cfg *apiConfig) getIngredientsId(ingredients []Ingredient, r *http.Request) ([]uuid.UUID, error) {

	ingrNames := make([]string, len(ingredients))
	for idx, ingr := range ingredients {
		ingrNames[idx] = ingr.Name
	}
	ingrQuery, err := cfg.Queries.GetIngredients(r.Context(), ingrNames)
	if err != nil {
		return nil, err
	}
	if len(ingrQuery) < len(ingrNames) {
		var ingrInDB []string
		for _, ingrQ := range ingrQuery {
			ingrInDB = append(ingrInDB, ingrQ.Name)
		}
		for _, ingrName := range ingrNames {
			if slices.Contains(ingrInDB, ingrName) {
				continue
			}
			//Add the ingredient
		}
	}
}

func (cfg *apiConfig) addCocktails(w http.ResponseWriter, r *http.Request) {
	var cocktail SendCocktail

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cocktail); err != nil {
		log.Printf("Failed to decode cocktail data")
		w.WriteHeader(500)
		return
	}
	img_name := sql.NullString{String: cocktail.ImgName, Valid: true}
	type_name := sql.NullString{String: cocktail.Type, Valid: true}

	created_cocktail, err := cfg.Queries.AddCocktail(r.Context(), database.AddCocktailParams{
		BaseSpirit:   cocktail.Base_spirit,
		CocktailType: cocktail.Cocktail_type,
		Name:         cocktail.Name,
		ImgName:      img_name,
		Type:         type_name})
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to insert cocktail into DB: %v", err)
		return
	}

	output_cocktail := ReturnCocktail{Name: created_cocktail.Name, ID: created_cocktail.ID.UUID}

	out, err := json.Marshal(output_cocktail)
	if err != nil {
		log.Printf("Failed to marshal the cocktail for a response")
		w.WriteHeader(500)
		return
	}
	w.Write(out)
	w.WriteHeader(201)
	w.Header().Set("Content-Type", "application/json")
}
