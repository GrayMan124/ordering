package main

import (
	// "database/sql"
	"encoding/json"
	"github.com/GrayMan124/ordering/internal/database"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type createOrder struct {
	CocktailID uuid.UUID `json:"cocktail_id"`
	OrderedBy  string    `json:"ordered_by"`
}

func (cfg *apiConfig) sendOrder(w http.ResponseWriter, r *http.Request) {
	var order createOrder
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&order); err != nil {
		log.Printf("Failed to decode order request")
		w.WriteHeader(500)
		return
	}
	createdOrder, err := cfg.Queries.SendOrder(r.Context(), database.SendOrderParams{CocktailID: order.CocktailID,
		OrderedBy: order.OrderedBy})
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to send order into DB: %v", err)
		return
	}

	out, err := json.Marshal(createdOrder)
	if err != nil {
		log.Printf("Failed to marshal the order for response")
		w.WriteHeader(500)
		return
	}
	w.Write(out)
	w.WriteHeader(201)
	w.Header().Set("Content-Type", "application/json")
}
