package bootstrap

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	var err error

	if Config.DBPath == "" {
		panic("no db path")
	}

	Db, err = gorm.Open(sqlite.Open(Config.DBPath), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		panic("failed to open database")
	}
}
