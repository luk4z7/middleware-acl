package role

import (
	"net/http"
	"middleware/service/role"
	"middleware/core/response"
	liberr "middleware/lib/error"
	"github.com/gorilla/mux"
	"middleware/service/model"
	"encoding/json"
)

func GetAll(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var responseStatus = http.StatusOK
	var errors = []liberr.Errors{}

	data := role.GetAll()
	roles := model.RolesResponse{}
	roles.Roles = data
	envelope := response.MountEnvelope(roles, responseStatus, r, errors)
	response.Header(w, responseStatus, envelope, response.Headers{})
}

func Create(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var responseStatus = http.StatusOK
	var errors = []liberr.Errors{}
	var message = []string{"Role já existe para esse usuário"}

	data := new(model.Role)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		errors = []liberr.Errors{
			0: {
				Message: "Um erro inesperado ocorreu, verifique o formato do seu json",
				Type:    "internal_error",
			},
		}
		responseStatus = 500
		message = []string{}
	}
	if (data.User == "" || data.Role == "") && err == nil {
		errors = []liberr.Errors{
			0: {
				Message: "Necessário usuário e role",
				Type:    "status_badrequest",
			},
		}
		responseStatus = 400
		message = []string{}
	}
 	if err == nil && data.User != "" && data.Role != "" {
		ok := role.Create(data)
		if ok {
			message = []string{"Role criado com sucesso para o usuário : " + data.User}
		}
	}
	envelope := response.MountEnvelope(message, responseStatus, r, errors)
	response.Header(w, responseStatus, envelope, response.Headers{})
}

func Get(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var responseStatus = http.StatusOK
	var errors = []liberr.Errors{}

	vars := mux.Vars(r)
	param := vars["name"]

	data := role.Get(param)
	envelope := response.MountEnvelope(data, responseStatus, r, errors)
	response.Header(w, responseStatus, envelope, response.Headers{})
}

func Update(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	var responseStatus = http.StatusOK
	var errors = []liberr.Errors{}
	var message = []string{"Usuário já possui essa permissão"}

	data := new(model.Role)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		errors = []liberr.Errors{
			0: {
				Message: "Um erro inesperado ocorreu, verifique o formato do seu json",
				Type:    "internal_error",
			},
		}
		responseStatus = 500
		message = []string{}
	}
	if (data.User == "" || data.Permission == "") && err == nil {
		errors = []liberr.Errors{
			0: {
				Message: "Necessário usuário e permissão",
				Type:    "status_badrequest",
			},
		}
		responseStatus = 400
		message = []string{}
	}
	if err == nil && data.User != "" && data.Permission != "" {
		ok := role.Update(data)
		if ok {
			message = []string{"Atualização efetuado com sucesso"}
		}
	}

	envelope := response.MountEnvelope(message, responseStatus, r, errors)
	response.Header(w, responseStatus, envelope, response.Headers{})
}

func Delete(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var responseStatus = http.StatusOK
	var errors = []liberr.Errors{}

	vars := mux.Vars(r)
	param := vars["name"]

	role.Delete(param)
	envelope := response.MountEnvelope(param, responseStatus, r, errors)
	response.Header(w, responseStatus, envelope, response.Headers{})
}