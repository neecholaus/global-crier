package headlines

import (
	"fmt"
	"time"
)

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
	Keywords []string
	Source   *Source
}

func PullAndProcessAllSources() {
	for _, s := range Sources {
		start := time.Now()

		tmpHeadlines, err := GetHeadlinesFromSource(s)
		if err != nil {
			fmt.Printf("ERR (%s - %s) pull failed (%s)\n", s.Publication, s.Name, err)
			continue
		}

		duration := time.Since(start)
		fmt.Printf("(%s - %s) pull took (%.1f) seconds to pull (%d) headlines\n", s.Publication, s.Name, duration.Seconds(), len(tmpHeadlines))

		ProcessNewHeadlines(tmpHeadlines)
	}
}
