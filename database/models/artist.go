package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Artists struct {
	ID          uint32        `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	ArtistName  string        `gorm:"column:artistName;varchar(127);" json:"artistName"`
	IsDeleted   uint16        `gorm:"column:isDeleted;not null;default:0" json:"isDeleted"`
	CreatedByID uint32        `gorm:"column:createdByID;not null" json:"createdById"`
	UpdatedByID sql.NullInt32 `gorm:"column:updatedByID;null" json:"updatedById"`
	CreatedAt   time.Time     `gorm:"column:createdAt;type:datetime;" json:"createdAt"`
	UpdatedAt   time.Time     `gorm:"column:updatedAt;type:datetime;" json:"updatedAt"`
}

/**
 *	This function is responsible for creating an Artist
 */
func CreateArtist(db *gorm.DB, artist *Artists) (err error) {
	err = db.Create(artist).Error
	if err != nil {
		return err
	}
	return nil
}

/**
 *	This function is responsible for getting Artists by name
 */
func GetArtists(db *gorm.DB, artists *[]Artists, artist string) (err error) {
	err = db.Find(artists, "artistName LIKE ?", "%"+artist+"%").Error
	if err != nil {
		return err
	}
	return nil
}
