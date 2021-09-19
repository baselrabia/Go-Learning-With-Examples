package main

import (
	"io"
	"os"
	"fmt"
	"net/http"
)

 

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	wr, err := io.Copy(os.Stdout ,resp.Body)
	// func Copy(dst Writer, src Reader) (written int64, err error)
	// Copy copies from src to dst until either EOF is reached on src or an error occurs. 
	// It returns the number of bytes copied and the first error encountered while copying, if any.
	fmt.Println()
	fmt.Println("WR", wr,"error:",err)
}

 