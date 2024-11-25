package nokocore

import "strings"

type RoleTyped string

const (
	RoleGuest      RoleTyped = "Guest"
	RoleUser       RoleTyped = "User"
	RoleAdmin      RoleTyped = "Admin"
	RoleSuperAdmin RoleTyped = "SuperAdmin"
	RoleOfficer    RoleTyped = "Officer"
	RoleAssistant  RoleTyped = "Assistant"
	RoleSupervisor RoleTyped = "Supervisor"
	RoleDeveloper  RoleTyped = "Developer"
)

type RoleTypedOrStringImpl interface {
	string | RoleTyped
}

func ToRoleTyped[T RoleTypedOrStringImpl](role T) RoleTyped {
	return RoleTyped(role)
}

func ToRoleString[T RoleTypedOrStringImpl](role T) string {
	return string(role)
}

func RolesPack[T RoleTypedOrStringImpl](roles []T) string {
	temp := Unwrap(CastArray[string](roles))
	return strings.Join(temp, ";")
}

func RolesUnpack(roles string) []string {
	var temp []string
	tokens := strings.Split(strings.TrimSpace(roles), ";")
	for i, role := range tokens {
		KeepVoid(i)

		// trying to convert role value to PascalCase
		if role = ToPascalCase(role); role != "" {
			temp = append(temp, role)
		}
	}
	return temp
}

// RoleIs method, unpack roles value ex. User;Admin;SuperAdmin;Officer;
func RoleIs[T RoleTypedOrStringImpl](roles string, expected ...T) bool {
	return RolesContains(RolesUnpack(roles), expected...)
}

func RolesContains[T1, T2 RoleTypedOrStringImpl](roles []T1, expected ...T2) bool {
	if len(expected) > 0 {
		for i, role2 := range expected {
			KeepVoid(i)

			found := false
			for j, role1 := range roles {
				KeepVoid(j)

				if strings.EqualFold(string(role1), string(role2)) {
					found = true
					break
				}
			}

			if !found {
				return false
			}
		}

		return true
	}

	return false
}

func RoleApply[T RoleTypedOrStringImpl](value string, role T) string {
	roles := RolesAdd(RolesUnpack(value), role)
	return strings.Join(roles, ";")
}

func RolesAppend[T1, T2 RoleTypedOrStringImpl](roles []T1, values ...T2) []T1 {
	temp := roles
	for i, role := range values {
		KeepVoid(i)

		temp = RolesAdd(temp, role)
	}

	return temp
}

func RolesAdd[T1, T2 RoleTypedOrStringImpl](roles []T1, role T2) []T1 {
	temp := roles
	found := false
	role2 := string(role)
	for i, value := range roles {
		KeepVoid(i)

		if role1 := string(value); role1 != "" {
			if role2 != "" && strings.EqualFold(role1, role2) {
				found = true
				break
			}
		}
	}

	if !found && role2 != "" {
		temp = append(temp, T1(role2))
	}

	return temp
}

func RolesRemove[T1, T2 RoleTypedOrStringImpl](roles []T1, role T2) []T1 {
	var temp []T1
	role2 := string(role)
	for i, value := range roles {
		KeepVoid(i)

		if role1 := string(value); role1 != "" {
			if role2 != "" && strings.EqualFold(role1, role2) {
				continue
			}

			temp = append(temp, T1(role1))
		}
	}

	return temp
}
