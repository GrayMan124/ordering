package server

import (
	// "github.com/GrayMan124/ordering/internal/database"
	"errors"
	"github.com/GrayMan124/ordering/internal/ui"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *ApiConfig) Login(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("ordering-bar-user")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			component := ui.Index(false)
			component.Render(r.Context(), w)
			return
		default:
			log.Println(err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
		}
		return
	}
	user_id, err := uuid.Parse(cookie.Value)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
	_, err = cfg.Queries.GetUserFromId(r.Context(), user_id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
	component := ui.Index(true)
	component.Render(r.Context(), w)
}
