package main

import (
	"fmt"

	"Student-Management-System/Config"
	"Student-Management-System/Models"
	"Student-Management-System/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	//Config.DB.AutoMigrate(&Models.Student{})

	Config.DB.AutoMigrate(&Models.Student{})

	r := Routes.SetupRouter()
	//running
	r.Run()
}
