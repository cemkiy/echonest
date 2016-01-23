package main

import "fmt"

// echonest returns  max 100 result
func main() {
	artistsByName, err := FindArtistByName("radiohead", "4") // name, results
	if err != nil {
		panic(err)
	}
	fmt.Println(*artistsByName)

	artistsByGenre, err := FindArtistByGenre("jazz", "4") // genre, results
	if err != nil {
		panic(err)
	}
	fmt.Println(*artistsByGenre)

	songsByName, err := FindSongByName("Mother", "4") // name, results
	if err != nil {
		panic(err)
	}
	fmt.Println(*songsByName)

	songsByMood, err := FindSongByMood("happy", "4") // mood, results
	if err != nil {
		panic(err)
	}
	fmt.Println(*songsByMood)
}
