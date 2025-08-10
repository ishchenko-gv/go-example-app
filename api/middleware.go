package api

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/ishchenko-gv/go-example-app/api/apictx"
	"github.com/ishchenko-gv/go-example-app/api/userapi"
	"github.com/ishchenko-gv/go-example-app/app/user"
)

type middleware struct {
	UserService user.Service
}

func NewMiddleware(userService user.Service) *middleware {
	return &middleware{
		UserService: userService,
	}
}

type middlewareFunc func(next http.Handler) http.Handler

func chainMiddlewares(middlewares ...middlewareFunc) middlewareFunc {
	return func(next http.Handler) http.Handler {
		curr := next
		for _, middleware := range middlewares {
			curr = middleware(curr)
		}

		return curr
	}
}

func (m *middleware) PanicRecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				switch t := r.(type) {
				case error:
					fmt.Printf("panic: %v", t.Error())
				default:
					fmt.Printf("panic: %v", t)
				}
				http.Error(w, "internal error", http.StatusInternalServerError)
				return
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (m *middleware) JsonContentMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func (m *middleware) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		reqDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Printf("failed to dump request: %s\n", err.Error())
			http.Error(rw, "internal error", http.StatusInternalServerError)
			return
		}

		fmt.Printf("%s\n\n", string(reqDump))

		next.ServeHTTP(rw, r)
	})
}

func (m *middleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		t, err := r.Cookie("a")
		if errors.Is(err, http.ErrNoCookie) {
			next.ServeHTTP(rw, r)
			return
		}

		if err != nil {
			fmt.Printf("failed to read cookie: %s\n", err.Error())
			http.Error(rw, "internal error", http.StatusInternalServerError)
			return
		}

		claims, err := userapi.VerifyJwt(t.Value)
		if err != nil {
			fmt.Printf("failed to verify jwt: %s\n", err.Error())
			http.Error(rw, "internal error", http.StatusInternalServerError)
			return
		}

		usr, err := m.UserService.GetUser(r.Context(), claims.ID)
		if err != nil {
			fmt.Printf("failed to fetch user: %s\n", err.Error())
			http.Error(rw, "internal error", http.StatusInternalServerError)
			return
		}

		r = apictx.SetUser(r, usr)

		next.ServeHTTP(rw, r)
	})
}
