package database

import (
	"backend/models"
	"backend/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.Users{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Faiiled")
	}

	fmt.Println("Migration Success")
}