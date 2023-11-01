package manager

import "redishop/usecase"

type UsecaseManager interface {
	//all usecase object goes here
	UserUsecase() usecase.UserUsecase
	UserCredUsecase() usecase.UserCredentialUsecase
	ProductUsecase() usecase.ProductUsecase
}

type usecaseManager struct {
	rm RepoManager
}

func (u usecaseManager) UserUsecase() usecase.UserUsecase {
	return usecase.NewUserUsecase(u.rm.UserRepo())
}

func (u usecaseManager) UserCredUsecase() usecase.UserCredentialUsecase {
	return usecase.NewUserCredentialUsecase(u.rm.UserCredRepo())
}

func (u usecaseManager) ProductUsecase() usecase.ProductUsecase {
	//TODO implement me
	return usecase.NewProductUsecase(u.rm.ProductRepo())
}

func NewUsecaseManager(rm RepoManager) UsecaseManager {
	return &usecaseManager{
		rm: rm,
	}
}
