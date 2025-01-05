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
		ExtractKeywordsFromTitle(h)

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

		// create
		prepared := bootstrap.Headline{
			Title:       h.Title,
			Description: h.Subtitle,
			URL:         h.URL,
			PulledAt:    h.PulledAt,
			Keywords:    h.Keywords,
		}

		res = bootstrap.Db.Create(&prepared)
		if res.Error != nil {
			fmt.Printf("ERR failed storing headline")
			continue
		}
		successCount++
	}

	duration := time.Since(start)
	fmt.Printf("GOOD (%d) stored, (%d) duplicates, (%d) total, (%.1f) seconds\n", successCount, existingCount, len(headlines), duration.Seconds())
}

func ExtractKeywordsFromTitle(headline *Headline) {
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
