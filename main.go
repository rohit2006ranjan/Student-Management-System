package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/rohit2006ranjan/Student-Management-System/Config"
	"github.com/rohit2006ranjan/Student-Management-System/Models"
	"github.com/rohit2006ranjan/Student-Management-System/Routes"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Student{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
