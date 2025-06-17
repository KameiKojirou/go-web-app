package routes

import (
	"net/http"
)


func APIRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", http.NotFoundHandler())
	mux.Handle("/home", http.HandlerFunc(HomeHandler))
	mux.Handle("/about", http.HandlerFunc(AboutHandler))
	return mux
}