package main

import (
	"bufio"
	"day6/part1"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	part1Solution := part1.Solution(bufio.NewScanner(file))

	fmt.Printf("part 1: %d", part1Solution)
}
