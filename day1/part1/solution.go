package part1

import (
	"bufio"
	"fmt"
	"strconv"
)

func Solution(scanner *bufio.Scanner) (int, error) {
	biggest, _, err := handleLine(scanner, 0, 0)
	if err != nil {
		return 0, fmt.Errorf("failed to handle line: %w", err)
	}

	return biggest, nil
}

func handleLine(scanner *bufio.Scanner, biggest, elfCalories int) (int, int, error) {
	if !scanner.Scan() {
		return biggest, elfCalories, nil
	}

	if scanner.Text() == "" {
		if elfCalories > biggest {
			biggest = elfCalories
		}

		return handleLine(scanner, biggest, 0)
	}

	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse number: %w", err)
	}

	return handleLine(scanner, biggest, elfCalories+num)
}
