package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnectTOPostgres() (*sql.DB, error) {
	connectionString := "postgres://postgres:postgres@localhost:5432/sample?sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	
	if err != nil {
		return nil, err
	}
	
	if err := db.Ping(); err != nil{
		return nil, err
	}
	defer db.Close()

	return db, nil
}
