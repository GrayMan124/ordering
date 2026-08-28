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
	var cockList []ui.CocktailUI
	for _, cockt := range cocs {
		cockList = append(cockList, ui.CocktailUI{
			Name:       cockt.Name,
			CockSpirit: cockt.BaseSpirit,
			ImgName:    cockt.ImgName.String,
			Is_new:     cockt.IsNew,
			Score:      cockt.Score,
		})
	}
	w.WriteHeader(200)
	component := ui.GetMenu(cockList)
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
	var final_cocs []database.GetAllCockRow
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

	var cockList []ui.CocktailUI
	for _, cockt := range cocs {
		cockList = append(cockList, ui.CocktailUI{
			Name:       cockt.Name,
			CockSpirit: cockt.BaseSpirit,
			ImgName:    cockt.ImgName.String,
			Is_new:     cockt.IsNew,
			Score:      cockt.Score,
		})
	}
	w.WriteHeader(200)
	component := ui.GetMenu(cockList)
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
	var cockList []ui.CocktailUI
	for _, cockt := range cocs {
		cockList = append(cockList, ui.CocktailUI{
			Name:       cockt.Name,
			CockSpirit: cockt.BaseSpirit,
			ImgName:    cockt.ImgName.String,
			Is_new:     cockt.IsNew,
			Score:      cockt.Score,
		})
	}
	w.WriteHeader(200)
	component := ui.GetMenu(cockList)
	component.Render(r.Context(), w)

}
