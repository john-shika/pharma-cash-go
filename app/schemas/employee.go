package schemas

import (
	"nokowebapi/apis/models"
	"nokowebapi/apis/schemas"
	"nokowebapi/sqlx"
	models2 "pharma-cash-go/app/models"
)

type EmployeeBody struct {
	FullName  string   `mapstructure:"full_name" json:"fullName" form:"full_name" validate:"ascii,omitempty"`
	Username  string   `mapstructure:"username" json:"username" form:"username" validate:"ascii"`
	Password  string   `mapstructure:"password" json:"password" form:"password" validate:"password"`
	Email     string   `mapstructure:"email" json:"email" form:"email" validate:"email,omitempty"`
	Phone     string   `mapstructure:"phone" json:"phone" form:"phone" validate:"phone,omitempty"`
	Admin     bool     `mapstructure:"admin" json:"admin" form:"admin" validate:"boolean,omitempty"`
	Roles     []string `mapstructure:"roles" json:"roles" form:"roles" validate:"omitempty"`
	Level     int      `mapstructure:"level" json:"level" form:"level" validate:"number,min=0,max=99,omitempty"` // FUTURE: can handle min=N,max=N
	Shift     string   `mapstructure:"shift" json:"shift" form:"shift" validate:"omitempty"`
	ShiftDate string   `mapstructure:"shift_date" json:"shiftDate" form:"shift_date" validate:"dateOnly,omitempty"`
}

func ToEmployeeModel(employee *EmployeeBody, user *models.User, shift *models2.Shift) *models2.Employee {
	return &models2.Employee{
		UserID:    user.ID,
		ShiftID:   shift.ID,
		ShiftDate: sqlx.ParseDateOnly(employee.ShiftDate),
		User:      *user,
		Shift:     *shift,
	}
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

func ToUserModel(employee *EmployeeBody) *models.User {
	return schemas.ToUserModel(ToUserBody(employee))
}

type EmployeeResult struct {
	schemas.UserResult
	Shift     string            `json:"shift"`
	ShiftDate sqlx.DateOnlyImpl `json:"shiftDate,omitempty"`
}

func ToEmployeeResult(employee *models2.Employee) EmployeeResult {
	if employee != nil {
		return EmployeeResult{
			UserResult: schemas.ToUserResult(&employee.User, nil),
			Shift:      employee.Shift.Name,
			ShiftDate:  employee.ShiftDate.DateOnly,
		}
	}

	return EmployeeResult{}
}
