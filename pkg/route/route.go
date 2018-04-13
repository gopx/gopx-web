/*
Package route provides interfaces to handle GoPX web app routings.
*/
package route

import (
	"net/http"
)

// GoPXRouter handles requested route, process it and hand over
// the specific controller.
type GoPXRouter struct{}

func (gr GoPXRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

// NewGoPXRouter returns a new GoPXRouter instance.
func NewGoPXRouter() *GoPXRouter {
	return &GoPXRouter{}
}
