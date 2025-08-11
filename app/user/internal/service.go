package internal

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ishchenko-gv/go-example-app/app/user"
	"github.com/ishchenko-gv/go-example-app/app/user/userid"
)

type Service struct {
	Repo Repository
}

var _ user.Service = (*Service)(nil)

func (s *Service) Register(ctx context.Context, user *user.User, password string) error {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}

	return s.Repo.Insert(ctx, user, hashedPassword)
}

func (s *Service) GetUser(ctx context.Context, userID userid.ID) (*user.User, error) {
	return s.Repo.Find(ctx, userID)
}

func (s *Service) AuthenticateByEmail(ctx context.Context, email string, password string) (*user.User, error) {
	usr, hashedPassword, err := s.Repo.FindByEmail(ctx, email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, user.ErrInvalidCredentials
	}

	if err != nil {
		return nil, err
	}

	if !checkPassword(hashedPassword, password) {
		return nil, user.ErrInvalidCredentials
	}

	return usr, nil
}
