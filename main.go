package main

import (
	"fmt"
	"os"

	"github.com/Rajeevnita1993/json-parser/validator"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: json-parser <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]
	err := validator.ValidateJSONFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println("Valid JSON")
}
