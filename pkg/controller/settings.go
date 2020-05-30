package controller

import (
	"io"
	"net/http"
	"strings"
)

// Settings handles current user account preferences
// e.g. profile, emails, notifications, billing,
// SSH or GPG keys etc.
func Settings(w http.ResponseWriter, r *http.Request) {
	route := strings.TrimPrefix(r.URL.Path, "/settings")

	if route == "" {
		route = "/"
	}

	switch route {
	case "/":
		http.Redirect(w, r, "/settings/profile", http.StatusFound)
	case "/profile":
		settingsProfile(w, r)
	case "/account":
		settingsAccount(w, r)
	default:
		Error404(w, r)
	}
}

func settingsProfile(w http.ResponseWriter, r *http.Request) {
	// TODO: Handle other request methods
	io.WriteString(w, "Settings: Profile")
}

func settingsAccount(w http.ResponseWriter, r *http.Request) {
	// TODO: Handle other request methods
	io.WriteString(w, "Settings: Account")
}
