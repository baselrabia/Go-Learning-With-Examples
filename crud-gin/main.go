package main

import (
 
 	"goapi/mappings"
 

	_ "github.com/go-sql-driver/mysql"
)


func main() {
	
	// dsn := "admin:mypassword@tcp(172.20.0.2:3306)/go_api"

	// db, err := sql.Open("mysql", dsn)
	// fmt.Println("db Connection")
	// fmt.Println(db)
	// fmt.Println(err)

	// if err != nil {
	// 	fmt.Println("err : " , err)
	// }
 	// dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	// err = dbmap.CreateTablesIfNotExists()
 	
	// if err != nil {
	// 	fmt.Println("err : " , err)
	// }

	// fmt.Println(dbmap)

	mappings.CreateUrlMappings()
	// Listen and server on 0.0.0.0:8080
	mappings.Router.Run(":8080")
}