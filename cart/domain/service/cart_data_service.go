package service

import (
	"cart/domain/model"
	"cart/domain/repository"
)

type CartDataService struct {
	CartRepository repository.ICartRepository
}

// NewCartDataService 创建CartDataService
func NewCartDataService(cartRepository repository.ICartRepository) ICartDataService {
	return &CartDataService{CartRepository: cartRepository}
}

func (u *CartDataService) AddCart(cart *model.Cart) (int64, error) {
	return  u.CartRepository.CreateCart(cart)
}

func (u *CartDataService) DeleteCart(i int64) error {
	return u.CartRepository.DeleteCart(i)
}
func (u *CartDataService) UpdateCart(cart *model.Cart) (err error) {
	return u.CartRepository.UpdateCart(cart)
}

func (u *CartDataService) FindById(id int64) (*model.Cart, error) {
	return u.CartRepository.FindCartById(id)
}

func (u *CartDataService) FindAllCart(userid int64) ([]model.Cart, error) {
	return u.CartRepository.FindAll(userid)
}

func (u *CartDataService)CleanCart(id int64) error{
	return u.CleanCart(id)
}

func (u *CartDataService)DecrNum(cartid int64, num int64) error{
	return u.DecrNum(cartid,num)
}

func (u *CartDataService)IncrNum(cartid int64, num int64) error{
	return u.IncrNum(cartid,num)
}

