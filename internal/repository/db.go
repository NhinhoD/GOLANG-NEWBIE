package repository

import (
	"log"
	"os"

	"HelloGolang/internal/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// 1. Load biến môi trường từ file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2. Lấy chuỗi kết nối
	dsn := os.Getenv("DB_URL")

	// 3. Mở kết nối đến Supabase (Postgres)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 4. Tự động tạo bảng (Auto Migrate) dựa trên Struct
	// Bước này giống Code First trong Entity Framework
	database.AutoMigrate(&model.ShortLink{})

	DB = database
	log.Println("Database connection established (Supabase)")
}
