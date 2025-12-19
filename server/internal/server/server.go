// Package server provides functionalities to start and manage the server.
package server

import (
	"fmt"

	"github.com/jirugutema/kaishare/internal/config"
	routes "github.com/jirugutema/kaishare/internal/router"
)

func StartServer(c *config.Config) {
	r := routes.Routes()

	err := r.Run(fmt.Sprintf(":%s", c.Port))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
