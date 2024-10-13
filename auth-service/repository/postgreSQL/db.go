package repository

import (
	"context"
	"log"

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

func (r *Repository) InitDB() (*pgx.Conn, error) {

	log.Printf("Init DB")
	db, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@postgres:5432/postgres")

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
