package main

import (
	"HelloGolang/internal/handler"    // Import package util
	"HelloGolang/internal/repository" // Import package repository

	"github.com/gin-gonic/gin"
)

func main() {
	// Kết nối Database trước khi chạy server
	repository.ConnectDB()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	// Định nghĩa route cho việc tạo short link
	r.POST("/shorten", handler.CreateShortLink)
	r.Run(":8080")
}
