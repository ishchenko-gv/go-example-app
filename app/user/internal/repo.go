package internal

import (
	"context"
	"database/sql"

	"github.com/ishchenko-gv/go-example-app/app/user"
)

type Repo struct {
	DB *sql.DB
}

type RepoInterface interface {
	Insert(ctx context.Context, user *user.User, password string) error
}

var _ RepoInterface = (*Repo)(nil)

func (r *Repo) Insert(ctx context.Context, user *user.User, password string) error {
	_, err := r.DB.Exec(
		"INSERT INTO users (id, email, password) VALUES ($1, $2, $3)",
		user.ID.String(),
		user.Email,
		password,
	)

	return err
}
