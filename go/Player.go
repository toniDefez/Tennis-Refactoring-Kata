package tenniskata

type Player struct {
	Name  string
	Score int
}

var scoreNames = []string{"Love", "Fifteen", "Thirty", "Forty"}

func (p *Player) WonPoint() {
	p.Score++
}

func (p *Player) ScoreName() string {
	if p.Score < len(scoreNames) {
		return scoreNames[p.Score]
	}
	return "Forty"
}

func (p *Player) RawScore() int {
	return p.Score
}
