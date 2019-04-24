package models

type RolePermission struct {
	Id              int `orm:"column(id);pk"`
	SysRoleId       string
	SysPermissionId string
}

func (t *RolePermission) TableName() string {
	return "sys_role_permission"
}
