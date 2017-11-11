package router

import (
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"middleware/controller/user"
)

func SetUser(router *mux.Router) *mux.Router {

	router.Handle("/v1/users/{user}/roles", negroni.New(
		negroni.HandlerFunc(user.GetUserRoles),
	)).Methods("GET")

	router.Handle("/v1/users/{user}/roles/{role}", negroni.New(
		negroni.HandlerFunc(user.DeleteRoleFromUser),
	)).Methods("DELETE")

	router.Handle("/v1/users/{subject}/resource/{object}/permission/{action}", negroni.New(
		negroni.HandlerFunc(user.CheckUserPermission),
	)).Methods("GET")

	return router
}