package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type cocktail struct {
	ID      uuid.UUID `json:"cocktail_id"`
	DataUrl string    `json: "data_url"`
	Name    string    `json:"name"`
}

func (cfg *apiConfig) getAllCocktails(w http.ResponseWriter, r *http.Request) {
	var cocktails []cocktail

	cocs, err := cfg.Queries.GetAllCock(r.Context())
	if err != nil {

		log.Printf("Failed to retrieve cocktails data ")
		w.WriteHeader(500)
		return
	}

	for _, cockt := range cocs {
		cocktails = append(cocktails, cocktail{ID: cockt.ID.UUID, DataUrl: cockt.DataUrl.String, Name: cockt.Name})
	}

	out, err := json.Marshal(cocktails)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to marshal cocktails")
		return
	}
	w.Write(out)
	w.WriteHeader(200)
}
