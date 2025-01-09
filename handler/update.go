package handler

import (
	"context"
	"github.com/daniellimadev/api-crud-in-go/database"
	"github.com/daniellimadev/api-crud-in-go/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	// Check if ID has 24 hexadecimal characters
	if len(id) != 24 {
		log.Println("Invalid ID: Incorrect length")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error converting ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"name":     product.Name,
			"price":    product.Price,
			"quantity": product.Quantity,
		},
	}

	result, err := database.Collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if err != nil {
		log.Println("Error updating product:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating product"})
		return
	}

	if result.MatchedCount == 0 {
		log.Println("Product not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	log.Println("Product updated successfully")
	product.ID = objID
	c.JSON(http.StatusOK, product)
}
