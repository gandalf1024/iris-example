package models

import (
	"time"
)

type ProductCategory struct {
	Id             int `orm:"column(category_id);auto" description:"分类ID"`
	CategoryName   string
	CategoryCode   string
	ParentId       uint16
	CategoryLevel  int8
	CategoryStatus int8
	ModifiedTime   time.Time
}

func (t *ProductCategory) TableName() string {
	return "por_product_category"
}
