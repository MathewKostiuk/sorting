package sorting

import (
	"html/template"
	"io"
	"sort"
	"strconv"
	"time"
)

// sortkey determines the primary, secondary, etc. sort keys.
// The most recently clicked column head becomes the primary
// sortkey (0).
type sortkey struct {
	name  string
	value int
}

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type TrackTable struct {
	Tracks []*Track
}

type bySortkey []*sortkey

func (a bySortkey) Len() int           { return len(a) }
func (a bySortkey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a bySortkey) Less(i, j int) bool { return a[i].value < a[j].value }

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

var sortkeys = []*sortkey{
	{"Title", 0},
	{"Artist", 1},
	{"Album", 2},
	{"Year", 3},
	{"Length", 4},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

// WriteTracks creates an HTML table for a TrackTable
func WriteTracks(tt TrackTable, wr io.Writer, tmpl *template.Template) {
	err := tmpl.ExecuteTemplate(wr, "index.html", tt)
	if err != nil {
		panic(err)
	}
}

// Makes the most recently clicked table header the primary sortkey
func UpdateSortKeys(n string) {
	for i, key := range sortkeys {
		if key.name == n {
			firstHalf := sortkeys[0:i]
			key.value = 0
			for _, fk := range firstHalf {
				fk.value++
			}
		}
	}
}

// SortTracks sorts the tracks according to the order of the sortkeys
func SortTracks(tracks []*Track) {
	sort.Sort(bySortkey(sortkeys))
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
}

func (s *sortkey) String() string {
	return "sortkey name=" + s.name + " value=" + strconv.Itoa(s.value)
}
