package routes

import (
	middle "main/middleware"
	"net/http"
)


func APIRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", http.NotFoundHandler())
	mux.Handle("/home", middle.ExampleMiddle(http.HandlerFunc(HomeHandler)))
	mux.Handle("/about", middle.ExampleMiddlewareStack(http.HandlerFunc(AboutHandler)))
	return mux
}