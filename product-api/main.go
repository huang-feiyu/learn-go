package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Hello, world!")
		d, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, "Oops, something went wrong!", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(writer, "Hello, %s!", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye, world!")
	})

	http.ListenAndServe(":9090", nil)
}
