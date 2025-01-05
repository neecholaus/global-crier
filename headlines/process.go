package headlines

import (
	"fmt"
	"nick/global-crier/bootstrap"
	"regexp"
	"slices"
	"strings"
	"time"
)

func ProcessHeadlines(headlines []*Headline) {
	start := time.Now()

	for _, h := range headlines {
		ExtractKeywords(h)
	}

	prepared := []bootstrap.Headline{}
	for _, h := range headlines {
		prepared = append(prepared, bootstrap.Headline{
			Title:       h.Title,
			Description: h.Subtitle,
			URL:         h.URL,
			PulledAt:    h.PulledAt,
			Keywords:    h.Keywords,
		})
	}

	res := bootstrap.Db.Create(&prepared)
	if res.Error != nil {
		panic(res.Error)
	}

	duration := time.Since(start)
	fmt.Printf("Done processing (%d) headlines in (%.1f) seconds\n", len(headlines), duration.Seconds())
}

func ExtractKeywords(headline *Headline) {
	keywords := []string{}

	exp := regexp.MustCompile(`[^a-zA-Z0-9\s\-]+`)
	sanitized := exp.ReplaceAll([]byte(headline.Title), []byte{})

	sanitizedTitle := strings.ToLower(string(sanitized))
	sanitizedTitle = strings.ReplaceAll(sanitizedTitle, "'", "")
	sanitizedTitle = strings.ReplaceAll(sanitizedTitle, "\"", "")

	allwords := strings.Split(sanitizedTitle, " ")

	for _, word := range allwords {
		if len(word) > 1 && !slices.Contains([]string{"the", "what", "a", "an", "and", "that", "in", "on", "around", "will", "be", "his", "her", "this", "must", "may", "as"}, word) {
			keywords = append(keywords, word)
		}
	}

	headline.Keywords = &keywords
}
