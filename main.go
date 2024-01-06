package main

import (
	"net/http"

	"github.com/aspumesh10/songsCollection/controllers"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	tracksRepo := controllers.New()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "System Working Good")
	})

	r.POST("/v1/addTrack", tracksRepo.AddTrack)
	r.GET("/v1/getTracks/:isrc", tracksRepo.GetTrackSingle)
	r.GET("/v1/getTracks", tracksRepo.GetTracks)

	return r
}

func main() {
	r := setupRouter()
	_ = r.Run(":3000")
}
