package usersdto

type CreateUserRequest struct {
	FullName string `json:"fullName" validate:"required"`
	UserName string `json:"userName" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

type UserResponse struct {
	Id       int    `json:"id"`
	FullName string `json:"fullName"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}