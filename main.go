package main

import (
	"log"

	"github.com/gin-gonic/gin"

	db "github.com/tusharhow/go-todo/db"
	co "github.com/tusharhow/go-todo/handlers"
)

func main() {
	r := gin.Default()
	db.Connect()

	r.POST("/todo", co.AddTodo)
	r.GET("/todos", co.GetTodos)
	r.PUT("/todo/:id", co.UpdateTodoByID)
	r.DELETE("/todo/:id", co.DeleteTodoByID)

	log.Fatal(r.Run(":8080"))
}
