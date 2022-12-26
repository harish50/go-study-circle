package main

import "github.com/gin-gonic/gin"

func main() {
	s := gin.New()
	s.POST("/api/songs", CreateSongHandler)
	s.GET("/api/songs", GetSongsHandler)
	s.PUT("/api/songs/:id", UpdateSong)
	s.Run("localhost:8080")
}
