package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit(){
	var err error 
	Conn := "root:@tcp(localhost:3306)/backend?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(Conn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected To Database")
}