package models

import (
	"time"
)

type OrderMaster struct {
	Id               int `orm:"column(order_id);auto" description:"订单ID"`
	OrderSn          uint64
	CustomerId       uint
	ShippingUser     string
	Province         int16
	City             int16
	District         int16
	Address          string
	PaymentMethod    int8
	OrderMoney       float64
	DistrictMoney    float64
	ShippingMoney    float64
	PaymentMoney     float64
	ShippingCompName string
	ShippingSn       string
	CreateTime       time.Time
	ShippingTime     time.Time
	PayTime          time.Time
	ReceiveTime      time.Time
	OrderStatus      int8
	OrderPoint       uint
	InvoiceTime      string
	ModifiedTime     time.Time
}

func (t *OrderMaster) TableName() string {
	return "por_order_master"
}
