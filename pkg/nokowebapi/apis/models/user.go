package models

import (
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"nokowebapi/nokocore"
	"strings"
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

func (u *User) CreateRoles(DB *gorm.DB) error {
	var err error
	var check Role
	nokocore.KeepVoid(err, check)
	for i, role := range u.Roles {
		nokocore.KeepVoid(i)

		// passing
		if role.ID != 0 {
			continue
		}

		// searching
		tx := DB.Where("role_name = ?", role.RoleName).Find(&check)
		if err = tx.Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// passing
		if check.ID != 0 {
			u.Roles[i] = check
			continue
		}

		// create new
		role.UUID = nokocore.NewUUID()
		if err = DB.Create(&role).Error; err != nil {
			return err
		}
	}

	return nil
}

func (u *User) RolesAppend(DB *gorm.DB, names ...string) error {
	for i, name := range names {
		nokocore.KeepVoid(i)
		found := false
		for j, role := range u.Roles {
			nokocore.KeepVoid(j)
			if strings.EqualFold(role.RoleName, name) {
				found = true
				break
			}
		}

		if !found {
			u.Roles = append(u.Roles, Role{RoleName: name})
		}
	}

	return u.CreateRoles(DB)
}

func (u *User) BeforeSave(DB *gorm.DB) (err error) {
	nokocore.KeepVoid(DB)

	// create user roles if not exists
	if err = u.CreateRoles(DB); err != nil {
		return err
	}

	// store the current user roles
	roles := u.Roles

	// remove all registered user roles, roles variable getting replaced
	if err = DB.Model(u).Association("Roles").Clear(); err != nil {
		return err
	}

	// get roles without registered
	u.Roles = roles

	password := nokocore.NewPassword(u.Password)
	if u.Password, err = password.Hash(); err != nil {
		return err
	}

	user := nokocore.ToRoleString(nokocore.RoleUser)
	if err = u.RolesAppend(DB, user); err != nil {
		return err
	}

	if u.Admin {
		admin := nokocore.ToRoleString(nokocore.RoleAdmin)
		if err = u.RolesAppend(DB, admin); err != nil {
			return err
		}
	}

	if u.SuperAdmin {
		superAdmin := nokocore.ToRoleString(nokocore.RoleSuperAdmin)
		if err = u.RolesAppend(DB, superAdmin); err != nil {
			return err
		}
	}

	return nil
}
