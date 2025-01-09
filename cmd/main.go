package main

import (
	"github.com/daniellimadev/api-crud-in-go/database"
	"github.com/daniellimadev/api-crud-in-go/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	database.SetupDatabase()

	r := gin.Default()

	r.POST("/products", handler.CreateProduct)
	r.GET("/products", handler.ListProducts)
	r.PUT("/products/:id", handler.UpdateProduct)
	r.DELETE("/products/:id", handler.DeleteProduct)

	r.Run(":8080")
	r.SetTrustedProxies(nil)

}
