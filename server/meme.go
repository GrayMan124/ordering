package server

import (
	"github.com/GrayMan124/ordering/internal/ui"
	"net/http"
)

func (cfg *ApiConfig) RespondWithMeme(w http.ResponseWriter, r *http.Request, meme string) {
	if meme == "cs?" {
		w.WriteHeader(200) //NOTE: I return 200 to make the HTMX swap the thing
		component := ui.MemeTempl("gomi", "Taksówka przyjechala już po ciebie, prosze wypierdalać")
		component.Render(r.Context(), w)
		return

	}
	// component := ui.ErrorRespons(errString, false)
	// component.Render(r.Context(), w)
}
