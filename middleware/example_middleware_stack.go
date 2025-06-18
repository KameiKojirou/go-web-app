package middle

import (
	"net/http"
)

func ExampleMiddlewareStack(next http.Handler) http.Handler {
	return CreateStack(
		ExampleMiddle,
		ExampleMiddleTwo,
	)(next)
}