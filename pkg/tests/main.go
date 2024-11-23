package main

import (
	"fmt"
	"nokowebapi/console"
	"nokowebapi/globals"
	"nokowebapi/nokocore"
)

type UserBody struct {
	Username string `json:"username" form:"username" validate:"ascii"`
	Password string `json:"password" form:"password" validate:"password"`
	Email    string `json:"email" form:"email" validate:"email,omitempty"`
	Phone    string `json:"phone" form:"phone" validate:"phone,omitempty"`
	Admin    bool   `json:"admin" form:"admin" validate:"boolean,omitempty"`
	Role     string `json:"role" form:"role" validate:"ascii,omitempty"`
	Level    string `json:"level" form:"level" validate:"numeric,min=0,max=80.12,omitempty"` // FUTURE: can handle min=N,max=N
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

	fmt.Printf("%T\n", nokocore.ValidateStruct(user))
	fmt.Printf("%+v\n", nokocore.ValidateStruct(user))

	console.Error(fmt.Sprintf("%+v", nokocore.ValidateStruct(user)))

	console.Dir(globals.GetTasksConfig())
}
