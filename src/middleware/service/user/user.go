package user

import (
	"middleware/service/authorization"
)

// curl -H "Content-Type: application/json" -X GET http://127.0.0.1:6060/v1/users/alice/roles | jq
// List user roles
func GetUserRoles(user string) []string {
	e :=  authorization.GetEnforcer()
	return e.GetRolesForUser(user)
}

// curl -H "Content-Type: application/json" -X DELETE http://127.0.0.1:6060/v1/users/alice/roles/visitante3 | jq
// Remove role from user
func DeleteRoleFromUser(user, role string) bool {
	e :=  authorization.GetEnforcer()
	return e.DeleteRoleForUser(user, role)
}

// curl -H "Content-Type: application/json" -X GET http://127.0.0.1:6060/v1/users/alice/resource/data1/permission/read | jq
// Check user permission
func CheckUserPermission(subject, object, action string) bool {
	e :=  authorization.GetEnforcer()
	return e.Enforce(subject, object, action)
}