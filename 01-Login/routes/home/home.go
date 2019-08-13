package home

import (
	templates ".."

	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, "home", nil)
}
