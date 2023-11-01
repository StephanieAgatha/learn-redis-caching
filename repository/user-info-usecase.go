package repository

import (
	"database/sql"
	"fmt"
	"redishop/model"
)

type UserRepo interface {
	CreateUser(user model.User) error
}

type userRepo struct {
	db *sql.DB
}

func (u userRepo) CreateUser(user model.User) error {
	//TODO implement me

	query := "insert into userr (name,age,address) values ($1, $2, $3)"
	_, err := u.db.Exec(query, user.Name, user.Age, user.Address)
	if err != nil {
		return fmt.Errorf("Failed to exec query %v", err.Error())
	}

	return nil
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}
