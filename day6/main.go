package main

import (
	"bufio"
	"day6/part1"
	"day6/part2"
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

	file, err = os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	part2Solution := part2.Solution(bufio.NewScanner(file))

	fmt.Printf("part 1: %d, part 2: %d", part1Solution, part2Solution)
}
