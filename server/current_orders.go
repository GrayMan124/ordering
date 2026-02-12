package server

import (
	"github.com/GrayMan124/ordering/internal/ui"
	"log"
	"net/http"
)

func (cfg *ApiConfig) GetCurrentOrders(w http.ResponseWriter, r *http.Request) {
	var orders []ui.OrderStruct
	ordersDB, err := cfg.Queries.GetFullOrders(r.Context())
	if err != nil {
		log.Printf("Failed to retrieve cocktails data ")
		w.WriteHeader(500)
		return
	}

	for _, ord := range ordersDB {
		ingrs, err := cfg.Queries.GetFullRecipie(r.Context(), ord.CocktailID)

		if err != nil {
			w.WriteHeader(500)
			log.Printf("Failed to get recipie: %v", err)
			return
		}
		var uiIngrs []ui.RecipeIngr
		for _, ingr := range ingrs {
			uiIngrs = append(uiIngrs, ui.RecipeIngr{Name: ingr.Name.String, Amount: int(ingr.Amount), Unit: ingr.Unit})
		}
		orders = append(orders, ui.OrderStruct{
			Cocktail:    ord.Name.String,
			Ingredients: uiIngrs,
			OrderId:     ord.ID,
			OrderedBy:   ord.OrderedBy})
	}
	w.WriteHeader(200)
	component := ui.Orders(orders)
	component.Render(r.Context(), w)
}
