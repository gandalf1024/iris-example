package models

type Permission struct {
	Id         int `orm:"column(id);pk" description:"主键"`
	Name       string
	Type       string
	Url        string
	Percode    string
	Parentid   int64
	Parentids  string
	Sortstring string
	Available  string
}

func (t *Permission) TableName() string {
	return "sys_permission"
}
