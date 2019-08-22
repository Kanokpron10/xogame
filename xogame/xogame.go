package xogame

type Game struct {
	Board  [3][3]string
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
		Board:  [3][3]string{},
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

func (game *Game) Play(player Player, row, column int) string {
	game.marking(player, row, column)
	winner := game.checkWin()
	game.switchTurn()
	return winner
}

func (game *Game) marking(player Player, row, column int) {
	game.Board[row][column] = player.symbol
}

func (game Game) checkWin() string {
	if game.Board[0][0] == "x" &&
		game.Board[0][1] == "x" &&
		game.Board[0][2] == "x" {
		return game.turn + " WIN Horizontal first line"
	}
	if game.Board[1][0] == "x" &&
		game.Board[1][1] == "x" &&
		game.Board[1][2] == "x" {
		return game.turn + " WIN Horizontal second line"
	}
	return "WINNER"
}

func (game *Game) switchTurn() {
	if game.turn == "x" {
		game.turn = "o"
	} else {
		game.turn = "x"
	}
}
