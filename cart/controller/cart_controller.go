package controller

import (
	"cart/common"
	"cart/domain/model"
	"cart/domain/service"
	"cart/proto/cart"
	"context"
)

type CartController struct {
	CartService service.ICartDataService
}

// AddCart 添加购物车
func (c *CartController)AddCart(ctx context.Context, request *cart.CartInfo, response *cart.ResponseAdd) (err error) {
	cart := &model.Cart{}
	common.SwapTo(request,cart)
	response.CartId,err = c.CartService.AddCart(cart)
	return err
}

// CleanCart 清空购物车
func (c *CartController)CleanCart(ctx context.Context, request  *cart.Clean, response *cart.Response) error  {
	if err := c.CartService.CleanCart(request.UserId);err!=nil{
		return err
	}
	response.Message = "购物车清空成功!"
	return nil
}

// Incr 增加商品数量
func (c *CartController)Incr(ctx context.Context, request *cart.Item, response *cart.Response) error  {
	if err := c.CartService.IncrNum(request.Id,request.ChangeNum);err!=nil{
		return err
	}
	response.Message = "购物车添加商品成功!"
	return nil
}

// Decr 减少商品数量
func (c *CartController)Decr(ctx context.Context, request *cart.Item, response *cart.Response) error  {
	if err := c.CartService.DecrNum(request.Id, request.ChangeNum);err!=nil{
		return err
	}
	response.Message = "购物车减少商品成功!"
	return nil
}

// DeleteItemById 删除商品
func (c *CartController)DeleteItemById(ctx context.Context, request *cart.CartId, response *cart.Response) error{
	if err:= c.CartService.DeleteCart(request.CartId);err!=nil{
		return err
	}
	response.Message = "购物车删除成功!"
	return nil
}

// GetAll 获取所有购物车
func (c *CartController)GetAll(ctx context.Context, request *cart.CartFindAll, response *cart.CartAll) error  {
	carts,err := c.CartService.FindAllCart(request.UserId)
	if err!=nil {
		return err
	}
	for _,v := range carts{
		cart := &model.Cart{}
		if err = common.SwapTo(v,cart);err!=nil{
			return err
		}
		response.Cart_Info = append(response.Cart_Info)
	}
	return nil
}