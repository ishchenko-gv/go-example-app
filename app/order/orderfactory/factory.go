package orderfactory

import "github.com/ishchenko-gv/go-example-app/app/order/internal"

func NewService(repo *internal.Repo) *internal.Service {
	return &internal.Service{
		Repo: repo,
	}
}

func NewRepo() *internal.Repo {
	return &internal.Repo{}
}
