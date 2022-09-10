package postgresql

import (
	"context"
	"fmt"
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
		func(i int) error {
			ctx, cancel := context.WithTimeout(ctx, timeDelay)
			defer cancel()

			pool, err = pgxpool.Connect(ctx, dns)
			if err != nil {
				// logger.Fatalf(strings.LogFatalConnectDb, err)
				return err
			} else {
				// logger.Infoln(strings.LogConnectSuccess)
			}

			return nil
		},
	)

	return pool, nil
}
