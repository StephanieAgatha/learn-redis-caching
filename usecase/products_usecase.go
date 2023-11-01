package usecase

import (
	"fmt"
	"redishop/model"
	"redishop/repository"
)

type ProductUsecase interface {
	CreateNewProduct(product model.Products) error
	GetAllProduct() ([]model.Products, error)
	GetProductByID(id int) (model.Products, error)
}

type productUsecase struct {
	prodRepo repository.ProductRepository
}

func (p productUsecase) CreateNewProduct(product model.Products) error {
	//TODO implement me

	if product.Name == "" {
		return fmt.Errorf("Name cannot be empty")
	} else if product.Price == 0 {
		return fmt.Errorf("Price cannot be zero")
	} else if product.Price < 0 {
		return fmt.Errorf("Price cannot under zero")
	}

	if err := p.prodRepo.CreateNewProduct(product); err != nil {
		return err
	}
	return nil
}

func (p productUsecase) GetAllProduct() ([]model.Products, error) {
	//TODO implement me

	products, err := p.prodRepo.GetAllProduct()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p productUsecase) GetProductByID(id int) (model.Products, error) {
	//TODO implement me

	product, err := p.prodRepo.GetProductByID(id)
	if err != nil {
		return model.Products{}, err
	}

	return product, nil
}

func NewProductUsecase(prodRepo repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		prodRepo: prodRepo,
	}
}
