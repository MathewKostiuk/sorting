package sorting

import (
	"fmt"
	"sort"
	"testing"
)

func TestPrintTracks(t *testing.T) {
	PrintTracks(tracks)
	sort.Sort(bySortkey(sortkeys))
	for _, key := range sortkeys {
		fmt.Println(key)
	}

	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		for _, key := range sortkeys {
			switch key.name {
			case "Title":
				if x.Title != y.Title {
					return x.Title < y.Title
				}
			case "Artist":
				if x.Artist != y.Artist {
					return x.Artist < y.Artist
				}
			case "Album":
				if x.Album != y.Album {
					return x.Album < y.Album
				}
			case "Year":
				if x.Year != y.Year {
					return x.Year < y.Year
				}
			case "Length":
				if x.Length != y.Length {
					return x.Length < y.Length
				}
			}
		}
		return false
	}})

	PrintTracks(tracks)
}
