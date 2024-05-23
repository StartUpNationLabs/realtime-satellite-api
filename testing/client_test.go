package testing

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/tsukoyachi/react-flight-tracker-satellite/spacetrack"
	"testing"
)

func TestFetchData(t *testing.T) {

	// load dotenv
	err := godotenv.Load()
	if err != nil {
		log.Warnf("Error loading .env file")
	}

	client, _ := spacetrack.New()

	err = client.Login()
	if err != nil {
		return
	}
	data, err := client.FetchData()

	log.Println(data, err)
}
