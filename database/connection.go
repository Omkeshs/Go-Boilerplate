package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	MYSQLDBUSER = "root"
	MYSQLDBPASS = "password"
	MYSQLDBADDR = "car_book"
)

//NewMysqlConnection for client connection
func NewMysqlConnection() *gorm.DB {

	connectionStr := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		MYSQLDBUSER,
		MYSQLDBPASS,
		MYSQLDBADDR,
	)

	//Open Connection for GORM instance
	client, err := gorm.Open("mysql", connectionStr)

	if err != nil {
		panic("Error In Create Client Connection")
	}
	return client

}
