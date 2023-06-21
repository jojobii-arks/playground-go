package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Please provide 2 numbers to add.")
	}

	x, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Fatal("First value is not valid.")
	}

	y, err := strconv.Atoi(os.Args[2])

	if err != nil {
		log.Fatal("Second value is not valid.")
	}

	var result int = x + y

	fmt.Println("The result is", result)

}
