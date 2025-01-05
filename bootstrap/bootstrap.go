package bootstrap

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open(sqlite.Open("crier.db"), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		panic("failed to open database")
	}
}
