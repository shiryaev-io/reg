package services

import (
	"errors"
	"reg/cmd/internal/server/models/db"
	"reg/cmd/internal/server/models/dtos"

	"golang.org/x/crypto/bcrypt"
)

// Интерфейс для работы с БД пользователей
type UserStorage interface {
	FindByEmail(email string) (*db.User, error)
	Create(email, password string) (int, error)
}

type UserService struct {
	UserStorage UserStorage
}

// Регистрация пользователя
func (service *UserService) Registration(
	email, password string,
) (*dtos.User, error) {
	// TODO: валидация пароля
	// Проверка, что пользователь уже существует
	user, err := service.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("Пользователь с почтовым адресом " + email + " уже существует")
	}

	hashPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}
	idUser, err := service.UserStorage.Create(
		email,
		string(hashPassword),
	)
	if err != nil {
		return nil, err
	}

	return &dtos.User{
		Id: idUser,
	}, nil
}

// Получает пользователя из БД
func (service *UserService) FindByEmail(email string) (*db.User, error) {
	return service.UserStorage.FindByEmail(email)
}
