package main

import (
	"xogame/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.POST("/xogame/player/1", handler.SetPlayerHandler)
	route.POST("/xogame/player/2", handler.SetPlayer2Handler)
	route.Run(":3001")
}
