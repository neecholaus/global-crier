package data

import "time"

type Source struct {
	Publication string
	Name        string
	URL         string
}

type Headline struct {
	Title    string
	Subtitle string
	URL      string
	Date     time.Time
}
