package model

type Products struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
