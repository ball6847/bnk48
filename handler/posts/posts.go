package posts

import (
	"net/http"
	"strconv"

	api "github.com/ball6847/bnk48/http"
	"github.com/labstack/echo"
)

// Posts handler for /posts
func Posts(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Error(err)
	}

	filtered := make([]api.Post, 0)

	posts, err := api.GetPosts()
	if err != nil {
		c.Error(err)
	}

	for _, post := range posts {
		if post.UserID == int(id) {
			filtered = append(filtered, post)
		}
	}

	return c.JSON(http.StatusOK, filtered)
}
