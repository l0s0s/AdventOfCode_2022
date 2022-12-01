package main

import (
	"bufio"
	"day1/part1"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("failed to read input: %w", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	answer, err := part1.Solution(scanner)
	if err != nil {
		log.Fatal("failed to solve part 1: %w", err)
	}

	fmt.Printf("part 1: %d", answer)
}
