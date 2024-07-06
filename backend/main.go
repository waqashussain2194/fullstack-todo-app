package main

import (
	"net/http"

	api "github.com/el10savio/TODO-Fullstack-App-Go-Gin-Postgres-React/backend/api"
	"github.com/el10savio/TODO-Fullstack-App-Go-Gin-Postgres-React/backend/middleware"
	"github.com/gin-gonic/gin"
)

// Function called for index
func indexView(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, gin.H{"message": "TODO APP"})
}

// Setup Gin Routes
func SetupRoutes() *gin.Engine {
	// Use Gin as router
	router := gin.Default()

	// Use CORS middleware
	router.Use(middleware.CORSMiddleware())

	// Set route for index
	router.GET("/", indexView)

	// Set routes for API
	router.GET("/items", api.TodoItems)
	router.POST("/item/create", api.CreateTodoItem)
	router.PATCH("/item/update/:id/:done", api.UpdateTodoItem)
	router.DELETE("/item/delete/:id", api.DeleteTodoItem)
	router.GET("/items/filter/:done", api.FilterTodoItems)

	// Set up Gin Server
	return router
}

// Main function
func main() {
	api.SetupPostgres()
	router := SetupRoutes()
	router.Run(":8081")
}
