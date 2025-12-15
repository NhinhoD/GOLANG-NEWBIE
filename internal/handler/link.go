package handler

import (
	"HelloGolang/internal/model"      // Import model
	"HelloGolang/internal/repository" // Import biến DB
	"HelloGolang/pkg/util"            // Import bộ sinh mã vừa tạo
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	// Kiểm tra tính hợp lệ của URL
	_, err := url.ParseRequestURI(req.OriginalURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL không đúng định dạng (phải có http:// hoặc https://)"})
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
	if err := repository.DB.Create(&link).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi database: " + err.Error()})
		return
	}

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080/"
	}

	// Trả về kết quả
	c.JSON(http.StatusOK, gin.H{
		"message":      "Thành công",
		"original_url": link.OriginalURL,
		"short_code":   link.ShortCode,
		"short_url":    baseURL + link.ShortCode,
	})
}

// Chức năng Redirect (Chuyển hướng + Đếm lượt click)
// GET /code
func RedirectToOriginal(c *gin.Context) {
	shortCode := c.Param("code") // Lấy mã từ URL

	var link model.ShortLink

	// Tìm link trong DB
	if err := repository.DB.Where("short_code = ?", shortCode).First(&link).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link không tồn tại"})
		return
	}

	// CẬP NHẬT SỐ LƯỢT CLICK +1
	repository.DB.Model(&link).Update("click_count", gorm.Expr("click_count + 1"))

	// Chuyển hướng người dùng về URL gốc (HTTP 302 Found)
	c.Redirect(http.StatusFound, link.OriginalURL)
}

// Chức năng Xem thông tin chi tiết Link
// GET /api/links/:code
func GetLinkInfo(c *gin.Context) {
	shortCode := c.Param("code")
	var link model.ShortLink

	if err := repository.DB.Where("short_code = ?", shortCode).First(&link).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link không tồn tại"})
		return
	}

	c.JSON(http.StatusOK, link)
}

// 4. Chức năng Liệt kê tất cả các Link (Admin)
// GET /api/links
func GetAllLinks(c *gin.Context) {
	var links []model.ShortLink

	// Lấy toàn bộ danh sách, sắp xếp mới nhất lên đầu
	repository.DB.Order("created_at desc").Find(&links)

	c.JSON(http.StatusOK, links)
}
