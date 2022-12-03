package part1

import (
	"bufio"
	"strings"
	"unicode"
)

func swapCase(s rune) rune {
	switch {
	case unicode.IsLower(s):
		return unicode.ToUpper(s)
	case unicode.IsUpper(s):
		return unicode.ToLower(s) - 6
	}

	return s

}

func alphaNum(letter rune) int {
	char := int(swapCase(letter))
	char -= 64

	return char
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
