package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "usage: go run . <input> <output>")
		os.Exit(1)
	}
	input, output := os.Args[1], os.Args[2]
	data, err := os.ReadFile(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Readding :", err)
	}
	text := string(data)
	result := transform(text)
	err = os.WriteFile(output, []byte(result), 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Writing :", err)
	}
}
