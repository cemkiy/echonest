package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindArtistByName(t *testing.T) {
	_, err := FindArtistByName("radiohead")
	require.Nil(t, err)
}

func TestFindArtistByGenre(t *testing.T) {
	_, err := FindArtistByGenre("jazz")
	require.Nil(t, err)
}

func TestFindSongByName(t *testing.T) {
	_, err := FindArtistByName("radiohead")
	require.Nil(t, err)
}

func TestFindSongByMood(t *testing.T) {
	_, err := FindSongByMood("happy")
	require.Nil(t, err)
}
