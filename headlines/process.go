package headlines

import (
	"fmt"
	"nick/global-crier/bootstrap"
	"regexp"
	"slices"
	"strings"
	"time"
)

func ReprocessExistingHeadlines() {
	bootstrap.Db.Where("1=1").Unscoped().Delete(&bootstrap.HeadlineRelation{})
	bootstrap.Db.Where("1=1").Unscoped().Delete(&bootstrap.Keyword{})

	headlines := []*bootstrap.Headline{}
	res := bootstrap.Db.Order("ID asc").Find(&headlines)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}

	for _, h := range headlines {
		h.Keywords = []string{}

		kw := getKeywordsFromString(h.Title)
		h.Keywords = kw

		err := storeKeywords(h)
		if err != nil {
			fmt.Printf("Error storing keywords: %s\n", err)
			return
		}

		identifyAndStoreHeadlineRelations(h)
	}
}

func ProcessNewHeadlines(tmpHeadlines []*TmpHeadline) {
	start := time.Now()

	existingCount := 0
	successCount := 0

	newHeadlines := []*bootstrap.Headline{}

	for _, th := range tmpHeadlines {
		kw := getKeywordsFromString(th.Title)
		th.Keywords = kw

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
		identifyAndStoreHeadlineRelations(h)
	}

	duration := time.Since(start)
	fmt.Printf("(%d) new, (%d) existing, (%d) total, (%.1f) seconds\n", successCount, existingCount, len(tmpHeadlines), duration.Seconds())
}

func getKeywordsFromString(text string) []string {
	keywords := []string{}

	exp := regexp.MustCompile(`[^a-zA-Z0-9\s\-]+`)
	sanitized := exp.ReplaceAll([]byte(text), []byte{})

	sanitizedTitle := strings.ToLower(string(sanitized))
	sanitizedTitle = strings.ReplaceAll(sanitizedTitle, "'", "")
	sanitizedTitle = strings.ReplaceAll(sanitizedTitle, "\"", "")

	allwords := strings.Split(sanitizedTitle, " ")

	for _, word := range allwords {
		if len(word) > 1 && !slices.Contains([]string{"the", "what", "a", "an", "and", "that", "in", "on", "around", "will", "be", "his", "her", "this", "must", "may", "as", "at", "of", "to", "not", "by"}, word) {
			keywords = append(keywords, word)
		}
	}

	return keywords
}

func storeHeadline(tmpHeadline *TmpHeadline) (*bootstrap.Headline, error) {
	// create headline record
	//
	prepared := &bootstrap.Headline{
		Title:       tmpHeadline.Title,
		Description: tmpHeadline.Subtitle,
		URL:         tmpHeadline.URL,
		Publication: tmpHeadline.Source.Publication,
		PulledAt:    tmpHeadline.PulledAt,
		Keywords:    tmpHeadline.Keywords,
	}
	res := bootstrap.Db.Create(&prepared)
	if res.Error != nil {
		fmt.Printf("ERR failed storing headline")
		return nil, res.Error
	}

	fmt.Printf("TRACE created headline #%d\n", prepared.ID)

	err := storeKeywords(prepared)
	if err != nil {
		return nil, err
	}

	return prepared, nil
}

func storeKeywords(h *bootstrap.Headline) error {
	var keywords []*bootstrap.Keyword
	for _, kw := range h.Keywords {
		keywords = append(keywords, &bootstrap.Keyword{
			HeadlineID:       h.ID,
			HeadlinePulledAt: h.PulledAt,
			Keyword:          kw,
		})
	}

	res := bootstrap.Db.Create(&keywords)
	if res.Error != nil {
		fmt.Printf("ERR failed storing keywords")
		return res.Error
	}

	fmt.Printf("TRACE created keywords for headline %d\n", h.ID)

	return nil
}

func identifyAndStoreHeadlineRelations(h *bootstrap.Headline) {
	var keywordMatches []*bootstrap.Keyword
	res := bootstrap.Db.
		Where("keyword in ?", h.Keywords).
		Where("headline_id != ?", h.ID).
		Where("headline_pulled_at >= ?", h.PulledAt.AddDate(0, -3, 0)). // compared against headlines from the recent past
		Where("headline_pulled_at <= ?", h.PulledAt).
		Find(&keywordMatches)
	if res.Error != nil {
		fmt.Printf("failed pulling keywords (%s)\n", res.Error)
	}

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

		fmt.Printf("TRACE created relation for headlines #%d and #%d\n", h.ID, matchedHeadlineID)
	}
}
