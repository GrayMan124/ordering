package server

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/GrayMan124/ordering/internal/database"
	"github.com/GrayMan124/ordering/internal/ui"

	"log"
	"net/http"
)

func (cfg *ApiConfig) GetCocktails(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Has("q") {
		cfg.CocktailSearch(w, r)
	} else if r.URL.Query().Has("filter") || r.URL.Query().Has("style") {
		cfg.CocktailsFiltered(w, r)
	} else {
		cfg.CocktailAllHandler(w, r)
	}
}

func (cfg *ApiConfig) CocktailAllHandler(w http.ResponseWriter, r *http.Request) {
	cocs, err := cfg.Queries.GetAllCock(r.Context())
	if err != nil {
		cfg.RespondWithError(w, r, 500)
		return
	}
	var cockNames []string
	var cockSpirits []string
	var imgNames []string
	var is_new_list []bool
	for _, cockt := range cocs {
		cockNames = append(cockNames, cockt.Name)
		cockSpirits = append(cockSpirits, cockt.BaseSpirit)
		imgNames = append(imgNames, cockt.ImgName.String)
		is_new_list = append(is_new_list, cockt.IsNew)
	}
	w.WriteHeader(200)
	component := ui.GetMenu(cockNames, cockSpirits, imgNames, is_new_list)
	component.Render(r.Context(), w)

}

func (cfg *ApiConfig) CocktailsFiltered(w http.ResponseWriter, r *http.Request) {
	cocs, err := cfg.Queries.GetAllCock(r.Context())
	if err != nil {
		log.Printf("Failed to retrieve cocktails data ")
		cfg.RespondWithError(w, r, 500)
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

func (cfg *ApiConfig) CocktailSearch(w http.ResponseWriter, r *http.Request) {
	quer := r.URL.Query().Get("q")
	if strings.Contains("Godfather", quer) && len(quer) >= 2 {
		cfg.RespondWithError(w, r, 403)
		return
	} else if quer == "cs?" {
		cfg.RespondWithMeme(w, r, "cs?")
		return
	} else if quer == "szlug?" {
		cfg.RespondWithMeme(w, r, "szlug")
		return
	}
	cocs, err := cfg.Queries.GetCockSearch(r.Context(), sql.NullString{String: fmt.Sprintf("%s", quer), Valid: true})
	if err != nil {
		cfg.RespondWithError(w, r, 500)
		return
	}
	var cockNames []string
	var cockSpirits []string
	var imgNames []string
	var is_new_list []bool
	for _, cockt := range cocs {
		cockNames = append(cockNames, cockt.Name)
		cockSpirits = append(cockSpirits, cockt.BaseSpirit)
		imgNames = append(imgNames, cockt.ImgName.String)
		is_new_list = append(is_new_list, cockt.IsNew)
	}
	w.WriteHeader(200)
	component := ui.GetMenu(cockNames, cockSpirits, imgNames, is_new_list)
	component.Render(r.Context(), w)

}
