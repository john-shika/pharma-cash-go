package models

import (
	"database/sql"
	"gorm.io/gorm"
	"nokowebapi/nokocore"
	"strings"
)

type User struct {
	Model
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
	
	if u.Password, err = nokocore.HashPassword(u.Password); err != nil {
		return err
	}

	if u.Admin {
		temp := make([]string, 0)
		roles := strings.Split(strings.TrimSpace(u.Role), ",")
		found := false
		for i, role := range roles {
			nokocore.KeepVoid(i)

			role = strings.TrimSpace(role)
			if strings.EqualFold(role, "Admin") {
				found = true
			}

			temp = append(temp, role)
		}

		if !found {
			temp = append(temp, "Admin")
		}

		u.Role = strings.Join(temp, ",")
	}

	return
}
