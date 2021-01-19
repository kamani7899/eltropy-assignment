package models

import (
  "fmt"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
  	dsn := "host=localhost user=babu password=babu DB.name=babu port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println("db",db)
	fmt.Println("err",err)
  	if err != nil {
    	panic("Failed to connect to database!")
  	}

    
  db.AutoMigrate(&Account{})
  db.AutoMigrate(&Customer{})
  db.AutoMigrate(&Employee{})
  db.AutoMigrate(&Transaction{})

  DB = db

}