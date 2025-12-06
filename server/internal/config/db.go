package config

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func ConnectDatabase() {
	fmt.Println("Connecting to Database")
	cfg := LoadConfig()
	dsn := ConstructDBString(cfg)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		fmt.Println("Error connecting to database")
		return
	}

	DB = db

	fmt.Println("Database Connected Successfully!")
}
