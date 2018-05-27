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
		Error404(w, r)
		if !os.IsNotExist(err) {
			log.Error("Error: %s", err)
		}
		return
	}

	// Prevent folder content and hidden file surfing.
	//
	// TODO: Detect a path is a hidden e.g. "/static/a/.b/c.txt", here ".b"
	// is hidden so, the path "/static/a/.b/c.txt" is also hidden path, so it can't be
	// accessed.
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
