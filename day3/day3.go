package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(file)

}
