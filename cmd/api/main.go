package main

import (
    "log"
    "gopastbin/internal/server"
)

func main() {

    srv := server.NewServer("3000")
    if err := srv.Start(); err != nil {
        log.Fatal(err)
    }
}
