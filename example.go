package main

import "fmt"

func main() {
	artistsByName, err := FindArtistByName("radiohead")
	if err != nil {
		panic(err)
	}
	fmt.Println(*artistsByName)

	artistsByGenre, err := FindArtistByGenre("jazz")
	if err != nil {
		panic(err)
	}
	fmt.Println(*artistsByGenre)

	songsByName, err := FindSongByName("Mother")
	if err != nil {
		panic(err)
	}
	fmt.Println(*songsByName)

	songsByMood, err := FindSongByMood("happy")
	if err != nil {
		panic(err)
	}
	fmt.Println(*songsByMood)
}
