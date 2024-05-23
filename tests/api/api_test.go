package tests

import (
	"log"
	"net/http"
	"testing"

	w "github.com/bjssacademy/go-play/httpwrapper"
	"github.com/stretchr/testify/assert"
)

type Post struct {
    UserID int    `json:"userId"`
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Body   string `json:"body"`
}

func TestAPIExample(t *testing.T) {

	httpClient := &w.HttpClient{HttpClient: &http.Client{}}
	
	url := "https://jsonplaceholder.typicode.com/posts"
	
	posts := &[]Post{}
	headers := map[string]string{"header_example": "header_value"}

	res, err := httpClient.Get(url, headers, posts)
	if err != nil {
		t.Fatal(err.Error())
	}

	// Check if response status code is OK
	if res.StatusCode != http.StatusOK {
		log.Fatalf("unexpected status code: %v", res.StatusCode)
	}

	assert.Greater(t, len(*posts), 0, "Number of posts should be greater than zero")

}