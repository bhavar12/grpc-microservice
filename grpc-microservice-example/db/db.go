package db

import (
	"encoding/json"
	"fmt"
	"grpc-microservice-example/models"
	"io/ioutil"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	fileName = "configs/app.json"
)

var DB *gorm.DB

type DBDetails struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	DBName   string `json:"dbName"`
	DBuser   string `json:"dbUser"`
	Password string `json:"password"`
}

func DatabaseConnection() {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error reading DB configuration from file..%s and error %v", fileName, err)
	}
	var db DBDetails
	json.Unmarshal(bytes, &db)

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		db.Host,
		db.Port,
		db.DBuser,
		db.DBName,
		db.Password,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(models.Movie{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connection successful...")
}
