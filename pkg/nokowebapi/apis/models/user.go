package models

import (
	"database/sql"
	"gorm.io/gorm"
	"nokowebapi/nokocore"
	"strings"
)

type User struct {
	BaseModel
	Username string         `db:"username" gorm:"unique;index;not null;" mapstructure:"username" json:"username"`
	Password string         `db:"password" gorm:"not null;" mapstructure:"password" json:"password"`
	FullName sql.NullString `db:"full_name" gorm:"unique;index;null;" mapstructure:"full_name" json:"fullName"`
	Email    sql.NullString `db:"email" gorm:"unique;index;null;" mapstructure:"email" json:"email"`
	Phone    sql.NullString `db:"phone" gorm:"unique;index;null;" mapstructure:"phone" json:"phone"`
	Admin    bool           `db:"admin" gorm:"not null;" mapstructure:"admin" json:"admin"`
	Roles    string         `db:"roles" gorm:"not null;" mapstructure:"roles" json:"roles"`
	Level    int            `db:"level" gorm:"not null;" mapstructure:"level" json:"level"`
	Sessions []Session      `db:"-" mapstructure:"sessions" json:"sessions"`
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
