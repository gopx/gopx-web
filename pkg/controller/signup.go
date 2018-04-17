package controller

import (
	"net/http"
	"strings"

	"gopx.io/gopx-web/pkg/log"
	"gopx.io/gopx-web/pkg/template"
)

// Signup handles HTTP request on "/signup" route.
func Signup(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	method := strings.ToUpper(r.Method)

	if path == "/signup" || path == "/signup/" {
		switch method {
		case "GET":
			signupGET(w, r)
		case "POST":
			signupPOST(w, r)
		default:
			Error405(w, r)
		}
	} else {
		Error404(w, r)
	}
}

func signupGET(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"title":   "GoPX | Signup",
		"message": "This is signup page",
	}
	err := template.Render("signup", w, data)
	if err != nil {
		Error500(w, r)
		log.Error("Error: %s", err)
	}
}

func signupPOST(w http.ResponseWriter, r *http.Request) {

}
