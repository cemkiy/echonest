package main

import "fmt"

func main() {
	artists, err := ArtistsByName("radiohead")
	if err != nil {
		panic(err)
	}
	fmt.Println(artists)
}
