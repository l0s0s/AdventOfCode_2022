package game

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

type GameItem interface {
	GetScore() int
}

func CompareGameItems(playerScore, opponentScore int) int {
	if playerScore == opponentScore {
		return playerScore + 3
	}

	if playerScore == opponentScore-1 || playerScore == opponentScore+2 {
		return playerScore
	}

	return playerScore + 6
}
