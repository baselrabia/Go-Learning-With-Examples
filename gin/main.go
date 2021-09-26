package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
 
)

func HomeAPI(c *gin.Context) {
		c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func main(){

fmt.Println("Hello World")
r := gin.Default()
r.GET("/", HomeAPI)

r.Run()


 
}
 