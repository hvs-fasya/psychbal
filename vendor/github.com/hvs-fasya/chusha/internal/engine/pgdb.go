package engine

import (
	"database/sql"
)

// DB - глобальная переменная для доступа к БД
var DB DBInterface

// PgDB engine PostgreSQL storage
type PgDB struct {
	Conn *sql.DB
}

// NewPgDB make PostgreSQL storage
func NewPgDB(dns string) (*PgDB, error) {
	conn, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	db := PgDB{conn}
	return &db, nil
}
