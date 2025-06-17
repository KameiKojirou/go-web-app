package main

import (
	"embed"
	"main/routes"
	"net/http"
)

//go:embed frontend/dist/*
//go:embed frontend/dist/**/*
var frontendFS embed.FS

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", routes.APIRouter()))
	mux.Handle("/", SPAHandler("index.html"))
	http.ListenAndServe(":1323", mux)
}