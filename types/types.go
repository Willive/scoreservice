package types

// Score Type used for select statements
type Score struct {
	FileName string `sql:"file_name"`
	Score    int    `sql:"seo_score"`
	Time     string `sql:"time_stamp"`
}

// Tags Valid tags and their weight
var Tags = map[string]int{
	"div":      3,
	"p":        1,
	"h1":       3,
	"h2":       2,
	"html":     5,
	"body":     5,
	"header":   10,
	"footer":   10,
	"font":     -1,
	"center":   -2,
	"big":      -2,
	"strike":   -1,
	"tt":       -2,
	"frameset": -5,
	"frame":    -5,
}
