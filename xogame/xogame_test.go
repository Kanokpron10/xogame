package xogame

import "testing"

func Test_PlayGame_PlayerOne_Position_1_Shoud_Be_WIN(t *testing.T) {
	expected := "WIN"
	board := xogame.Board{
		Board:          nil,
		PlayerOneScore: 0,
		PlayerTwoScore: 0,
		DrawScore:      0,
	}
	player := xogame.Player{}

	actual := player.PlayGame()

}
