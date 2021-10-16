package sorting

import (
	"fmt"
	"html/template"
	"os"
	"testing"
)

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func TestWriteTracks(t *testing.T) {
	var tt TrackTable
	tt.Tracks = tracks
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	WriteTracks(tt, os.Stdout, tmpl)
}

func TestSortTracks(t *testing.T) {
	var tt TrackTable
	tt.Tracks = tracks
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}

	WriteTracks(tt, os.Stdout, tmpl)
}

func TestSortKeys(t *testing.T) {
	for _, key := range sortkeys {
		fmt.Println(key)
	}
	UpdateSortKeys("Album")
	for _, key := range sortkeys {
		fmt.Println(key)
	}
	UpdateSortKeys("Length")
	for _, key := range sortkeys {
		fmt.Println(key)
	}
	SortTracks(tracks)
	for _, key := range sortkeys {
		fmt.Println(key)
	}
}
