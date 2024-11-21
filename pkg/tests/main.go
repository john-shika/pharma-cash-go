package main

import (
	"fmt"
	"nokowebapi/apis/validators"
)

type UserBody struct {
	Username string `json:"username" yaml:"username" form:"username" validate:"ascii"`
	Password string `json:"password" yaml:"password" form:"password" validate:"password"`
	Email    string `json:"email" yaml:"email" form:"email" validate:"email,omitempty"`
	Phone    string `json:"phone" yaml:"phone" form:"phone" validate:"phone,omitempty"`
	Admin    bool   `json:"admin" yaml:"admin" form:"admin" validate:"boolean,omitempty"`
	Role     string `json:"role" yaml:"role" form:"role" validate:"ascii,omitempty"`
	Level    string `json:"level" yaml:"level" form:"level" validate:"numeric,min=0,max=80.12,omitempty"` // FUTURE: can handle min=N,max=N
}

func main() {

	user := &UserBody{
		//Username: "admin",
		Password: "Admin@1234",
		Email:    "admin@",
		Phone:    "+64 823-2653-0277x",
		Admin:    true,
		Role:     "admin",
		Level:    "-12",
	}

	fmt.Printf("%T\n", validators.ValidateStruct(user))
	fmt.Printf("%+v\n", validators.ValidateStruct(user))
}
