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

func CreateEmployee(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	userRepository := repositories.NewUserRepository(DB)
	employeeRepository := repositories2.NewEmployeeRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var user *models.User
		var employee *models2.Employee
		nokocore.KeepVoid(err, user, employee)

		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		if !utils.RoleIsAdmin(jwtAuthInfo) {
			return extras.NewMessageBodyUnauthorized(ctx, "Unauthorized access attempt.", nil)
		}

		employeeBody := new(schemas2.EmployeeBody)
		if err = ctx.Bind(employeeBody); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))

			return extras.NewMessageBodyInternalServerError(ctx, "Invalid request body.", nil)
		}

		if err = ctx.Validate(employeeBody); err != nil {
			return err
		}

		if strings.TrimSpace(employeeBody.Password) == "" {
			return extras.NewMessageBodyUnprocessableEntity(ctx, "Password is required.", nil)
		}

		username := strings.TrimSpace(employeeBody.Username)
		employeeBody.Username = username
		if username != "" {
			if user, err = userRepository.SafeFirst("username = ?", employeeBody.Username); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to get user.", nil)
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
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to get user.", nil)
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
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to get user.", nil)
			}

			if user != nil {
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Phone already exists.", nil)
			}
		}

		err = DB.Transaction(func(tx *gorm.DB) error {
			userRepository := repositories.NewUserRepository(tx)
			employeeRepository := repositories2.NewEmployeeRepository(tx)
			shiftRepository := repositories2.NewShiftRepository(tx)

			// normalize shift name
			shiftName := utils2.ToShiftNameNorm(employeeBody.Shift)

			var shift *models2.Shift
			if shift, err = shiftRepository.SafeFirst("name = ?", shiftName); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return errors.New("shift not found")
			}

			if shift == nil {
				return errors.New("shift not found")
			}

			user = schemas2.ToUserModelFromEmployeeBody(employeeBody)
			if err = userRepository.Create(user); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return errors.New("failed to create user")
			}

			// parsing shift date
			var shiftDate sqlx.NullDateOnly
			if employeeBody.ShiftDate != "" {
				shiftDate = sqlx.ParseDateOnly(employeeBody.ShiftDate)
			}

			employee = &models2.Employee{
				UserID:    user.ID,
				ShiftID:   shift.ID,
				ShiftDate: shiftDate,
			}

			if err = employeeRepository.Create(employee); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return errors.New("failed to create employee")
			}

			return nil
		})

		if err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyUnprocessableEntity(ctx, "Failed to create a new user.", nil)
		}

		preloads := []string{"Shift", "User", "User.Roles"}
		if employee, err = employeeRepository.SafePreFirst(preloads, "user_id = ?", user.ID); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Unable to get employee.", nil)
		}

		shift := schemas2.ToShiftResult(&employee.Shift)
		userResult := schemas.ToUserResult(user)
		employeeResult := schemas2.ToEmployeeResult(employee)
		return extras.NewMessageBodyOk(ctx, "Successfully create a new employee.", &nokocore.MapAny{
			"employee": employeeResult,
			"user":     userResult,
			"shift":    shift,
		})
	}
}

func UpdateEmployee(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	userRepository := repositories.NewUserRepository(DB)
	employeeRepository := repositories2.NewEmployeeRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var userID string
		var employeeID string
		var user *models.User
		var employee *models2.Employee
		nokocore.KeepVoid(err, userID, employeeID, user, employee)

		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		if !utils.RoleIsAdmin(jwtAuthInfo) {
			return extras.NewMessageBodyUnauthorized(ctx, "Unauthorized access attempt.", nil)
		}

		userID = extras.ParseQueryToString(ctx, "user_id")
		employeeID = extras.ParseQueryToString(ctx, "employee_id")

		if userID == "" && employeeID == "" {
			return extras.NewMessageBodyUnprocessableEntity(ctx, "Required parameter 'user_id' or 'employee_id' is missing.", nil)
		}

		if userID != "" {
			if err = sqlx.ValidateUUID(userID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Invalid parameter 'user_id'.", nil)
			}
		}

		if employeeID != "" {
			if err = sqlx.ValidateUUID(employeeID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Invalid parameter 'employee_id'.", nil)
			}
		}

		employeeBody := new(schemas2.EmployeeBody)
		if err = ctx.Bind(employeeBody); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))

			return extras.NewMessageBodyInternalServerError(ctx, "Invalid request body.", nil)
		}

		if err = ctx.Validate(employeeBody); err != nil {
			return err
		}

		if userID != "" {
			preloads := []string{"Roles"}
			if user, err = userRepository.SafePreFirst(preloads, "uuid = ?", userID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to get user.", nil)
			}
		}

		if user == nil && employeeID != "" {
			preloads := []string{"Shift", "User", "User.Roles"}
			if employee, err = employeeRepository.SafePreFirst(preloads, "uuid = ?", employeeID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to get user.", nil)
			}

			if employee == nil {
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Employee not found.", nil)
			}

			// inject with current user
			user = &employee.User
		}

		if user == nil {
			return extras.NewMessageBodyUnprocessableEntity(ctx, "User not found.", nil)
		}

		err = DB.Transaction(func(tx *gorm.DB) error {
			userRepository := repositories.NewUserRepository(tx)
			employeeRepository := repositories2.NewEmployeeRepository(tx)
			shiftRepository := repositories2.NewShiftRepository(tx)

			// normalize shift name
			shiftName := utils2.ToShiftNameNorm(employeeBody.Shift)

			var shift *models2.Shift
			if shift, err = shiftRepository.SafeFirst("name = ?", shiftName); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return errors.New("shift not found")
			}

			if shift == nil {
				return errors.New("shift not found")
			}

			var user2 *models.User
			if user2, err = utils.CopyBaseModel(schemas2.ToUserModelFromEmployeeBody(employeeBody), user); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return errors.New("failed to copy base model")
			}

			// fill password from previous user password
			if strings.TrimSpace(user2.Password) == "" {
				user2.Password = user.Password
			}

			if err = userRepository.SafeUpdate(user2, "id = ?", user2.ID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return errors.New("failed to update user")
			}

			// parsing shift date
			var shiftDate sqlx.NullDateOnly
			if employeeBody.ShiftDate != "" {
				if shiftDate, err = sqlx.SafeParseDateOnly(employeeBody.ShiftDate); err != nil {
					console.Error(fmt.Sprintf("panic: %s", err.Error()))
					return errors.New("failed to parse shift date")
				}
			}

			preloads := []string{"Shift", "User", "User.Roles"}
			if employee, err = employeeRepository.SafePreFirst(preloads, "user_id = ?", user.ID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return errors.New("failed to get employee")
			}

			if employee == nil {
				return errors.New("employee not found")
			}

			employee.ShiftDate = shiftDate
			if err = employeeRepository.SafeUpdate(employee, "id = ?", employee.ID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return errors.New("failed to update employee")
			}

			return nil
		})

		if err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyUnprocessableEntity(ctx, "Failed to update user.", nil)
		}

		// get new employee data
		preloads := []string{"Shift", "User", "User.Roles"}
		if employee, err = employeeRepository.SafePreFirst(preloads, "id = ?", employee.ID); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return errors.New("failed to get employee")
		}

		var shift schemas2.ShiftResult
		if employee != nil {
			shift = schemas2.ToShiftResult(&employee.Shift)
		}

		userResult := schemas.ToUserResult(user)
		employeeResult := schemas2.ToEmployeeResult(employee)
		return extras.NewMessageBodyOk(ctx, "Successfully update employee.", &nokocore.MapAny{
			"employee": employeeResult,
			"user":     userResult,
			"shift":    shift,
		})
	}
}

func GetAllUsers(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	userRepository := repositories.NewUserRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var users []models.User
		nokocore.KeepVoid(err, users)

		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		if !utils.RoleIsAdmin(jwtAuthInfo) {
			return extras.NewMessageBodyUnauthorized(ctx, "Unauthorized access attempt.", nil)
		}

		// get pagination request
		pagination := extras.NewURLQueryPaginationFromEchoContext(ctx)

		preloads := []string{"Roles"}
		if users, err = userRepository.SafePreMany(preloads, pagination.Offset, pagination.Limit, "1=1"); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to get users.", nil)
		}

		size := len(users)
		userResults := make([]schemas.UserResult, size)
		for i, user := range users {
			nokocore.KeepVoid(i)

			userResult := schemas.ToUserResult(&user)
			userResults[i] = userResult
		}

		return extras.NewMessageBodyOk(ctx, "Successfully retrieved.", &nokocore.MapAny{
			"users": userResults,
		})
	}
}

func GetAllEmployees(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	employeeRepository := repositories2.NewEmployeeRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var employees []models2.Employee
		nokocore.KeepVoid(err, employees)

		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		if !utils.RoleIsAdmin(jwtAuthInfo) {
			return extras.NewMessageBodyUnauthorized(ctx, "Unauthorized access attempt.", nil)
		}

		// get pagination request
		pagination := extras.NewURLQueryPaginationFromEchoContext(ctx)

		preloads := []string{"Shift", "User", "User.Roles"}
		if employees, err = employeeRepository.SafePreMany(preloads, pagination.Offset, pagination.Limit, "1=1"); err != nil {
			console.Error(fmt.Sprintf("panic: %s", err.Error()))
			return extras.NewMessageBodyInternalServerError(ctx, "Unable to get employees.", nil)
		}

		size := len(employees)
		employeeResults := make([]schemas2.EmployeeResult, size)
		for i, employee := range employees {
			nokocore.KeepVoid(i)

			employeeResult := schemas2.ToEmployeeResult(&employee)
			employeeResults[i] = employeeResult
		}

		return extras.NewMessageBodyOk(ctx, "Successfully retrieved.", &nokocore.MapAny{
			"employees": employeeResults,
		})
	}
}

func DeleteUser(DB *gorm.DB) echo.HandlerFunc {
	nokocore.KeepVoid(DB)

	userRepository := repositories.NewUserRepository(DB)
	employeeRepository := repositories2.NewEmployeeRepository(DB)

	return func(ctx echo.Context) error {
		var err error
		var userID string
		var employeeID string
		var user *models.User
		var employee *models2.Employee
		nokocore.KeepVoid(err, userID, employeeID, user, employee)

		jwtAuthInfo := extras.GetJwtAuthInfoFromEchoContext(ctx)

		forced := extras.ParseQueryToBool(ctx, "forced")

		if !utils.RoleIsAdmin(jwtAuthInfo) {
			return extras.NewMessageBodyUnauthorized(ctx, "Unauthorized access attempt.", nil)
		}

		userID = extras.ParseQueryToString(ctx, "user_id")
		employeeID = extras.ParseQueryToString(ctx, "employee_id")

		if userID == "" && employeeID == "" {
			return extras.NewMessageBodyUnprocessableEntity(ctx, "Required parameter 'user_id' or 'employee_id' is missing.", nil)
		}

		if userID != "" {
			if err = sqlx.ValidateUUID(userID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Invalid parameter 'user_id'.", nil)
			}
		}

		if employeeID != "" {
			if err = sqlx.ValidateUUID(employeeID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Invalid parameter 'employee_id'.", nil)
			}
		}

		if userID != "" {
			preloads := []string{"Roles"}
			if user, err = userRepository.PreFirst(preloads, "uuid = ?", userID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to get user.", nil)
			}
		}

		if user == nil && employeeID != "" {
			preloads := []string{"Shift", "User", "User.Roles"}
			if employee, err = employeeRepository.PreFirst(preloads, "uuid = ?", employeeID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to get user.", nil)
			}

			if employee == nil {
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Employee not found.", nil)
			}

			// inject with current user
			user = &employee.User
		}

		if user == nil {
			return extras.NewMessageBodyUnprocessableEntity(ctx, "User not found.", nil)
		}

		data := &nokocore.MapAny{
			"user": schemas.ToUserResult(user),
		}

		if !forced && user.DeletedAt.Valid {
			return extras.NewMessageBodyOk(ctx, "User already deleted.", data)
		}

		if employee == nil {
			if employee, err = employeeRepository.SafeFirst("user_id = ?", user.ID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to get employee.", nil)
			}
		}

		if employee != nil {
			if forced {
				if err = employeeRepository.Delete(employee, "id = ?", employee.ID); err != nil {
					console.Error(fmt.Sprintf("panic: %s", err.Error()))
					return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to delete employee.", nil)
				}

			} else {
				if err = employeeRepository.SafeDelete(employee, "id = ?", employee.ID); err != nil {
					console.Error(fmt.Sprintf("panic: %s", err.Error()))
					return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to delete employee.", nil)
				}
			}
		}

		if forced {
			if err = userRepository.Delete(user, "uuid = ?", userID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to delete user.", nil)
			}

		} else {
			if err = userRepository.SafeDelete(user, "uuid = ?", userID); err != nil {
				console.Error(fmt.Sprintf("panic: %s", err.Error()))
				return extras.NewMessageBodyUnprocessableEntity(ctx, "Unable to delete user.", nil)
			}
		}

		return extras.NewMessageBodyOk(ctx, "Successfully deleted.", data)
	}
}

func AdminController(group *echo.Group, DB *gorm.DB) *echo.Group {

	group.POST("/user", CreateEmployee(DB))
	group.PUT("/user", UpdateEmployee(DB))
	group.GET("/users", GetAllUsers(DB))
	group.GET("/employees", GetAllEmployees(DB))
	group.DELETE("/user", DeleteUser(DB))

	return group
}
