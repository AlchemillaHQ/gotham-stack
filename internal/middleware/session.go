package middleware

import (
	"github.com/AlchemillaHQ/gotham-stack/internal/config"
	"github.com/gorilla/sessions"
	"net/http"
)

var Store *sessions.CookieStore

func InitializeSessionStore(secret string) {
	if config.ParsedConfig.SessionSecret == "" {
		panic("Session secret cannot be empty")
	}
	Store = sessions.NewCookieStore([]byte(secret))
}

func EnsureAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := Store.Get(r, "session")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(w, r, "/auth/login", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func SetSession(w http.ResponseWriter, r *http.Request, email string, userID int) error {
	session, err := Store.Get(r, "session")
	if err != nil {
		return err
	}
	session.Values["authenticated"] = true
	session.Values["email"] = email
	session.Values["userID"] = userID
	return session.Save(r, w)
}

func ClearSession(w http.ResponseWriter, r *http.Request) error {
	session, err := Store.Get(r, "session")
	if err != nil {
		return err
	}
	session.Values["authenticated"] = false
	session.Values["email"] = ""
	return session.Save(r, w)
}
