package handler

import (
	"HelloGolang/internal/service" // <--- Handler chỉ gọi Service
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

// Request DTO
type CreateLinkRequest struct {
	OriginalURL string `json:"original_url" binding:"required"`
}

// Tạo Link
func CreateShortLink(c *gin.Context) {
	var req CreateLinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Vui lòng gửi original_url"})
		return
	}

	// Validate URL
	if _, err := url.ParseRequestURI(req.OriginalURL); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL không đúng định dạng"})
		return
	}

	// GỌI SERVICE
	link, err := service.CreateShortLink(req.OriginalURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi hệ thống: " + err.Error()})
		return
	}

	// Lấy Base URL từ env
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080/"
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Thành công",
		"original_url": link.OriginalURL,
		"short_code":   link.ShortCode,
		"short_url":    baseURL + link.ShortCode,
	})
}

// Redirect
func RedirectToOriginal(c *gin.Context) {
	shortCode := c.Param("code")

	// GỌI SERVICE
	originalURL, err := service.GetOriginalURL(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link không tồn tại"})
		return
	}

	c.Redirect(http.StatusFound, originalURL)
}

// Xem Info
func GetLinkInfo(c *gin.Context) {
	shortCode := c.Param("code")

	// GỌI SERVICE
	link, err := service.GetLinkInfo(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link không tồn tại"})
		return
	}

	c.JSON(http.StatusOK, link)
}

// Lấy tất cả
func GetAllLinks(c *gin.Context) {
	// GỌI SERVICE
	links, err := service.GetAllLinks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, links)
}
