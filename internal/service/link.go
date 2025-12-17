package service

import (
	"HelloGolang/internal/model"
	"HelloGolang/internal/repository"
	"HelloGolang/pkg/util"

	"gorm.io/gorm"
)

// Tạo Link rút gọn
func CreateShortLink(originalURL string) (*model.ShortLink, error) {
	// Sinh mã ngẫu nhiên
	shortCode := util.GenerateShortCode(6)

	// Tạo object
	link := model.ShortLink{
		OriginalURL: originalURL,
		ShortCode:   shortCode,
	}

	// Gọi Repository để lưu vào DB
	if err := repository.DB.Create(&link).Error; err != nil {
		return nil, err
	}

	return &link, nil
}

// Lấy Link gốc & Tăng view
func GetOriginalURL(shortCode string) (string, error) {
	var link model.ShortLink

	// Tìm trong DB
	if err := repository.DB.Where("short_code = ?", shortCode).First(&link).Error; err != nil {
		return "", err
	}

	// Tăng lượt click
	repository.DB.Model(&link).Update("click_count", gorm.Expr("click_count + 1"))

	return link.OriginalURL, nil
}

// Lấy thông tin chi tiết
func GetLinkInfo(shortCode string) (*model.ShortLink, error) {
	var link model.ShortLink
	if err := repository.DB.Where("short_code = ?", shortCode).First(&link).Error; err != nil {
		return nil, err
	}
	return &link, nil
}

// Lấy tất cả
func GetAllLinks() ([]model.ShortLink, error) {
	var links []model.ShortLink
	repository.DB.Order("created_at desc").Find(&links)
	return links, nil
}
