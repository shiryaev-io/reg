package postgresql

import (
	"context"
	"reg/cmd/internal/server/adapters/db/postgresql/queries"
	"reg/cmd/internal/server/config"
	"reg/cmd/internal/server/models/db"
	"reg/cmd/pkg/logging"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Стурктура для БД
type registrationDatabase struct {
	pool   *pgxpool.Pool
	logger *logging.Logger
}

// Инициализация сткуртуры RegistrationDatabase
func NewRegistrationDatabase(
	ctx context.Context,
	config *config.ConfigDb,
	logger *logging.Logger,
) (*registrationDatabase, error) {
	pool, err := connectDb(ctx, config, logger)
	if err != nil {
		return nil, err
	}

	return &registrationDatabase{
		pool:   pool,
		logger: logger,
	}, err
}

// Находит пользователя в БД
func (storage *registrationDatabase) FindByEmail(email string) (*db.User, error) {
	user := &db.User{}
	query := queries.QuerySelectUserByEmail
	err := storage.pool.
		QueryRow(context.Background(), query, email).
		Scan(
			user.Id,
			user.Email,
			user.Password,
		)
	if err != nil {
		return nil, err
	}
	return user, err
}

// Создает пользователя в БД
func (storage *registrationDatabase) Create(email, password string) error {
	query := queries.QueryInsertUser
	err := storage.pool.
		QueryRow(context.Background(), query, email, password).
		Scan()
	if err != nil {
		return err
	}
	return nil
}
