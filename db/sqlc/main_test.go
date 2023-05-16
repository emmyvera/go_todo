package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq" // Postgres Driver
)

const (
	DBDriver = "postgres"
	DBSource = "postgresql://root:secret@localhost:5432/todo_db?sslmode=disable"
)

var testQueries *Queries

var testDB *sql.DB

func TestMain(m *testing.M) {
	conn, err := sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Fatal("Cannot open database: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
