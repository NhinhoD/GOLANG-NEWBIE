package handler

import (
	"HelloGolang/internal/model"      // Import model
	"HelloGolang/internal/repository" // Import biến DB
	"HelloGolang/pkg/util"            // Import bộ sinh mã
	"net/http"

	"github.com/gin-gonic/gin"
)

// Request Body (DTO)
type CreateLinkRequest struct {
	OriginalURL string `json:"original_url" binding:"required"`
}

func CreateShortLink(c *gin.Context) {
	var req CreateLinkRequest

	// Hứng dữ liệu JSON gửi lên
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Vui lòng gửi original_url"})
		return
	}

	// Sinh mã ngẫu nhiên 6 ký tự
	shortCode := util.GenerateShortCode(6)

	// Tạo object để lưu (Mapping)
	link := model.ShortLink{
		OriginalURL: req.OriginalURL,
		ShortCode:   shortCode,
	}

	// Lưu vào Database thông qua GORM
	// repository.DB là biến toàn cục mình đã kết nối ở bài trước
	if err := repository.DB.Create(&link).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi database: " + err.Error()})
		return
	}

	// Trả về kết quả
	c.JSON(http.StatusOK, gin.H{
		"message":      "Thành công",
		"original_url": link.OriginalURL,
		"short_code":   link.ShortCode,
		"short_url":    "http://localhost:8080/" + link.ShortCode,
	})
}
