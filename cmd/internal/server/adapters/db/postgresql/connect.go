package postgresql

import (
	"context"
	"fmt"
	"reg/cmd/internal/res/strings"
	"reg/cmd/internal/server/config"
	"reg/cmd/pkg/logging"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/tinrab/retry"
)

const (
	dbUrl     = "postgresql://%s:%s@%s:%s/%s"
	timeDelay = 5 * time.Second
)

// Подключение к БД
func connectDb(
	ctx context.Context,
	dbConfig *config.ConfigDb,
	logger *logging.Logger,
) (pool *pgxpool.Pool, err error) {
	dns := fmt.Sprintf(
		dbUrl,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)

	retry.ForeverSleep(
		timeDelay,
		func(attempt int) error {
			logger.Infof(strings.LogAttemptConnectDb, attempt)

			ctx, cancel := context.WithTimeout(ctx, timeDelay)
			defer cancel()

			logger.Infoln(strings.LogTryConnectDb)

			pool, err = pgxpool.Connect(ctx, dns)
			if err != nil {
				logger.Errorf(strings.LogFatalConnectDb, err)
				return err
			} else {
				logger.Infoln(strings.LogConnectSuccess)
			}

			return nil
		},
	)

	return pool, nil
}
