package router

import (
	"go-echo-v2/internal/handlers/index"
	"go-echo-v2/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo) {
	e.GET("/", index.Index)

	v1 := e.Group("/api/v1")
	v1.GET("/", index.Index, middleware.AuthMiddleware)
}
