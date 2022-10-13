/*
	To create a connection to database
*/

package config

import (
	"08_rest_api_project/structs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	dest := "root:@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dest), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.Person{})
	return db
}
