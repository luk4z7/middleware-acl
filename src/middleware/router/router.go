package router

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetRole(router)
	router = SetUser(router)
	return router
}