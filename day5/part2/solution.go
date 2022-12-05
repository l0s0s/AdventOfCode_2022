package part2

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

		stacks[to-1].Push(stacks[from-1].PopMany(move)...)
	}

	result := ""

	for _, stack := range stacks {
		result += stack.Pop()
	}

	return result, nil
}
