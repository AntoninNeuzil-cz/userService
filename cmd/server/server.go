package main

import (
	"log"

	"github.com/labstack/echo/v4"

	"userService/internal/db"
	"userService/internal/gen"
	"userService/internal/handler"
)

func main() {
	// Initialize the database
	db.InitDB()

	// Create Echo instance
	e := echo.New()

	// Create handler
	h := handler.NewHandler(db.DB)
	apiHandler := &handler.APIHandler{Handler: h}

	// Register API handlers
	api.RegisterHandlers(e, apiHandler)

	// Start the server
	port := "8080" // Replace with your desired port
	log.Printf("Starting server on port %s", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
