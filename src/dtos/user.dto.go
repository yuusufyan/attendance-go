package dtos

type UserCreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserName string `json:"user_name"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
}
