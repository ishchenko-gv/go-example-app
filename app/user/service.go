package user

import (
	"context"

	"github.com/ishchenko-gv/go-example-app/app/user/userid"
)

type Service interface {
	CreateUser(ctx context.Context, user *User, password string) error
	GetUser(ctx context.Context, userID userid.ID) (*User, error)
	AuthenticateByEmail(ctx context.Context, email string, password string) (*User, error)
}
