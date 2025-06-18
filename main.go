package main

import (
	"main/routes"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", routes.APIRouter()))
	mux.Handle("/", SPAHandler("index.html"))
	http.ListenAndServe(":1323", mux)
}