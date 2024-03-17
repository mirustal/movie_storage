package database

import (
	"context"
	"database/sql"
	"fmt"
	"movie_storage/pkg/configs"
	_ "github.com/lib/pq"
)

type Client interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Database struct {
	db *sql.DB
}

func NewDatabase(cfg *configs.ConfigPostgressDB) (*Database, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.Name, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (db *Database) Close() {
	db.db.Close()
}

func (db *Database) GetDB() *sql.DB {
	return db.db
}
