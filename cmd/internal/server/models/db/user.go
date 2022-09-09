package db

// Данные пользователя, которые хранятся в БД
type User struct {
	Id       int
	Email    string
	Password string
}
