package main

import (
	"encoding/json"

	"github.com/cemkiy/echonest/utils"
	"github.com/parnurzeal/gorequest"
)

// Artist json struct
type Artist struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ArtistsResponse envelope struct
type ArtistsResponse struct {
	Response struct {
		Artists []Artist `json:"artists"`
	} `json:"response"`
}

// FindArtistByName retrieves artists related to given name
func FindArtistByName(name string) (*[]Artist, error) {
	response := ArtistsResponse{}
	request := gorequest.New()
	_, bodyBytes, _ := request.Get(utils.APIURL + utils.APIARTISTSEARCH).
		Query("api_key=" + utils.APIKEY).
		Query("name=" + name).
		Query("format=json").
		EndBytes()

	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return nil, err
	}

	return &response.Response.Artists, nil
}

// FindArtistByGenre retrieves artists related to given genre
func FindArtistByGenre(genre string) (*[]Artist, error) {
	response := ArtistsResponse{}
	request := gorequest.New()
	_, bodyBytes, _ := request.Get(utils.APIURL + utils.APIARTISTSEARCH).
		Query("api_key=" + utils.APIKEY).
		Query("genre=" + genre).
		Query("format=json").
		EndBytes()

	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return nil, err
	}

	return &response.Response.Artists, nil
}

// Song json struct
type Song struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	ArtistID   string `json:"artist_id"`
	ArtistName string `json:"artist_name"`
}

// SongsResponse envelope struct
type SongsResponse struct {
	Response struct {
		Songs []Song `json:"songs"`
	} `json:"response"`
}

// FindSongByName retrieves songs related to given name
func FindSongByName(title string) (*[]Song, error) {
	response := SongsResponse{}
	request := gorequest.New()
	_, bodyBytes, _ := request.Get(utils.APIURL + utils.APISONGSEARCH).
		Query("api_key=" + utils.APIKEY).
		Query("title=" + title).
		Query("format=json").
		EndBytes()

	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return nil, err
	}

	return &response.Response.Songs, nil
}

// FindSongByMood retrieves songs related to given name
func FindSongByMood(mood string) (*[]Song, error) {
	response := SongsResponse{}
	request := gorequest.New()
	_, bodyBytes, _ := request.Get(utils.APIURL + utils.APISONGSEARCH).
		Query("api_key=" + utils.APIKEY).
		Query("mood=" + mood).
		Query("format=json").
		EndBytes()

	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return nil, err
	}

	return &response.Response.Songs, nil
}
