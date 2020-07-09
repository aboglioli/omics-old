package auth

import "omics/pkg/models"

type LoginRequest struct {
	UsernameOrEmail string `json:"username"`
	Password        string `json:"password"`
}

type LoginResponse struct {
	AuthToken string `json:"auth_token"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Password string `json:"password"`
}

func (req *RegisterRequest) Validate() error {
	return nil
}

type ChangePasswordRequest struct {
	UserID      models.ID `json:"user_id"`
	OldPassword string    `json:"old_password"`
	NewPassword string    `json:"new_password"`
}

func (req *ChangePasswordRequest) Validate() error {
	return nil
}

type ValidateRequest struct {
	UserID models.ID `json:"user_id"`
	Code   string    `json:"code"`
}
