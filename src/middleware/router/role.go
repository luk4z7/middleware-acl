package router

import (
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"middleware/controller/role"
)

func SetRole(router *mux.Router) *mux.Router {

	router.Handle("/v1/roles", negroni.New(
		negroni.HandlerFunc(role.GetAll),
	)).Methods("GET")

	router.Handle("/v1/roles", negroni.New(
		negroni.HandlerFunc(role.Create),
	)).Methods("POST")

	router.Handle("/v1/roles/{name}", negroni.New(
		negroni.HandlerFunc(role.Get),
	)).Methods("GET")

	router.Handle("/v1/roles", negroni.New(
		negroni.HandlerFunc(role.Update),
	)).Methods("PUT")

	router.Handle("/v1/roles/{name}", negroni.New(
		negroni.HandlerFunc(role.Delete),
	)).Methods("DELETE")

	return router
}