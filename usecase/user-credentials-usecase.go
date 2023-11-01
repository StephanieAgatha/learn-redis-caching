package usecase

import (
	"fmt"
	"os"
	"redishop/model"
	"redishop/repository"
	"redishop/util/helper"
)

type UserCredentialUsecase interface {
	Register(userCred model.UserCredentials) error
	Login(userCred model.UserCredentials) (string, error)
	FindUserEMail(email string) (userCred model.UserCredentials, err error)
}

type userCredentialUsecase struct {
	usercredRepo repository.UserCredential
}

func (u userCredentialUsecase) Register(userCred model.UserCredentials) error {
	//TODO implement me

	//generate uuid for user id
	userCred.ID = helper.GenerateUUID()

	if userCred.Email == "" {
		return fmt.Errorf("Username is required")
	}

	if userCred.Password == "" {
		return fmt.Errorf("Password is required")
	}

	//is email alr valid?
	if err := helper.IsEmailValid(userCred); err != nil {
		return err
	}

	/*
		password requirement
	*/
	if len(userCred.Password) < 6 {
		return fmt.Errorf("Password must contain at least six number")
	}
	if !helper.PasswordContainsUppercase(userCred.Password) {
		return fmt.Errorf("Password must contain at least one uppercase letter")
	}

	if !helper.PasswordContainsSpecialChar(userCred.Password) {
		return fmt.Errorf("Password must contain at least one special character")
	}

	if !helper.PasswordConstainsOneNumber(userCred.Password) {
		return fmt.Errorf("Password must contain at least one number")
	}

	//generate password in here
	hashedPass, err := helper.HashPassword(userCred.Password)
	if err != nil {
		return err
	}

	userCred.Password = hashedPass

	if err = u.usercredRepo.Register(userCred); err != nil {
		return err
	}

	return nil
}

func (u userCredentialUsecase) Login(userCred model.UserCredentials) (string, error) {
	//TODO implement me

	if userCred.Email == "" {
		return "", fmt.Errorf("Email is required")
	} else if userCred.Password == "" {
		return "", fmt.Errorf("Password is required")
	}

	userHashedPass, err := u.usercredRepo.Login(userCred)
	if err != nil {

	}
	//compare password
	if err = helper.ComparePassword(userHashedPass, userCred.Password); err != nil {
		return "", fmt.Errorf("Invalid Password")
	}

	//generate paseto or jwt in here
	symetricKey := os.Getenv("PASETO_SECRET")
	//jwtsecret := os.Getenv("JWT_SECRET")
	pasetoToken := helper.GeneratePaseto(userCred.Email, symetricKey)

	return pasetoToken, nil
}

func (u userCredentialUsecase) FindUserEMail(email string) (userCred model.UserCredentials, err error) {
	//TODO implement me

	if email == "" {
		return model.UserCredentials{}, fmt.Errorf("Email is required")
	}

	user, err := u.usercredRepo.FindUserEMail(email)
	if err != nil {
		return model.UserCredentials{}, err
	}

	return user, nil
}

func NewUserCredentialUsecase(uc repository.UserCredential) UserCredentialUsecase {
	return &userCredentialUsecase{
		usercredRepo: uc,
	}
}
