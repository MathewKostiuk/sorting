package sorting

import (
	"fmt"
	"os"
	"text/tabwriter"
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
	{"Title", 3},
	{"Artist", 1},
	{"Album", 2},
	{"Year", 0},
	{"Length", 4},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func PrintTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "-----", "-----")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "-----", "-----")
	tw.Flush()
}
