package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func InitDB() (*pgx.Conn, error) {

	log.Printf("Init DB")
	db, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/postgres")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	_, err = db.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS tasks (id TEXT PRIMARY KEY, status TEXT)")
	if err != nil {
		return nil, err
	}
	return db, nil
}
