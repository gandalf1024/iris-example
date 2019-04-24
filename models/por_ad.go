package models

import (
	"time"
)

type Ad struct {
	Id           int `orm:"column(id);auto"`
	AdPositionId uint16
	MediaType    uint8
	Name         string
	Link         string
	ImageUrl     string
	Content      string
	EndTime      time.Time
	Enabled      uint8
}

func (t *Ad) TableName() string {
	return "por_ad"
}
