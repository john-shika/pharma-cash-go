package apis

import (
	"nokowebapi/apis/extras"
	"nokowebapi/apis/models"
	"nokowebapi/nokocore"
	"strings"
)

type UserOrJwtAuthInfoImpl interface {
	*extras.JwtAuthInfo | *models.User
}

func IsUser[T UserOrJwtAuthInfoImpl](userOrJwtAuthInfo T) bool {
	return GetUser(userOrJwtAuthInfo) != nil
}

func GetUser[T UserOrJwtAuthInfoImpl](userOrJwtAuthInfo T) *models.User {
	temp := nokocore.Unwrap(nokocore.CastAny(userOrJwtAuthInfo))
	switch t := temp.(type) {
	case *extras.JwtAuthInfo:
		return t.User
	case *models.User:
		return t
	default:
		return nil
	}
}

func IsJwtAuthInfo[T UserOrJwtAuthInfoImpl](userOrJwtAuthInfo T) bool {
	return GetJwtAuthInfo(userOrJwtAuthInfo) != nil
}

func GetJwtAuthInfo[T UserOrJwtAuthInfoImpl](userOrJwtAuthInfo T) *extras.JwtAuthInfo {
	temp := nokocore.Unwrap(nokocore.CastAny(userOrJwtAuthInfo))
	switch t := temp.(type) {
	case *extras.JwtAuthInfo:
		return t
	default:
		return nil
	}
}

type RoleTypedOrStringImpl interface {
	string | nokocore.RoleTyped
}

func GetRole[T RoleTypedOrStringImpl](role T) nokocore.RoleTyped {
	return nokocore.RoleTyped(role)
}

func GetRoleString[T RoleTypedOrStringImpl](role T) string {
	return string(role)
}

func RoleIs[T UserOrJwtAuthInfoImpl, V nokocore.RoleTypedOrStringImpl](userOrJwtAuthInfo T, roleTyped V) bool {
	user := GetUser(userOrJwtAuthInfo)
	role := nokocore.ToRoleString(roleTyped)
	if user != nil {
		roles := nokocore.RolesUnpack(user.Roles)
		for i, value := range roles {
			nokocore.KeepVoid(i)

			if strings.EqualFold(value, role) {
				return true
			}
		}
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
