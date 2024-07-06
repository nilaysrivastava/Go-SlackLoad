package main

import (
	"Go-SlackLoad/handler"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	router := gin.Default()
	router.POST("/upload", handler.UploadFileHandler)
	router.Run(":8080")
}
