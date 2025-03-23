package tenniskata

type Player struct {
	Score int
	Name  string
}

func (p *Player) WonPoint() {
	p.Score++
}

func (p *Player) GetScore() string {
	score := ""
	switch p.Score {
	case 0:
		score += "Love"
	case 1:
		score += "Fifteen"
	case 2:
		score += "Thirty"
	case 3:
		score += "Forty"
	}
	return score

}
