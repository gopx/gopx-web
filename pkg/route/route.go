package route

import (
	"net/http"

	"gopx.io/gopx-web/pkg/controller"
	"gopx.io/gopx-web/pkg/log"
)

// GoPXRouter handles requested HTTP route, process it and hand over
// the specific controller.
type GoPXRouter struct{}

func (gr GoPXRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Info("New connection from %s", r.RemoteAddr)
	controller.IndexGet(w, r)
}

// NewGoPXRouter returns a new GoPXRouter instance.
func NewGoPXRouter() *GoPXRouter {
	return &GoPXRouter{}
}
