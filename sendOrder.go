package main

import (
	// "database/sql"
	// "encoding/json"
	"github.com/GrayMan124/ordering/internal/database"
	"github.com/GrayMan124/ordering/internal/ui"
	// "github.com/google/uuid"
	"log"
	"net/http"
)

func (cfg *apiConfig) sendOrder(w http.ResponseWriter, r *http.Request) {
	cocktail_name := r.FormValue("cocktail")
	user_name := r.FormValue("customerName")
	cocktail, err := cfg.Queries.GetCocktail(r.Context(), cocktail_name)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to retrieve data from DB: %v", err)
		return
	}
	_, err = cfg.Queries.SendOrder(r.Context(), database.SendOrderParams{CocktailID: cocktail.ID.UUID,
		OrderedBy: user_name})
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to send order into DB: %v", err)
		return
	}
	w.WriteHeader(201)
	component := ui.OrderSend()
	component.Render(r.Context(), w)
}
