package models

type Role struct {
	BaseModel
	RoleName string `db:"role_name" gorm:"unique;index;not null;" mapstructure:"role_name" json:"roleName"`
}

func NewRole(name string) Role {
	return Role{
		RoleName: name,
	}
}

func (u *Role) TableName() string {
	return "roles"
}

type UserRoles struct {
	UserID int `db:"user_id" gorm:"index;not null;" mapstructure:"user_id" json:"userId"`
	RoleID int `db:"role_id" gorm:"index;not null;" mapstructure:"role_id" json:"roleId"`
}

func (u *UserRoles) TableName() string {
	return "user_roles"
}
