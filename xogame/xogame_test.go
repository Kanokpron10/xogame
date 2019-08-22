package xogame_test

import (
	"testing"
	"xogame/xogame"
)

func Test_PlayerOne_Win_Horizontal_FirstLine_(t *testing.T) {
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

func Test_PlayerOne_Win_Horizontal_SecondLine(t *testing.T) {
	expected := "x WIN Horizontal second line"
	player1 := xogame.NewPlayer("Mo", "x")
	player2 := xogame.NewPlayer("Praw", "o")
	game := xogame.NewGame(player1, player2, "x")

	game.Play(player1, 1, 0)
	game.Play(player2, 0, 0)
	game.Play(player1, 1, 1)
	game.Play(player2, 0, 1)
	actual := game.Play(player1, 1, 2)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_PlayerOne_Win_Horizontal_ThirdLine(t *testing.T) {
	expected := "x WIN Horizontal third line"
	player1 := xogame.NewPlayer("Mo", "x")
	player2 := xogame.NewPlayer("Praw", "o")
	game := xogame.NewGame(player1, player2, "x")

	game.Play(player1, 2, 0)
	game.Play(player2, 0, 0)
	game.Play(player1, 2, 1)
	game.Play(player2, 0, 1)
	actual := game.Play(player1, 2, 2)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_PlayerOne_Win_TopLeftDiagonal(t *testing.T) {
	expected := "x WIN Top left diagonal"
	player1 := xogame.NewPlayer("Mo", "x")
	player2 := xogame.NewPlayer("Praw", "o")
	game := xogame.NewGame(player1, player2, "x")

	game.Play(player1, 0, 0)
	game.Play(player2, 0, 1)
	game.Play(player1, 1, 1)
	game.Play(player2, 0, 2)
	actual := game.Play(player1, 2, 2)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_PlayerOne_Win_DownLeftDiagonal(t *testing.T) {
	expected := "x WIN Down left diagonal"
	player1 := xogame.NewPlayer("Mo", "x")
	player2 := xogame.NewPlayer("Praw", "o")
	game := xogame.NewGame(player1, player2, "x")

	game.Play(player1, 0, 2)
	game.Play(player2, 0, 1)
	game.Play(player1, 1, 1)
	game.Play(player2, 0, 0)
	actual := game.Play(player1, 2, 0)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_PlayerTwo_Win_Horizontal_FirstLine(t *testing.T) {
	expected := "o WIN Horizontal first line"
	player1 := xogame.NewPlayer("Mo", "x")
	player2 := xogame.NewPlayer("Praw", "o")
	game := xogame.NewGame(player1, player2, "x")

	game.Play(player1, 2, 0)
	game.Play(player2, 0, 0)
	game.Play(player1, 1, 1)
	game.Play(player2, 0, 1)
	game.Play(player1, 1, 2)
	actual := game.Play(player2, 0, 2)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}
