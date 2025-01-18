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

	newHeadlines := []*bootstrap.Headline{}

	for _, th := range tmpHeadlines {
		extractKeywordsFromTitle(th)

		// check if headline is already in system
		//
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

		newHeadlines = append(newHeadlines, headline)
		successCount++
	}

	for _, h := range newHeadlines {
		createKeywordStreamRelations(h)
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

	headline.Keywords = keywords
}

func storeHeadline(tmpHeadline *TmpHeadline) (*bootstrap.Headline, error) {
	// create headline record
	//
	prepared := &bootstrap.Headline{
		Title:       tmpHeadline.Title,
		Description: tmpHeadline.Subtitle,
		URL:         tmpHeadline.URL,
		PulledAt:    tmpHeadline.PulledAt,
		Keywords:    tmpHeadline.Keywords,
	}
	res := bootstrap.Db.Create(&prepared)
	if res.Error != nil {
		fmt.Printf("ERR failed storing headline")
		return nil, res.Error
	}

	// todo log

	// create keyword records
	//
	var keywords []*bootstrap.Keyword
	for _, kw := range tmpHeadline.Keywords {
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

	// todo log

	return prepared, nil
}

func createKeywordStreamRelations(h *bootstrap.Headline) {
	var keywordMatches []*bootstrap.Keyword
	res := bootstrap.Db.
		Where("keyword in ?", h.Keywords).
		Where("headline_id != ?", h.ID).
		Find(&keywordMatches)
	if res.Error != nil {
		fmt.Printf("failed pulling keywords (%s)\n", res.Error)
	}

	// todo query the correct days

	matches := make(map[uint][]string, 0)

	// assemble map of shared keywords by headline id
	//
	for _, kword := range keywordMatches {
		if _, ok := matches[kword.HeadlineID]; !ok {
			matches[kword.HeadlineID] = []string{}
		}
		matches[kword.HeadlineID] = append(matches[kword.HeadlineID], kword.Keyword)
	}

	// create headline relation record for each strong match
	//
	for matchedHeadlineID, matchedKeywords := range matches {
		if len(matchedKeywords) < 4 {
			continue
		}

		// todo log

		prepared := bootstrap.HeadlineRelation{
			ExistingHeadlineID: matchedHeadlineID,
			NewHeadlineID:      h.ID,
			KeywordMatches:     matchedKeywords,
		}

		res := bootstrap.Db.Create(&prepared)
		if res.Error != nil {
			fmt.Printf("ERR failed storing headline match for #%d and #%d\n - %s", h.ID, matchedHeadlineID, res.Error)
			continue
		}
	}
}
