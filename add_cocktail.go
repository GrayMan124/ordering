package main

import (
	"database/sql"
	"encoding/json"
	"github.com/GrayMan124/ordering/internal/database"
	// "github.com/google/uuid"
	"log"
	"net/http"
)

type SendCocktail struct {
	Data_url      sql.NullString `json:"data_url"`
	Base_spirit   string         `json:"base_spirit"`
	Cocktail_type string         `json:"cocktail_type"`
	Name          string         `json:"name"`
}

func (cfg *apiConfig) addCocktails(w http.ResponseWriter, r *http.Request) {
	var cocktail SendCocktail

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cocktail); err != nil {
		log.Printf("Failed to decode cocktail data")
		w.WriteHeader(500)
		return
	}

	created_cocktail, err := cfg.Queries.AddCocktail(r.Context(), database.AddCocktailParams{
		DataUrl:      cocktail.Data_url,
		BaseSpirit:   cocktail.Base_spirit,
		CocktailType: cocktail.Cocktail_type,
		Name:         cocktail.Name})
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to insert cocktail into DB: %v", err)
		return
	}

	out, err := json.Marshal(created_cocktail)
	if err != nil {
		log.Printf("Failed to marshal the cocktail for a response")
		w.WriteHeader(500)
		return
	}
	w.Write(out)
	w.WriteHeader(201)
	w.Header().Set("Content-Type", "application/json")
}
