package xogame

import (
	"testing"
)

func Test_PlayerOne_Marking_Position_0_0(t *testing.T) {
	expected := [3][3]string{{"x", "", ""}, {"", "", ""}, {"", "", ""}}
	player1 := NewPlayer("Mo", "x")
	player2 := NewPlayer("Praw", "o")
	game := NewGame(player1, player2, "x")

	game.marking(player1, 0, 0)
	actual := game.Board

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_PlayerOne_Marking_Position_0_0_And_Position_0_1(t *testing.T) {
	expected := [3][3]string{{"x", "x", ""}, {"", "", ""}, {"", "", ""}}
	player1 := NewPlayer("Mo", "x")
	player2 := NewPlayer("Praw", "o")
	game := NewGame(player1, player2, "x")

	game.marking(player1, 0, 0)
	game.marking(player1, 0, 1)
	actual := game.Board

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CheckWin_Win_Horizontal_FirstLine(t *testing.T) {
	expected := "x WIN Horizontal"
	player1 := NewPlayer("Mo", "x")
	player2 := NewPlayer("Praw", "o")
	game := NewGame(player1, player2, "x")

	game.marking(player1, 0, 0)
	game.marking(player1, 0, 1)
	game.marking(player1, 0, 2)
	actual := game.checkWin()

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CheckWin_Win_Horizontal_SecondLine(t *testing.T) {
	expected := "x WIN Horizontal"
	player1 := NewPlayer("Mo", "x")
	player2 := NewPlayer("Praw", "o")
	game := NewGame(player1, player2, "x")

	game.marking(player1, 1, 0)
	game.marking(player1, 1, 1)
	game.marking(player1, 1, 2)
	actual := game.checkWin()

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_SwitchTurn_Should_Be_O(t *testing.T) {
	expected := "o"
	player1 := NewPlayer("Mo", "x")
	player2 := NewPlayer("Praw", "o")
	game := NewGame(player1, player2, "")

	game.Play(player1, 0, 0)
	game.switchTurn()
	actual := game.turn

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_SwitchTurn_Should_Be_X(t *testing.T) {
	expected := "x"
	player1 := NewPlayer("Mo", "x")
	player2 := NewPlayer("Praw", "o")
	game := NewGame(player1, player2, "")

	game.Play(player1, 0, 0)
	game.Play(player2, 1, 0)
	game.switchTurn()
	actual := game.turn

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_CheckFullBoard_Should_Be_True(t *testing.T) {
	expected := true
	player1 := NewPlayer("Mo", "x")
	player2 := NewPlayer("Praw", "o")
	game := NewGame(player1, player2, "x")

	game.marking(player1, 0, 0)
	game.marking(player2, 2, 2)
	game.marking(player1, 0, 1)
	game.marking(player2, 0, 2)
	game.marking(player1, 1, 2)
	game.marking(player2, 1, 1)
	game.marking(player1, 2, 0)
	game.marking(player2, 1, 0)
	game.marking(player1, 2, 1)
	actual := game.checkFullBoard()

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CheckFullBoard_Should_Be_False(t *testing.T) {
	expected := false
	player1 := NewPlayer("Mo", "x")
	player2 := NewPlayer("Praw", "o")
	game := NewGame(player1, player2, "x")

	game.marking(player1, 0, 0)
	game.marking(player2, 2, 2)
	game.marking(player2, 1, 1)
	actual := game.checkFullBoard()

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}
