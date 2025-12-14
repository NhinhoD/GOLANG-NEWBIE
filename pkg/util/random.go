package util

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// khởi tạo seed ~ new Random()
func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateShortCode tạo chuỗi ngẫu nhiên độ dài n
func GenerateShortCode(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
