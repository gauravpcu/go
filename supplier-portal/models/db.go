package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	//dsn := "root:Procurement!2023@tcp(127.0.0.1:3306)/sp_test?charset=utf8mb4&parseTime=True&loc=Local"
	//server=benchmarking.mysql.database.azure.com;database=sp_test;user=sproot;password=Procurement!2023

	dsn := "sproot:Procurement!2023@tcp(benchmarking.mysql.database.azure.com)/sp_test?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		panic("Failed to connect to database!")
	}

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&GoUsers{})

	DB = database
}
