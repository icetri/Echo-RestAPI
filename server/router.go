package server

import (
	"echo-restapi/server/handlers"
	"github.com/labstack/echo"
)

func NewRouter(router *echo.Echo, h *handlers.Handlers) {
	router.POST("/users", h.SaveUser)
	router.GET("/users", h.GetUsers)
	router.GET("/users/:id", h.GetUser)
	router.PUT("/users/:id", h.UpdateUser)
	router.DELETE("/users/:id", h.DeleteUser)
}
