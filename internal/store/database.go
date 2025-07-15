package store

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func Open() (*sql.DB, error) {
	db,err := sql.Open("pgx", "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable")
	if err!= nil {
		return nil, fmt.Errorf("db: open %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db: ping %w", err)
	}

	// Configure connection pool settings
	db.SetMaxOpenConns(25)									// Maximum number of open connections
	db.SetMaxIdleConns(25)									// Maximum number of idle connections
	db.SetConnMaxLifetime(5 * time.Minute)	// Maximum connection lifetime
	db.SetConnMaxIdleTime(5 * time.Minute)	// Maximum idle time for connections

	fmt.Println("Connected to Database...")
	return db, nil
}