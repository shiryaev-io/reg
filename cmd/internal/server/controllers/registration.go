package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, exceptions.BadRequest("Необходимо передавать Body в запрос", err)
	}

	userRequest := &requests.User{}
	err = json.Unmarshal(body, userRequest)
	if err != nil {
		return nil, exceptions.BadRequest("Неверный формат Body", err)
	}

	err = controller.UserService.Registration(
		userRequest.Email,
		userRequest.Password,
	)
	if err != nil {
		return nil, exceptions.BadRequest("Не удалось зарегистрироваться", err)
	}

	return &responses.Common{
		Status: http.StatusNoContent,
		Body:   nil,
	}, err
}
