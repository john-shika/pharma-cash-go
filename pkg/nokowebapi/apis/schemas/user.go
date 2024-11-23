package schemas

import (
	"nokowebapi/apis/models"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
)

type UserBody struct {
	FullName string   `json:"fullName" yaml:"full_name" form:"full_name" validate:"ascii,omitempty"`
	Username string   `json:"username" yaml:"username" form:"username" validate:"ascii"`
	Password string   `json:"password" yaml:"password" form:"password" validate:"password"`
	Email    string   `json:"email" yaml:"email" form:"email" validate:"email,omitempty"`
	Phone    string   `json:"phone" yaml:"phone" form:"phone" validate:"phone,omitempty"`
	Admin    bool     `json:"admin" yaml:"admin" form:"admin" validate:"boolean,omitempty"`
	Roles    []string `json:"roles" yaml:"roles" form:"roles" validate:"omitempty"`
	Level    int      `json:"level" yaml:"level" form:"level" validate:"number,min=0,max=99,omitempty"` // FUTURE: can handle min=N,max=N
}

func ToUserModel(user *UserBody) *models.User {
	return &models.User{
		FullName: sqlx.NewString(user.FullName),
		Username: user.Username,
		Password: user.Password,
		Email:    sqlx.NewString(user.Email),
		Phone:    sqlx.NewString(user.Phone),
		Admin:    user.Admin,
		Roles:    nokocore.RolesPack(user.Roles),
		Level:    user.Level,
	}
}

type UserResult struct {
	Username  string          `json:"username" yaml:"username" form:"username"`
	Email     string          `json:"email" yaml:"email" form:"email"`
	Phone     string          `json:"phone" yaml:"phone" form:"phone"`
	Admin     bool            `json:"admin" yaml:"admin" form:"admin"`
	Roles     []string        `json:"roles" yaml:"roles" form:"roles"`
	Level     int             `json:"level" yaml:"level" form:"level"`
	CreatedAt string          `json:"createdAt" yaml:"created_at" form:"created_at"`
	UpdatedAt string          `json:"updatedAt" yaml:"updated_at" form:"updated_at"`
	Sessions  []SessionResult `json:"sessions" yaml:"sessions"`
}

func ToUserResult(user *models.User, sessions []SessionResult) UserResult {
	return UserResult{
		Username:  user.Username,
		Email:     user.Email.String,
		Phone:     user.Phone.String,
		Admin:     user.Admin,
		Roles:     nokocore.RolesUnpack(user.Roles),
		Level:     user.Level,
		CreatedAt: nokocore.ToTimeUtcStringISO8601(user.CreatedAt),
		UpdatedAt: nokocore.ToTimeUtcStringISO8601(user.UpdatedAt),
		Sessions:  sessions,
	}
}
