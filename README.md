# CRUD API in Go

This is a CRUD API (Create, Read, Update, Delete) developed in Go, using the Gin framework and MongoDB as a database. The API allows the manipulation of a product catalog, offering endpoints to add, list, update and remove products.

## Technologies

- **Golang:** Programming language used to develop the API.
- **Gin:** Web framework for managing routes and middlewares.
- **MongoDB:** NoSQL database for storing product data.
- **Postman:** Tool for testing and validating API endpoints.

## Features

- **Product Creation:** Add new products to the database.
- **Product Listing:** List all registered products.
- **Product Update:** Update the information of an existing product.
- **Product Removal:** Delete a product from the database.

## Project Structure

- `handler/`: Contains the route handlers for CRUD operations.
- `create.go`: Handles the creation of new products.
- `delete.go`: Handles the deletion of products.
- `list.go`: Handles the listing of products.
- `update.go`: Handles the updating of products.
- `database/`: Configures the connection to MongoDB.
- `database.go`: Manages the connection to the database.
- `models/`: Defines the structure of the `Product` model.
- `product.go`: Structure of the product model.
- `main.go`: Main file that starts the server and configures the routes.
- `go.mod`: Go module file.
- `go.sum`: Summary of Go dependencies.

## Configuration

1. **Clone the Repository:**

```bash
git clone https://github.com/daniellimadev/api-crud-in-go
cd api-crud-in-go
```
Install the Dependencies:

Copy code:
```bash
go mod tidy
```

Configure the Database:

Replace the MongoDB connection URI in database/mongo.go with your own MongoDB URI.

Run the Server:

Copy code:
```bash
go run main.go
```

The server will be available at http://localhost:8080.

Testing the API
You can use Postman to test the following endpoints:

POST /products: Adds a new product.
```bash
{
    "name" : "product1",
    "price" : 99.56,
    "quantity" : 10
}
```

GET /products: Lists all products.

PUT /products/
: Updates an existing product.
DELETE /products/
: Removes a product by its ID.

<h3>Author</h3>

<a href="https://www.linkedin.com/in/danielpereiralima/">
 <img style="border-radius: 50%;" src="https://avatars.githubusercontent.com/u/96916005?v=4" width="100px;" alt=""/>

Made by Daniel Pereira Lima 👋🏽 Contact!

[![Linkedin Badge](https://img.shields.io/badge/-Daniel-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/danielpereiralima/)](https://www.linkedin.com/in/danielpereiralima/)