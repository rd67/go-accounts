package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_DRIVER            = ""
	DB_CONNECTION_STRING = ""
)

func init() {
	// Load variables from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	DB_DRIVER = os.Getenv("DB_DRIVER")

	//	"${DB_DRIVER}://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):${DB_PORT})/$(DB_NAME)"
	DB_CONNECTION_STRING = fmt.Sprintf("%s://%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

}

func main() {
	// Print variables
	fmt.Println("env variables loaded....")
}
