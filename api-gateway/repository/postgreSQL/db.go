package repository

import (
	"context"
	"log"

	repository "github.com/KolesnikM8O/distributed-task-system/api-gateway/repository"
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

func (r *Repository) InitDB() (*pgx.Conn, error) {

	log.Printf("Init DB")
	db, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@postgres:5432/postgres")

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
