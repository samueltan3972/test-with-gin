package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/http/httptest"
	"test/gin-test/database"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
)

type Fruit struct {
	ID            int64   `db:"id, primarykey" json:"id"`
	Fruit_label   string  `db:"fruit_label" json:"email"`
	Fruit_name    string  `db:"fruit_name" json:"fruit_name"`
	Fruit_subtype string  `db:"fruit_subtype" json:"fruit_subtype"`
	Mass          float64 `db:"mass" json:"mass"`
	Width         float64 `db:"width" json:"width"`
	Height        float64 `db:"height" json:"height"`
	Color_score   float64 `db:"color_score" json:"color_score"`
}

func TestHelloWorld(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	// The setupServer method, that we previously refactored
	// is injected into a test server
	ts := httptest.NewServer(setupRouter())
	// Shut down the server and block until all requests have gone through
	defer ts.Close()

	// Make a request to our server with the {base url}/ping
	resp, err := http.Get(fmt.Sprintf("%s/hello", ts.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	assert.Equal(t, string(body), "Hello World")
}

func TestDatabaseOperation(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	database.Init()

	ts := httptest.NewServer(setupRouter())

	defer ts.Close()

	// Make a request to our server with the {base url}/ping
	resp, err := http.Get(fmt.Sprintf("%s/database", ts.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	var fruit []Fruit

	err = json.Unmarshal(body, &fruit)

	if err != nil {
		fmt.Println(err)
	}

	assert.Equal(t, len(fruit), 59)
}

func TestFibonacci(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	ts := httptest.NewServer(setupRouter())

	defer ts.Close()

	// Make a request to our server with the {base url}/ping
	resp, err := http.Get(fmt.Sprintf("%s/fibonacci", ts.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	var numbers []big.Int

	err = json.Unmarshal(body, &numbers)

	if err != nil {
		fmt.Println(err)
	}

	assert.Equal(t, len(numbers), 5000)
}
