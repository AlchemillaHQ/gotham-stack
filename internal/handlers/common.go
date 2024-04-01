package handlers

import "net/http"

func SendHTMLResponse(w http.ResponseWriter, html string) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}
