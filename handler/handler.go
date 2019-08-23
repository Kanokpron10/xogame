package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlayerRequest struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

func SetPlayerHandler(context *gin.Context) {
	var player1 PlayerRequest
	context.BindJSON(&player1)
	fmt.Println(player1)
	context.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"name":   player1.Name,
		"symbol": player1.Symbol,
	})
}

func SetPlayer2Handler(context *gin.Context) {
	var player2 PlayerRequest
	context.BindJSON(&player2)
	context.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"name":   player2.Name,
		"symbol": player2.Symbol,
	})
}
