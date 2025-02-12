package bootstrap

import (
	"os"

	"github.com/joho/godotenv"
)

var Config *config

type config struct {
	DBPath string
}

func init() {
	Config = &config{}

	err := godotenv.Load()
	if err != nil {
		panic("coult not load .env")
	}

	Config.DBPath = os.Getenv("DBPATH")
}
