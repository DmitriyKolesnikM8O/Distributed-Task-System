package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/KolesnikM8O/distributed-task-system/auth-service/internal/config/config"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	db *pgx.Conn
}

func New() *Repository {
	return &Repository{
		db: nil,
	}
}

func (r *Repository) InitDB(cfg *config.StorageConfig) (*pgx.Conn, error) {

	log.Printf("Init DB")
	//db, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@postgres:5432/postgres")
	dbString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := pgx.Connect(context.Background(), dbString)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	r.db = db
	_, err = db.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) NOT NULL UNIQUE,
			email VARCHAR(100) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			role VARCHAR(10) NOT NULL
		);
	`)
	if err != nil {
		return nil, err
	}
	return db, nil
}
