package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func NewPostgresDB() (*sql.DB, error) {
	p1 := fmt.Sprintf("%s:%s", os.Getenv("DBUSER"), os.Getenv("DBPASS"))
	p2 := fmt.Sprintf("%s:%s", "localhost", "5433")
	p3 := fmt.Sprintf("%s?sslmode=disable", os.Getenv("DBNAME"))

	// PostgreSQL connection string
	psqlInfo := fmt.Sprintf("postgres://%s@%s/%s", p1, p2, p3)
	fmt.Println(psqlInfo)
	// Open a connection to the database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Check the connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}
	fmt.Println("[!] - Database connected: POSTGRES")
	return db, nil
}
