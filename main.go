package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// User struct
type Message struct {
    Message string `json:"message"`
}

type User struct {
    Name string `json:"name"`
    Age int `json:"age"`
}

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
    
    // e.GET("/api/v1/", default_url)
    // e.GET("/api/v1/user/:name", getUser)
    // e.POST("/api/v1/user", createUser)

    httpPort := os.Getenv("PORT")
    if httpPort == "" {
        httpPort = "9080"
    }
    
    // e.GET("/swagger/*", echoSwagger.WrapHandler)

    e.Logger.Fatal(e.Start(":" + httpPort))
}


/*
// @Summary Default URL
// @Description Default URL
// @Accept json
// @Produce json
// @Success 200 {object} Message
// @Router / [get]
func default_url(c echo.Context) error {
    // return c.String(http.StatusOK, "Health is OK1!!")
    
    message := new(Message)
    message.message = "test"
    fmt.Println("@$%% '- ", c.JSONPretty(http.StatusOK, *message, "Health is OK1!!"))
    return c.JSONPretty(http.StatusOK, *message, "Health is OK1!!")
}

// @Summary Get user
// @Description Get user's info
// @Accept json
// @Produce json
// @Param name path string true "name of the user"
// @Success 200 {object} User
// @Router /user/{name} [get]
func getUser(c echo.Context) error {
    user := new(User)
    // 
    return c.JSONPretty(http.StatusOK, *user, "  ")
}

// @Summary Create user
// @Description Create new user
// @Accept json
// @Produce json
// @Param userBody body User true "User Info Body"
// @Success 200 {object} User
// @Router /user [post]
func createUser(c echo.Context) error {
    user := new(User)
    // 
    return c.JSONPretty(http.StatusOK, *user, "  ")
}
*/