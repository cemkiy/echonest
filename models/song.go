package models

import (
	"encoding/json"
	"fmt"

	"github.com/cemkiy/go-musicsearch/utils"
	"github.com/parnurzeal/gorequest"
)

// Song Struct
type Song struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	ArtistID   string `json:"artist_id"`
	ArtistName string `json:"artist_name"`
}

// Songs list
type Songs []Song

// SongByName song list
func SongByName(name string) ([]Song, error) {
	request := gorequest.New()
	_, body, _ := request.Get(utils.APIURL + "song/search?api_key=" + utils.APIKEY + "&format=json&title=" + name).End()

	var result Result
	err := json.Unmarshal([]byte(body), &result)
	if err == nil {
		fmt.Println(err)
		return nil, err
	}
	return result.Response.Songs, nil
}
