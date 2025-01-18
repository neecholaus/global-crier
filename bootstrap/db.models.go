package bootstrap

import (
	"time"

	"gorm.io/gorm"
)

type Headline struct {
	gorm.Model
	Title       string
	Description string
	URL         string
	Publication string
	PulledAt    time.Time
	Keywords    []string `gorm:"serializer:json"`
}

type Keyword struct {
	gorm.Model
	HeadlineID       uint
	Keyword          string
	HeadlinePulledAt time.Time
}

// Idea being that new headlines are related to headlines that have
// already been pulled. By specifying which was first we can more
// easily step through time to see the "stream".
type HeadlineRelation struct {
	gorm.Model
	ExistingHeadlineID uint
	NewHeadlineID      uint
	KeywordMatches     []string `gorm:"serializer:json"`
}

func init() {
	Db.AutoMigrate(&Headline{}, &Keyword{}, &HeadlineRelation{})
}
