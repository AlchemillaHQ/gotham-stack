package handlers

import (
	"github.com/AlchemillaHQ/gotham-stack/internal"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, templ *template.Template) {
	data := struct{ Title string }{Title: "Home"}
	internal.RenderTemplate(w, "dashboard", data, templ)
}
