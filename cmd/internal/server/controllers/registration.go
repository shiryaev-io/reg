package controllers

import (
	"encoding/json"
	"net/http"
	"reg/cmd/internal/res/strings"
	"reg/cmd/internal/server/models/requests"
	"reg/cmd/internal/server/models/responses"
	"reg/cmd/internal/server/services"
	"reg/cmd/pkg/logging"
)

// Контроллер для регистрации пользователей
type RegistrationController struct {
	RegService *services.RegistrationService
	Logger     *logging.Logger
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

		return nil, err
	}

	controller.Logger.Infoln(strings.LogRegistration)

	err = controller.RegService.Registration(
		userRequest.Email,
		userRequest.Password,
	)
	if err != nil {
		controller.Logger.Errorf(strings.LogErrorRegistration, err)

		return nil, err
	}

	return &responses.Common{
		Status: http.StatusNoContent,
		Body:   nil,
	}, nil
}
