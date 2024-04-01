package handlers

import (
	"github.com/AlchemillaHQ/gotham-stack/internal"
	"github.com/AlchemillaHQ/gotham-stack/internal/middleware"
	"github.com/AlchemillaHQ/gotham-stack/internal/services"
	"html/template"
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request, templ *template.Template, authService *services.AuthService) {
	aType := r.URL.Path[6:]

	switch aType {
	case "":
		http.Redirect(w, r, "/auth/login", http.StatusFound)
	case "login":
		handleLogin(w, r, templ, authService)
	case "logout":
		handleLogout(w, r)
	default:
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request, templ *template.Template, authService *services.AuthService) {
	switch r.Method {
	case "GET":
		internal.RenderTemplate(w, "auth/login", struct{ Title string }{"Login"}, templ)
	case "POST":
		email := r.FormValue("email")
		password := r.FormValue("password")
		_, ok := authService.VerifyCredentials(email, password)
		if !ok {
			http.Redirect(w, r, "/auth/login?error=invalid-credentials", http.StatusFound)
			return
		}
		if err := middleware.SetSession(w, r, email, 1); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	if err := middleware.ClearSession(w, r); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/auth/login?success=logged-out", http.StatusFound)
}
