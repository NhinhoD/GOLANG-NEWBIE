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
	//
	r.GET("/:code", handler.RedirectToOriginal)

	// 3. Xem thông tin link
	r.GET("/api/links/:code", handler.GetLinkInfo)

	// 4. Xem danh sách tất cả link
	r.GET("/api/links", handler.GetAllLinks)

	r.Run(":8080")
}
