package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

  db :=  DbConn()
 
  fmt.Println(*db)

 
	user := User{
		Name:     "jinzhu",
		Age:      18,
		Email:	"jhjhjk@gmail.com",
	}

  db.AutoMigrate(&User{})



	fmt.Println(user)
	result := db.Create(&user) // pass pointer of data to Create
	fmt.Println(result)




}

func DbConn() *gorm.DB {
  // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
  dsn := "admin:mypassword@tcp(127.0.0.1:3308)/go_api?charset=utf8mb4&parseTime=True&loc=Local"
  
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    log.Fatal(err)
	  panic(err)
  }
 

	//db.SingularTable(true)
	return db
}



type User struct {
  gorm.Model
  ID           uint
  Name         string
  Email        string
  Age          uint8
  Birthday     *time.Time
  MemberNumber sql.NullString
  Activat sql.NullTime
  Newreqord   sql.NullString
  CreatedAt    time.Time
  UpdatedAt    time.Time
}