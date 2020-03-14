package main

import (
	"github.com/alochym01/ftp-api/src"
)

func main() {
	// fmt.Println("hello world")
	router := src.InitRouter()
	router.Run()
}
