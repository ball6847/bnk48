package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Post post
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// Posts posts
type Posts struct {
	Collection []Post
}

// GetPosts get posts from jsonplaceholder
func GetPosts() ([]Post, error) {
	res, _ := http.Get("https://jsonplaceholder.typicode.com/posts")
	body, _ := ioutil.ReadAll(res.Body)

	items := make([]Post, 0)

	err := json.Unmarshal(body, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}
