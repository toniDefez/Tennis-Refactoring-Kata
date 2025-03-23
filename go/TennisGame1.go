package tenniskata

import "fmt"

type tennisGame1 struct {
	player1 *Player
	player2 *Player
}

func TennisGame1(player1Name, player2Name string) *tennisGame1 {
	return &tennisGame1{
		player1: &Player{Name: player1Name},
		player2: &Player{Name: player2Name},
	}
}

func (game *tennisGame1) WonPoint(playerName string) {
	if playerName == game.player1.Name {
		game.player1.WonPoint()
	} else if playerName == game.player2.Name {
		game.player2.WonPoint()
	}
}

func (game *tennisGame1) GetScore() string {
	p1 := game.player1
	p2 := game.player2

	if p1.Score == p2.Score {
		return getEqualScore(p1.Score)
	}

	if p1.Score >= 4 || p2.Score >= 4 {
		return getAdvantageOrWin(p1, p2)
	}

	return fmt.Sprintf("%s-%s", p1.ScoreName(), p2.ScoreName())
}

func getEqualScore(score int) string {
	switch score {
	case 0:
		return "Love-All"
	case 1:
		return "Fifteen-All"
	case 2:
		return "Thirty-All"
	default:
		return "Deuce"
	}
}

func getAdvantageOrWin(p1, p2 *Player) string {
	diff := p1.Score - p2.Score
	switch {
	case diff == 1:
		return "Advantage " + p1.Name
	case diff == -1:
		return "Advantage " + p2.Name
	case diff >= 2:
		return "Win for " + p1.Name
	default:
		return "Win for " + p2.Name
	}
}
