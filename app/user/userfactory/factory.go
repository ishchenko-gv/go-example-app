package userfactory

import (
	"database/sql"

	"github.com/ishchenko-gv/go-example-app/app/user/internal"
)

func NewService(repo internal.Repository) *internal.Service {
	return &internal.Service{
		Repo: repo,
	}
}

func NewRepo(db *sql.DB) *internal.Repo {
	return &internal.Repo{
		DB: db,
	}
}
