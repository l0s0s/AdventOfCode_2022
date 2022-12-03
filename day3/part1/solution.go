package part1

import (
	"bufio"
	"strings"
	"unicode"
)

func alphaNum(s rune) int {
	switch {
	case unicode.IsLower(s):
		return int(unicode.ToUpper(s)) - 64
	case unicode.IsUpper(s):
		return int(unicode.ToLower(s) - 70)
	}

	return 0

}

func Solution(scanner *bufio.Scanner) (int, error) {
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		part1 := line[:len(line)/2]
		part2 := line[len(line)/2:]

		for _, char := range part1 {
			if strings.Contains(part2, string(char)) {
				score := alphaNum(char)

				total += score
				break
			}
		}
	}

	return total, nil
}
