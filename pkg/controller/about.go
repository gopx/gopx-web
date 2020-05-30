package controller

import (
	"net/http"
	"strings"

	"gopx.io/gopx-web/pkg/log"
	"gopx.io/gopx-web/pkg/template"
)

// About handles HTTP request on "/about" route.
func About(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	method := strings.ToUpper(r.Method)

	if path == "/about" {
		switch method {
		case "GET":
			aboutGET(w, r)
		default:
			Error405(w, r)
		}
	} else {
		Error404(w, r)
	}
}

func aboutGET(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"title":   "GoPX | About",
		"message": "This is about page",
	}
	err := template.Render("about", w, data)
	if err != nil {
		Error500(w, r)
		log.Error("Error: %s", err)
	}
}
