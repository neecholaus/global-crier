package headlines

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
	"time"
)

func ProcessHeadlines(headlines []*Headline) {
	start := time.Now()

	for _, h := range headlines {
		ExtractKeywords(h)

		fmt.Printf("%s\n%s\n\n", h.Title, *h.Keywords)
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
