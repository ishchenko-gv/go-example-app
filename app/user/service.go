package user

import (
	"context"
)

type Service interface {
	CreateUser(ctx context.Context, user *User, password string) error
	AuthenticateByEmail(ctx context.Context, email string, password string) (*User, error)
}
