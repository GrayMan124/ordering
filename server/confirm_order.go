package server

import (
	"net/http"

	"github.com/GrayMan124/ordering/internal/ui"
)

func (cfg *ApiConfig) ConfirmOrder(w http.ResponseWriter, r *http.Request) {
	cocktail_name := r.FormValue("cocktail")

	w.WriteHeader(201)
	component := ui.ConfirmOrder(cocktail_name)
	component.Render(r.Context(), w)
}

func (cfg *ApiConfig) CheckGodfather(w http.ResponseWriter, r *http.Request) {
	choice := r.FormValue("choice")
	key := r.FormValue("key")
	if choice != key {
		component := ui.CancelOrder()
		component.Render(r.Context(), w)
		return
	}
	cfg.ConfirmGodfather(w, r)
}

func (cfg *ApiConfig) ConfirmGodfather(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
	component := ui.ConfirmOrder("Godfather")
	component.Render(r.Context(), w)
}
