package models

import (
	"database/sql"
	"gorm.io/gorm"
	"nokowebapi/nokocore"
	"strings"
)

type User struct {
	BaseModel
	Username string         `db:"username" gorm:"unique;index;not null;" mapstructure:"username" json:"username" yaml:"username"`
	Password string         `db:"password" gorm:"not null;" mapstructure:"password" json:"password" yaml:"password"`
	FullName sql.NullString `db:"full_name" gorm:"unique;index;null;" mapstructure:"full_name" json:"fullName" yaml:"full_name"`
	Email    sql.NullString `db:"email" gorm:"unique;index;null;" mapstructure:"email" json:"email,omitempty" yaml:"email"`
	Phone    sql.NullString `db:"phone" gorm:"unique;index;null;" mapstructure:"phone" json:"phone,omitempty" yaml:"phone"`
	Admin    bool           `db:"admin" gorm:"not null;" mapstructure:"admin" json:"admin" yaml:"admin"`
	Roles    string         `db:"roles" gorm:"not null;" mapstructure:"roles" json:"roles" yaml:"roles"`
	Level    int            `db:"level" gorm:"not null;" mapstructure:"level" json:"level" yaml:"level"`
	Sessions []Session      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"-" json:"-" yaml:"-"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	nokocore.KeepVoid(db)

	password := nokocore.NewPassword(u.Password)
	if u.Password, err = password.Hash(); err != nil {
		return err
	}

	if u.Admin {
		roles := nokocore.RolesUnpack(u.Roles)
		found := false
		for i, role := range roles {
			nokocore.KeepVoid(i)

			if strings.EqualFold(role, string(nokocore.RoleAdmin)) {
				found = true
				break
			}
		}

		if !found {
			roles = append(roles, string(nokocore.RoleAdmin))
			u.Roles = nokocore.RolesPack(roles)
		}
	}

	return nil
}
