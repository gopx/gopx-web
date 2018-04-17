package controller

import (
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"

	"gopx.io/gopx-web/pkg/config"

	"gopx.io/gopx-web/pkg/log"
	"gopx.io/gopx-web/pkg/template"
)

// Package handles HTTP request on package and sub package routes.
func Package(w http.ResponseWriter, r *http.Request) {
	route := strings.TrimPrefix(r.URL.Path, "/")
	method := strings.ToUpper(r.Method)
	var pkg, subPkg string

	if idx := strings.Index(route, "/"); idx == -1 {
		pkg = route
		subPkg = "/"
	} else {
		pkg = route[:idx]
		subPkg = route[idx:]
	}

	// Responds with appropriate go-get meta info irrespective
	// on request method e.g. GET, POST.
	if goGet := r.FormValue("go-get"); goGet == "1" {
		packageGoGetMeta(pkg, subPkg, w, r)
		return
	}

	if subPkg == "/" {
		switch method {
		case "GET":
			packageGET(pkg, w, r)
		default:
			Error405(w, r)
		}
	} else {
		switch method {
		case "GET":
			subPackageGET(pkg, subPkg, w, r)
		default:
			Error405(w, r)
		}
	}
}

func packageGET(pkg string, w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Package: "+pkg)
}

func subPackageGET(pkg, subPkg string, w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Package: "+pkg+", sub package: "+subPkg)
}

func packageGoGetMeta(pkg, subPkg string, w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"goImport": packageGoImportMeta(pkg),
		"goSource": packageGoSourceMeta(pkg),
		"message":  fmt.Sprintf("go get %s/%s", config.Host.Web, path.Join(pkg, subPkg)),
	}
	err := template.Render("go-get", w, data)
	if err != nil {
		Error500(w, r)
		log.Error("Error: %s", err)
	}
}

func packageGoImportMeta(pkg string) string {
	return fmt.Sprintf("%s/%s git %s://%s/%s", config.Host.Web, pkg, "http", config.Host.VCS, pkg)
}

func packageGoSourceMeta(pkg string) string {
	return ""
}
