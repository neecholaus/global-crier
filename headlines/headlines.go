package headlines

import "time"

type Source struct {
	Publication string
	Name        string
	URL         string
	Parser      parser
}

type parser interface {
	ParseBytes(s []byte) ([]*TmpHeadline, error)
}

type TmpHeadline struct {
	Title    string
	Subtitle string
	URL      string
	PulledAt time.Time
	Keywords *[]string
}
