package router

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"student-planner/data"
	"student-planner/db"
	"student-planner/usecases"
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
	userRoute.Handle("/user/ask-gemini", http.HandlerFunc(userHandler.AskGemini)).Methods("POST")
}

func (r *Router) Run(addr string, router *mux.Router) error {
	corsConfig := handlers.AllowedOrigins([]string{"*"})
	corsHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	handler := handlers.CORS(corsConfig, corsHeaders, corsMethods)(router)

	log.Println("Server running on port: ", addr)
	return http.ListenAndServe(addr, handler)
}
