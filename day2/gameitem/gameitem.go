package gameitem

type GameItem struct {
	score int
}

func New(score int) *GameItem {
	return &GameItem{score: score}
}

func (g *GameItem) GetScore() int {
	return g.score
}
