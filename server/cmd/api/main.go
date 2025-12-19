package main

import (
	"log"

	"github.com/fatih/color"
	"github.com/jirugutema/kaishare/internal/config"
	"github.com/jirugutema/kaishare/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load()
	config.ConnectDatabase()
	c := config.LoadConfig()
	if e != nil {
		log.Fatal("Error loading .env file")
	}
	color.Cyan("Server Starting...")

	server.StartServer(c)
}
