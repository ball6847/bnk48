package main

import (
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type accesToken struct {
	Token     string `json:"accessToken"`
	ExpiresIn int    `json:"expiresIn"`
}

type signupPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func signup(c echo.Context) (err error) {
	payload := new(signupPayload)

	if err = c.Bind(payload); err != nil {
		c.Error(err)
	}

	if payload.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Email is required",
		})
	}

	session, err := mgo.Dial("localhost")
	if err != nil {
		c.Error(err)
	}

	collection := session.DB("odds").C("credentials")

	err = collection.Upsert(bson.M{"email": payload.Email}, bson.M{"$set": payload})
	if err != nil {
		c.Error(err)
	}

	return c.JSON(http.StatusOK, payload)
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.POST("/signup", signup)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
