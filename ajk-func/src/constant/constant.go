package constant

import (
	"log"
	"os"
	"strconv"
)

var (
	// ResponseHeaders is http response headers
	ResponseHeaders = map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "origin,Accept,Authorization,Content-Type",
		"Content-Type":                 "application/json",
	}

	// Concurrency is number of workers
	Concurrency int
)

func init() {
	c, err := strconv.Atoi(os.Getenv("CONCURRENCY"))
	if err != nil {
		log.Fatal(err)
	}
	Concurrency = c
}
