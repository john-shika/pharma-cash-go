package controllers

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"nokowebapi/apis/extras"
	"nokowebapi/apis/models"
	"nokowebapi/apis/repositories"
	"nokowebapi/apis/schemas"
	"nokowebapi/apis/utils"
	"nokowebapi/console"
	"nokowebapi/nokocore"
	"nokowebapi/sqlx"
	models2 "pharma-cash-go/app/models"
	repositories2 "pharma-cash-go/app/repositories"
	schemas2 "pharma-cash-go/app/schemas"
	utils2 "pharma-cash-go/app/utils"
	"strings"
)

func CreateUserHandler(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	userRepository := repositories.NewUserRepository(DB)
	employeeRepository := repositories2.NewEmployeeRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var user *models.User
		var employee *models2.Employee
		nokocore.KeepVoid(err, user, employee)

		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		if utils.RoleIsAdmin(jwtAuthInfo) {

			employeeBody := new(schemas2.EmployeeBody)
			if err = ctx.Bind(employeeBody); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))

				return extras.NewMessageBodyInternalServerError(ctx, "Invalid request body.", nil)
			}

			if err = ctx.Validate(employeeBody); err != nil {
				return err
			}

			username := strings.TrimSpace(employeeBody.Username)
			employeeBody.Username = username
			if username != "" {
				if user, err = userRepository.SafeFirst("username = ?", employeeBody.Username); err != nil {
					console.Error(fmt.Sprintf("panic: %s", err.Error()))
					return extras.NewMessageBodyInternalServerError(ctx, "Internal server error.", nil)
				}

				if user != nil {
					return extras.NewMessageBodyUnprocessableEntity(ctx, "Username already exists.", nil)
				}
			}

			email := strings.TrimSpace(employeeBody.Email)
			employeeBody.Email = email
			if email != "" {
				if user, err = userRepository.SafeFirst("email = ?", email); err != nil {
					console.Error(fmt.Sprintf("panic: %s", err.Error()))
					return extras.NewMessageBodyInternalServerError(ctx, "Internal server error.", nil)
				}

				if user != nil {
					return extras.NewMessageBodyUnprocessableEntity(ctx, "Email already exists.", nil)
				}
			}

			phone := strings.TrimSpace(employeeBody.Phone)
			employeeBody.Phone = phone
			if phone != "" {
				if user, err = userRepository.SafeFirst("phone = ?", phone); err != nil {
					console.Error(fmt.Sprintf("panic: %s", err.Error()))
					return extras.NewMessageBodyInternalServerError(ctx, "Internal server error.", nil)
				}

				if user != nil {
					return extras.NewMessageBodyUnprocessableEntity(ctx, "Phone already exists.", nil)
				}
			}

			err = DB.Transaction(func(tx *gorm.DB) error {

				// create repositories with open transactions
				userRepository := repositories.NewUserRepository(tx)
				employeeRepository := repositories2.NewEmployeeRepository(tx)
				shiftRepository := repositories2.NewShiftRepository(tx)

				user = schemas2.ToUserModel(employeeBody)
				if err = userRepository.SafeCreate(user); err != nil {
					console.Error(fmt.Sprintf("panic: %s", err.Error()))
					return errors.New("failed to create a new user")
				}

				// normalize shift name
				shiftName := utils2.ToShiftName(employeeBody.Shift)

				var shift *models2.Shift
				if shift, err = shiftRepository.SafeFirst("name = ?", shiftName); err != nil {
					console.Error(fmt.Sprintf("panic: %s", err.Error()))
					return errors.New("shift not found")
				}

				if shift != nil {
					// immediately work
					shiftDate := sqlx.NewDateOnly(nokocore.GetTimeUtcNow())
					employee := &models2.Employee{
						UserID:    user.ID,
						ShiftID:   shift.ID,
						ShiftDate: shiftDate,
					}

					if err = employeeRepository.SafeCreate(employee); err != nil {
						console.Error(fmt.Sprintf("panic: %s", err.Error()))
						return errors.New("failed to create a new employee")
					}

					return nil
				}

				return errors.New("shift not found")
			})

			if err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Failed to create a new user.", nil)
			}

			preloads := []string{"Shift"}
			if employee, err = employeeRepository.SafePreFirst(preloads, "user_id = ?", user.ID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyInternalServerError(ctx, "Unable to get employee.", nil)
			}

			var shift schemas2.ShiftResult
			if employee != nil {
				shift = schemas2.ToShiftResult(&employee.Shift)
			}

			return extras.NewMessageBodyOk(ctx, "Successfully created a new user.", &nokocore.MapAny{
				"user":  schemas.ToUserResult(user, nil),
				"shift": shift,
			})
		}

		return extras.NewMessageBodyUnauthorized(ctx, "Unauthorized access attempt.", nil)
	}
}

func GetAllUsersHandler(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	userRepository := repositories.NewUserRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var users []models.User
		nokocore.KeepVoid(err, users)

		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		if utils.RoleIsAdmin(jwtAuthInfo) {
			pagination := extras.NewURLQueryPaginationFromEchoContext(ctx)
			if users, err = userRepository.SafeMany(pagination.Offset, pagination.Limit, "1=1"); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyInternalServerError(ctx, "Unable to get users.", nil)
			}

			var userResults []schemas.UserResult
			for i, user := range users {
				nokocore.KeepVoid(i)

				userResults = append(userResults, schemas.ToUserResult(&user, nil))
			}

			return extras.NewMessageBodyOk(ctx, "Successfully retrieved.", &nokocore.MapAny{
				"users": userResults,
			})
		}

		return extras.NewMessageBodyUnauthorized(ctx, "Unauthorized access attempt.", nil)
	}
}

func GetAllEmployeesHandler(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	employeeRepository := repositories2.NewEmployeeRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var employees []models2.Employee
		nokocore.KeepVoid(err, employees)

		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		if utils.RoleIsAdmin(jwtAuthInfo) {
			pagination := extras.NewURLQueryPaginationFromEchoContext(ctx)
			if employees, err = employeeRepository.SafeMany(pagination.Offset, pagination.Limit, "1=1"); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyInternalServerError(ctx, "Unable to get employees.", nil)
			}

			var employeesResult []schemas2.EmployeeResult
			for i, employee := range employees {
				nokocore.KeepVoid(i)

				employeesResult = append(employeesResult, schemas2.ToEmployeeResult(&employee))
			}

			return extras.NewMessageBodyOk(ctx, "Successfully retrieved.", &nokocore.MapAny{
				"employees": employeesResult,
			})
		}

		return extras.NewMessageBodyUnauthorized(ctx, "Unauthorized access attempt.", nil)
	}
}

func DeleteUserHandler(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	userRepository := repositories.NewUserRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var user *models.User
		nokocore.KeepVoid(err, user)

		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		if utils.RoleIsAdmin(jwtAuthInfo) {
			if userId := extras.ParseQueryToString(ctx, "user_id"); userId != "" {
				if user, err = userRepository.First("uuid = ?", userId); err != nil {
					console.Error(fmt.Sprintf("panic: %s", err.Error()))
					return extras.NewMessageBodyInternalServerError(ctx, "Unable to get user.", nil)
				}

				if user != nil {
					data := &nokocore.MapAny{
						"user": schemas.ToUserResult(user, nil),
					}

					if !user.DeletedAt.Valid {
						if err = userRepository.SafeDelete(user, "uuid = ?", userId); err != nil {
							console.Error(fmt.Sprintf("panic: %s", err.Error()))
							return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to delete user.", nil)
						}

						return extras.NewMessageBodyOk(ctx, "Successfully deleted.", data)
					}

					return extras.NewMessageBodyOk(ctx, "User already deleted.", data)
				}

				return extras.NewMessageBodyNotFound(ctx, "User not found.", nil)
			}

			return extras.NewMessageBodyUnprocessableEntity(ctx, "Required parameter 'user_id' is missing.", nil)
		}

		return extras.NewMessageBodyUnauthorized(ctx, "Unauthorized access attempt.", nil)
	}
}

func AdminController(group *echo.Group, DB *gorm.DB) *echo.Group {

	group.POST("/user", CreateUserHandler(DB))
	group.GET("/users", GetAllUsersHandler(DB))
	group.GET("/employees", GetAllEmployeesHandler(DB))
	group.DELETE("/user", DeleteUserHandler(DB))

	return group
}
