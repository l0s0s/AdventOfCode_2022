package main

import (
	"bufio"
	"day5/part1"
	"day5/part2"
	"day5/stack"
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

	s, err := stack.NewFromFile(scanner)
	if err != nil {
		log.Fatal(err)
	}

	part1Result, err := part1.Solution(scanner, s)
	if err != nil {
		log.Fatal(err)
	}

	file, err = os.Open("input.txt")
	if err != nil {
		log.Fatal("failed to read input: %w", err)
	}

	defer file.Close()

	scanner = bufio.NewScanner(file)

	s, err = stack.NewFromFile(scanner)
	if err != nil {
		log.Fatal(err)
	}

	part2Result, err := part2.Solution(scanner, s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("part 1: %s, part 2: %s", part1Result, part2Result)
}
