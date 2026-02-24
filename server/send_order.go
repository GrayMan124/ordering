package server

import (
	"log"
	"net/http"

	"github.com/GrayMan124/ordering/internal/database"
	"github.com/GrayMan124/ordering/internal/ui"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) SendOrder(w http.ResponseWriter, r *http.Request) {
	cocktail_name := r.FormValue("cocktail")
	cookie, err := r.Cookie("ordering-bar-user")
	if err != nil {
		log.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	userUUID, err := uuid.Parse(cookie.Value)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	user, err := cfg.Queries.GetUserFromId(r.Context(), userUUID)
	cocktail, err := cfg.Queries.GetCocktail(r.Context(), cocktail_name)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to retrieve data from DB: %v", err)
		return
	}
	_, err = cfg.Queries.SendOrder(r.Context(), database.SendOrderParams{CocktailID: cocktail.ID.UUID,
		OrderedBy: user.Name})
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to send order into DB: %v", err)
		return
	}
	w.WriteHeader(201)
	component := ui.OrderSend()
	component.Render(r.Context(), w)
}
