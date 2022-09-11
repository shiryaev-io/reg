package routes

import (
	"reg/cmd/internal/server/controllers"
	"reg/cmd/internal/server/middlewares"
	"reg/cmd/internal/server/services"
	"reg/cmd/pkg/logging"

	"github.com/gorilla/mux"
)

const (
	get  = "GET"
	post = "POST"

	// Путь для регистрации пользователя
	urlRegistration = "/registration"
)

type ApiRoute struct {
	Router     *mux.Router
	RegService *services.RegistrationService
	Logger     *logging.Logger
}

func (route *ApiRoute) Init() {
	regController := &controllers.RegistrationController{
		UserService: route.RegService.UserService,
		Logger:      route.Logger,
	}

	route.handlerFunc(
		urlRegistration,
		regController.Registration,
	).Methods(post)

	route.Router.Use(middlewares.HeaderMiddleware)
}

func (route *ApiRoute) handlerFunc(
	path string,
	handler middlewares.ErrorHandlerFunc,
) *mux.Route {
	return route.Router.HandleFunc(
		path,
		middlewares.ErrorMiddleware(handler).ServeHTTP,
	)
}
