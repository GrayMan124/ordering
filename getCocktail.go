package main

import (
	"github.com/GrayMan124/ordering/internal/database"
	"github.com/GrayMan124/ordering/internal/ui"
	"log"
	"net/http"
)

func (cfg *apiConfig) getCocktailAPI(w http.ResponseWriter, r *http.Request) {
	var ingredients []database.Ingredient
	var err error
	cockName := r.URL.Query().Get("cocktail")

	ingredients, err = cfg.Queries.GetRecipie(r.Context(), cockName)
	if err != nil {
		log.Printf("Failed to retrieve cocktail data ")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	component := ui.ExpandCocktail(cockName, ingredients)
	component.Render(r.Context(), w)

}
