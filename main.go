package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	a := App{}

	//try to load .env file if environment variables aren't set
	if os.Getenv("APP_DB_USERNAME") == "" {
		// load .env file from given path
		// we keep it empty it will load .env from current directory
		err := godotenv.Load(".env")

		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	port, _ := strconv.Atoi(os.Getenv("APP_DB_PORT"))

	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		port)

	a.Run(":8010")
}
