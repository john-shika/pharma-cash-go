package utils

import (
	"nokowebapi/apis/extras"
	"nokowebapi/apis/models"
	"nokowebapi/nokocore"
)

type UserOrJwtAuthInfoImpl interface {
	*extras.JwtAuthInfo | *models.User
}

func IsUser[T UserOrJwtAuthInfoImpl](userOrJwtAuthInfo T) bool {
	return GetUser(userOrJwtAuthInfo) != nil
}

func GetUser[T UserOrJwtAuthInfoImpl](userOrJwtAuthInfo T) *models.User {
	value := nokocore.Unwrap(nokocore.CastAny(userOrJwtAuthInfo))
	switch val := value.(type) {
	case *extras.JwtAuthInfo:
		return val.User
	case *models.User:
		return val
	default:
		return nil
	}
}

func IsJwtAuthInfo[T UserOrJwtAuthInfoImpl](userOrJwtAuthInfo T) bool {
	return GetJwtAuthInfo(userOrJwtAuthInfo) != nil
}

func GetJwtAuthInfo[T UserOrJwtAuthInfoImpl](userOrJwtAuthInfo T) *extras.JwtAuthInfo {
	value := nokocore.Unwrap(nokocore.CastAny(userOrJwtAuthInfo))
	switch val := value.(type) {
	case *extras.JwtAuthInfo:
		return val
	default:
		return nil
	}
}
