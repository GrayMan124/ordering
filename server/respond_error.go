package server

import (
	"github.com/GrayMan124/ordering/internal/ui"
	"net/http"
)

func (cfg *ApiConfig) RespondWithError(w http.ResponseWriter, r *http.Request, err_code int) {
	var errString string
	switch err_code {
	case 400:
		errString = "Error code 400\n Bad Request"
	case 401:
		errString = "Error code 401\n Unauthorized"
	case 403:
		errString = "Error code 403\n Forbidden"
	case 404:
		errString = "Error code 404\n Page Not Found"
	case 500:
		errString = "Error code 500\n Internal Server Error "
	case 501:
		errString = "Error code 501\n Not implemented"
	case 503:
		errString = "Error code 503\n Service Anavailable"
	default:
		errString = "知らないエッラが現れた。"
	}
	w.WriteHeader(200) //NOTE: I return 200 to make the HTMX swap the thing
	if err_code == 403 {
		component := ui.ErrorRespons(errString, true)
		component.Render(r.Context(), w)
		return
	}
	component := ui.ErrorRespons(errString, false)
	component.Render(r.Context(), w)
}
