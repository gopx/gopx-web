package controller

import (
	"net/http"

	"gopx.io/gopx-web/pkg/log"
	"gopx.io/gopx-web/pkg/template"
)

// IndexGet handles HTTP GET request on home page.
func IndexGet(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"title":   "GoPX | Home",
		"message": "GoPX: Package registry for GoLang",
	}
	err := template.Render("index", w, data)
	if err != nil {
		log.Error("Error: %s", err)
	}
}
