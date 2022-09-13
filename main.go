package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var name string
	var number string

	for {
		fmt.Println("Enter your name:")
		fmt.Scanln(&name)
		if _, err := strconv.ParseFloat(name, 64); err == nil {
			fmt.Printf("%q looks like a number.\n", name)
		} else {
			break

		}
	}

	for {
		fmt.Println("Enter your number:")
		fmt.Scanln(&number)
		if _, err := strconv.ParseFloat(number, 64); err != nil {
			fmt.Printf("%q not a number.\n", number)
		} else {
			break

		}
	}

	var nameUPPER string = strings.ToUpper(name)
	new, _ := strconv.ParseFloat(number, 64)
	fmt.Printf("%v%v%03d", nameUPPER, "-", int(new))

}
