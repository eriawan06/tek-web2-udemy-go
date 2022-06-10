package database

import (
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase() *gorm.DB {
	// Loan ENV
	err := godotenv.Load()

	// Check if there is error when load ENV
	if err != nil {
		log.Fatal("Failed to Load .env")
	}

	// Assign ENV value to Variable
	//dbUsername := os.Getenv("POSTGRESQL_USERNAME")
	//dbPassword := os.Getenv("POSTGRESQL_PASSWORD")
	//dbHost := os.Getenv("POSTGRESQL_HOST")
	//dbPort := os.Getenv("POSTGRESQL_PORT")
	//dbName := os.Getenv("POSTGRESQL_DBNAME")
	dbUrl := os.Getenv("POSTGRESQL_URL")

	// Database URL ( Data Source Name )
	//dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	// Connect to Database
	// db, err := sql.Open("mysql", dsn)
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})

	// Check if there is error when connect to database
	if err != nil {
		log.Fatal("Failed connect to database: ", err.Error())
	}

	// Double Check Database Connection ( As Recommended by Official Documentation )
	// if errorPing := db.Ping(); err != nil {
	// 	log.Fatal("Failed connect to database: ", errorPing.Error())
	// }

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)

	DB = db
	return db
}
