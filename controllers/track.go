package controllers

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/aspumesh10/songsCollection/database"
	"github.com/aspumesh10/songsCollection/database/models"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type TracksRepo struct {
	Db *gorm.DB
}

func New() *TracksRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Artists{})
	db.AutoMigrate(&models.TracksArtists{})
	db.AutoMigrate(&models.Track{})
	return &TracksRepo{Db: db}
}

type Isrc struct {
	Isrc string `json:"isrc" form:"isrc"`
}

type AddTrackResponse struct {
	TrackId uint32
}

/**
 *	This function is responsible for getting GetTrackArtists by ids
 */
func (repository *TracksRepo) AddTrack(c *gin.Context) {
	var track models.Track
	var artist models.Artists
	var trackArtists models.TracksArtists
	var response AddTrackResponse
	var isrcBody Isrc
	//using bind as preferring the form encoded
	//use BindJSON if the data is coming in json form
	c.Bind(&isrcBody)
	log.Println(c.Request.Body)
	log.Println(isrcBody)

	track.Title = "Numb"
	track.Popularity = 100
	track.CreatedByID = 1
	track.Isrc = "TEST"
	track.AlbumName = "Meteora"
	track.AlbumReleaseDate, _ = time.Parse("2006-01-02", "2003-09-02")

	log.Println("creating track")
	err := models.CreateTrack(repository.Db, &track)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	log.Println("creating artist")

	artist.ArtistName = "Linkin Park"
	errArtist := models.CreateArtist(repository.Db, &artist)
	if errArtist != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	log.Println("creating artist track relation")

	trackArtists.ArtistID = artist.ID
	trackArtists.TrackID = track.ID
	errTrackArtists := models.CreateTrackArtists(repository.Db, &trackArtists)
	if errTrackArtists != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	log.Println("done all")
	response.TrackId = track.ID
	c.JSON(http.StatusOK, response)
}

/**
 *	This function is responsible for getting GetTrackArtists by ids
 */
func (repository *TracksRepo) GetTracks(c *gin.Context) {
	var tracks []models.Track
	var artists []models.Artists
	var trackArtists []models.TracksArtists

	artist := c.Query("artist")

	err := models.GetArtists(repository.Db, &artists, artist)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Record not found")
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	artistIds := []uint32{}
	for _, val := range artists {
		artistIds = append(artistIds, val.ID)
	}

	if len(artistIds) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Artists not found"})
		return
	}

	errTrackArtists := models.GetTrackArtists(repository.Db, &trackArtists, artistIds)
	if errTrackArtists != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	trackIds := []uint32{}
	for _, val := range trackArtists {
		trackIds = append(trackIds, val.TrackID)
	}

	if len(artistIds) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Tracks not found"})
		return
	}

	errTrack := models.GetTracksByIds(repository.Db, &tracks, trackIds)
	if errTrack != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, tracks)
}

/**
 *	This function is responsible for getting GetTrackArtists by ids
 */
func (repository *TracksRepo) GetTrackSingle(c *gin.Context) {
	isrc := c.Param("isrc")
	var track models.Track
	err := models.GetTrack(repository.Db, &track, isrc)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Record not found")
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Record Not Found"})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, track)
}

// update user
func (repository *TracksRepo) UpdateUser(c *gin.Context) {
	var user models.User
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetUser(repository.Db, &user, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&user)
	err = models.UpdateUser(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

// delete user
func (repository *TracksRepo) DeleteUser(c *gin.Context) {
	var user models.User
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteUser(repository.Db, &user, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
