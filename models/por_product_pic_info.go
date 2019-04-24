package models

import (
	"time"
)

type ProductPicInfo struct {
	Id           int `orm:"column(product_pic_id);auto" description:"商品图片ID"`
	ProductId    uint
	PicDesc      string
	PicUrl       string
	IsMaster     int8
	PicOrder     int8
	PicStatus    int8
	ModifiedTime time.Time
}

func (t *ProductPicInfo) TableName() string {
	return "por_product_pic_info"
}
