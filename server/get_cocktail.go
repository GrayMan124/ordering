package server

import (
	"log"
	"math/rand/v2"
	"net/http"

	"github.com/GrayMan124/ordering/internal/ui"
)

func (cfg *ApiConfig) GetCocktailData(w http.ResponseWriter, r *http.Request) {
	var err error
	cockName := r.URL.Query().Get("cocktail")
	if cockName == "Godfather" {
		cfg.GetGodfather(w, r)
		return
	}
	ingredients, err := cfg.Queries.GetRecipie(r.Context(), cockName)
	if err != nil {
		log.Printf("Failed to retrieve cocktail data ")
		cfg.RespondWithError(w, r, 500)
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

func (cfg *ApiConfig) GetGodfather(w http.ResponseWriter, r *http.Request) {
	key := rand.IntN(6) + 1
	w.WriteHeader(200)
	component := ui.ExpandGodfather(key)
	component.Render(r.Context(), w)
}
