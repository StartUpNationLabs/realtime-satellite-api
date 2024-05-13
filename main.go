package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/tsukoyachi/react-flight-tracker-satellite/spacetrack"
)

func loadEnv() (username, password string) {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file:", err)
		os.Exit(1)
	}

	//Load .env file variables
	username = os.Getenv("identity")
	password = os.Getenv("password")

	return username, password
}

func main() {
	username, password := loadEnv()
	credentials := spacetrack.Credentials{Identity: username, Password: password}
	loggedInCredentials, err := spacetrack.Login(credentials)
	if err != nil {
		fmt.Println(err)
	}
	spacetrack.FetchData("./tle.json", loggedInCredentials)
}
