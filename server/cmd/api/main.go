package main

import (
	"github.com/fatih/color"
	"github.com/jirugutema/kaishare/internal/server"
)

func main() {
	color.Cyan("Server Starting...")
	server.StartServer()
}
