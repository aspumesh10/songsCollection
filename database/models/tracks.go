package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Track struct {
	ID               uint32        `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Title            string        `gorm:"column:title;varchar(127);" json:"title"`
	Isrc             string        `gorm:"column:isrc;varchar(127);" json:"isrc"`
	Popularity       uint32        `gorm:"column:popularity;default:0" json:"popularity"`
	AlbumName        string        `gorm:"column:name;varchar(127);" json:"name"`
	AlbumReleaseDate time.Time     `gorm:"column:releaseDate;default:null" json:"releaseDate"`
	IsActive         uint16        `gorm:"column:isActive;not null;default:1" json:"isActive"`
	IsDeleted        uint16        `gorm:"column:isDeleted;not null;default:0" json:"isDeleted"`
	CreatedByID      uint32        `gorm:"column:createdByID;not null" json:"createdById"`
	UpdatedByID      sql.NullInt32 `gorm:"column:updatedByID;null" json:"updatedById"`
	CreatedAt        time.Time     `gorm:"column:createdAt;type:datetime;" json:"createdAt"`
	UpdatedAt        time.Time     `gorm:"column:updatedAt;type:datetime;" json:"updatedAt"`
}

/**
 *	This function is responsible for creating a track
 */
func CreateTrack(db *gorm.DB, Track *Track) (err error) {
	err = db.Create(Track).Error
	if err != nil {
		return err
	}
	return nil
}

/**
 *	This function is responsible for getting single track by isrc
 */
func GetTrack(db *gorm.DB, track *Track, isrc string) (err error) {
	err = db.Where("isrc = ?", isrc).First(track).Error
	if err != nil {
		return err
	}
	return nil
}

/**
 *	This function is responsible for getting GetTrack by ids
 */
func GetTracksByIds(db *gorm.DB, tracks *[]Track, idList []uint32) (err error) {
	err = db.Find(tracks, "id IN ?", idList).Error
	if err != nil {
		return err
	}
	return nil
}
