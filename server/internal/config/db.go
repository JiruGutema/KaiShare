package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func ConnectDatabase() {
	fmt.Println("Connecting to Database...")
	cfg := LoadConfig()
	dsn := ConstructDBString(*cfg)

	if dsn == "" {
		log.Fatal(" ERROR: DSN is empty! Check your environment variables or ConstructDBString().")
	}


	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf(" Failed to parse DSN: %v\n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf(" Database connection failed: %v\nDSN: %s\n", err, dsn)
	}

	DB = db
	fmt.Println(" Database Connected Successfully!")
}
