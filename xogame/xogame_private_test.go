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
