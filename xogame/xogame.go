package xogame

type Game struct {
	board  [3][3]string
	player []Player
	turn   string
}

type Player struct {
	name   string
	score  int
	symbol string
}

func NewGame(player1, player2 Player, turn string) Game {
	return Game{
		board:  [3][3]string{},
		player: []Player{player1, player2},
		turn:   turn,
	}
}

func NewPlayer(name, symbol string) Player {
	return Player{
		name:   name,
		score:  0,
		symbol: symbol,
	}
}

func (game Game) Play(player Player, row, column int) string {
	game.marking(player, row, column)
	winner := game.checkWin()
	game.switchTurn()
	return winner
}

func (game Game) marking(player Player, row, column int) {

}

func (game Game) checkWin() string {
	return "x WINNER"
}

func (game Game) switchTurn() {

}
