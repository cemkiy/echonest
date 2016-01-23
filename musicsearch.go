package main

import (
	"encoding/json"
	"fmt"

	"github.com/cemkiy/go-musicsearch/utils"
	"github.com/parnurzeal/gorequest"
)

// Status Struct
type Status struct {
	Version string `json:"id"`
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// Response Struct
type Response struct {
	Status  Status   `json:"status"`
	Artists []Artist `json:"artists"`
	Songs   []Song   `json:"songs"`
}

// Result Struct
type Result struct {
	Response Response `json:"response"`
}

// Artist Struct
type Artist struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Artists list
type Artists []Artist

// ArtistByName artist list
func ArtistByName(name string) ([]Artist, error) {
	request := gorequest.New()
	_, body, _ := request.Get(utils.APIURL + "artist/search?api_key=" + utils.APIKEY + "&format=json&name=" + name).End()

	var result Result
	err := json.Unmarshal([]byte(body), &result)
	if err == nil {
		fmt.Println(err)
		return nil, err
	}
	return result.Response.Artists, nil
}

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
