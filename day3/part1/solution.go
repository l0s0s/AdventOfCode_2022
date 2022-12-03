package part1

import (
	"bufio"
	"day3/alphabet"
	"strings"
)

func Solution(scanner *bufio.Scanner) (int, error) {
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		part1 := line[:len(line)/2]
		part2 := line[len(line)/2:]

		for _, char := range part1 {
			if strings.Contains(part2, string(char)) {
				score := alphabet.Num(char)

				total += score
				break
			}
		}
	}

	return total, nil
}
