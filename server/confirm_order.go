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
