package controller

import (
	"net/http"
	"strings"

	"gopx.io/gopx-web/pkg/log"
	"gopx.io/gopx-web/pkg/template"
)

// Index handles HTTP request on "/" route.
func Index(w http.ResponseWriter, r *http.Request) {
	switch method := strings.ToUpper(r.Method); method {
	case "GET":
		indexGET(w, r)
	default:
		Error405(w, r)
	}
}

func indexGET(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"title":   "GoPX | Home",
		"message": "GoPX: Package registry for GoLang",
	}
	err := template.Render("index", w, data)
	if err != nil {
		Error500(w, r)
		log.Error("Error: %s", err)
	}
}
