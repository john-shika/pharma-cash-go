package schemas

import "nokowebapi/apis/schemas"

type EmployeeBody struct {
	FullName string   `mapstructure:"full_name" json:"fullName" form:"full_name" validate:"ascii,omitempty"`
	Username string   `mapstructure:"username" json:"username" form:"username" validate:"ascii"`
	Password string   `mapstructure:"password" json:"password" form:"password" validate:"password"`
	Email    string   `mapstructure:"email" json:"email" form:"email" validate:"email,omitempty"`
	Phone    string   `mapstructure:"phone" json:"phone" form:"phone" validate:"phone,omitempty"`
	Admin    bool     `mapstructure:"admin" json:"admin" form:"admin" validate:"boolean,omitempty"`
	Roles    []string `mapstructure:"roles" json:"roles" form:"roles" validate:"omitempty"`
	Level    int      `mapstructure:"level" json:"level" form:"level" validate:"number,min=0,max=99,omitempty"` // FUTURE: can handle min=N,max=N
	Shift    string   `mapstructure:"shift" json:"shift" form:"shift" validate:"omitempty"`
}

func ToUserBody(employee *EmployeeBody) *schemas.UserBody {
	if employee != nil {
		return &schemas.UserBody{
			FullName: employee.FullName,
			Username: employee.Username,
			Password: employee.Password,
			Email:    employee.Email,
			Phone:    employee.Phone,
			Admin:    employee.Admin,
			Roles:    employee.Roles,
			Level:    employee.Level,
		}
	}
	
	return nil
}
