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

func (Role) TableName() string {
	return "roles"
}
