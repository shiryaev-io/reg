package services

import (
	"errors"
	"fmt"
	"reg/cmd/internal/res/strings"
	"reg/cmd/internal/server/models/db"
	"reg/cmd/pkg/logging"

	"github.com/jackc/pgx/v4"
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
	service.Logger.Infoln(strings.LogFingUserByEmail)
	// Проверка, что пользователь уже существует
	user, err := service.FindByEmail(email)
	if err != nil && err != pgx.ErrNoRows {
		service.Logger.Errorf(strings.LogErrorFindUser, err)

		return err
	}
	if user != nil {
		service.Logger.Errorln(strings.LogErrorUserAlreadyExists)

		errorMessage := fmt.Sprintf(strings.ErrorUserWithEmailExists, email)
		return errors.New(errorMessage)
	}

	service.Logger.Infoln(strings.LogGenerateHashedPassword)

	hashPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		service.Logger.Errorln(strings.LogErrorGenerateHashedPassword)

		return err
	}

	service.Logger.Infoln(strings.LogCreateUserInDb)

	err = service.UserStorage.Create(
		email,
		string(hashPassword),
	)
	if err != nil {
		service.Logger.Errorf(strings.LogErrorCreateUserInDb, err)

		return err
	}

	return nil
}

// Получает пользователя из БД
func (service *UserService) FindByEmail(email string) (*db.User, error) {
	return service.UserStorage.FindByEmail(email)
}
