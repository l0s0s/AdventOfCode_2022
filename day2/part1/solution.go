package part1

import (
	"bufio"
	"day2/game"
	"day2/gameitem"
	"strings"
)

var (
	Scores = map[string]int{
		"A": 1,
		"X": 1,
		"B": 2,
		"Y": 2,
		"C": 3,
		"Z": 3,
	}
)

func Solution(scanner *bufio.Scanner) (int, error) {
	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()

		scores := strings.Split(line, " ")

		score := game.CompareGameItems(gameitem.New(Scores[scores[1]]), gameitem.New(Scores[scores[0]]))
		totalScore += score
	}

	return totalScore, nil
}
