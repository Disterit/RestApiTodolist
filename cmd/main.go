package main

import (
	RestApiTodolist "ResTApiTodolist"
	"ResTApiTodolist/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(RestApiTodolist.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while runing http server: %s", err.Error())
	}
}
