package server

import (
	// "github.com/GrayMan124/ordering/internal/database"
	// "errors"
	"github.com/GrayMan124/ordering/internal/ui"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *ApiConfig) MyOrders(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("ordering-bar-user")
	if err != nil {
		log.Println(err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
	user_id, err := uuid.Parse(cookie.Value)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
	user_name, err := cfg.Queries.GetUserFromId(r.Context(), user_id)
	orders, err := cfg.Queries.GetMyOrders(r.Context(), user_name.Name)
	if err != nil {
		log.Println(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
	var outputOrders []ui.MyOrders
	for _, ord := range orders {
		isCancel := ord.CanceledAt.Valid
		outputOrders = append(outputOrders, ui.MyOrders{
			OrderId:      ord.ID,
			OrderedAt:    ord.CreatedAt,
			IsCancelled:  isCancel,
			CocktailName: ord.Name.String,
			Finished:     ord.Finished.Bool,
			ImgName:      ord.ImgName.String,
		})
	}

	component := ui.MyOrdersUI(outputOrders, user_name.Name)
	component.Render(r.Context(), w)
}
