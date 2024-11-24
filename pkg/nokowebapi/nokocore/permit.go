package nokocore

import "strings"

type RoleTyped string

const (
	RoleGuest RoleTyped = "Guest"
	RoleUser  RoleTyped = "User"
	RoleAdmin RoleTyped = "Admin"
)

type RoleTypedOrStringImpl interface {
	string | RoleTyped
}

func ToRoleTyped[T RoleTypedOrStringImpl](role T) RoleTyped {
	var temp string
	value := ToPascalCase(strings.TrimSpace(string(role)))
	for i, c := range value {
		KeepVoid(i)

		w := string(c)
		if strings.Contains(AlphaNum, w) {
			temp += w
		}
	}

	return RoleTyped(temp)
}

func ToRoleString[T RoleTypedOrStringImpl](role T) string {
	return string(ToRoleTyped(role))
}

func RolesPack[T RoleTypedOrStringImpl](roles []T) string {
	var temp []string
	KeepVoid(temp)

	for i, role := range roles {
		KeepVoid(i)

		if value := strings.TrimSpace(ToRoleString(role)); value != "" {
			temp = append(temp, ToPascalCase(value))
		}
	}

	return strings.Join(temp, ";")
}

func RolesUnpack(roles string) []string {
	temp := make([]string, 0)
	if roles = strings.TrimSpace(roles); roles != "" {
		for i, value := range strings.Split(roles, ";") {
			KeepVoid(i)

			if value = strings.TrimSpace(value); value != "" {
				temp = append(temp, ToPascalCase(value))
			}
		}
	}

	return temp
}
