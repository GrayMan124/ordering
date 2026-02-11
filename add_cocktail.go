package main

import (
	// "database/sql"
	"database/sql"
	"encoding/json"
	// "fmt"
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

func (cfg *apiConfig) getIngredientsId(ingredients []Ingredient, r *http.Request) (map[string]uuid.UUID, error) {

	ingrNames := make([]string, len(ingredients))
	for idx, ingr := range ingredients {
		ingrNames[idx] = ingr.Name
	}
	ingrQuery, err := cfg.Queries.GetIngredients(r.Context(), ingrNames)
	if err != nil {
		return nil, err
	}
	ingrQueryNames := make([]string, len(ingrQuery))
	for idx, ingrQ := range ingrQuery {
		ingrQueryNames[idx] = ingrQ.Name
	}
	ingrMap := make(map[string]uuid.UUID, len(ingrNames))
	for _, ingr := range ingredients {
		if slices.Contains(ingrQueryNames, ingr.Name) {
			for _, ingrQ := range ingrQuery {
				if ingrQ.Name == ingr.Name {
					ingrMap[ingr.Name] = ingrQ.ID
					break
				}
			}
		} else {
			ingrNew, err := cfg.Queries.AddIngredient(r.Context(), database.AddIngredientParams{
				Name: ingr.Name,
				Abv:  ingr.ABV,
			})
			if err != nil {
				return nil, err
			}
			ingrMap[ingr.Name] = ingrNew.ID
			continue
		}
	}
	return ingrMap, nil
}

func (cfg *apiConfig) addRecipie(cocktail database.Cocktail, ingredients []Ingredient, ingredientMap map[string]uuid.UUID, r *http.Request) error {
	for _, ingr := range ingredients {
		_, err := cfg.Queries.AddRecipIngr(r.Context(), database.AddRecipIngrParams{
			CocktailID:   cocktail.ID.UUID,
			IngredientID: ingredientMap[ingr.Name],
			Amount:       ingr.Amount,
			Unit:         ingr.Unit,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (cfg *apiConfig) addCocktail(w http.ResponseWriter, r *http.Request) {
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
	ingrMap, err := cfg.getIngredientsId(cocktail.Ingredients, r)
	if err != nil {
		log.Printf("Failed to retrieve the ingredient Ids: %v", err)
		w.WriteHeader(500)
		return
	}
	//Add all the ingredients into the db
	err = cfg.addRecipie(created_cocktail, cocktail.Ingredients, ingrMap, r)
	if err != nil {
		log.Printf("Failed to add parts of recipie into recipies: %v", err)
		w.WriteHeader(500)
		return
	}
	output_cocktail := ReturnCocktail{Name: created_cocktail.Name, ID: created_cocktail.ID.UUID}

	out, err := json.Marshal(output_cocktail)
	if err != nil {
		log.Printf("Failed to marshal the cocktail for a response")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(201)
	w.Write(out)
	w.Header().Set("Content-Type", "application/json")
}
