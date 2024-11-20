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
	Password string         `db:"password" gorm:"not null" mapstructure:"password" json:"password" yaml:"password"`
	Email    sql.NullString `db:"email" gorm:";" mapstructure:"email" json:"email,omitempty" yaml:"email"`
	Phone    sql.NullString `db:"phone" gorm:";" mapstructure:"phone" json:"phone,omitempty" yaml:"phone"`
	Admin    bool           `db:"admin" gorm:"index;not null;" mapstructure:"admin" json:"admin" yaml:"admin"`
	Role     string         `db:"role" gorm:"index;not null;" mapstructure:"role" json:"role" yaml:"role"`
	Level    int            `db:"level" gorm:"index;not null;" mapstructure:"level" json:"level" yaml:"level"`
	Sessions []Session      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" mapstructure:"-" json:"-" yaml:"-"`
}

func (User) TableName() string {
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

	separator := ";" // Guest;User;Admin;TeamKit;Enterprise;Developer
	roles := strings.Split(strings.TrimSpace(u.Role), separator)
	temp := make([]string, 0)

	found := false
	for i, role := range roles {
		nokocore.KeepVoid(i)

		role = nokocore.ToCamelCase(strings.TrimSpace(role))
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

	u.Role = strings.Join(temp, separator)
	return
}
