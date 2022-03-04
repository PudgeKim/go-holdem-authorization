package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5050
	user     = "postgres"
	password = "mypassword"
	dbname   = "goholdem"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

func NewPostgresDB() (*sqlx.DB, error) {
	conn, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
