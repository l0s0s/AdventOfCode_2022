package part1

import (
	"bufio"
	"day2/game"
	"strings"
)

func Solution(scanner *bufio.Scanner) (int, error) {
	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()

		scores := strings.Split(line, " ")

		score := game.CompareGameItems(game.Scores[scores[1]], game.Scores[scores[0]])
		totalScore += score
	}

	return totalScore, nil
}
