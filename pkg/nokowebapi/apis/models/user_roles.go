package models

type UserRoles struct {
	UserID uint `db:"user_id" gorm:"index;not null;" mapstructure:"user_id" json:"userId"`
	RoleID uint `db:"role_id" gorm:"index;not null;" mapstructure:"role_id" json:"roleId"`
	User   User `db:"-" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"user" json:"user"`
	Role   Role `db:"-" gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"role" json:"role"`
}

func (u *UserRoles) TableName() string {
	return "user_roles"
}
