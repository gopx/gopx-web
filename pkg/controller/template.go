package controller

import (
	"html/template"
)

var templates = template.Must(template.ParseGlob("./we"))
