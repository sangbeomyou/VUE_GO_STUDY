package main

import (
	"goBack/handlers" //  함수가 있는 패키지 임포트
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/test", handlers.GetUsersHandler)

	e.POST("/user/login", handlers.GetUsersLoginHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
