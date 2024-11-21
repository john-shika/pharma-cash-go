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
	FullName sql.NullString `db:"full_name" gorm:"index;not null;" mapstructure:"full_name" json:"fullName" yaml:"full_name"`
	Email    sql.NullString `db:"email" gorm:"index;" mapstructure:"email" json:"email,omitempty" yaml:"email"`
	Phone    sql.NullString `db:"phone" gorm:"index;" mapstructure:"phone" json:"phone,omitempty" yaml:"phone"`
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

	roleApply := string(nokocore.RoleUser)

	if u.Admin {
		roleApply = string(nokocore.RoleAdmin)
	}

	roles := u.GetRoles()
	temp := make([]string, 0)

	found := false
	for i, role := range roles {
		nokocore.KeepVoid(i)

		// Guest;User;Admin;Enterprise;Developer
		role = nokocore.ToPascalCase(strings.TrimSpace(role))
		if role != "" {
			if strings.EqualFold(role, roleApply) {
				found = true
			}

			temp = append(temp, role)
		}
	}

	if !found {
		temp = append(temp, roleApply)
	}

	u.Roles = strings.Join(temp, ";")
	return
}

func (u *User) GetRoles() []string {
	temp := strings.Split(strings.TrimSpace(u.Roles), ";")
	for i, role := range temp {
		nokocore.KeepVoid(i)

		// Guest;User;Admin;Enterprise;Developer
		role = nokocore.ToPascalCase(strings.TrimSpace(role))
		if role != "" {
			temp[i] = role
		}
	}
	return temp
}
