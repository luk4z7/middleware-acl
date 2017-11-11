package user

import (
	"net/http"
	"middleware/service/user"
	"github.com/gorilla/mux"
	"middleware/core/response"
	liberr "middleware/lib/error"
	"middleware/service/model"
)

func GetUserRoles(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var responseStatus = http.StatusOK
	var errors = []liberr.Errors{}

	vars := mux.Vars(r)
	param := vars["user"]

	data := user.GetUserRoles(param)
	roles := model.RolesResponse{}
	roles.Roles = data

	envelope := response.MountEnvelope(roles, responseStatus, r, errors)
	response.Header(w, responseStatus, envelope, response.Headers{})
}

func DeleteRoleFromUser(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var responseStatus = http.StatusOK
	var errors = []liberr.Errors{}

	vars := mux.Vars(r)
	userParam := vars["user"]
	roleParam := vars["role"]
	user.DeleteRoleFromUser(userParam, roleParam)

	envelope := response.MountEnvelope([]string{}, responseStatus, r, errors)
	response.Header(w, responseStatus, envelope, response.Headers{})
}

func CheckUserPermission(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var responseStatus = http.StatusOK
	var errors = []liberr.Errors{}
	var message = []string{"Sem permissão"}

	vars := mux.Vars(r)
	subject := vars["subject"]
	object := vars["object"]
	action := vars["action"]
	ok := user.CheckUserPermission(subject, object, action)
	if ok {
		message = []string{"Permissão concedida"}
	}
	envelope := response.MountEnvelope(message, responseStatus, r, errors)
	response.Header(w, responseStatus, envelope, response.Headers{})
}