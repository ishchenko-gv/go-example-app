package api

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/ishchenko-gv/go-example-app/api/apictx"
	"github.com/ishchenko-gv/go-example-app/app/user"
)

type middleware func(next http.Handler) http.Handler

func chainMiddlewares(middlewares ...middleware) middleware {
	return func(next http.Handler) http.Handler {
		curr := next
		for _, middleware := range middlewares {
			curr = middleware(curr)
		}

		return curr
	}
}

func jsonResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		reqDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(rw, err.Error(), 500)
		}

		fmt.Printf("%s", string(reqDump))

		next.ServeHTTP(rw, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		user := user.NewUser()
		r = apictx.SetUser(r, user)

		next.ServeHTTP(rw, r)
	})
}
