package services

import (
	"errors"
	"net/mail"
	"reg/cmd/internal/res/strings"
	"reg/cmd/pkg/logging"
)

// Сервис для валидации данных
type ValidationService struct {
	Logger *logging.Logger
}

// Валидация почты и пароля
func (service *ValidationService) Validate(email, password string) error {
	isEmailValid := service.validateEmail(email)
	if !isEmailValid {
		return errors.New(strings.ErrorEmailValid)
	}
	isPasswordValid := service.validatePassword(password)
	if !isPasswordValid {
		return errors.New(strings.ErrorPasswordValid)
	}
	return nil
}

// Валидация почты
func (service *ValidationService) validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// Валидация пароля
func (service *ValidationService) validatePassword(password string) bool {
	// TODO: реализовать нормальую валидацию пароля
	return len(password) != 0
}
