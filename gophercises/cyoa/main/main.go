package main

import (
	"fmt"
	"net/http"

	"huangblog.com/cyoa"
)

func main() {
	h := cyoa.NewHandler(cyoa.ReadJSONFile("../gopher.json"))
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", h)
}
