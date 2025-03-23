package tenniskata

type tennisGame1 struct {
	Player1 Player
	Player2 Player
	StatusGame
}

type StatusGame int

const (
	InProcess StatusGame = iota
	Finished
	Tie
	Unknown
)

func TennisGame1(player1Name string, player2Name string) TennisGame {
	game := &tennisGame1{
		Player1: Player{
			Name:  player1Name,
			Score: 0,
		},
		Player2: Player{
			Name:  player2Name,
			Score: 0,
		},
		StatusGame: Unknown,
	}

	return game
}

func (game *tennisGame1) WonPoint(playerName string) {
	if playerName == game.Player1.Name {
		game.player1WonPoint()
	} else {
		game.player2WonPoint()
	}
}

func (game *tennisGame1) playerSameScore() bool {
	return game.Player1.Score == game.Player2.Score
}

func (game *tennisGame1) player1WonPoint() {
	game.Player1.WonPoint()
}

func (game *tennisGame1) player2WonPoint() {
	game.Player2.WonPoint()
}

func (game *tennisGame1) player1Score() string {
	return game.Player1.GetScore()
}

func (game *tennisGame1) player2Score() string {
	return game.Player2.GetScore()
}

func (game *tennisGame1) getInProcesScore() string {
	return game.player1Score() + "-" + game.player2Score()
}

func (game *tennisGame1) getTieScore() string {
	score := ""
	switch game.Player1.Score {
	case 0:
		score = "Love-All"
	case 1:
		score = "Fifteen-All"
	case 2:
		score = "Thirty-All"
	default:
		score = "Deuce"
	}

	return score

}

func (game *tennisGame1) getOverScore() string {
	score1 := game.Player1.Score
	score2 := game.Player2.Score
	score := ""
	minusResult := score1 - score2
	if minusResult == 1 {
		score = "Advantage player1"
	} else if minusResult == -1 {
		score = "Advantage player2"
	} else if minusResult >= 2 {
		score = "Win for player1"
	} else {
		score = "Win for player2"
	}
	return score

}

func (game *tennisGame1) playerHighScore() bool {
	return game.Player1.Score >= 4 || game.Player2.Score >= 4
}

func (game *tennisGame1) getStatus() StatusGame {
	if !game.playerSameScore() {

		if !game.playerHighScore() {
			return Finished
		}
		return InProcess
	}
	return Tie

}

func (game *tennisGame1) GetScore() string {

	statusGame := game.getStatus()
	switch statusGame {
	case InProcess:
		return game.getInProcesScore()
	case Finished:
		return game.getOverScore()
	case Tie:
		return game.getTieScore()
	default:
		return ""
	}
}
