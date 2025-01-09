package handler

import (
	"context"
	"github.com/daniellimadev/api-crud-in-go/database"
	"github.com/daniellimadev/api-crud-in-go/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func CreateProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	product.ID = primitive.NewObjectID()
	_, err := database.Collection.InsertOne(context.TODO(), product)
	if err != nil {
		log.Printf("Error inserting product into database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, product)
}
