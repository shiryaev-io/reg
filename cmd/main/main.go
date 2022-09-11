package main

import (
	"context"
	"os"
	"os/signal"
	"reg/cmd/internal/res/strings"
	"reg/cmd/internal/server"
	"reg/cmd/internal/server/adapters/db/postgresql"
	"reg/cmd/internal/server/config"
	"reg/cmd/internal/server/routes"
	"reg/cmd/internal/server/services"
	"reg/cmd/pkg/logging"
	"syscall"

	"github.com/gorilla/mux"
)

func main() {
	logger := logging.GetLogger()
	
	// Канал для сигналов
	sig := make(chan bool)
	// Основной канал
	loop := make(chan error)

	// Мониторинг сигналов
	go listenerSignal(sig, logger)

	for quit := false; !quit; {
		go func() {
			initAndRunServer(logger)
			loop <- nil
		}()

		// Блокировка программы при получении сигнала
		select {
		// Прерывается выполнение программы
		case quit = <-sig:
		// Продолжается выполлнение программы
		case <-loop:
		}
	}
}

func listenerSignal(sig chan bool, logger *logging.Logger) {
	var quit bool

	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)

	for signal := range c {
		logger.Infof(strings.LogGetSignalSuccess, signal.String())

		switch signal {
		case syscall.SIGINT, syscall.SIGTERM:
			quit = true
		case syscall.SIGHUP:
			quit = false
		}

		if quit {
			quit = false
			// TODO: closeDB(), closeLog()
		}

		// Оповещение о прекращении работы
		sig <- quit
	}
}

func initAndRunServer(logger *logging.Logger) {
	// Получение конфигурации
	dbConfig := config.NewConfigDb()
	// Коннект к БД
	registrationDatabase, err := postgresql.NewRegistrationDatabase(
		context.Background(),
		dbConfig,
		logger,
	)
	if err != nil {
		logger.Errorln(strings.LogGetDatabaseError)
	}

	userService := &services.UserService{
		UserStorage: registrationDatabase,
		Logger:      logger,
	}
	regService := &services.RegistrationService{
		UserService: userService,
		Logger:      logger,
	}

	router := mux.NewRouter()

	apiRouter := &routes.ApiRoute{
		Router:     router,
		RegService: regService,
		Logger:     logger,
	}

	// Инициализация сервера
	serv := server.NewServer(
		apiRouter,
		regService,
		logger,
	)

	// Запуск сервера
	serv.Run()
}
