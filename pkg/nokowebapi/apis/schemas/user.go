package schemas

import (
	"nokowebapi/apis/models"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
)

type UserBody struct {
	FullName string   `mapstructure:"full_name" json:"fullName" form:"full_name" validate:"ascii,omitempty"`
	Username string   `mapstructure:"username" json:"username" form:"username" validate:"ascii"`
	Password string   `mapstructure:"password" json:"password" form:"password" validate:"password"`
	Email    string   `mapstructure:"email" json:"email" form:"email" validate:"email,omitempty"`
	Phone    string   `mapstructure:"phone" json:"phone" form:"phone" validate:"phone,omitempty"`
	Admin    bool     `mapstructure:"admin" json:"admin" form:"admin" validate:"boolean,omitempty"`
	Roles    []string `mapstructure:"roles" json:"roles" form:"roles" validate:"omitempty"`
	Level    int      `mapstructure:"level" json:"level" form:"level" validate:"number,min=0,max=99,omitempty"` // FUTURE: can handle min=N,max=N
}

func ToUserModel(user *UserBody) *models.User {
	if user != nil {
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

	return nil
}

type UserResult struct {
	FullName  string          `mapstructure:"full_name" json:"fullName"`
	Username  string          `mapstructure:"username" json:"username"`
	Email     string          `mapstructure:"email" json:"email"`
	Phone     string          `mapstructure:"phone" json:"phone"`
	Admin     bool            `mapstructure:"admin" json:"admin"`
	Roles     []string        `mapstructure:"roles" json:"roles"`
	Level     int             `mapstructure:"level" json:"level"`
	CreatedAt string          `mapstructure:"created_at" json:"createdAt"`
	UpdatedAt string          `mapstructure:"updated_at" json:"updatedAt"`
	Sessions  []SessionResult `mapstructure:"sessions" json:"sessions,omitempty"`
}

func ToUserResult(user *models.User, sessions []SessionResult) UserResult {
	if user != nil {
		return UserResult{
			FullName:  user.FullName.String,
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

	return UserResult{}
}
