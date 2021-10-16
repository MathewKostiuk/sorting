package sorting

import (
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
