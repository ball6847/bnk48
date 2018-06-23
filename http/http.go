package http

import (
	"crypto/tls"
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

// GetInsecureResource from p'yod machine
func GetInsecureResource() ([]byte, error) {
	url := "http://192.168.100.3:8080/thai"

	req, _ := http.NewRequest("GET", url, nil)

	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body, nil
}
