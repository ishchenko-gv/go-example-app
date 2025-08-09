package user

import "github.com/ishchenko-gv/go-example-app/app/user/userid"

type User struct {
	ID    userid.ID `json:"id"`
	Email string    `json:"email"`
}

func NewUser(email string) *User {
	return &User{
		ID:    userid.New(),
		Email: email,
	}
}
