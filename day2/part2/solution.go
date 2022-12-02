package part2

import (
	"bufio"
	"day2/game"
	"strings"
)

func getScore(opponentScore, resultScore int) int {
	if resultScore == 1 && opponentScore == 1 {
		return 3
	}

	if resultScore == 1 {
		return opponentScore - 1
	}

	if resultScore == 2 {
		return opponentScore
	}

	if opponentScore == 3 {
		return 1
	}

	return opponentScore + 1
}

func Solution(scanner *bufio.Scanner) (int, error) {
	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()

		scores := strings.Split(line, " ")

		opponentScore := game.Scores[scores[0]]

		playerScore := getScore(opponentScore, game.Scores[scores[1]])

		score := game.CompareGameItems(playerScore, opponentScore)

		totalScore += score
	}

	return totalScore, nil
}
