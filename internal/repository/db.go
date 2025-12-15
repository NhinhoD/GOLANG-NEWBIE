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
	// Load biến môi trường từ file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Lấy chuỗi kết nối
	dsn := os.Getenv("DB_URL")

	dbConfig := postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}

	database, err := gorm.Open(postgres.New(dbConfig), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Tự động tạo bảng (Auto Migrate) dựa trên Struct
	// Bước này giống Code First trong Entity Framework
	database.AutoMigrate(&model.ShortLink{})

	DB = database
	log.Println("Database connection established (Supabase)")
}
