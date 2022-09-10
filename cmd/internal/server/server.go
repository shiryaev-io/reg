package server

import (
	"log"
	"net/http"
	"os"
	"reg/cmd/internal/server/routes"
	"reg/cmd/internal/server/services"
	"reg/cmd/pkg/logging"
)

const (
	serviceHost = "SERVICE_HOST"
	servicePort = "SERVICE_PORT"
)

// Структура сервера
type server struct {
	router     *routes.ApiRoute
	regService *services.RegistrationService
	logger     *logging.Logger
}

// Инициалиазция сервера
func NewServer(
	router *routes.ApiRoute,
	regService *services.RegistrationService,
	logger     *logging.Logger,
) *server {
	router.Init()
	return &server{
		router:     router,
		regService: regService,
		logger:     logger,
	}
}

// Запускает сервер
func (server *server) Run() {
	host := os.Getenv(serviceHost)
	port := os.Getenv(servicePort)
	serviceUrl := host + ":" + port

	log.Fatal(http.ListenAndServe(serviceUrl, server.router.Router))
}
