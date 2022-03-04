package main

import (
	"fmt"
	"github.com/Pudgekim/db"
	"github.com/Pudgekim/handlers"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	name VARCHAR(50) UNIQUE NOT NULL,
	email VARCHAR(50) UNIQUE NOT NULL,
	balance BIGSERIAL NOT NULL
);
`

func main() {

	conn, err := db.NewPostgresDB()
	if err != nil {
		panic(fmt.Sprintf("postgres connection fail: %s", err.Error()))
	}

	err = conn.Ping()
	if err != nil {
		panic(fmt.Sprintf("postgres ping error: %s", err.Error()))
	}

	fmt.Println("ping success!")

	conn.MustExec(schema)

	handler := handlers.NewHandler(conn)
	router := handler.Routes()

	router.Run(":3000")

}
