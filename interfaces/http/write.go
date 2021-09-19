package main

import (
	"io"
	"os"
	"fmt"
	"net/http"
)

type logWriter struct {}


 

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}


	lw := logWriter{}

	n,err := io.Copy(lw, resp.Body)
	fmt.Println("Bytes copied:", n)
}

 func (logWriter) Write(bs []byte) (n int, err error) {
	fmt.Println("logWriter Write", string(bs))
	return len(bs), nil
}