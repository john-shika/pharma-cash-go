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
	var temp string
	value := strings.TrimSpace(string(role))
	for i, character := range value {
		KeepVoid(i)

		w := string(character)
		if !strings.Contains(AlphaNum, w) {
			break
		}

		temp += w
	}

	temp = ToPascalCase(temp)
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

		if value := ToRoleString(role); value != "" {
			temp = append(temp, value)
		}
	}

	return strings.Join(temp, ";")
}

func RolesUnpack(values string) []string {
	var temp []string
	if values = strings.TrimSpace(values); values != "" {
		for i, role := range strings.Split(values, ";") {
			KeepVoid(i)

			if role = ToRoleString(role); role != "" {
				temp = append(temp, role)
			}
		}
	}

	return temp
}

func RolesIs[T RoleTypedOrStringImpl](value string, expected ...T) bool {
	roles := RolesUnpack(value)
	return RolesContains(roles, expected...)
}

func RolesContains[T1, T2 RoleTypedOrStringImpl](roles []T1, expected ...T2) bool {
	if len(expected) > 0 {
		for i, role2 := range expected {
			KeepVoid(i)

			found := false
			for j, role1 := range roles {
				KeepVoid(j)

				value := ToRoleString(role2)

				// maybe role1 has been normalized
				if strings.EqualFold(string(role1), value) {
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

func RolesApply[T RoleTypedOrStringImpl](value string, role T) string {
	roles := RolesUnpack(value)
	roles = RolesAppend(roles, role)
	return strings.Join(roles, ";")
}

func RolesAppend[T1, T2 RoleTypedOrStringImpl](roles []T1, role T2) []T1 {
	var temp []T1
	found := false
	role2 := ToRoleString(role)
	for i, value := range roles {
		KeepVoid(i)

		// maybe value has been normalized
		if role1 := string(value); role1 != "" {
			if strings.EqualFold(role1, role2) {
				found = true
			}

			temp = append(temp, T1(role1))
		}
	}

	if !found && role2 != "" {
		temp = append(temp, T1(role2))
	}

	return temp
}

func RolesRemove[T1, T2 RoleTypedOrStringImpl](roles []T1, role T2) []T1 {
	var temp []T1
	role2 := ToRoleString(role)
	for i, value := range roles {
		KeepVoid(i)

		// maybe value has been normalized
		if role1 := string(value); role1 != "" {
			if strings.EqualFold(role1, role2) {
				continue
			}

			temp = append(temp, T1(role1))
		}
	}

	return temp
}
