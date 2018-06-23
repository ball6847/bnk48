package handler

import (
	"net/http"

	flag "github.com/ball6847/bnk48/flag"
	p "github.com/ball6847/bnk48/payload"
	u "github.com/ball6847/bnk48/utils"
	"github.com/globalsign/mgo"
	"github.com/labstack/echo"
)

// Signup handler for signup
func Signup(c echo.Context) (err error) {
	payload := new(p.Signup)

	if err = c.Bind(payload); err != nil {
		c.Error(err)
	}

	// validate email
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

	// has passowrd before saving to database
	payload.HashPassword()

	err = collection.Insert(payload)
	if err != nil {
		c.Error(err)
	}

	token, err := u.GenerateToken(payload, *flag.Secret)

	if err != nil {
		c.Error(err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
