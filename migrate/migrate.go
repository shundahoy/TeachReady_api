package main

import (
	"fmt"

	"github.com/shundahoy/techready/db"
	"github.com/shundahoy/techready/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Question{}, &model.Tag{}, &model.Answer{})
}
