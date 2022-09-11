package strings

const (
	LogGetSignalSuccess  = "Сигнал получен!"
	LogGetDatabaseError  = "База данных не получена"
	LogGetHostAndPortEnv = "Получение Host и Port из файла переменных окружения .env"
	LogRunServer         = "Запуск сервера на URL: %s"

	LogAttemptConnectDb = "Попытка подключения к БД: %s"
	LogTryConnectDb     = "Подключение к БД"
	LogFatalConnectDb   = "Не удалось подключиться к БД, ошибка: %v"
	LogConnectSuccess   = "База данных успешно подключена"

	LogGettingRequestBody            = "Получение тела запроса Body"
	LogErrorInvalidRequestBodyFormat = "Неверный формат тела запроса: %v"
	LogRegistration                  = "Регистрация пользователя"
	LogErrorRegistration             = "Не удалось зарегистрироваться: %v"

	LogFingUserByEmail             = "Поиск пользователя по email"
	LogErrorFindUser               = "Ошибка поиска пользователя в БД: %v"
	LogErrorUserAlreadyExists      = "Пользователь уже существует в БД"
	LogGenerateHashedPassword      = "Генерация захешированного пароля"
	LogErrorGenerateHashedPassword = "Ошибка генерации захешированного пароля: %v"
	LogCreateUserInDb              = "Создание пользователя в БД"
	LogErrorCreateUserInDb         = "Ошибка создания пользователя в БД: %v"

	LogStartValidData    = "Начало валидации данных"
	LogErrorValidData    = "Ошибка валидации данных: %v"
	LogStartRegistration = "Начало регистрации пользователя"
)
