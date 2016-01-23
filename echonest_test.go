package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindArtistByName(t *testing.T) {
	artistsByName, err := FindArtistByName("radiohead", "1")
	require.Nil(t, err)
	require.Equal(t, len(*artistsByName), 1)
}

func TestFindArtistByGenre(t *testing.T) {
	artistsByGenre, err := FindArtistByGenre("jazz", "1")
	require.Nil(t, err)
	require.Equal(t, len(*artistsByGenre), 1)
}

func TestFindSongByName(t *testing.T) {
	songsByName, err := FindArtistByName("radiohead", "1")
	require.Nil(t, err)
	require.Equal(t, len(*songsByName), 1)
}

func TestFindSongByMood(t *testing.T) {
	songsByMood, err := FindSongByMood("happy", "1")
	require.Nil(t, err)
	require.Equal(t, len(*songsByMood), 1)
}
