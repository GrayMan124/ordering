package server

import (
	// "database/sql"
	"log"
	"net/http"

	"github.com/GrayMan124/ordering/internal/ui"
	"github.com/google/uuid"
)

type Cancellation struct {
	OrderID uuid.UUID `json:"order_id"`
}

func (cfg *ApiConfig) CancelOrder(w http.ResponseWriter, r *http.Request) {
	order_id := r.URL.Query().Get("ordId")
	log.Printf("Got request to cancel order:\n%v", order_id)
	ordId, err := uuid.Parse(order_id)
	if err != nil {
		w.WriteHeader(500)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	_, err = cfg.Queries.CancelOrder(r.Context(), ordId)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to send cancellation into DB: %v", err)
		return
	}
	w.WriteHeader(200)
	component := ui.CancelOrder()
	err = component.Render(r.Context(), w)
	if err != nil {
		log.Printf("Failed to render cancel order: %s", err)
	}
}
