package main

import (
	"student-planner/router"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	r := router.NewRouter(route)
	r.RegisterRoute()
	r.Run(":8080", route)
}
