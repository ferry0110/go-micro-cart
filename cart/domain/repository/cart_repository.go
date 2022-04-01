package repository

import (
	"cart/domain/model"
	"errors"
	"github.com/jinzhu/gorm"
)

// CartRepository Cart数据库实体类
type CartRepository struct {
	mysqlDb *gorm.DB
}

// NewCartRepository 创建CartRepository
func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{mysqlDb: db}
}

// InitTable 初始化表
func (u *CartRepository)InitTable() error{
	return u.mysqlDb.CreateTable(&model.Cart{}).Error
}

// FindCartById 查询购物车byId
func (u *CartRepository)FindCartById(id int64) (cart *model.Cart, err error) {
	cart = &model.Cart{}
	return cart,u.mysqlDb.First(cart,id).Error
}

// FindCartByName 通过购物车名查询购物车
func (u *CartRepository)FindCartByName(name string) (cart *model.Cart, err error)  {
	cart = &model.Cart{}
	return cart,u.mysqlDb.Where("cart_name = ?",name).Find(cart).Error
}

// CreateCart 创建新购物车
func (u *CartRepository)CreateCart(cart *model.Cart) (id int64, err error) {
	db:= u.mysqlDb.FirstOrCreate(cart,model.Cart{ProductId: cart.ProductId,SizeId: cart.SizeId,Userid: cart.Userid})
	if db.Error !=nil  {
		return 0,db.Error
	}
	if db.RowsAffected == 0 {
		return 0,errors.New("购物车创造失败")
	}
	return cart.Id,nil
}

// DeleteCart 删除购物车
func (u *CartRepository)DeleteCart(id int64) error {
	return u.mysqlDb.Where("id = ?",id).Delete(&model.Cart{}).Error
}

// UpdateCart 更新购物车信息
func (u *CartRepository)UpdateCart(cart *model.Cart) error{
	return u.mysqlDb.Model(cart).Update(&cart).Error
}

// FindAll 查找所有购物车
func  (u *CartRepository)FindAll(userId int64) (carts []model.Cart, err error)  {
	return carts,u.mysqlDb.Where("user_id = ?",userId).Find(&carts).Error
}

// CleanCart  根据userId清空购物车
func (u *CartRepository)CleanCart(userId int64) error  {
	return u.mysqlDb.Where("user_id = ?",userId).Delete(&model.Cart{}).Error
}

// IncrNum 添加商品数量
func (u *CartRepository)IncrNum(cartId int64, num int64) error {
	cart:= &model.Cart{Id: cartId}
	return u.mysqlDb.Model(cart).UpdateColumn("num",gorm.Expr("num + ?",num)).Error
}

// DecrNum 减少商品数量
func (u *CartRepository)DecrNum(cartId int64, num int64) error {
	cart:= &model.Cart{Id: cartId}
	db := u.mysqlDb.Model(cart).Where("num >= ?",num).UpdateColumn("num",gorm.Expr("num - ?",num))
	if db.Error !=nil  {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("商品数量减少失败")
	}
	return nil
}


















