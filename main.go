package main

import (
	"github.com/BigNutJaa/generateName"
	"log"
)

func main() {

	result, err := generateName.Generate("bignut", "12.2")
	if err != nil {
		log.Println("error:", err)
	}
	log.Println(result)
}
