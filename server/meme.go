package server

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/GrayMan124/ordering/internal/ui"
)

func (cfg *ApiConfig) RespondWithMeme(w http.ResponseWriter, r *http.Request, meme string) {
	if meme == "cs?" {
		w.WriteHeader(200)
		component := ui.MemeTempl("gomi", "Taksówka przyjechala już po ciebie, prosze wypierdalać")
		component.Render(r.Context(), w)
		return

	} else if meme == "szlug" {
		outcome := rand.Int31n(4) + 1
		var text_message string
		switch outcome {
		case 1:
			text_message = "SZLUG?"
		case 2:
			text_message = "UUUUUU SZLUG???"
		case 3:
			text_message = "Krystian zaprasza na szluga"
		case 4:
			text_message = "Hmm?"
		case 5:
			text_message = "Znowu na szluga?"

		}

		w.WriteHeader(200)
		component := ui.MemeTempl(fmt.Sprintf("szlug_v%v", outcome), text_message)
		component.Render(r.Context(), w)
		return

	}
	// component := ui.ErrorRespons(errString, false)
	// component.Render(r.Context(), w)
}
