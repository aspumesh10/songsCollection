package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type TracksArtists struct {
	ID          uint32        `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	TrackID     uint32        `gorm:"column:trackId;ForeignKey:trackId;references:id" json:"trackId"`
	ArtistID    uint32        `gorm:"column:artistId;ForeignKey:artistId;references:id" json:"artistId"`
	IsDeleted   uint16        `gorm:"column:isDeleted;not null;default:0" json:"isDeleted"`
	CreatedByID uint32        `gorm:"column:createdByID;not null" json:"createdById"`
	UpdatedByID sql.NullInt32 `gorm:"column:updatedByID;null" json:"updatedById"`
	CreatedAt   time.Time     `gorm:"column:createdAt;type:datetime;" json:"createdAt"`
	UpdatedAt   time.Time     `gorm:"column:updatedAt;type:datetime;" json:"updatedAt"`
}

/**
 *	This function is responsible for creating a Artist Track relation ship
 */
func CreateTrackArtists(db *gorm.DB, tracksArtist *TracksArtists) (err error) {
	err = db.Create(tracksArtist).Error
	if err != nil {
		return err
	}
	return nil
}

/**
 *	This function is responsible for getting GetTrackArtists by ids
 */
func GetTrackArtists(db *gorm.DB, tracksArtist *[]TracksArtists, idList []uint32) (err error) {
	err = db.Find(tracksArtist, "artistId IN ?", idList).Error
	if err != nil {
		return err
	}
	return nil
}
