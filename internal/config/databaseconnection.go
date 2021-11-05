package config

import (
	"authe/internal/database"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var datastore *gorm.DB

func IniiateDatabaseConnection()  {
	
	
}

func Mapping()  {
	CreateDatabaseConnection().AutoMigrate(&database.Account{},&database.User{})
	
}

func CreateDatabaseConnection() *gorm.DB {
	if datastore != nil{
		return datastore
	}

	var err error
	var cfg Config = GetConfig()
	databaseParameters := fmt.Sprintf("host=%s port=%s user= %s dbname=%s password=%s",cfg.DatabaseHost,cfg.DatabasePort,cfg.DatabaseUser,cfg.DatabaseName,cfg.DatabasePassword)


	datastore,err := gorm.Open(postgres.Open(databaseParameters),&gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v\n",err))
	}

	return datastore
	
}
