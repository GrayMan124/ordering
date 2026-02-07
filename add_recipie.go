package main

import (
	// "database/sql"
	"encoding/json"
	"github.com/GrayMan124/ordering/internal/database"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type SendRecipie struct {
	Name        string    `json:"name"`
	Quantity    int32     `json:"quantity"`
	Abv         float32   `json:"abv"`
	Unit        string    `json:"unit"`
	Cockatil_id uuid.UUID `json:"cocktail_id"`
}

func (cfg *apiConfig) addRecipieFunc(w http.ResponseWriter, r *http.Request) {
	var sendRec SendRecipie

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&sendRec); err != nil {
		log.Printf("Failed to decode recipie data: %v", err)
		w.WriteHeader(500)
		return
	}

	created_recipie, err := cfg.Queries.AddRecipie(r.Context(), database.AddRecipieParams{
		Name:       sendRec.Name,
		Quantity:   sendRec.Quantity,
		Abv:        sendRec.Abv,
		CocktailID: sendRec.Cockatil_id})

	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to insert recipie into DB: %v", err)
		return
	}

	out, err := json.Marshal(created_recipie)
	if err != nil {
		log.Printf("Failed to marshal the recipie for a response")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(201)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
