package part2

import (
	"bufio"
	"day3/alphabet"
	"strings"
)

func Solution(scanner *bufio.Scanner) (int, error) {
	total := 0

	for {
		group := make([]string, 3)

		for i := 0; i < 3 && scanner.Scan(); i++ {
			text := scanner.Text()
			group[i] = text
		}

		if group[2] == "" {
			break
		}

		for _, char := range group[0] {
			if strings.Contains(group[1], string(char)) && strings.Contains(group[2], string(char)) {
				score := alphabet.Num(char)

				total += score
				break
			}
		}
	}

	return total, nil
}
