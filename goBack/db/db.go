package db

import (
	"fmt"
	"goBack/ent"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	// .env 파일에서 환경 변수를 로드합니다.
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

func ConnectDB() (*ent.Client, error) {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	log.Printf(dsn)

	client, err := ent.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to sqlite: %v", err)
	}

	return client, nil
}
