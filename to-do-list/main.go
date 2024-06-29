package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean room", Completed: false},
	{ID: "2", Item: "Do laundry", Completed: false},
	{ID: "3", Item: "Learn Go", Completed: false},
}

func getTodo(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

// func addTodo (context *gin.Context) {
// 	var newTodo todo
// }

func main() {
	router := gin.Default()
	router.GET("/todos", getTodo) 
	router.Run("localhost:8000")
}
