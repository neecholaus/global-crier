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

	existingCount := 0
	successCount := 0

	for _, h := range headlines {
		extractKeywordsFromTitle(h)

		// existence check
		res := bootstrap.Db.
			Where("title = ?", h.Title).
			Find(&bootstrap.Headline{})
		if res.Error != nil {
			fmt.Println(res.Error)
		}
		if res.RowsAffected > 0 {
			existingCount++
			continue
		}

		err := storeNewHeadline(h)
		if err == nil {
			successCount++
		}
	}

	duration := time.Since(start)
	fmt.Printf("(%d) new, (%d) existing, (%d) total, (%.1f) seconds\n", successCount, existingCount, len(headlines), duration.Seconds())
}

func extractKeywordsFromTitle(headline *Headline) {
	keywords := []string{}

	exp := regexp.MustCompile(`[^a-zA-Z0-9\s\-]+`)
	sanitized := exp.ReplaceAll([]byte(headline.Title), []byte{})

	sanitizedTitle := strings.ToLower(string(sanitized))
	sanitizedTitle = strings.ReplaceAll(sanitizedTitle, "'", "")
	sanitizedTitle = strings.ReplaceAll(sanitizedTitle, "\"", "")

	allwords := strings.Split(sanitizedTitle, " ")

	for _, word := range allwords {
		if len(word) > 1 && !slices.Contains([]string{"the", "what", "a", "an", "and", "that", "in", "on", "around", "will", "be", "his", "her", "this", "must", "may", "as", "at", "of", "to", "not", "by"}, word) {
			keywords = append(keywords, word)
		}
	}

	headline.Keywords = &keywords
}

func storeNewHeadline(headline *Headline) error {
	prepared := bootstrap.Headline{
		Title:       headline.Title,
		Description: headline.Subtitle,
		URL:         headline.URL,
		PulledAt:    headline.PulledAt,
		Keywords:    headline.Keywords,
	}

	res := bootstrap.Db.Create(&prepared)
	if res.Error != nil {
		fmt.Printf("ERR failed storing headline")
		return res.Error
	}

	return nil
}
