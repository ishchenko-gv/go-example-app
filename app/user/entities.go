package user

import (
	"encoding/json"

	"github.com/google/uuid"
)

type UserID uuid.UUID

func (u UserID) MarshalJSON() ([]byte, error) {
	return json.Marshal(uuid.UUID(u).String())
}

func NewUserID() UserID {
	return UserID(uuid.New())
}

type User struct {
	ID    UserID `json:"id"`
	Email string `json:"email"`
}

func NewUser() *User {
	return &User{
		ID: NewUserID(),
	}
}
