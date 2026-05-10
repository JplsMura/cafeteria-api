package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func NewPostgresConnection() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, fmt.Errorf("Erro ao abrir conexão: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Erro ao pingar no banco: %v", err)
	}

	return db, nil
}
