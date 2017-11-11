package main

import (
	"github.com/codegangsta/negroni"
	"net/http"
	"github.com/rs/cors"
	"runtime"
	"middleware/router"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	rt := router.InitRoutes()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"accept", "authorization", "content-type", "*"},
		AllowedMethods: []string{"PUT", "POST", "DELETE", "OPTIONS", "GET", "HEAD"},
	})
	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(rt)
	http.ListenAndServe(":6060", n)
}