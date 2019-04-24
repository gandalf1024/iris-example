package models

import (
	"time"
)

type CustomerAddr struct {
	Id           int `orm:"column(customer_addr_id);auto" description:"自增主键ID"`
	CustomerId   uint
	Zip          int16
	Province     int16
	City         int16
	District     int16
	Address      string
	IsDefault    int8
	ModifiedTime time.Time
}

func (t *CustomerAddr) TableName() string {
	return "por_customer_addr"
}
