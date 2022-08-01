package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"rest-ws/handlers"
	"rest-ws/server"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"rest-ws/middleware"
)

func main() {
	log.Println("loading envs")
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("gets envs")
	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	log.Println("creating server")
	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DataBaseUrl: DATABASE_URL,
	})
	log.Println("Server created")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("start server")
	s.Start(BindRouter)
}

func BindRouter(s server.Server, r *mux.Router) {
	api := r.PathPrefix("/api/v1").Subrouter()

	api.Use(middleware.CheckAuthMiddleware(s))

	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)
	api.HandleFunc("/me", handlers.MeHandler(s)).Methods(http.MethodGet)
	api.HandleFunc("/posts", handlers.InsertPostHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/posts/{id}", handlers.GetPostByIdHandler(s)).Methods(http.MethodGet)
	api.HandleFunc("/posts/{id}", handlers.UpdatePostHandler(s)).Methods(http.MethodPut)
	api.HandleFunc("/posts/{id}", handlers.DeletePostHandler(s)).Methods(http.MethodDelete)
	r.HandleFunc("/posts", handlers.ListPostHandler(s)).Methods(http.MethodGet)

	r.HandleFunc("/ws", s.Hub().HandleWebSocket)
}
