package main

import (
	"net/http"
	"strconv"

	"github.com/GrayMan124/ordering/internal/ui"
)

func GetRecipieForm(w http.ResponseWriter, r *http.Request) {
	num_ingredients := r.URL.Query().Get("num")
	var numIngr int
	if num_ingredients == "" {
		numIngr = 0
	} else {
		numIngr, _ = strconv.Atoi(num_ingredients)
	}
	w.WriteHeader(200)
	component := ui.RecipieForm(numIngr)
	component.Render(r.Context(), w)

}
