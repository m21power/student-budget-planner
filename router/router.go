package router

import (
	"log"
	"net/http"
	"student-planner/data"
	"student-planner/db"
	"student-planner/usecases"

	"github.com/gorilla/mux"
)

type Router struct {
	route *mux.Router
}

func NewRouter(r *mux.Router) *Router {
	return &Router{route: r}
}
func (r *Router) RegisterRoute() {
	db, err := db.ConnectDB()
	if err != nil {
		log.Println("cannot connect to db")
		return
	}
	userStore := data.NewUserStore(db)
	userUsecase := usecases.NewUserUsecase(userStore)
	userHandler := data.NewUserHandler(*userUsecase)
	userRoute := r.route.PathPrefix("/api/v1").Subrouter()
	userRoute.Handle("/user", http.HandlerFunc(userHandler.GetUser)).Methods("GET")
	userRoute.Handle("/user/login", http.HandlerFunc(userHandler.Login)).Methods("POST")
	userRoute.Handle("/user/register", http.HandlerFunc(userHandler.Register)).Methods("POST")
	userRoute.Handle("/user/update-badge", http.HandlerFunc(userHandler.UpdateBadge)).Methods("PUT")

}
func (r *Router) Run(addr string, router *mux.Router) error {
	log.Println("Server running on port: ", addr)
	return http.ListenAndServe(addr, router)
}
