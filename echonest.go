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

// ArtistsByName retrieves artists related to given name
func ArtistsByName(name string) (*[]Artist, error) {
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
