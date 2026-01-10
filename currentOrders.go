package main

import (
	"github.com/GrayMan124/ordering/internal/ui"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func (cfg *apiConfig) getCurrentOrders(w http.ResponseWriter, r *http.Request) {
	var orders []ui.Recipie
	ordersDB, err := cfg.Queries.GetCurrentOrders(r.Context())
	if err != nil {
		log.Printf("Failed to retrieve cocktails data ")
		w.WriteHeader(500)
		return
	}
	for _, ord := range ordersDB {
		cocktailId := ord.CocktailID
		cocktailFull, err := cfg.Queries.GetCocktailName(r.Context(), uuid.NullUUID{UUID: cocktailId, Valid: true})
		if err != nil {
			w.WriteHeader(500)
			log.Printf("Failed to get Cocktail name: %v", err)
			return
		}
		ingr, err := cfg.Queries.GetRecipie(r.Context(), cocktailFull.Name)
		if err != nil {
			w.WriteHeader(500)
			log.Printf("Failed to get recipie: %v", err)
			return
		}
		orders = append(orders, ui.Recipie{
			Cocktail:    cocktailFull.Name,
			Ingredients: ingr,
			OrderId:     ord.ID,
			OrderedBy:   ord.OrderedBy})
	}
	w.WriteHeader(200)
	component := ui.Orders(orders)
	component.Render(r.Context(), w)
}
