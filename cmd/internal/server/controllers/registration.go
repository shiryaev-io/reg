package controllers

import (
	"encoding/json"
	"net/http"
	"reg/cmd/internal/res/strings"
	"reg/cmd/internal/server/exceptions"
	"reg/cmd/internal/server/models/requests"
	"reg/cmd/internal/server/models/responses"
	"reg/cmd/internal/server/services"
	"reg/cmd/pkg/logging"
)

// Контроллер для регистрации пользователей
type RegistrationController struct {
	UserService *services.UserService
	Logger      *logging.Logger
}

// Сценарий регистрации пользователей
func (controller *RegistrationController) Registration(
	response http.ResponseWriter,
	request *http.Request,
) (*responses.Common, error) {
	controller.Logger.Infoln(strings.LogGettingRequestBody)
	
	userRequest := &requests.User{}
	err := json.NewDecoder(request.Body).Decode(userRequest)
	if err != nil {
		controller.Logger.Errorf(strings.LogErrorInvalidRequestBodyFormat, err)

		return nil, exceptions.BadRequest("Не удалось зарегистрироваться. Попробуйте позже", err)
	}

	controller.Logger.Infoln(strings.LogRegistration)

	err = controller.UserService.Registration(
		userRequest.Email,
		userRequest.Password,
	)
	if err != nil {
		controller.Logger.Errorf(strings.LogErrorRegistration, err)

		return nil, exceptions.BadRequest("Не удалось зарегистрироваться", err)
	}

	return &responses.Common{
		Status: http.StatusNoContent,
		Body:   nil,
	}, err
}
