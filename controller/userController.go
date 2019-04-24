package controller

import "time"

type UserController struct {
}

func (c *UserController) Get() string {
	return "get: " + time.Now().Format("Mon, 01 Jan 2006 15:04:05 GMT")
}

func (c *UserController) GetOther() string {
	return "other: " + time.Now().Format("Mon, 01 Jan 2006 15:04:05 GMT")
}
