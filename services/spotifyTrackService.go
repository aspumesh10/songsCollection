package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type ArtistDetails struct {
	Name string `json:"name"`
}

type AlbumImages struct {
	Url string `json:"url"`
}

type AlbumDetails struct {
	ReleaseDate string        `json:"release_date"`
	TotalTracks uint32        `json:"totalTracks"`
	Images      []AlbumImages `json:"images"`
	Name        string        `json:"name"`
}

type Isrc struct {
	Isrc string `json:"isrc"`
}

type TrackItem struct {
	ID         string          `json:"id"`
	Popularity uint32          `json:"popularity"`
	Name       string          `json:"name"`
	ExternalID Isrc            `json:"external_ids"`
	Album      AlbumDetails    `json:"album"`
	Artists    []ArtistDetails `json:"artists"`
}

type TrackDetailedInfo struct {
	Items []TrackItem `json:"items"`
}

type TrackResponse struct {
	TrackInfo TrackDetailedInfo `json:"tracks"`
}

/**
 *	This function is responsible for getting Track response by isrc and
 *  Also this function gets the track with highest popularity
 *  return specific track from TrackItem object
 */
func FetchTrackDetails(isrc string) (TrackItem, error) {
	var trackItem TrackItem
	var trackResponse TrackResponse

	//fetch access token
	token := GetAccessToken()
	log.Println(token)

	url := "https://api.spotify.com/v1/search?type=track&q=isrc%3A" + isrc
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Println(err)
		return trackItem, err
	}
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return trackItem, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return trackItem, err
	}

	errMarshal := json.Unmarshal(body, &trackResponse)
	if errMarshal != nil {
		return trackItem, err
	}
	maxPopularity := uint32(0)
	maxIndex := 0
	log.Println("count ---- ,", len(trackResponse.TrackInfo.Items))
	if len(trackResponse.TrackInfo.Items) > 0 {
		for index, val := range trackResponse.TrackInfo.Items {
			if val.Popularity > maxPopularity {
				maxPopularity = val.Popularity
				maxIndex = index
			}
			// log.Println(val)
			// log.Println("MAXpopularity ", maxPopularity, " currentPopularity ", val.Popularity, " maxIndex ", maxIndex, " Index ", index)
		}
		return trackResponse.TrackInfo.Items[maxIndex], nil
	} else {
		return trackItem, errors.New("No valid track details found")
	}
}
