package schemas

import (
	"nokowebapi/apis/models"
	"time"
)

type UserBody struct {
	Username string        `json:"username" yaml:"username" form:"username" validate:"ascii"`
	Password string        `json:"password" yaml:"password" form:"password" validate:"password"`
	Email    string        `json:"email" yaml:"email" form:"email" validate:"email,omitempty"`
	Phone    string        `json:"phone" yaml:"phone" form:"phone" validate:"phone,omitempty"`
	Admin    bool          `json:"admin" yaml:"admin" form:"admin" validate:"boolean,omitempty"`
	Roles    []string      `json:"roles" yaml:"roles" form:"roles" validate:"ascii,omitempty"`
	Level    int           `json:"level" yaml:"level" form:"level" validate:"number,min=0,max=99,omitempty"` // FUTURE: can handle min=N,max=N
	Sessions []SessionBody `json:"-" yaml:"-" validate:"-"`
}

type UserResp struct {
	Username  string        `json:"username" yaml:"username" form:"username"`
	Email     string        `json:"email" yaml:"email" form:"email"`
	Phone     string        `json:"phone" yaml:"phone" form:"phone"`
	Admin     bool          `json:"admin" yaml:"admin" form:"admin"`
	Roles     []string      `json:"roles" yaml:"roles" form:"roles"`
	Level     int           `json:"level" yaml:"level" form:"level"`
	CreatedAt time.Time     `json:"createdAt" yaml:"created_at" form:"created_at"`
	UpdatedAt time.Time     `json:"updatedAt" yaml:"updated_at" form:"updated_at"`
	Sessions  []SessionResp `json:"-" yaml:"-"`
}

func ToUserResp(user *models.User) UserResp {
	return UserResp{
		Username:  user.Username,
		Email:     user.Email.String,
		Phone:     user.Phone.String,
		Admin:     user.Admin,
		Roles:     user.GetRoles(),
		Level:     user.Level,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,

		// TODO: sessions
	}
}
