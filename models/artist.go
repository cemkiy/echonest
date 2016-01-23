package models

import (
	"encoding/json"
	"fmt"

	"github.com/cemkiy/go-musicsearch/utils"
	"github.com/parnurzeal/gorequest"
)

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
