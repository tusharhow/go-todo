package handlers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/tusharhow/go-todo/db"
	mod "github.com/tusharhow/go-todo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateTodoByID(c *gin.Context) {
	var todo mod.TodoModel
	todoId := c.Param("id")
	c.BindJSON(&todo)
	collection := db.MGI.Db.Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"_id": todoId}
	update := bson.D{
		{"$set", bson.D{
			{"title", todo.Title},
			{"description", todo.Description},
			{"created_at", time.Now()},
		}},
	}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{"status": "success", "message": "Todo updated successfully", "todo": todo})

}
