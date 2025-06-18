package middle

import (
	"fmt"
	"net/http"
)

func ExampleMiddle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("example middleware")
		next.ServeHTTP(w, r)
	})
}