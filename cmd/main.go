package main

import (
	"echo-restapi/server"
	"echo-restapi/server/handlers"
	"echo-restapi/service"
)

func main() {
	srv := service.NewService()
	handl := handlers.NewHandlers(srv)
	server.StartServer(handl)
}
