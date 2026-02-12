package server

import (
	"github.com/GrayMan124/ordering/internal/ui"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func (cfg *ApiConfig) FinishOrder(w http.ResponseWriter, r *http.Request) {
	ordID := r.FormValue("orderId")
	CockName := r.FormValue("CocktailName")
	ordBy := r.FormValue("OrderedBy")
	orderID, err := uuid.Parse(ordID)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to Parse UUID for Order: %v", err)
		return
	}
	_, err = cfg.Queries.FinishOrder(r.Context(), orderID)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to send Finised order into DB: %v", err)
		return
	}
	w.WriteHeader(200)
	component := ui.FinishOrder(CockName, ordBy)
	err = component.Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to Render Object: %v", err)
	}
}
