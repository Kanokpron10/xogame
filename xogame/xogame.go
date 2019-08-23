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
	if game.checkWinHorizontal() {
		return game.turn + " WIN Horizontal"
	}
	if game.checkWinVertical() {
		return game.turn + " WIN Vertical"
	}
	if game.checkWinDiagonal() {
		return game.turn + " WIN Diagonal"
	}
	if game.checkFullBoard() {
		return "Tie"
	}
	return "No winner"
}

func (game Game) checkWinHorizontal() bool {
	for index := range game.Board {
		if game.Board[index][0] == game.turn &&
			game.Board[index][1] == game.turn &&
			game.Board[index][2] == game.turn {
			return true
		}
	}
	return false
}

func (game Game) checkWinVertical() bool {
	for index := range game.Board {
		if game.Board[0][index] == game.turn &&
			game.Board[1][index] == game.turn &&
			game.Board[2][index] == game.turn {
			return true
		}
	}
	return false
}

func (game Game) checkWinDiagonal() bool {
	if game.Board[0][0] == game.turn &&
		game.Board[1][1] == game.turn &&
		game.Board[2][2] == game.turn {
		return true
	}
	if game.Board[0][2] == game.turn &&
		game.Board[1][1] == game.turn &&
		game.Board[2][0] == game.turn {
		return true
	}
	return false
}

func (game Game) checkFullBoard() bool {
	for index := range game.Board {
		for index2 := range game.Board {
			if game.Board[index][index2] == "" {
				return false
			}
		}
	}
	return true
}

func (game *Game) switchTurn() {
	if game.turn == "x" {
		game.turn = "o"
	} else {
		game.turn = "x"
	}
}
