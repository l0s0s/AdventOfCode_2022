package part1

import (
	"bufio"
	"day5/stack"
	"fmt"
	"regexp"
	"strconv"
)

func Solution(scanner *bufio.Scanner, stacks []stack.Stack) (string, error) {
	for scanner.Scan() {
		line := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`).FindAllString(scanner.Text(), -1)

		move, err := strconv.Atoi(string(line[0]))
		if err != nil {
			return "", fmt.Errorf("failed to parse move value: %w", err)
		}

		from, err := strconv.Atoi(string(line[1]))
		if err != nil {
			return "", fmt.Errorf("failed to parse from value: %w", err)
		}

		to, err := strconv.Atoi(string(line[2]))
		if err != nil {
			return "", fmt.Errorf("failed to parse to value: %w", err)
		}

		for i := 0; i < move; i++ {
			stacks[to-1].Push(stacks[from-1].Pop())
		}
	}

	result := ""

	for _, stack := range stacks {
		result += stack.Pop()
	}

	return result, nil
}
