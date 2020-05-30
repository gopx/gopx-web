package controller

import (
	"net/http"
	"strings"

	"gopx.io/gopx-web/pkg/log"
	"gopx.io/gopx-web/pkg/template"
)

// Login handles HTTP request on "/login" route.
func Login(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	method := strings.ToUpper(r.Method)

	if path == "/login" {
		switch method {
		case "GET":
			loginGET(w, r)
		case "POST":
			loginPOST(w, r)
		default:
			Error405(w, r)
		}
	} else {
		Error404(w, r)
	}
}

func loginGET(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"title":   "GoPX | Login",
		"message": "This is login page",
	}
	err := template.Render("login", w, data)
	if err != nil {
		Error500(w, r)
		log.Error("Error: %s", err)
	}
}

func loginPOST(w http.ResponseWriter, r *http.Request) {

}
