package controllers

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/aspumesh10/songsCollection/database"
	"github.com/aspumesh10/songsCollection/database/models"
	"github.com/aspumesh10/songsCollection/services"
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
	var response AddTrackResponse
	var isrcBody Isrc

	//using bind as preferring the form encoded
	//use BindJSON if the data is coming in json form
	c.Bind(&isrcBody)
	log.Println(c.Request.Body)
	log.Println(isrcBody)

	trackItem, errTrackDetails := services.FetchTrackDetails(isrcBody.Isrc)
	log.Println(trackItem)
	if errTrackDetails != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errTrackDetails.Error()})
		return
	}
	track.Title = trackItem.Name
	track.Popularity = trackItem.Popularity
	track.CreatedByID = 1
	track.Isrc = trackItem.ExternalID.Isrc
	track.AlbumName = trackItem.Album.Name
	track.AlbumReleaseDate, _ = time.Parse("2006-01-02", trackItem.Album.ReleaseDate)
	if len(trackItem.Album.Images) > 0 {
		track.Image = trackItem.Album.Images[0].Url
	}
	log.Println("creating track")
	err := models.CreateTrack(repository.Db, &track)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Println("creating artist")

	for _, artistDetails := range trackItem.Artists {
		var artist models.Artists
		var newArtist models.Artists
		var trackArtists models.TracksArtists

		artistId := 0
		err := models.GetArtistByName(repository.Db, &artist, artistDetails.Name)
		createArtist := false
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				createArtist = true
			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
		if createArtist == true {
			newArtist.ArtistName = artistDetails.Name
			errArtist := models.CreateArtist(repository.Db, &newArtist)
			log.Println(errArtist)

			if errArtist != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errArtist.Error()})
				return
			}
			artistId = int(newArtist.ID)
		} else {
			artistId = int(artist.ID)
		}

		log.Println("creating artist track relation")

		trackArtists.ArtistID = uint32(artistId)
		trackArtists.TrackID = track.ID
		errTrackArtists := models.CreateTrackArtists(repository.Db, &trackArtists)
		if errTrackArtists != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errTrackArtists.Error()})
			return
		}
	}

	log.Println("done all")
	response.TrackId = track.ID
	c.JSON(http.StatusOK, gin.H{"data": response})
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
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record Not Found"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	artistIds := []uint32{}
	for _, val := range artists {
		artistIds = append(artistIds, val.ID)
	}

	if len(artistIds) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"data": tracks, "totalCount": len(tracks), "error": "Artists not found"})
		return
	}

	errTrackArtists := models.GetTrackArtists(repository.Db, &trackArtists, artistIds)
	if errTrackArtists != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	trackIds := []uint32{}
	for _, val := range trackArtists {
		trackIds = append(trackIds, val.TrackID)
	}

	if len(artistIds) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"data": tracks, "totalCount": len(tracks), "error": "Tracks not found"})
		return
	}

	errTrack := models.GetTracksByIds(repository.Db, &tracks, trackIds)
	if errTrack != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tracks, "totalCount": len(tracks), "error": ""})
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
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"data": nil, "error": "Record Not Found"})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": track, "error": ""})
}
