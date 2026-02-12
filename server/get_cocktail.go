package server

import (
	"github.com/GrayMan124/ordering/internal/ui"
	"log"
	"net/http"
)

func (cfg *ApiConfig) GetCocktailData(w http.ResponseWriter, r *http.Request) {
	var err error
	cockName := r.URL.Query().Get("cocktail")

	ingredients, err := cfg.Queries.GetRecipie(r.Context(), cockName)
	if err != nil {
		log.Printf("Failed to retrieve cocktail data ")
		w.WriteHeader(500)
		return
	}
	var ingrNames []string
	for _, ingr := range ingredients {
		ingrNames = append(ingrNames, ingr.String)
	}
	w.WriteHeader(200)
	component := ui.ExpandCocktail(cockName, ingrNames)
	component.Render(r.Context(), w)
}
