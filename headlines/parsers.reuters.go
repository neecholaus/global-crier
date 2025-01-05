package headlines

import (
	"encoding/json"
	"time"
)

type reutersVideoParser struct{}

func (t reutersVideoParser) ParseBytes(s []byte) ([]*Headline, error) {
	type shape struct {
		Result struct {
			Videos []map[string]interface{} `json:"videos"`
		} `json:"result"`
	}

	js := shape{}

	err := json.Unmarshal(s, &js)
	if err != nil {
		return nil, err
	}

	response := []*Headline{}

	for _, v := range js.Result.Videos {
		response = append(response, &Headline{
			Title:    v["title"].(string),
			Subtitle: v["description"].(string),
			// URL:      v["url"].(string),
			Date: time.Now(),
		})
	}

	return response, nil
}

type reutersBigStoryParser struct{}

func (t reutersBigStoryParser) ParseBytes(s []byte) ([]*Headline, error) {
	type shape struct {
		Result struct {
			Articles []map[string]interface{} `json:"articles"`
		} `json:"result"`
	}

	js := shape{}

	err := json.Unmarshal(s, &js)
	if err != nil {
		return nil, err
	}

	response := []*Headline{}

	for _, v := range js.Result.Articles {
		response = append(response, &Headline{
			Title:    v["title"].(string),
			Subtitle: v["description"].(string),
			// URL:      v["url"].(string),
			Date: time.Now().UTC(),
		})
	}

	return response, nil
}
