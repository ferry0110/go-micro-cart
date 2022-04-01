package repository

import (
	"cart/domain/model"
)

// ICartRepository 操作用户数据库接口
type ICartRepository interface {
	InitTable() error

	CreateCart(*model.Cart) (int64, error)

	FindCartByName(string) (*model.Cart, error)

	FindCartById(int64) (*model.Cart, error)

	DeleteCart(int64) error

	UpdateCart(*model.Cart) error

	FindAll(int64) ([]model.Cart, error)

	CleanCart(int64) error

	IncrNum(int64, int64) error

	DecrNum(int64, int64) error
}
