package server

import (
	// "database/sql"
	// "fmt"

	// "github.com/GrayMan124/ordering/internal/database"
	"github.com/GrayMan124/ordering/internal/ui"

	"net/http"
)

func (cfg *ApiConfig) IngredientMenu(w http.ResponseWriter, r *http.Request) {
	ingrs, err := cfg.Queries.GetIngr(r.Context())
	if err != nil {
		cfg.RespondWithError(w, r, 500)
		return
	}
	var ingrOut []ui.Ingredient
	for _, ingr := range ingrs {
		ingrOut = append(ingrOut, ui.Ingredient{Name: ingr.Name, Available: ingr.IsAvailable})
	}
	w.WriteHeader(200)
	component := ui.IngredientMenu(ingrOut)
	component.Render(r.Context(), w)

}

func (cfg *ApiConfig) IngredientExpand(w http.ResponseWriter, r *http.Request) {
	ingrName := r.URL.Query().Get("q")
	availability := r.URL.Query().Get("av")
	var avBool bool
	if availability == "t" {
		avBool = true
	} else {
		avBool = false
	}
	ingr := ui.Ingredient{
		Name:      ingrName,
		Available: avBool,
	}

	cockNames, err := cfg.Queries.GetCocksIngr(r.Context(), ingrName)
	if err != nil {
		cfg.RespondWithError(w, r, 500)
		return
	}
	var cocksOut []string
	for _, cock := range cockNames {
		cocksOut = append(cocksOut, cock.String)
	}
	w.WriteHeader(200)
	component := ui.IngredientDesc(ingr, cocksOut)
	component.Render(r.Context(), w)

}
