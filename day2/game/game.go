package game

type GameItem interface {
	GetScore() int
}

func CompareGameItems(playerItem, opponent GameItem) int {
	if playerItem.GetScore() == opponent.GetScore() {
		return playerItem.GetScore() + 3
	}

	if playerItem.GetScore() == opponent.GetScore()-1 || playerItem.GetScore() == opponent.GetScore()+2 {
		return playerItem.GetScore()
	}

	return playerItem.GetScore() + 6
}
