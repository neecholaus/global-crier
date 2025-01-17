package headlines

import (
	"fmt"
	"nick/global-crier/bootstrap"
	"regexp"
	"slices"
	"strings"
	"time"
)

func ProcessHeadlines(tmpHeadlines []*TmpHeadline) {
	start := time.Now()

	existingCount := 0
	successCount := 0

	headlines := []*bootstrap.Headline{}

	for _, th := range tmpHeadlines {
		extractKeywordsFromTitle(th)

		// existence check
		res := bootstrap.Db.
			Where("title = ?", th.Title).
			Find(&bootstrap.Headline{})
		if res.Error != nil {
			fmt.Println(res.Error)
		}
		if res.RowsAffected > 0 {
			existingCount++
			continue
		}

		headline, err := storeHeadline(th)
		if err != nil {
			continue
		}

		headlines = append(headlines, headline)
		successCount++
	}

	for _, h := range headlines {
		createKeywordStreamRelations(h, []time.Time{time.Now(), time.Now().AddDate(0, 0, -1)})
	}

	duration := time.Since(start)
	fmt.Printf("(%d) new, (%d) existing, (%d) total, (%.1f) seconds\n", successCount, existingCount, len(tmpHeadlines), duration.Seconds())
}

func extractKeywordsFromTitle(headline *TmpHeadline) {
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

func storeHeadline(tmpHeadline *TmpHeadline) (*bootstrap.Headline, error) {
	// create headline record
	prepared := &bootstrap.Headline{
		Title:       tmpHeadline.Title,
		Description: tmpHeadline.Subtitle,
		URL:         tmpHeadline.URL,
		PulledAt:    tmpHeadline.PulledAt,
	}
	res := bootstrap.Db.Create(&prepared)
	if res.Error != nil {
		fmt.Printf("ERR failed storing headline")
		return nil, res.Error
	}

	// create keyword records
	var keywords []*bootstrap.Keyword
	for _, kw := range *tmpHeadline.Keywords {
		keywords = append(keywords, &bootstrap.Keyword{
			HeadlineID: prepared.ID,
			Keyword:    kw,
		})
	}
	res = bootstrap.Db.Create(&keywords)
	if res.Error != nil {
		fmt.Printf("ERR failed storing keywords")
		return nil, res.Error
	}

	return prepared, nil
}

func createKeywordStreamRelations(h *bootstrap.Headline, days []time.Time) {
	// pull all keywords on days and determine all headlines that match
}
