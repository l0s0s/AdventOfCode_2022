package part2

import (
	"bufio"
	"fmt"
	"strconv"
)

func Solution(scanner *bufio.Scanner) (int, error) {
	biggest, _, err := handleLine(scanner, [3]int{}, 0)
	if err != nil {
		return 0, fmt.Errorf("failed to handle line: %w", err)
	}

	sum := 0

	for _, num := range biggest {
		sum += num
	}

	return sum, nil
}

func handleLine(scanner *bufio.Scanner, biggest [3]int, elfCalories int) ([3]int, int, error) {
	if !scanner.Scan() {
		return biggest, elfCalories, nil
	}

	if scanner.Text() == "" {
		if elfCalories > biggest[0] {
			biggest[0], elfCalories = elfCalories, biggest[0]
		}

		if elfCalories > biggest[1] && elfCalories < biggest[0] {
			biggest[1], elfCalories = elfCalories, biggest[1]
		}

		if elfCalories > biggest[2] && elfCalories < biggest[1] {
			biggest[2] = elfCalories
		}

		return handleLine(scanner, biggest, 0)
	}

	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return biggest, 0, fmt.Errorf("failed to parse number: %w", err)
	}

	return handleLine(scanner, biggest, elfCalories+num)
}
