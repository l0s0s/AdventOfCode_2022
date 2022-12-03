package main

import (
	"bufio"
	"day3/part1"
	"day3/part2"
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

	part1Total, err := part1.Solution(scanner)
	if err != nil {
		log.Fatal("failed to solve part 1: %w", err)
	}

	file, err = os.Open("input.txt")
	if err != nil {
		log.Fatal("failed to read input: %w", err)
	}

	defer file.Close()

	scanner = bufio.NewScanner(file)

	part2Total, err := part2.Solution(scanner)
	if err != nil {
		log.Fatal("failed to solve part 2: %w", err)
	}

	fmt.Printf("part 1: %d, part 2: %d", part1Total, part2Total)
}
