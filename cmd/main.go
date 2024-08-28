package main

import (
	"log"

	"github.com/bonsoirelliot/todo-app"
	"github.com/bonsoirelliot/todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal("error occured while runnung http server: %s", err.Error())
	}
}
