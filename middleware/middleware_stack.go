package middle

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func CreateStack(middleware ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for _, m := range middleware {
			next = m(next)
		}
		return next
	}
}

