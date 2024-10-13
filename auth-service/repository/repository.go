package repository

import "github.com/jackc/pgx/v5"

type Repository interface {
	InitDB() (*pgx.Conn, error)
}
