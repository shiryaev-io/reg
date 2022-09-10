package services

import (
	"errors"
	"reg/cmd/internal/server/models/db"
	"reg/cmd/pkg/logging"

	"golang.org/x/crypto/bcrypt"
)

// Интерфейс для работы с БД пользователей
type UserStorage interface {
	FindByEmail(email string) (*db.User, error)
	Create(email, password string) error
}

type UserService struct {
	UserStorage UserStorage
	Logger      *logging.Logger
}

// Регистрация пользователя
func (service *UserService) Registration(
	email, password string,
) error {
	// TODO: валидация пароля
	// Проверка, что пользователь уже существует
	user, err := service.FindByEmail(email)
	if err != nil {
		return err
	}
	if user != nil {
		return errors.New("Пользователь с почтовым адресом " + email + " уже существует")
	}

	hashPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	err = service.UserStorage.Create(
		email,
		string(hashPassword),
	)
	if err != nil {
		return err
	}

	return nil
}

// Получает пользователя из БД
func (service *UserService) FindByEmail(email string) (*db.User, error) {
	return service.UserStorage.FindByEmail(email)
}
