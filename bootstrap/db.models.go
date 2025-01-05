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
	PulledAt    time.Time
	Keywords    *[]string `gorm:"serializer:json"`
}

func init() {
	Db.AutoMigrate(&Headline{})
}
