package main

import (
	"fmt"
	"log"
	"net/http"
)

type List struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func makeList() {

}

func main() {
	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Todo List")
	})

	fmt.Println("Server started on https://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
