package testing

import (
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/tsukoyachi/react-flight-tracker-satellite/spacetrack"
	"os"
	"testing"
)

func TestFetchData(t *testing.T) {
	client := resty.New()

	creds, _ := spacetrack.Login(client, spacetrack.Credentials{
		Identity: os.Getenv("identity"),
		Password: os.Getenv("password"),
	})

	data, err := spacetrack.FetchData(client, "data.json", creds)

	log.Println(data, err)
}
