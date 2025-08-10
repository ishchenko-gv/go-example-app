package internal

import (
	"context"
	"database/sql"

	"github.com/ishchenko-gv/go-example-app/app/user"
	"github.com/ishchenko-gv/go-example-app/app/user/userid"
)

type Repo struct {
	DB *sql.DB
}

type Repository interface {
	Insert(ctx context.Context, user *user.User, password string) error
	Find(ctx context.Context, id userid.ID) (*user.User, error)
	FindByEmail(ctx context.Context, email string) (*user.User, string, error)
}

var _ Repository = (*Repo)(nil)

func (r *Repo) Insert(ctx context.Context, user *user.User, password string) error {
	_, err := r.DB.Exec(
		"INSERT INTO users (id, email, password) VALUES ($1, $2, $3)",
		user.ID.String(),
		user.Email,
		password,
	)

	return err
}

func (r *Repo) Find(ctx context.Context, userID userid.ID) (*user.User, error) {
	usr := &user.User{}

	q := "SELECT id, email FROM users WHERE id = $1"
	err := r.DB.QueryRow(q, userID).Scan(&usr.ID, &usr.Email)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (r *Repo) FindByEmail(ctx context.Context, email string) (*user.User, string, error) {
	usr := &user.User{}
	var pwd string

	q := "SELECT id, email, password FROM users WHERE email = $1"
	err := r.DB.QueryRow(q, email).Scan(&usr.ID, &usr.Email, &pwd)
	if err != nil {
		return nil, "", err
	}

	return usr, pwd, nil
}
