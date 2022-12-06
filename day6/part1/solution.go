package part1

import (
	"bufio"
	"day6/set"
)

func Solution(scanner *bufio.Scanner) int {
	if !scanner.Scan() {
		return 0
	}

	line := scanner.Text()

	marker := ""

	for i, s := range line {
		marker += string(s)

		if i%4 == 0 && i != 0 {
			if len(set.New(marker)) < len(marker) {
				marker = ""

				continue
			}

			return i - 2
		}
	}

	return 0
}
