package usecase

import (
	"fmt"
	"redishop/model"
	"redishop/repository"
)

type UserUsecase interface {
	CreateUser(user model.User) error
}

type userUsecase struct {
	//isinya adalah object nya
	userRepo repository.UserRepo
}

func (u userUsecase) CreateUser(user model.User) error {
	//TODO implement me

	if user.Name == "" {
		return fmt.Errorf("Name is required")
	}
	if user.Age == "" {
		return fmt.Errorf("Age is required")
	}

	//repository
	if err := u.userRepo.CreateUser(user); err != nil {
		return fmt.Errorf("Failed to insert %v", err.Error())
	}
	return nil
}

func NewUserUsecase(userRepo repository.UserRepo) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}
