package usecase

import (
	"attendance-go/src/dtos"
	"attendance-go/src/models"
	"attendance-go/src/repositories"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userRepository *repositories.UserRepository
}

func NewUserUseCase(userRepository *repositories.UserRepository) *UserUseCase {
	return &UserUseCase{userRepository: userRepository}
}

func (u *UserUseCase) Create(request *dtos.UserCreateRequest) error {
	_, err := u.userRepository.GetByEmail(request.Email)
	if err == nil {
		return errors.New("email already registered")
	}

	_, err = u.userRepository.GetByUsername(request.UserName)
	if err == nil {
		return errors.New("username already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.UserModels{
		Email:    request.Email,
		Password: string(hashedPassword),
		Username: request.UserName,
		IsActive: true,
	}
	err = u.userRepository.Create(user)
	if err != nil {
		return err
	}
	return nil
}
