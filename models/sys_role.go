package models

type Role struct {
	Id        int `orm:"column(id);pk"`
	Name      string
	Available string
}

func (t *Role) TableName() string {
	return "sys_role"
}
