package controller

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"gopx.io/gopx-web/pkg/config"
	"gopx.io/gopx-web/pkg/log"
)

// Static handles HTTP request on static resources.
func Static(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	path := strings.TrimPrefix(r.URL.Path, "/static")
	resPath := filepath.Join(config.Web.StaticAssetPath, path)

	stat, err := os.Stat(resPath)
	if err != nil {
		if os.IsNotExist(err) {
			Error404(w, r)
		} else {
			Error500(w, r)
			log.Error("Error: %s", err)
		}
		return
	}

	if mode := stat.Mode(); !mode.IsRegular() || strings.HasPrefix(filepath.Base(resPath), ".") {
		Error403(w, r)
		return
	}

	switch method {
	case "GET":
		http.ServeFile(w, r, resPath)
	default:
		Error405(w, r)
	}
}
