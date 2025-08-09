package apictx

import (
	"context"
	"net/http"

	"github.com/ishchenko-gv/go-example-app/app/user"
)

type ctxKey string

const userCtxKey ctxKey = "user"

func User(r *http.Request) *user.User {
	return r.Context().Value(userCtxKey).(*user.User)
}

func SetUser(r *http.Request, user *user.User) *http.Request {
	ctx := context.WithValue(r.Context(), userCtxKey, user)
	return r.WithContext(ctx)
}
