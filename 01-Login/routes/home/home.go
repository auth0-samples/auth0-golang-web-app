package home

import (
	"net/http"

	"01-Login/routes/templates"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, "home", nil)
}
