// package main

// import (
// 	"fmt"
// 	"time"
// )

// func printNumbers(){
// 	for i:=0; i<10; i++{
// 		fmt.Println(i)
// 		time.Sleep(100 * time.Millisecond)
// 	}
// }

// func printLetters(){
//     for i:='a'; i<='e'; i++ {
//         fmt.Printf("%c\n", i)
//         time.Sleep(500 * time.Millisecond)
//     }
// }


// func main() {
// 	go printNumbers()
// 	go printLetters()

// 	time.Sleep(3 * time.Second)
// 	fmt.Println("Done")
// }


package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

type Response struct {
    Message string `json:"message"`
}

func processRequest(ch chan string) {
    // Simulate a long-running task
    time.Sleep(2 * time.Second)
    ch <- "Task completed"
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
    ch := make(chan string)

    go processRequest(ch) // Start the task in a goroutine

    result := <-ch // Wait for the result from the goroutine

    response := Response{Message: result}
    w.Header().Set("Content-Type", "application/json")// Set the response header
    json.NewEncoder(w).Encode(response) // Send JSON response
}

func main() {
    http.HandleFunc("/api", apiHandler)
    fmt.Println("Server started at:8080")
    http.ListenAndServe(":8080", nil)
}
