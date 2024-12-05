package utils

import (
	"errors"
	"gorm.io/gorm"
	"nokowebapi/apis/models"
	"nokowebapi/nokocore"
)

type RoleTypedOrStringImpl interface {
	string | nokocore.RoleTyped
}

func GetRole[T RoleTypedOrStringImpl](role T) nokocore.RoleTyped {
	return nokocore.RoleTyped(role)
}

func GetRoleString[T RoleTypedOrStringImpl](role T) string {
	return string(role)
}

func RoleIs[T UserOrJwtAuthInfoImpl, V nokocore.RoleTypedOrStringImpl](userOrJwtAuthInfo T, roles ...V) bool {
	return nokocore.RolesContains(GetUserRoles(userOrJwtAuthInfo), roles...)
}

func RoleApply[T UserOrJwtAuthInfoImpl, V nokocore.RoleTypedOrStringImpl](userOrJwtAuthInfo T, roles ...V) []string {
	return nokocore.RolesAppend(GetUserRoles(userOrJwtAuthInfo), roles...)
}

func RoleIsGuest[T UserOrJwtAuthInfoImpl](userOrJwtAuthInfo T) bool {
	return RoleIs(userOrJwtAuthInfo, nokocore.RoleGuest)
}

func RoleIsUser[T UserOrJwtAuthInfoImpl](userOrJwtAuthInfo T) bool {
	return RoleIs(userOrJwtAuthInfo, nokocore.RoleUser)
}

func RoleIsAdmin[T UserOrJwtAuthInfoImpl](userOrJwtAuthInfo T) bool {
	return RoleIs(userOrJwtAuthInfo, nokocore.RoleAdmin)
}

func ToUserRoles(values []string) []models.Role {
	var temp []models.Role
	for i, role := range values {
		nokocore.KeepVoid(i)
		temp = append(temp, models.Role{RoleName: role})
	}

	return temp
}

func ToUserRolesArrayString(roles []models.Role) []string {
	var temp []string
	for i, role := range roles {
		nokocore.KeepVoid(i)
		temp = append(temp, role.RoleName)
	}

	return temp
}

func GetUserRoles[T UserOrJwtAuthInfoImpl](userOrJwtAuthInfo T) []string {
	if user := GetUser(userOrJwtAuthInfo); user != nil {
		return ToUserRolesArrayString(user.Roles)
	}

	return []string{}
}

func SetUserRoles[T UserOrJwtAuthInfoImpl](DB *gorm.DB, userOrJwtAuthInfo T, roles []string) error {
	var err error
	nokocore.KeepVoid(err)

	if user := GetUser(userOrJwtAuthInfo); user != nil {
		temp := ToUserRoles(roles)
		if err = DB.Model(user).Association("Roles").Replace(temp); err != nil {
			return err
		}

		return nil
	}

	return errors.New("user not found")
}
