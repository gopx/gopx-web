package controller

import (
	"io"
	"net/http"
	"strings"
)

// User handles HTTP request on user route e.g. /@user.
func User(w http.ResponseWriter, r *http.Request) {
	route := strings.TrimPrefix(r.URL.Path, "/@")
	var user string

	if idx := strings.Index(route, "/"); idx == -1 {
		user = route
		route = "/"
	} else {
		user = route[:idx]
		route = route[idx:]
	}

	if route == "/" {
		userProfile(user, w, r)
	} else {
		Error404(w, r)
	}
}

// userProfile handles user info e.g. public profile, packages, organizations etc.
// It uses path query to switch between user routes,
// e.g. for packages -> @gopx?q=packages
// 			for organizations -> @gopx?q=orgs
func userProfile(user string, w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is about "+user)
}
