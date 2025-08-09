package internal

import (
	"context"

	"github.com/ishchenko-gv/go-example-app/app/user"
)

type Service struct {
	Repo RepoInterface
}

var _ user.Service = (*Service)(nil)

func (s *Service) CreateUser(ctx context.Context, user *user.User, password string) error {
	return s.Repo.Insert(ctx, user, password)
}
