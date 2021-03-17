package server

import (
	"echo-restapi/server/handlers"
	"github.com/labstack/echo"
)

func StartServer(handl *handlers.Handlers) {
	e := echo.New()
	NewRouter(e, handl)
	e.Logger.Fatal(e.Start(":8080"))
}
