package main

import (
	First_server "github.com/headrush95/first_server"
	"github.com/headrush95/first_server/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(First_server.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
