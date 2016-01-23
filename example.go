package main

import "fmt"

func main() {
	artists, err := FindArtist("radiohead")
	if err != nil {
		panic(err)
	}
	fmt.Println(*artists)

	songs, err := FindSong("Mother")
	if err != nil {
		panic(err)
	}
	fmt.Println(*songs)
}
