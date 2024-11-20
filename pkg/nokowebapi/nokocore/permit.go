package nokocore

type RoleTyped string

const (
	RoleGuest RoleTyped = "Guest"
	RoleUser  RoleTyped = "User"
	RoleAdmin RoleTyped = "Admin"
)

// TODO: role string can be snippet with multiple roles
