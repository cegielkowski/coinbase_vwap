package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	WebServiceAddress string
	ProductIds        []string
	ValueLimitSize    int16
}

// LoadConfig will load config from environment variable.
func LoadConfig() (config *Config) {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	valueLimitSize, err := strconv.Atoi(os.Getenv("VALUE_LIMIT_SIZE"))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return &Config{
		ValueLimitSize:    int16(valueLimitSize),
		WebServiceAddress: os.Getenv("WEB_SERVICE_ADDRESS"),
		ProductIds:        strings.Split(os.Getenv("PRODUCT_IDS"), ","),
	}
}
