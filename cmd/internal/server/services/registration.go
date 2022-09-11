package services

import (
	"reg/cmd/internal/res/strings"
	"reg/cmd/pkg/logging"
)

// Сервис регистрации
// Содержит другие сервисы
type RegistrationService struct {
	UserService       *UserService
	ValidationService *ValidationService
	Logger            *logging.Logger
}

// Регистрация пользователя
func (service *RegistrationService) Registration(email, password string) error {
	service.Logger.Infoln(strings.LogStartValidData)

	err := service.ValidationService.Validate(email, password)
	if err != nil {
		service.Logger.Errorf(strings.LogErrorValidData, err)

		return err
	}

	service.Logger.Infoln(strings.LogStartRegistration)

	return service.UserService.Registration(email, password)
}
