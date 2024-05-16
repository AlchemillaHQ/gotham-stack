package handlers

import (
	"github.com/AlchemillaHQ/gotham-stack/internal/middleware"
	"net/http"
	"reflect"
)

func SendHTMLResponse(w http.ResponseWriter, html string) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}

func GetPageData(data interface{}, r *http.Request) map[string]interface{} {
	v := reflect.ValueOf(data)
	dataMap := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		dataMap[v.Type().Field(i).Name] = v.Field(i).Interface()
	}

	dataMap["Lang"] = middleware.GetSessionLanguage(r)

	return dataMap
}
