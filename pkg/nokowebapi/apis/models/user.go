package models

import (
	"database/sql"
	"gorm.io/gorm"
	"nokowebapi/nokocore"
)

type User struct {
	BaseModel
	Username   string         `db:"username" gorm:"unique;index;not null;" mapstructure:"username" json:"username"`
	Password   string         `db:"password" gorm:"not null;" mapstructure:"password" json:"password"`
	FullName   sql.NullString `db:"full_name" gorm:"unique;index;null;" mapstructure:"full_name" json:"fullName"`
	Email      sql.NullString `db:"email" gorm:"unique;index;null;" mapstructure:"email" json:"email"`
	Phone      sql.NullString `db:"phone" gorm:"unique;index;null;" mapstructure:"phone" json:"phone"`
	Admin      bool           `db:"admin" gorm:"not null;" mapstructure:"admin" json:"admin"`
	SuperAdmin bool           `db:"super_admin" gorm:"not null;" mapstructure:"super_admin" json:"superAdmin"`
	Level      int            `db:"level" gorm:"not null;" mapstructure:"level" json:"level"`

	Roles    []Role    `gorm:"many2many:user_roles;" mapstructure:"roles" json:"roles,omitempty"`
	Sessions []Session `db:"-" mapstructure:"sessions" json:"sessions,omitempty"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) SaveRoles(DB *gorm.DB) error {
	for i, role := range u.Roles {
		nokocore.KeepVoid(i)

		// find first or create one
		tx := DB.Where("role_name = ?", role.RoleName).FirstOrCreate(&role)
		if err := tx.Error; err != nil {
			return err
		}

		// assign new role
		u.Roles[i] = role
	}
	return nil
}

func (u *User) BeforeCreate(DB *gorm.DB) (err error) {
	nokocore.KeepVoid(DB)

	password := nokocore.NewPassword(u.Password)
	if u.Password, err = password.Hash(); err != nil {
		return err
	}

	if u.Admin {
		admin := nokocore.ToRoleString(nokocore.RoleAdmin)
		user := nokocore.ToRoleString(nokocore.RoleUser)
		roles := []Role{
			NewRole(admin),
			NewRole(user),
		}

		u.Roles = roles
		if err = u.SaveRoles(DB); err != nil {
			return err
		}
	}

	return nil
}

func (u *User) BeforeUpdate(DB *gorm.DB) (err error) {
	nokocore.KeepVoid(DB)

	if u.Admin {
		var roles []Role
		association := DB.Model(u).Association("Roles")
		if err = association.Find(&roles); err != nil {
			return err
		}

		var temp []string
		for i, role := range roles {
			nokocore.KeepVoid(i)
			temp = append(temp, role.RoleName)
		}

		admin := nokocore.ToRoleString(nokocore.RoleAdmin)
		if found := nokocore.RolesContains(temp, admin); !found {
			roles = append(roles, Role{RoleName: admin})
			if err = association.Append(&roles); err != nil {
				return err
			}
		}
	}

	return nil
}
