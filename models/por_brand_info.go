package models

import (
	"time"
)

type BrandInfo struct {
	Id           int `orm:"column(brand_id);auto" description:"品牌ID"`
	BrandName    string
	Telephone    string
	BrandWeb     string
	BrandLogo    string
	BrandDesc    string
	BrandStatus  int8
	BrandOrder   int8
	ModifiedTime time.Time
}

func (t *BrandInfo) TableName() string {
	return "por_brand_info"
}
