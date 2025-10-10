package main

import (
	// "database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type Cancellation struct {
	OrderID uuid.UUID `json:"order_id"`
}

func (cfg *apiConfig) cancelOrder(w http.ResponseWriter, r *http.Request) {
	var canc Cancellation
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&canc); err != nil {
		log.Printf("Failed to decode cancel order request")
		w.WriteHeader(500)
		return
	}
	canceled_order, err := cfg.Queries.CancelOrder(r.Context(), canc.OrderID)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to send cancellation into DB: %v", err)
		return
	}

	out, err := json.Marshal(canceled_order)
	if err != nil {
		log.Printf("Failed to marshal the cancellation for response")
		w.WriteHeader(500)
		return
	}
	w.Write(out)
	w.WriteHeader(204)
	w.Header().Set("Content-Type", "application/json")
}
