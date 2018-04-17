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

// User handles HTTP request on user info and user scoped package route.
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
		scopedPackage(user, route, w, r)
	}
}

// userProfile handles user info e.g. public profile, packages, organizations etc.
// It uses path query to switch between user routes,
// orgs for packages -> @gopx?tab=packages
// 			for organizations -> @gopx?tab=orgs
func userProfile(user string, w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is about "+user)
}

// It should not have any reserved scoped package name,
// the user has full freedom on package naming on their own scope.
func scopedPackage(user, pkg string, w http.ResponseWriter, r *http.Request) {
	route := strings.TrimPrefix(pkg, "/")
	method := strings.ToUpper(r.Method)
	var subPkg string

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
		scopedPackageGoGetMeta(user, pkg, subPkg, w, r)
		return
	}

	if subPkg == "/" {
		switch method {
		case "GET":
			scopedPackageGET(user, pkg, w, r)
		default:
			Error405(w, r)
		}
	} else {
		switch method {
		case "GET":
			scopedSubPackageGET(user, pkg, subPkg, w, r)
		default:
			Error405(w, r)
		}
	}
}

func scopedPackageGET(user, pkg string, w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("Scoped package: @%s/%s", user, pkg))
}

func scopedSubPackageGET(user, pkg, subPkg string, w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("Scoped sub-package: @%s/%s%s", user, pkg, subPkg)) // subPkg starts with "/"
}

func scopedPackageGoGetMeta(user, pkg, subPkg string, w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"goImport": scopedPackageGoImportMeta(user, pkg),
		"goSource": scopedPackageGoSourceMeta(user, pkg),
		"message":  fmt.Sprintf("go get %s/%s", config.Host.Web, path.Join("@"+user, pkg, subPkg)),
	}
	err := template.Render("go-get", w, data)
	if err != nil {
		Error500(w, r)
		log.Error("Error: %s", err)
	}
}

func scopedPackageGoImportMeta(user, pkg string) string {
	return fmt.Sprintf("%s/@%s/%s git %s://%s/@%s/%s", config.Host.Web, user, pkg, "http", config.Host.VCS, user, pkg)
}

func scopedPackageGoSourceMeta(user, pkg string) string {
	return ""
}
