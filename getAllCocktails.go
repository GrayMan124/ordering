package main

import (
	"github.com/GrayMan124/ordering/internal/database"
	"github.com/GrayMan124/ordering/internal/ui"
	// "github.com/a-h/templ"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type cocktail struct {
	ID         uuid.UUID `json:"cocktail_id"`
	DataUrl    string    `json: "data_url"`
	Name       string    `json:"name"`
	BaseSpirit string    `json:"base_spirit"`
}

func (cfg *apiConfig) getAllCocktails(w http.ResponseWriter, r *http.Request) {
	cocs, err := cfg.Queries.GetAllCock(r.Context())
	if err != nil {
		log.Printf("Failed to retrieve cocktails data ")
		w.WriteHeader(500)
		return
	}
	filter := r.URL.Query().Get("filter")
	cock_style := r.URL.Query().Get("style")
	var final_cocs []database.Cocktail
	if filter != "" || cock_style != "" {
		if filter != "" {
			for _, coc := range cocs {
				if coc.BaseSpirit == filter {
					final_cocs = append(final_cocs, coc)
				}
			}
		} else {
			for _, coc := range cocs {
				if coc.Type.String == cock_style {
					final_cocs = append(final_cocs, coc)
				}
			}
		}
	} else {
		final_cocs = cocs
	}

	var cockNames []string
	var cockSpirits []string
	var imgNames []string
	var is_new_list []bool
	for _, cockt := range final_cocs {
		cockNames = append(cockNames, cockt.Name)
		cockSpirits = append(cockSpirits, cockt.BaseSpirit)
		imgNames = append(imgNames, cockt.ImgName.String)
		is_new_list = append(is_new_list, cockt.IsNew)
	}
	w.WriteHeader(200)
	component := ui.GetMenu(cockNames, cockSpirits, imgNames, is_new_list)
	component.Render(r.Context(), w)

}
