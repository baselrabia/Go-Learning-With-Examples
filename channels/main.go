package main

import (
	"fmt"
	"net/http"
	"os"
)

func main(){



	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://golang.com",
		"http://stackoverflow.com",
		"http://fb.com",

	}



	for _, link := range links {
		go checklink(link)
	}

}

func checklink(link string){

	_,err := http.Get(link)

	if err != nil {
		fmt.Println("error:",link ,"might be down!")
		os.Exit(1)
	}
	
	fmt.Println(link ,"is up!")

}