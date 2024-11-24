package utils

import (
	"nokowebapi/nokocore"
	"strings"
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
	user := GetUser(userOrJwtAuthInfo)
	if user != nil && len(roles) > 0 {
		rolesUnpack := nokocore.RolesUnpack(user.Roles)
		for i, roleExpected := range roles {
			nokocore.KeepVoid(i)

			for j, roleUnpack := range rolesUnpack {
				nokocore.KeepVoid(j)

				value := nokocore.ToRoleString(roleExpected)
				if !strings.EqualFold(roleUnpack, value) {
					return false
				}
			}
		}

		return true
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
