package models

import (
	"time"
)

type OrderCart struct {
	Id            int `orm:"column(cart_id);auto" description:"购物车ID"`
	CustomerId    uint
	ProductId     uint
	ProductAmount int
	Price         float64
	AddTime       time.Time
	ModifiedTime  time.Time
}

func (t *OrderCart) TableName() string {
	return "por_order_cart"
}
