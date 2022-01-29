package database

import (
	"log"
	"os"

	"github.com/MbuguaCaleb/go_rest_api_one/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//creating a database Instance
//i will use a struct---way of representing customtypes
type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

//DB Migrate and Connection Utility
func ConnectDb (){

	db, err :=gorm.Open(sqlite.Open("api.db"),&gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the Database! \n",  err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	//Running Migrations
	db.AutoMigrate(&models.User{},&models.Product{},&models.Order{})

	Database = DbInstance{Db:db}
}

