package schemas

import (
	"github.com/google/uuid"
	"nokowebapi/apis/models"
	"nokowebapi/apis/schemas"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
	models2 "pharma-cash-go/app/models"
)

type EmployeeBody struct {
	FullName  string   `mapstructure:"full_name" json:"fullName" form:"full_name" validate:"ascii,omitempty"`
	Username  string   `mapstructure:"username" json:"username" form:"username" validate:"ascii"`
	Password  string   `mapstructure:"password" json:"password" form:"password" validate:"password,omitempty"`
	Email     string   `mapstructure:"email" json:"email" form:"email" validate:"email,omitempty"`
	Phone     string   `mapstructure:"phone" json:"phone" form:"phone" validate:"phone,omitempty"`
	Admin     bool     `mapstructure:"admin" json:"admin" form:"admin" validate:"boolean,omitempty"`
	Roles     []string `mapstructure:"roles" json:"roles" form:"roles" validate:"alphabet,omitempty"`
	Role      string   `mapstructure:"role" json:"role" form:"role" validate:"alphabet,omitempty"`
	Level     int      `mapstructure:"level" json:"level" form:"level" validate:"number,min=0,max=99,omitempty"`
	Shift     string   `mapstructure:"shift" json:"shift" form:"shift" validate:"omitempty"`
	ShiftDate string   `mapstructure:"shift_date" json:"shiftDate" form:"shift_date" validate:"dateOnly,omitempty"`
}

func ToEmployeeModel(employee *EmployeeBody, user *models.User, shift *models2.Shift) *models2.Employee {
	if employee != nil {
		return &models2.Employee{
			UserID:    user.ID,
			ShiftID:   shift.ID,
			ShiftDate: sqlx.ParseDateOnly(employee.ShiftDate),
			User:      *user,
			Shift:     *shift,
		}
	}
	return nil
}

func ToUserBodyFromEmployeeBody(employee *EmployeeBody) *schemas.UserBody {
	if employee != nil {
		return &schemas.UserBody{
			FullName: employee.FullName,
			Username: employee.Username,
			Password: employee.Password,
			Email:    employee.Email,
			Phone:    employee.Phone,
			Admin:    employee.Admin,
			Roles:    employee.Roles,
			Role:     employee.Role,
			Level:    employee.Level,
		}
	}

	return nil
}

func ToUserModelFromEmployeeBody(employee *EmployeeBody) *models.User {
	return schemas.ToUserModel(ToUserBodyFromEmployeeBody(employee))
}

type EmployeeResult struct {
	UUID      uuid.UUID     `mapstructure:"uuid" json:"uuid"`
	UserID    uuid.UUID     `mapstructure:"user_id" json:"userId"`
	FullName  string        `mapstructure:"full_name" json:"fullName"`
	Username  string        `mapstructure:"username" json:"username"`
	Email     string        `mapstructure:"email" json:"email"`
	Phone     string        `mapstructure:"phone" json:"phone"`
	Admin     bool          `mapstructure:"admin" json:"admin"`
	Level     int           `mapstructure:"level" json:"level"`
	CreatedAt string        `mapstructure:"created_at" json:"createdAt"`
	UpdatedAt string        `mapstructure:"updated_at" json:"updatedAt"`
	DeletedAt string        `mapstructure:"deleted_at" json:"deletedAt,omitempty"`
	Roles     []string      `mapstructure:"roles" json:"roles"`
	Role      string        `mapstructure:"role" json:"role"` // MAPPING Officer -> Apoteker / Assistant -> TTK
	ShiftDate sqlx.DateOnly `json:"shiftDate,omitempty"`
	Shift     ShiftResult   `json:"shift"`
}

func ToEmployeeResult(employee *models2.Employee) EmployeeResult {
	if employee != nil {
		userResult := schemas.ToUserResult(&employee.User)
		shiftResult := ToShiftResult(&employee.Shift)
		shiftDate := employee.ShiftDate.DateOnly
		createdAt := nokocore.ToTimeUtcStringISO8601(employee.CreatedAt)
		updatedAt := nokocore.ToTimeUtcStringISO8601(employee.UpdatedAt)
		var deletedAt string
		if employee.DeletedAt.Valid {
			deletedAt = nokocore.ToTimeUtcStringISO8601(employee.DeletedAt.Time)
		}
		return EmployeeResult{
			UUID:      employee.UUID,
			UserID:    userResult.UUID,
			FullName:  userResult.FullName,
			Username:  userResult.Username,
			Email:     userResult.Email,
			Phone:     userResult.Phone,
			Admin:     userResult.Admin,
			Level:     userResult.Level,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			DeletedAt: deletedAt,
			Roles:     userResult.Roles,
			Role:      userResult.Role,
			ShiftDate: shiftDate,
			Shift:     shiftResult,
		}
	}

	return EmployeeResult{}
}
