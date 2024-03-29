package auth

import (
	"context"

	"omics/internal/domain/auth"
	"omics/internal/domain/roles"
)

type Service interface {
	Register(ctx context.Context, req *RegisterRequest) error
	Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error)
	ChangePassword(ctx context.Context, req *ChangePasswordRequest) error
	Logout(ctx context.Context) error
	Validate(ctx context.Context, req *ValidateRequest) error
}

type service struct {
	userRepo        auth.UserRepository
	authRepo        auth.AuthRepository
	roleRepo        roles.Repository
	passwordCrypter auth.PasswordCrypter
	tokenEncoder    auth.Encoder
	authServ        auth.Service
}

func NewService(
	userRepo auth.UserRepository,
	authRepo auth.AuthRepository,
	roleRepo roles.Repository,
	passwordCrypter auth.PasswordCrypter,
	tokenEncoder auth.Encoder,
	authServ auth.Service,
) Service {
	return &service{
		userRepo:        userRepo,
		authRepo:        authRepo,
		roleRepo:        roleRepo,
		passwordCrypter: passwordCrypter,
		tokenEncoder:    tokenEncoder,
		authServ:        authServ,
	}
}

func (s *service) Register(ctx context.Context, req *RegisterRequest) error {
	if err := req.Validate(); err != nil {
		return ErrNull
	}

	if user, err := s.userRepo.FindByUsername(ctx, req.Username); user != nil || err == nil {
		if user, err := s.userRepo.FindByEmail(ctx, req.Email); user != nil || err == nil {
			return ErrNull
		}
	}

	authUser := auth.NewUser(req.Username, req.Email)
	authUser.ChangePassword(s.passwordCrypter, "", req.Password)

	role, err := s.roleRepo.FindByCode(ctx, "user")
	if err != nil {
		return ErrNull
	}

	permissions := make([]auth.Permission, 0)
	for _, perm := range role.Permissions {
		permissions = append(permissions, auth.Permission{
			Permission: perm.Permission,
			Module:     perm.Module.Code,
		})
	}

	authUser.Role = auth.Role{
		Code:        role.Code,
		Permissions: permissions,
	}

	if err := s.userRepo.Save(ctx, authUser); err != nil {
		return ErrNull
	}

	return nil
}

func (s *service) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	user, err := s.userRepo.FindByUsername(ctx, req.UsernameOrEmail)
	if err != nil {
		user, err = s.userRepo.FindByEmail(ctx, req.UsernameOrEmail)
		if err != nil {
			return nil, ErrNull
		}
	}

	if !user.ValidatePassword(s.passwordCrypter, req.Password) {
		return nil, ErrNull
	}

	token := auth.NewToken()
	encodedToken, err := token.Encode(s.tokenEncoder)
	if err != nil {
		return nil, ErrNull
	}

	if err := s.authRepo.Save(ctx, token.ID, user); err != nil {
		return nil, ErrNull
	}

	user.Login()

	return &LoginResponse{
		AuthToken: encodedToken,
	}, nil
}

func (s *service) ChangePassword(ctx context.Context, req *ChangePasswordRequest) error {
	user, err := s.userRepo.FindByID(ctx, req.UserID)
	if err != nil {
		return ErrNull
	}

	if err := user.ChangePassword(s.passwordCrypter, req.OldPassword, req.NewPassword); err != nil {
		return ErrNull
	}

	if err := s.userRepo.Save(ctx, user); err != nil {
		return ErrNull
	}

	return nil
}

func (s *service) Logout(ctx context.Context) error {
	token, err := s.authServ.GetToken(ctx)
	if err != nil {
		return ErrNull
	}

	if err := s.authRepo.Delete(ctx, token.ID); err != nil {
		return ErrNull
	}

	return nil
}

func (s *service) Validate(ctx context.Context, req *ValidateRequest) error {
	user, err := s.userRepo.FindByID(ctx, req.UserID)
	if err != nil {
		return ErrNull
	}

	// TODO: validate code
	user.Validate()

	if err := s.userRepo.Save(ctx, user); err != nil {
		return ErrNull
	}

	return nil
}
