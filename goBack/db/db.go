package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

func ConnectDB() (*gorm.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	log.Printf(dsn)

	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 로깅 활성화
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	// AutoMigrate를 사용하여 모델의 테이블을 생성하거나 수정할 수 있습니다.
	// db.AutoMigrate(&models.YourModel{})

	return client, nil
}

// package db

// import (
// 	"fmt"
// 	"goBack/ent"
// 	"log"
// 	"os"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/joho/godotenv"
// )

// func init() {
// 	// .env 파일에서 환경 변수를 로드합니다.
// 	if err := godotenv.Load(); err != nil {
// 		log.Printf("No .env file found")
// 	}
// }

// func ConnectDB() (*ent.Client, error) {
// 	dbUser := os.Getenv("DB_USER")
// 	dbPass := os.Getenv("DB_PASS")
// 	dbHost := os.Getenv("DB_HOST")
// 	dbPort := os.Getenv("DB_PORT")
// 	dbName := os.Getenv("DB_NAME")
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
// 	log.Printf(dsn)

// 	client, err := ent.Open("mysql", dsn)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed opening connection to sqlite: %v", err)
// 	}

// 	return client, nil
// }
