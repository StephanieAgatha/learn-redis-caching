package repository

import (
	"database/sql"
	"fmt"
	"github.com/gookit/slog"
	"redishop/model"
)

type ProductRepository interface {
	CreateNewProduct(product model.Products) error
	GetAllProduct() ([]model.Products, error)
	GetProductByID(id int) (model.Products, error)
}

type productRepository struct {
	db *sql.DB
}

func (p productRepository) CreateNewProduct(product model.Products) error {
	//TODO implement me
	q := "insert into products (name,price) values ($1, $2)"
	_, err := p.db.Exec(q, product.Name, product.Price)
	if err != nil {
		return fmt.Errorf("Failed to exec query, err : %v", err.Error())
	}
	return nil
}

func (p productRepository) GetAllProduct() ([]model.Products, error) {
	//TODO implement me
	var products []model.Products
	q := "select * from products"

	rows, err := p.db.Query(q)
	if err != nil {
		return nil, fmt.Errorf("Failed to exec query, err : %v", err.Error())
	}

	//row next
	for rows.Next() {
		var product model.Products
		//row scan
		if err = rows.Scan(&product.Id, &product.Name, &product.Price); err != nil {
			return nil, fmt.Errorf("Failed to scan row, err : %v", err.Error())
		}

		products = append(products, product)
	}

	return products, nil
}

func (p productRepository) GetProductByID(id int) (model.Products, error) {
	//TODO implement me
	var product model.Products

	q := "select * from products where id = $1"

	if err := p.db.QueryRow(q, id).Scan(&product.Id, &product.Name, &product.Price); err != nil {
		slog.Errorf("Failed to exec query ")
		return model.Products{}, nil
	}

	return product, nil
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}
