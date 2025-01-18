package headlines

import (
	"fmt"
	"nick/global-crier/bootstrap"
	"regexp"
	"slices"
	"strings"
)

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
