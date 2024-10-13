package repository

import (
	"github.com/KolesnikM8O/distributed-task-system/api-gateway/internal/config/config"
	"github.com/jackc/pgx/v5"
)

type Repository interface {
	InitDB(cfg *config.StorageConfig) (*pgx.Conn, error)
}
