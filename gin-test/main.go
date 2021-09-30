package main

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

 
func main() { 
  dsn := "admin:mypassword@tcp(127.0.0.1:3308)/go_api?charset=utf8mb4&parseTime=True&loc=Local"

  db, err := sql.Open("mysql", dsn)
  if err != nil {
      fmt.Print(err.Error())
  }
  defer db.Close()
  // make sure connection is available
  err = db.Ping()
  if err != nil {
      fmt.Print(err.Error())
  }
    fmt.Println("connected success")

  type Person struct {
      Id         int
      First_Name string
      Last_Name  string
  }
        
  r := gin.Default()
  

// POST new person details
    r.POST("/person", func(c *gin.Context) {
        var buffer bytes.Buffer
        first_name := c.PostForm("first_name")
        last_name := c.PostForm("last_name")
        fmt.Println(first_name)
        fmt.Println(last_name)


         query := `CREATE TABLE IF NOT EXISTS person(
           person_id int primary key auto_increment, 
           first_name text,
           last_name text,
           created_at datetime default CURRENT_TIMESTAMP,
           updated_at datetime default CURRENT_TIMESTAMP)`
        ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancelfunc()
        _, err := db.ExecContext(ctx, query)
        if err != nil {
             fmt.Printf("Error %s when creating person table", err)
         }
        

        stmt, err := db.Prepare("insert into person (first_name, last_name) values(?,?);")
        if err != nil {
            fmt.Print(err.Error())
        }
        _, err = stmt.Exec(first_name, last_name)

        if err != nil {
            fmt.Print(err.Error())
        }

        // Fastest way to append strings
        buffer.WriteString(first_name)
        buffer.WriteString(" ")
        buffer.WriteString(last_name)
        defer stmt.Close()
        name := buffer.String()
        c.JSON(http.StatusOK, gin.H{
            "message": fmt.Sprintf(" %s successfully created", name),
        })
    })


    	r.Run() // listen and serve on 0.0.0.0:8080



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