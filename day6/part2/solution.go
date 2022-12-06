package part2

import (
	"bufio"
	"day6/set"
)

func Solution(scanner *bufio.Scanner) int {
	if !scanner.Scan() {
		return 0
	}

	line := scanner.Text()

	for i := range line {
		if len(set.New(line[i:i+14])) == 14 {
			return i + 14
		}
	}

	return 0
}
