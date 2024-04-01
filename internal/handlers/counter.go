package handlers

import (
	"fmt"
	"github.com/AlchemillaHQ/gotham-stack/internal"
	"github.com/AlchemillaHQ/gotham-stack/internal/services"
	"net/http"
)

func CounterHandler(w http.ResponseWriter, r *http.Request, counterService *services.CounterService) {
	userId := internal.GetUserID(r)

	if userId == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	switch r.Method {
	case "GET":
		handleGetCounter(w, r, counterService, userId)
	default:
		handleModifyCounter(w, r, counterService, userId)
	}
}

func handleGetCounter(w http.ResponseWriter, r *http.Request, counterService *services.CounterService, userId int) {
	count, err := counterService.GetCount(userId)
	if err != nil {
		http.Error(w, "Failed to get count", http.StatusInternalServerError)
		return
	}

	html := fmt.Sprintf(`<div class="text-2xl mb-4" id="counter">%d</div>`, count)
	SendHTMLResponse(w, html)
}

func handleModifyCounter(w http.ResponseWriter, r *http.Request, counterService *services.CounterService, userId int) {
	aType := r.PathValue("type")
	var count int
	var err error

	switch aType {
	case "increment":
		count, err = counterService.IncrementCount(userId)
	case "decrement":
		count, err = counterService.DecrementCount(userId)
	default:
		http.Error(w, "Invalid counter type", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "Failed to modify count", http.StatusInternalServerError)
		return
	}

	html := fmt.Sprintf(`<div class="text-2xl mb-4" id="counter">%d</div>`, count)
	SendHTMLResponse(w, html)
}
