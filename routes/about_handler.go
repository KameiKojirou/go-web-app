package routes

import (
	"encoding/json"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	http.Header.Set(w.Header(), "Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "About Page",
	})
}
