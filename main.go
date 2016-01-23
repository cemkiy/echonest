package main

import "fmt"

func main() {
	artists, _ := ArtistByName("imagine")
	fmt.Println(artists)
	// songs, _ := models.SongByName("Love")
	// fmt.Println(songs)
}
