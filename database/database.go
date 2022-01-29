package database

import (
	"log"
	"os"

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

func ConnectDb (){

	db, err :=gorm.Open(sqlite.Open("api.db"),&gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the Database! \n",  err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	//TODO ADD Migrations

	Database = DbInstance{Db:db}
}

