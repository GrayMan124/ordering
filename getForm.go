package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/GrayMan124/ordering/internal/ui"
)

func (cfg *apiConfig) GetRecipieForm(w http.ResponseWriter, r *http.Request) {
	num_ingredients := r.URL.Query().Get("num")
	var numIngr int
	if num_ingredients == "" {
		numIngr = 0
	} else {
		numIngr, _ = strconv.Atoi(num_ingredients)
	}
	cockTypes, err := cfg.Queries.GetCocktailTypes(r.Context())
	if err != nil {
		log.Printf("Failed to retrieve cocktail types: %v", err)
		w.WriteHeader(500)
		return
	}
	imgLabels, err := cfg.Queries.GetCocktailImgs(r.Context())
	if err != nil {
		log.Printf("Failed to retrieve cocktail imgs: %v", err)
		w.WriteHeader(500)
		return
	}
	options := make(map[string][]string)
	options["CocktailType"] = make([]string, 0)
	for _, cockType := range cockTypes {
		options["CocktailType"] = append(options["CocktailType"], cockType.String)
	}

	options["CocktailIMG"] = make([]string, 0)
	for _, cockImg := range imgLabels {
		options["CocktailIMG"] = append(options["CocktailIMG"], cockImg.String)
	}
	w.WriteHeader(200)
	component := ui.RecipieForm(numIngr, options)
	component.Render(r.Context(), w)

}
