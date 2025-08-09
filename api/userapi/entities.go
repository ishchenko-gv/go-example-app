package userapi

type UserCreateBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
