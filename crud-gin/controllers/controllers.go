package controllers

import (
	"database/sql"
	
	"fmt"
	"goapi/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
 )


var dbmap = initDb()



func initDb() *gorp.DbMap {
	dsn := "admin:mypassword@tcp(172.20.0.2:3306)/go_api"
	db, err := sql.Open("mysql", dsn)

	checkErr(err, "sql.Open failed")
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	dbmap.AddTableWithName(models.User{}, "user").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
	return dbmap
}


func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}


func Cors() gin.HandlerFunc {
	println("Cors")
	
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

 

func GetUser(c *gin.Context) {

	var user []models.User
	_, err := dbmap.Select(&user, "select * from users")
	if err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{"error": "user not found | " + err.Error()})
	}
}


func GetUserDetail(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := dbmap.SelectOne(&user, "SELECT * FROM user WHERE id=? LIMIT 1", id)
	if err == nil {
		user_id, _ := strconv.ParseInt(id, 0, 64)
		content := &models.User{
			Id:        user_id,
			Username:  user.Username,
			Password:  user.Password,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}





func Login(c *gin.Context) {
	var input,user models.User
	 c.Bind(&input)
	 userstr,errstr := c.GetPostForm("Username")
	fmt.Println(userstr,errstr)

	// fmt.Println("not1", input ,user)
 	err := dbmap.SelectOne(&user, "select * from user where Username=? LIMIT 1", input.Username)
	// fmt.Println("not2",  input ,user)
 
	if err == nil {
		if input.Password == user.Password {
			user_id := user.Id
			content := &models.User{
				Id:        user_id,
				Username:  user.Username,
				Password:  user.Password,
				Firstname: user.Firstname,
				Lastname:  user.Lastname,
			}

			c.JSON(200, gin.H{"status": "success", "data": content})
		} else {
			c.JSON(401, gin.H{"status": "error", "message": "Password is wrong"})
		}	
	} else {
		c.JSON(404, gin.H{"status": "error", "message": "User not found"})
	}
	
	// if err == nil {
	// 	user_id := user.Id
	// 	content := &models.User{
	// 		Id:        user_id,
	// 		Username:  user.Username,
	// 		Password:  user.Password,
	// 		Firstname: user.Firstname,
	// 		Lastname:  user.Lastname,
	// 	}
	// 	c.JSON(200, content)
	// } else {
	// 	c.JSON(404, gin.H{"error": "user not found"})
	// }
}


func PostUser(c *gin.Context) {
	var user models.User
	c.Bind(&user)
			fmt.Println("not1")
			// fmt.Println("not13232")

	if user.Username != "" && user.Password != "" && user.Firstname != "" && user.Lastname != "" {
					fmt.Println("not2")

		insert, err := dbmap.Exec(`INSERT INTO user (Username, Password, Firstname, Lastname) VALUES (?, ?, ?, ?)`, user.Username, user.Password, user.Firstname, user.Lastname)
		fmt.Println("insert " , insert)
		fmt.Println("err " , err)

		if err != nil {
 			c.JSON(404, gin.H{"error": err })
		}  



		if insert != nil {
			
			fmt.Println("not3")

			if err != nil {
				checkErr(err, "Insert failed")
			}
			fmt.Println("4")


			user_id, err := insert.LastInsertId()
			if err == nil {
				content := &models.User{
					Id:        user_id,
					Username:  user.Username,
					Password:  user.Password,
					Firstname: user.Firstname,
					Lastname:  user.Lastname,
				}
				c.JSON(201, content)
			} else {
				checkErr(err, "Insert failed")
			}
		}
		
		


	} else {
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}
}


func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := dbmap.SelectOne(&user, "SELECT * FROM user WHERE id=?", id)
	if err == nil {
		var json models.User
		c.Bind(&json)
		user_id, _ := strconv.ParseInt(id, 0, 64)
		user := models.User{
			Id:        user_id,
			Username:  user.Username,
			Password:  user.Password,
			Firstname: json.Firstname,
			Lastname:  json.Lastname,
		}
		if user.Firstname != "" && user.Lastname != "" {
			_, err = dbmap.Update(&user)
			if err == nil {
				c.JSON(200, user)
			} else {
				checkErr(err, "Updated failed")
			}
		} else {
			c.JSON(400, gin.H{"error": "fields are empty"})
		}
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}