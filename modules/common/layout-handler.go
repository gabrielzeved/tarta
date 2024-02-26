package common

import (
	"net/http"

	"github.com/a-h/templ"
	"tarta.com/modules/pages"
)

func LayoutHandler(page templ.Component) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		useLayout := req.Header.Get("hx-request") != "true"
		if useLayout {
			pages.Layout(page).Render(req.Context(), rw)
		} else {
			page.Render(req.Context(), rw)
		}
	})

}
