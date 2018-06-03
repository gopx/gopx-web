package route

import (
	"net/http"
	"path"
	"strings"

	"gopx.io/gopx-web/pkg/controller"
	"gopx.io/gopx-web/pkg/log"
)

// GoPXWebRouter handles requested HTTP route, process it and hand over
// the specific controller.
type GoPXWebRouter struct{}

func (gr GoPXWebRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Info("%s %s", strings.ToUpper(r.Method), r.RequestURI)
	processRoute(w, r)
}

func processRoute(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = sanitizeRoute(r.URL.Path)
	path := r.URL.Path

	switch {
	case matchRoute(path, "/"):
		controller.Index(w, r)
	case matchRoute(path, "/static"):
		controller.Static(w, r)
	case strings.HasPrefix(path, "/@"):
		controller.User(w, r)
	case matchRoute(path, "/settings"):
		controller.Settings(w, r)
	case matchRoute(path, "/about"):
		controller.About(w, r)
	case matchRoute(path, "/login"):
		controller.Login(w, r)
	case matchRoute(path, "/signup"):
		controller.Signup(w, r)
	default:
		controller.Package(w, r)
	}
}

// Here requested route needs to be converted to lower case,
// which enables "/About" is equivalent to "/about" etc.
// and finally cleans the path e.g. end slashes would be removed from path
// e.g. "/about/" -> "/about" etc.
func sanitizeRoute(route string) string {
	return path.Clean(strings.ToLower(route))
}

// Here param path and route should be cleaned.
func matchRoute(path, route string) bool {
	if route == "/" {
		return path == "/"
	}
	return route == path || strings.HasPrefix(path, route+"/")
}

// NewGoPXWebRouter returns a new GoPXWebRouter instance.
func NewGoPXWebRouter() *GoPXWebRouter {
	return &GoPXWebRouter{}
}
