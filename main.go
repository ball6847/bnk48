package main

import (
	"net/http"

	flag "github.com/ball6847/bnk48/flag"
	handler "github.com/ball6847/bnk48/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var secret *string

type accesToken struct {
	Token     string `json:"accessToken"`
	ExpiresIn int    `json:"expiresIn"`
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func init() {
	flag.Init()
}

func main() {
	flag.Parse()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.POST("/signup", handler.Signup)
	e.GET("/protected", hello, middleware.JWT([]byte(*flag.Secret)))

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
