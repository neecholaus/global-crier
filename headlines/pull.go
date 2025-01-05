package headlines

import (
	"io"
	"net/http"
)

var pullClient = &http.Client{}

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
