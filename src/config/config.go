package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	APPPORT  = 0
	DBDRIVER = ""
	DBURL    = ""
)

func Load() {
	if err := godotenv.Load(".env"); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			log.Printf("config: .env: %v (continuing with process environment)", err)
		}
	}

	var err error
	APPPORT, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil || APPPORT == 0 {
		APPPORT = 8000
	}
	DBDRIVER = os.Getenv("DB_DRIVER")
	if DBDRIVER == "" {
		DBDRIVER = "mysql"
	}
	DBURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}
