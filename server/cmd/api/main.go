package main

import (
	"fmt"
	"log"

	"github.com/jirugutema/kaishare/internal/config"
	routes "github.com/jirugutema/kaishare/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Main thread started!")
	config.ConnectDatabase()

	e := godotenv.Load()
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
