package db

import (
	"DMP2S/internal/core/domain"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Global DB instance

// InitDatabase initializes the database connection
func InitDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	user := os.Getenv("DB_USER")
	database := os.Getenv("DB_DATABASE")
	password := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s sslmode=require", host, user, password, database, port)
	//dsn := "host=db.vljolfsisxcvbjozejji.supabase.co user=postgres password=SU7pSDLDxlqCqnfd dbname=postgres port=5432 sslmode=require"
	fmt.Println(dsn)
	dialector := postgres.Open(dsn)
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// making database instance Global
	DB = db

	// Auto-migrate models
	err = DB.AutoMigrate(
		&domain.Pipeline{},
		&domain.Stage{},
		&domain.PipelineExecution{},
		&domain.StageExecution{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("Database connected and migrated successfully!")
}
