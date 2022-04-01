package model

type Cart struct {
	Id        int64 `gorm:"primary_key;not_null;auto_increment" json:"id"`
	ProductId int64 `gorm:"not_null" json:"product_id"`
	Num       int64 `gorm:"not_null" json:"num"`
	SizeId    int64 `gorm:"not_null" json:"size_id"`
	Userid    int64 `gorm:"not_null" json:"user_id"`
}
