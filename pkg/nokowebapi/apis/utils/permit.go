package utils

import (
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
	if user := GetUser(userOrJwtAuthInfo); user != nil {
		return nokocore.RolesIs(user.Roles, roles...)
	}

	return false
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
