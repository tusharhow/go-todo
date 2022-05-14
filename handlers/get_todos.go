package handlers

import (
	"context"
	"time"
	"github.com/gin-gonic/gin"
	db "github.com/tusharhow/go-todo/db"
	mod "github.com/tusharhow/go-todo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTodos(c *gin.Context) {
	query := bson.D{{}}
	collection := db.MGI.Db.Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, query)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var results []*mod.TodoModel
	for cur.Next(ctx) {
		var elem mod.TodoModel
		err := cur.Decode(&elem)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		results = append(results, &elem)

	}
	c.JSON(200, results)

}
