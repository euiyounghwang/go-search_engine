package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

/*
# Unit Test (*_test.go) ==> ./tests/
go test ./tests
*/

func main() {
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World1!")
    })

    e.GET("/health", func(c echo.Context) error {
        return c.String(http.StatusOK, "Health is OK1!!")
    })

    httpPort := os.Getenv("PORT")
    if httpPort == "" {
        httpPort = "9080"
    }

    e.Logger.Fatal(e.Start(":" + httpPort))
}
