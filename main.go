package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aspumesh10/songsCollection/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	dir, errDir := os.Executable()
	if errDir != nil {
		log.Println("unable to find env path")
	}

	absPath := filepath.Dir(dir)
	godotenv.Load(absPath + "/.env")

	r := setupRouter()
	_ = r.Run(":3000")
}
