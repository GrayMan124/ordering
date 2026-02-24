package server

import (
	"github.com/GrayMan124/ordering/internal/ui"
	"log"
	"net/http"
)

func (cfg *ApiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("UserName")
	if userName == "" {
		log.Printf("Retrieved empty UserName")
		w.WriteHeader(400)
		return
	}
	user, err := cfg.Queries.CreateUser(r.Context(), userName)
	if err != nil {
		log.Printf("Failed to add user to DB")
		w.WriteHeader(500)
		return
	}
	log.Printf("Created user: %v with id: %v", user.Name, user.ID)
	cookie := http.Cookie{
		Name:     "ordering-bar-user",
		Value:    user.ID.String(),
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	log.Printf("Cookie has been send, cookie name: %s", cookie.Name)
	w.WriteHeader(201)
	component := ui.Index(true)
	component.Render(r.Context(), w)
}
