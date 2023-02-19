package main

import (
	"net/http"
	"example.com/packages/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
