package controller

import "time"

type AdController struct {
}

func (c *AdController) Get() string {
	return "get: " + time.Now().Format("Mon, 01 Jan 2006 15:04:05 GMT")
}

func (c *AdController) GetOther() string {
	return "other: " + time.Now().Format("Mon, 01 Jan 2006 15:04:05 GMT")
}
