package manager

import "redishop/repository"

type RepoManager interface {
	UserRepo() repository.UserRepo
	UserCredRepo() repository.UserCredential
	ProductRepo() repository.ProductRepository
}

type repoManager struct {
	im InfraManager
}

func (r *repoManager) UserRepo() repository.UserRepo {
	//return constructor from repository user
	return repository.NewUserRepo(r.im.Connect())
}

func (r *repoManager) UserCredRepo() repository.UserCredential {
	return repository.NewUserCredentials(r.im.Connect())
}

func (r *repoManager) ProductRepo() repository.ProductRepository {
	//TODO implement me
	return repository.NewProductRepository(r.im.Connect())
}

func NewRepoManager(im InfraManager) RepoManager {
	return &repoManager{
		im: im,
	}
}
