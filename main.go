package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world....!--\n first page")
}

func handleFunction3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world....!--\n Login page")
}

func handleFunction4(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world....!--\n Admin panel")
}

func handleFunction2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world....!--\n home page")
}

func main() {
	http.HandleFunc("/", handleFunction)
	http.HandleFunc("/home", handleFunction2)
	http.HandleFunc("/login", handleFunction3)
	http.HandleFunc("/admin", handleFunction4)

	fmt.Printf("Starting server at port 8082\n")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal(err)
	}
}
