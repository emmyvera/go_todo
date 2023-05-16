package main

import (
	"database/sql"
	"log"

	"github.com/emmyvera/go_todo/api"
	db "github.com/emmyvera/go_todo/db/sqlc"
	_ "github.com/lib/pq" // Postgres Driver
)

const (
	DBDriver      = "postgres"
	DBSource      = "postgresql://root:secret@localhost:5432/todo_db?sslmode=disable"
	serverAddress = "0.0.0.0:5000"
)

func main() {
	conn, err := sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Fatal("Cannot open database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
