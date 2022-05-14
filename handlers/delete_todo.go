package handlers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/tusharhow/go-todo/db"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteTodoByID(c *gin.Context) {
	todoId := c.Param("id")
	collection := db.MGI.Db.Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{"_id": todoId}
	defer cancel()
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{"status": "success", "message": "Todo deleted successfully"})

}
