package services

import "reg/cmd/pkg/logging"

// Сервис регистрации
// Содержит другие сервисы
type RegistrationService struct {
	UserService *UserService
	Logger      *logging.Logger
}
