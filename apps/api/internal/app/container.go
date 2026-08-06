package app

import (
	"log/slog"

	_ "github.com/jackc/pgx/v5/stdlib" // PostgreSQL driver (pgx) for database/sql
	"github.com/jmoiron/sqlx"

	"app/api/internal/config"
	"app/api/internal/repository"
	"app/api/migrations"
)

type Container struct {
	Config *config.Config
	DB     *sqlx.DB
	Log    *slog.Logger
	Repo   *repository.Repository
}

func NewContainer(cfg *config.Config, log *slog.Logger) (*Container, error) {
	db, err := sqlx.Connect("pgx", cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	migrator := migrations.NewMigrator(cfg.DatabaseURL)
	if err := migrator.MigrateUp(); err != nil {
		return nil, err
	}

	repo := repository.NewRepository(db)

	return &Container{
		Config: cfg,
		DB:     db,
		Log:    log,
		Repo:   repo,
	}, nil
}

func (c *Container) Close() error {
	return c.DB.Close()
}
