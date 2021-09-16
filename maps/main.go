package main

import "fmt"
 

func main() {

 colors := map[string]string{
	 "red": "#999",
	 "blue": "#333",
 }
//  delete(colors,"red")
//  fmt.Println(colors)
 printMap(colors)
}

 func printMap(c map[string]string) {
	 for color,hex := range c {
		 fmt.Println("hex code for",color,"is",hex)
	 }
 }