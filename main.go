package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage : go run . 'input file' 'output file' ")
		return
	}

	input := os.Args[1]
	output := os.Args[2]

	text, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error,go look again", err)
		return
	}
	content := string(text)
	// fmt.Println(content)

	err = os.WriteFile(output, []byte(content), 0644)
	// fmt.Println("I am building a text-tool")
	if err != nil {
		fmt.Println("Error writing output file")
		return
	}

}
