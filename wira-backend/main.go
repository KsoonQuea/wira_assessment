package main

import (
	"fmt"
	"log"
	"net/http"

	// "github.com/gin-gonic/gin"

	"wira-backend/connection"
	"wira-backend/routers"
)

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Or specify frontend URL
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}
// 		c.Next()
// 	}
// }

func main() {
	fmt.Println("Server Started")

	// r := gin.Default()
	// r.Use(CORSMiddleware())

	// Initialize database
	connection.Init()

	router := routers.SetupRoutes()

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

