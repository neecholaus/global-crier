package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Puller struct{}

func (p Puller) Pull(s Source) ([]*Headline, error) {
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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	type t struct {
		Result struct {
			Videos []map[string]interface{} `json:"videos"`
		} `json:"result"`
	}

	js := t{}

	err = json.Unmarshal(body, &js)
	if err != nil {
		return nil, err
	}

	for _, v := range js.Result.Videos {
		fmt.Printf("%v\n", v["title"])
	}

	duration := time.Since(start)
	fmt.Printf("(%s) pull took (%.1f) seconds\n", s.Name, duration.Seconds())
	return nil, nil
}
