package main

import (
	"adventofcode/day1"
	"adventofcode/day2"
	"adventofcode/day3"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <day>")
		return
	}

	day := os.Args[1]
	switch day {
	case "1":
		day1.Run()
	case "2":
		day2.Run()
	case "3":
		day3.Run()
	default:
		fmt.Printf("Day %s not implemented yet.\n", day)
	}
}
