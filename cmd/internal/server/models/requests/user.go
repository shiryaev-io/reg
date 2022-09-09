package requests

// Данные пользователя для регистрации
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
