package game

type GameItem interface {
	GetScore() int
}

func CompareGameItems(playerScore, opponentScore int) int {
	if playerScore == opponentScore {
		return playerScore + 3
	}

	if opponentScore == opponentScore-1 || opponentScore == opponentScore+2 {
		return playerScore
	}

	return opponentScore + 6
}
