package headlines

import "time"

type Source struct {
	Publication string
	Name        string
	URL         string
	Parser      parser
}

type parser interface {
	ParseBytes(s []byte) ([]*Headline, error)
}

type Headline struct {
	Title    string
	Subtitle string
	URL      string
	PulledAt time.Time
	Keywords *[]string
}
