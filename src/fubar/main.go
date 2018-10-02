package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		if os.Getenv("WORKING") == "" {
			res.Write([]byte("Not so fast, I'm still broken :(\n"))
		} else {
			res.Write([]byte("You "))
			res.Write([]byte("win!\n"))
		}
	})
	port := "8080"
	if os.Getenv("FUBAR_PORT") != "" {
		port = os.Getenv("FUBAR_PORT")
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
