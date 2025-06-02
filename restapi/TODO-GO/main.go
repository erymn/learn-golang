package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{Id: "1", Title: "Clean the room", Completed: false},
	{Id: "2", Title: "Read book", Completed: false},
	{Id: "3", Title: "Write code", Completed: false},
	{Id: "4", Title: "Create Video", Completed: false},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func addTodo(c *gin.Context) {
	var newTodo todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.Id == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func getTodo(c *gin.Context) {
	id := c.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(c *gin.Context) {
	id := c.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed
	c.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.POST("/todos", addTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.Run("localhost:9090")
}
