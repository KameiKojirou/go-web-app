package routes

import "net/http"

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About"))
}