package usecase

import (
	"errors"
	"nongki/internal/domain"
	"nongki/internal/repository"
	"nongki/internal/request"
	"nongki/pkg/constant"
)

type UserUsecase interface {
	GetMe(userID string) (*domain.User, error)
	UpdateUser(userID string, req request.UpdateUserRequest) (*domain.User, error)
	DeleteUser(userID string, deletedBy string) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) GetMe(userID string) (*domain.User, error) {
	data, err := u.userRepo.GetMe(userID)
	if data == nil {
		return nil, errors.New(constant.DATA_NOT_FOUND)
	}

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *userUsecase) UpdateUser(userID string, req request.UpdateUserRequest) (*domain.User, error) {
	existingUser, err := u.userRepo.GetMe(userID)
	if err != nil {
		return nil, err
	}

	if existingUser == nil {
		return nil, errors.New(constant.DATA_NOT_FOUND)
	}

	existingUser.Name = req.Name
	existingUser.Email = req.Email
	existingUser.Address = req.Address
	existingUser.Gender = req.Gender
	existingUser.MaritalStatus = req.MaritalStatus
	data, err := u.userRepo.UpdateUser(userID, *existingUser)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *userUsecase) DeleteUser(userID string, deletedBy string) error {
	err := u.userRepo.DeleteUser(userID, deletedBy)
	if err != nil {
		return err
	}

	return nil
}
