package middle

import (
	"fmt"
	"net/http"
)

func ExampleMiddleTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("example middleware two")
		next.ServeHTTP(w, r)
	})
}