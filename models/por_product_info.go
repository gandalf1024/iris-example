package models

import (
	"time"
)

type ProductInfo struct {
	Id              int `orm:"column(product_id);auto" description:"商品ID"`
	ProductCore     string
	ProductName     string
	BarCode         string
	BrandId         uint
	OneCategoryId   uint16
	TwoCategoryId   uint16
	ThreeCategoryId uint16
	SupplierId      uint
	Price           float64
	AverageCost     float64
	PublishStatus   int8
	AuditStatus     int8
	Weight          float32
	Length          float32
	Height          float32
	Width           float32
	ColorType       string
	ProductionDate  time.Time
	ShelfLife       int
	Descript        string
	Indate          time.Time
	ModifiedTime    time.Time
}

func (t *ProductInfo) TableName() string {
	return "por_product_info"
}
