package schemas

import (
	"github.com/google/uuid"
	"nokowebapi/apis/models"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
)

type UserBody struct {
	FullName   string   `mapstructure:"full_name" json:"fullName" form:"full_name" validate:"ascii,omitempty"`
	Username   string   `mapstructure:"username" json:"username" form:"username" validate:"ascii"`
	Password   string   `mapstructure:"password" json:"password" form:"password" validate:"password"`
	Email      string   `mapstructure:"email" json:"email" form:"email" validate:"email,omitempty"`
	Phone      string   `mapstructure:"phone" json:"phone" form:"phone" validate:"phone,omitempty"`
	Admin      bool     `mapstructure:"admin" json:"admin" form:"admin" validate:"boolean,omitempty"`
	SuperAdmin bool     `mapstructure:"super_admin" json:"superAdmin" form:"super_admin" validate:"boolean,omitempty"`
	Level      int      `mapstructure:"level" json:"level" form:"level" validate:"number,min=0,max=99,omitempty"`
	Roles      []string `mapstructure:"roles" json:"roles" form:"roles" validate:"alphanum,min=1,omitempty"`
	Role       string   `mapstructure:"role" json:"role" form:"role" validate:"alphanum,omitempty"`
}

// TODO: Explicit for pharma cash app
// normRoleDW method, modified explicit for pharma cash app
func normRoleDW(role string) models.Role {
	roleModel := models.Role{}
	switch nokocore.ToPascalCase(role) {
	case "Apoteker":
		roleModel.RoleName = nokocore.ToRoleString(nokocore.RoleOfficer)
	case "Ttk":
		roleModel.RoleName = nokocore.ToRoleString(nokocore.RoleAssistant)
	case "Supervisor":
		roleModel.RoleName = nokocore.ToRoleString(nokocore.RoleSupervisor)
	default:
		roleModel.RoleName = nokocore.ToRoleString(nokocore.RoleUser)
	}
	return roleModel
}

// TODO: Explicit for pharma cash app
// normRoleSingleDW method, modified explicit for pharma cash app
func normRoleSingleDW(roles []string) string {
	var role string
	found := false
	for i, value := range roles {
		nokocore.KeepVoid(i)
		switch nokocore.ToPascalCase(value) {
		case nokocore.ToRoleString(nokocore.RoleOfficer):
			role = "apoteker"
			found = true
			break

		case nokocore.ToRoleString(nokocore.RoleAssistant):
			role = "ttk"
			found = true
			break

		case nokocore.ToRoleString(nokocore.RoleSupervisor):
			role = "supervisor"
			found = true
			break
		}
	}

	if !found {
		role = roles[0]
	}
	return role
}

func ToUserModel(user *UserBody) *models.User {
	if user != nil {
		var roles []models.Role
		for i, role := range user.Roles {
			nokocore.KeepVoid(i)
			if role = nokocore.ToPascalCase(role); role != "" {
				roleModel := normRoleDW(role)
				roles = append(roles, roleModel)
			}
		}
		if role := nokocore.ToPascalCase(user.Role); role != "" {
			roleModel := normRoleDW(role)
			roles = append(roles, roleModel)
		}
		return &models.User{
			FullName:   sqlx.NewString(user.FullName),
			Username:   user.Username,
			Password:   user.Password,
			Email:      sqlx.NewString(user.Email),
			Phone:      sqlx.NewString(user.Phone),
			Admin:      user.Admin,
			SuperAdmin: user.SuperAdmin,
			Roles:      roles,
			Level:      user.Level,
		}
	}

	return nil
}

type UserResult struct {
	UUID      uuid.UUID       `mapstructure:"uuid" json:"uuid"`
	FullName  string          `mapstructure:"full_name" json:"fullName"`
	Username  string          `mapstructure:"username" json:"username"`
	Email     string          `mapstructure:"email" json:"email"`
	Phone     string          `mapstructure:"phone" json:"phone"`
	Admin     bool            `mapstructure:"admin" json:"admin"`
	Level     int             `mapstructure:"level" json:"level"`
	CreatedAt string          `mapstructure:"created_at" json:"createdAt"`
	UpdatedAt string          `mapstructure:"updated_at" json:"updatedAt"`
	DeletedAt string          `mapstructure:"deleted_at" json:"deletedAt,omitempty"`
	Sessions  []SessionResult `mapstructure:"sessions" json:"sessions,omitempty"`
	Roles     []string        `mapstructure:"roles" json:"roles"`
	Role      string          `mapstructure:"role" json:"role"`
}

func ToUserResult(user *models.User) UserResult {
	if user != nil {
		var roles []string
		for i, role := range user.Roles {
			nokocore.KeepVoid(i)
			roles = append(roles, role.RoleName)
		}
		createdAt := nokocore.ToTimeUtcStringISO8601(user.CreatedAt)
		updatedAt := nokocore.ToTimeUtcStringISO8601(user.UpdatedAt)
		var deletedAt string
		if user.DeletedAt.Valid {
			deletedAt = nokocore.ToTimeUtcStringISO8601(user.DeletedAt.Time)
		}
		var role string
		if len(roles) > 0 {
			role = normRoleSingleDW(roles)
		}
		return UserResult{
			UUID:      user.UUID,
			FullName:  user.FullName.String,
			Username:  user.Username,
			Email:     user.Email.String,
			Phone:     user.Phone.String,
			Admin:     user.Admin,
			Level:     user.Level,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			DeletedAt: deletedAt,
			Roles:     roles,
			Role:      role,
		}
	}

	return UserResult{}
}
