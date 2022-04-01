package service

import "cart/domain/model"

type ICartDataService interface {
	AddCart(*model.Cart) (int64, error)

	DeleteCart(int64) error

	UpdateCart(cart *model.Cart) (err error)

	FindById(id int64) (*model.Cart, error)

	FindAllCart(id int64) ([]model.Cart, error)

	CleanCart(id int64) error

	DecrNum(int64, int64) error

	IncrNum(int64, int64) error
}
