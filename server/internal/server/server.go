// Package server provides functionalities to start and manage the server.
package server

import (
	"fmt"
	"log"

	"github.com/jirugutema/kaishare/internal/config"
	routes "github.com/jirugutema/kaishare/internal/router"
	"github.com/joho/godotenv"
)

func StartServer() {
	e := godotenv.Load()
	config.ConnectDatabase()
	if e != nil {
		log.Fatal("Error loading .env file")
	}

	c := config.LoadConfig()

	r := routes.Routes()

	err := r.Run(fmt.Sprintf(":%s", c.Port))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
