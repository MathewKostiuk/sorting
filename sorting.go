package sorting

import (
	"html/template"
	"io"
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

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

var sortkeys = []*sortkey{
	{"Title", 0},
	{"Artist", 0},
	{"Album", 0},
	{"Year", 0},
	{"Length", 0},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

// WriteTracks creates an HTML table for a TrackTable
func WriteTracks(tt TrackTable, wr io.Writer, tmpl *template) {
	err := tmpl.ExecuteTemplate(wr, "index.html", tt)
	if err != nil {
		panic(err)
	}
}
