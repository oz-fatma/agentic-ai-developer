package auth

const (
	RoleAdmin  = "admin"
	RoleReader = "reader"
)

func CanAccess(role, required string) bool {
	if role == required {
		return true
	}
	return role == RoleAdmin
}
