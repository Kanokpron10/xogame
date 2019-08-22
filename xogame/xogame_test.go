package xogame_test

import (
	"testing"
	"xogame/xogame"
)

func Test_PlayerOne_Win_Horizontal_FirstLine(t *testing.T) {
	expected := "x WIN Horizontal"
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
	expected := "x WIN Horizontal"
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
	expected := "x WIN Horizontal"
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

func Test_PlayerOne_Win_Vertical_FirstLine(t *testing.T) {
	expected := "x WIN Vertical"
	player1 := xogame.NewPlayer("Mo", "x")
	player2 := xogame.NewPlayer("Praw", "o")
	game := xogame.NewGame(player1, player2, "x")

	game.Play(player1, 0, 0)
	game.Play(player2, 2, 0)
	game.Play(player1, 1, 0)
	game.Play(player2, 1, 1)
	actual := game.Play(player1, 2, 0)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_PlayerOne_Win_Vertical_SecondLine(t *testing.T) {
	expected := "x WIN Vertical"
	player1 := xogame.NewPlayer("Mo", "x")
	player2 := xogame.NewPlayer("Praw", "o")
	game := xogame.NewGame(player1, player2, "x")

	game.Play(player1, 0, 1)
	game.Play(player2, 2, 0)
	game.Play(player1, 1, 1)
	game.Play(player2, 1, 0)
	actual := game.Play(player1, 2, 1)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_PlayerOne_Win_Vertical_ThirdLine(t *testing.T) {
	expected := "x WIN Vertical"
	player1 := xogame.NewPlayer("Mo", "x")
	player2 := xogame.NewPlayer("Praw", "o")
	game := xogame.NewGame(player1, player2, "x")

	game.Play(player1, 0, 2)
	game.Play(player2, 2, 0)
	game.Play(player1, 1, 2)
	game.Play(player2, 1, 1)
	actual := game.Play(player1, 2, 2)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_PlayerOne_Win_TopLeftDiagonal(t *testing.T) {
	expected := "x WIN Diagonal"
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
	expected := "x WIN Diagonal"
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
	expected := "o WIN Horizontal"
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

func Test_PlayFullBoard_Should_Be_Tie(t *testing.T) {
	expected := "Tie"
	player1 := xogame.NewPlayer("Mo", "x")
	player2 := xogame.NewPlayer("Praw", "o")
	game := xogame.NewGame(player1, player2, "x")

	game.Play(player1, 0, 0)
	game.Play(player2, 2, 2)
	game.Play(player1, 0, 1)
	game.Play(player2, 0, 2)
	game.Play(player1, 1, 2)
	game.Play(player2, 1, 1)
	game.Play(player1, 2, 0)
	game.Play(player2, 1, 0)
	actual := game.Play(player1, 2, 1)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}
