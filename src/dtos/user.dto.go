package dtos

type UserCreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserName string `json:"user_name"`
}
