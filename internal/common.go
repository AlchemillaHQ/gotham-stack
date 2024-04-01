package internal

import (
	"github.com/AlchemillaHQ/gotham-stack/internal/logger"
	"github.com/AlchemillaHQ/gotham-stack/internal/middleware"
	"go.uber.org/zap"
	"html/template"
	"net/http"
)

func GetUserID(r *http.Request) int {
	session, err := middleware.Store.Get(r, "session")
	if err != nil {
		return 0
	}

	if userID, ok := session.Values["userID"].(int); ok {
		return userID
	}

	return 0
}

func RenderTemplate(w http.ResponseWriter, tmplName string, data interface{}, templ *template.Template) {
	if err := templ.ExecuteTemplate(w, tmplName, data); err != nil {
		logger.Error("Error rendering template", zap.Error(err))
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
