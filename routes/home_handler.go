package routes

import (
"net/http"
"encoding/json"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Header.Set(w.Header(), "Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Home Page",
	})
}