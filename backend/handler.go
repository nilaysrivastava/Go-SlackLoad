package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
)

func UploadFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open file"})
		return
	}
	defer src.Close()

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	params := slack.FileUploadParameters{
		Channels: []string{"#general"}, // Change this to your channel
		File:     src,
		Filename: file.Filename,
	}

	_, err = api.UploadFile(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to upload file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "File uploaded successfully"})
}
