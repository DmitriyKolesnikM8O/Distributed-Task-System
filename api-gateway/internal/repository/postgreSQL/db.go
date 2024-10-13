package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/KolesnikM8O/distributed-task-system/api-gateway/internal/config/config"
	repository "github.com/KolesnikM8O/distributed-task-system/api-gateway/internal/repository"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	db *pgx.Conn
}

func New() repository.Repository {
	return &Repository{
		db: nil,
	}
}

func (r *Repository) InitDB(cfg *config.StorageConfig) (*pgx.Conn, error) {

	log.Printf("Init DB")

	connectString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := pgx.Connect(context.Background(), connectString)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	r.db = db
	_, err = db.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS tasks (id TEXT PRIMARY KEY, status TEXT)")
	if err != nil {
		return nil, err
	}
	return db, nil
}
