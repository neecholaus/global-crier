package headlines

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var pullClient = &http.Client{}

// tmp function
func Test() {
	for _, s := range sources {
		start := time.Now()
		headlines, err := GetHeadlinesFromSource(*s)
		if err != nil {
			fmt.Printf("ERR (%s - %s) pull failed (%s)\n", s.Publication, s.Name, err)
			continue
		}
		duration := time.Since(start)
		fmt.Printf("(%s - %s) pull took (%.1f) seconds to pull (%d) headlines\n", s.Publication, s.Name, duration.Seconds(), len(headlines))

		ProcessHeadlines(headlines)
	}
}

func GetHeadlinesFromSource(s Source) ([]*Headline, error) {
	req, err := http.NewRequest("GET", s.URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36")

	res, err := pullClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	headlines, err := s.Parser.ParseBytes(raw)
	if err != nil {
		return nil, err
	}

	return headlines, nil
}
