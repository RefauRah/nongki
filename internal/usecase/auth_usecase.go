package usecase

import (
	"errors"
	"nongki/internal/domain"
	"nongki/internal/repository"
	"nongki/internal/request"
	"nongki/pkg/constant"
	"nongki/pkg/jwt"
	"time"
)

type AuthUsecase interface {
	Register(req request.RegisterRequest) (*domain.User, string, error)
	Login(req request.LoginRequest) (*domain.User, string, error)
}

type authUsecase struct {
	userRepo repository.UserRepository
}

func NewAuthUsecase(userRepo repository.UserRepository) AuthUsecase {
	return &authUsecase{userRepo: userRepo}
}

func (u *authUsecase) Register(req request.RegisterRequest) (*domain.User, string, error) {
	existingUser, _ := u.userRepo.FindByEmail(req.Email)
	if existingUser != nil {
		return nil, "", errors.New(constant.EMAIL_EXIST)
	}

	now := time.Now()
	var user = &domain.User{
		Name:          req.Name,
		Email:         req.Email,
		Address:       req.Address,
		Gender:        req.Gender,
		MaritalStatus: req.MaritalStatus,
	}

	err := user.HashPassword(req.Password)
	if err != nil {
		return nil, "", err
	}

	user.CreatedAt = now
	user.UpdatedAt = &now
	user.CreatedBy = &req.Email
	user.UpdatedBy = &req.Email

	err = u.userRepo.Create(user)
	if err != nil {
		return nil, "", err
	}

	token, err := jwt.GenerateJWT(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (u *authUsecase) Login(req request.LoginRequest) (*domain.User, string, error) {
	user, err := u.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, "", errors.New(constant.DATA_NOT_FOUND)
	}

	if err := user.CheckPassword(req.Password); err != nil {
		return nil, "", errors.New(constant.INVALID_PASSWORD)
	}

	token, err := jwt.CreateToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
