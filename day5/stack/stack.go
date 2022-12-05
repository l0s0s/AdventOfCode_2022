package stack

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func removeSymbols(str string, symbols ...string) string {
	for _, s := range symbols {
		str = strings.ReplaceAll(str, s, "")
	}

	return str
}

func reverseSlice(input []string) []string {
	inputLen := len(input)
	output := make([]string, inputLen)

	for i, n := range input {
		j := inputLen - i - 1

		output[j] = n
	}

	return output
}

func NewFromFile(scanner *bufio.Scanner) ([]Stack, error) {

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		lines = append(lines, line)
	}

	stacks := make([]Stack, len(removeSymbols(lines[len(lines)-1], " ")))

	for i, stackIndex := range lines[len(lines)-1] {
		if string(stackIndex) == " " {
			continue
		}

		stack := Stack{}

		lines := lines[:len(lines)-1]

		for _, line := range reverseSlice(lines) {
			symbol := string(line[i])
			if symbol == " " {
				continue
			}

			stack.Push(string(line[i]))
		}

		i, err := strconv.Atoi(string(stackIndex))
		if err != nil {
			return nil, fmt.Errorf("failed to parse stack index: %w", err)
		}

		stacks[i-1] = stack
	}

	return stacks, nil
}

type Stack struct {
	stages []string
}

func (s *Stack) Push(stage ...string) {
	s.stages = append(s.stages, stage...)
}

func (s *Stack) Pop() string {
	if len(s.stages) == 0 {
		return ""
	}

	stage := s.stages[len(s.stages)-1]
	s.stages = s.stages[:len(s.stages)-1]

	return stage
}

func (s *Stack) PopMany(n int) []string {
	if len(s.stages) == 0 {
		return []string{}
	}

	stage := s.stages[len(s.stages)-n:]
	s.stages = s.stages[:len(s.stages)-n]

	return stage
}
