package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"rest-ws/database"
	"rest-ws/repository"
	"rest-ws/websocket"
)

type Config struct {
	Port        string
	JWTSecret   string
	DataBaseUrl string
}

type Server interface {
	Config() *Config
	Hub() *websocket.Hub
}

type Broker struct {
	config *Config
	router *mux.Router
	hub    *websocket.Hub
}

func (b *Broker) Config() *Config {
	return b.config
}

func (b *Broker) Hub() *websocket.Hub {
	return b.hub
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}

	if config.JWTSecret == "" {
		return nil, errors.New("secret is required")
	}

	if config.DataBaseUrl == "" {
		return nil, errors.New("database url is required")
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
		hub:    websocket.NewHub(),
	}

	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	log.Println("Creating router")
	b.router = mux.NewRouter()
	binder(b, b.router)
	handler := cors.Default().Handler(b.router)
	log.Println("repositori db")
	repo, err := database.NewPostgresRepository(b.config.DataBaseUrl)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Start hub")
	go b.hub.Run()

	log.Println("set repository")
	repository.SetRepository(repo)
	binder(b, b.router)
	log.Println("Starting server on port", b.Config().Port)

	if err := http.ListenAndServe(b.config.Port, handler); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
