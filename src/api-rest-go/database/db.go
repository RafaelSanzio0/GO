package database

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	server   string = "\\SQLEXPRESS2017"
	database string = "UdemyDB"
	user     string = "sa"
	password string = "root"
	port     int    = 1433
	DB       *gorm.DB
	err      error
)

func ConnectDataBase() {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d; database=%s", server, user, password, port, database)

	DB, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})

	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	fmt.Println("database is reachable:", DB.Name())
}
