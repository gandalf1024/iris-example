package models

type UserRole struct {
	Id        int `orm:"column(id);pk"`
	SysUserId string
	SysRoleId string
}

func (t *UserRole) TableName() string {
	return "sys_user_role"
}
