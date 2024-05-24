package testing

import (
	"github.com/StartUpNationLabs/react-flight-tracker-satellite/spacetrack"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
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
