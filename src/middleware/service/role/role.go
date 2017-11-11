package role

import (
	"middleware/service/authorization"
	"middleware/service/model"
)

// curl -H "Content-Type: application/json" -X GET http://127.0.0.1:6060/v1/roles | jq
// List Role
func GetAll() []string {
	e := authorization.GetEnforcer()
	roles := e.GetAllRoles()
	return roles
}

// curl -H "Content-Type: application/json" -X POST -d '{"user":"alice", "role":"visitante"}' http://127.0.0.1:6060/v1/roles | jq
// Create Role and User
// Casbin only stores the user-role mapping.
// Do not use the same name for a user and a role inside a RBAC system,
// because Casbin recognizes users and roles as strings, and there's no way for
// Casbin to know whether you are specifying user alice or role alice.
// You can simply solve it by using role_alice.
func Create(data *model.Role) bool {
	e := authorization.GetEnforcer()
	return e.AddRoleForUser(data.User, data.Role)
}

// curl -H "Content-Type: application/json" -X GET http://127.0.0.1:6060/v1/roles/alice | jq
// Get role and its permissions
func Get(user string) [][]string {
	e := authorization.GetEnforcer()
	return e.GetPermissionsForUser(user)
}

// curl -H "Content-Type: application/json" -X PUT -d '{"user":"alice", "permission": "read"}' http://127.0.0.1:6060/v1/roles | jq
// Update role (add permissions)
func Update(data *model.Role) bool {
	e := authorization.GetEnforcer()
	return e.AddPermissionForUser(data.User, data.Permission)
}

// curl -H "Content-Type: application/json" -X DELETE http://127.0.0.1:6060/v1/roles/administrador | jq
// Delete role
func Delete(role string) {
	e := authorization.GetEnforcer()
	e.DeleteRole(role)
}

// Delete permission
func DeletePermissionForUser(user, permission string) bool {
	e := authorization.GetEnforcer()
	return e.DeletePermissionForUser(user, permission)
}