package users

type RegisterRequest struct {
	Username string  `json:"username" validate:"required,min=4,max=64"`
	Password string  `json:"password" validate:"required,min=8"`
	Email    string  `json:"email" validate:"required,email"`
	Name     *string `json:"name"`
	Lastname *string `json:"lastname"`
}

type UserDTO struct {
	Username string `json:"username" validate:"required,min=4,max=64"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}
