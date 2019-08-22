package xogame_test

import (
	"testing"
	"xogame/xogame"
)

func Test_PlayerOne_Win_Horizontal_Line_First(t *testing.T) {
	expected := "x WIN Horizontal first line"
	player1 := xogame.NewPlayer("Mo", "x")
	player2 := xogame.NewPlayer("Praw", "o")
	game := xogame.NewGame(player1, player2, "x")

	game.Play(player1, 0, 0)
	game.Play(player2, 1, 0)
	game.Play(player1, 0, 1)
	game.Play(player2, 1, 1)
	actual := game.Play(player1, 0, 2)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}
