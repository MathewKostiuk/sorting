package sorting

import (
	"fmt"
	"html/template"
	"os"
	"testing"
)

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
	SortTracks()
	for _, key := range sortkeys {
		fmt.Println(key)
	}
}
