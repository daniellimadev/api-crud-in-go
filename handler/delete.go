package handler

import (
	"context"
	"github.com/daniellimadev/api-crud-in-go/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error converting ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	filter := bson.M{"_id": objectId}

	result, err := database.Collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println("Error deleting product:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting product"})
		return
	}

	if result.DeletedCount == 0 {
		log.Println("Product not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	log.Println("Product deleted successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
