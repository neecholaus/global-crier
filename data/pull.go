package data

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func PullAll() {
	for _, s := range sources {
		PullSource(*s)
	}
}

func PullSource(s Source) ([]*Headline, error) {
	start := time.Now()

	req, err := http.NewRequest("GET", s.URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36")

	client := &http.Client{}
	res, err := client.Do(req)
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

	for _, h := range headlines {
		fmt.Printf("%s\n%s\n%s\n\n", h.Title, h.Date.UTC(), h.Subtitle)
	}

	duration := time.Since(start)
	fmt.Printf("(%s) pull took (%.1f) seconds\n", s.Name, duration.Seconds())
	return nil, nil
}
