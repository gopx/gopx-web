package controller

import (
	"net/http"

	"gopx.io/gopx-web/pkg/log"
	"gopx.io/gopx-web/pkg/template"
)

// Error403 handles forbidden request.
func Error403(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
	data := map[string]string{
		"title":   "403 Forbidden",
		"message": r.URL.Path,
	}
	err := template.Render("error", w, data)
	if err != nil {
		log.Error("Error: %s", err)
	}
}

// Error404 handles HTTP request on non-existing routes.
func Error404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	data := map[string]string{
		"title":   "404 Not Found",
		"message": r.URL.Path,
	}
	err := template.Render("error", w, data)
	if err != nil {
		log.Error("Error: %s", err)
	}
}

// Error405 handles not allowed http method.
func Error405(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	data := map[string]string{
		"title":   "405 Method Not Allowed",
		"message": r.URL.Path,
	}
	err := template.Render("error", w, data)
	if err != nil {
		log.Error("Error: %s", err)
	}
}

// Error500 handles internal server error.
func Error500(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	data := map[string]string{
		"title":   "500 Internal Server Error",
		"message": r.URL.Path,
	}
	err := template.Render("error", w, data)
	if err != nil {
		log.Error("Error: %s", err)
	}
}
