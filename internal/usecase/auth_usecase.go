package usecase

import (
	"simpleorder/internal/domain"
	"simpleorder/internal/repository"
	"simpleorder/pkg/utils"
)

type AuthUsecase interface {
	Register(req *domain.UserRegistrationRequest) error
	Login(req *domain.UserLoginRequest, secret string, expHours int) (*domain.UserLoginResponse, error)
}

type authUsecase struct {
	userRepo repository.UserRepository
}

func NewAuthUsecase(ur repository.UserRepository) AuthUsecase {
	return &authUsecase{
		userRepo: ur,
	}
}

func (u *authUsecase) Register(req *domain.UserRegistrationRequest) error {
	// Check if user exists
	_, err := u.userRepo.FindByUsername(req.Username)
	if err == nil {
		return domain.ErrUsernameExists
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return domain.ErrInternalServer
	}

	role := "user"
	if req.Role != "" {
		role = req.Role
	}

	user := &domain.User{
		Username: req.Username,
		Password: hashedPassword,
		Role:     role,
	}

	return u.userRepo.Create(user)
}

func (u *authUsecase) Login(req *domain.UserLoginRequest, secret string, expHours int) (*domain.UserLoginResponse, error) {
	user, err := u.userRepo.FindByUsername(req.Username)
	if err != nil {
		return nil, domain.ErrInvalidPassword // hide actual error
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, domain.ErrInvalidPassword
	}

	token, err := utils.GenerateToken(user.ID, user.Username, user.Role, secret, expHours)
	if err != nil {
		return nil, domain.ErrInternalServer
	}

	return &domain.UserLoginResponse{
		Token: token,
		User:  *user,
	}, nil
}
