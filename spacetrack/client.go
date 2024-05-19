package spacetrack

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
)
import "github.com/go-resty/resty/v2"

import (
	log "github.com/sirupsen/logrus"
)

func Login(client *resty.Client, credentials Credentials) (LoggedInCredentials, error) {

	loginPage, _ := client.R().
		Get(SPACE_TRACK_LOGIN_URL)

	spacecraftCookie := ""
	chocolatechip := ""

	for _, cookie := range loginPage.Cookies() {
		if cookie.Name == "spacetrack_csrf_cookie" {
			spacecraftCookie = cookie.Value
		} else if cookie.Name == "chocolatechip" {
			chocolatechip = cookie.Value
		}
	}

	if chocolatechip == "" || spacecraftCookie == "" {
		return LoggedInCredentials{}, fmt.Errorf("the cookie aren't fectched successfully :/")
	}

	// Create form data
	formData := url.Values{}
	formData.Set("identity", credentials.Identity)
	formData.Set("password", credentials.Password)
	formData.Set("spacetrack_csrf_token", spacecraftCookie)
	formData.Set("btnLogin", "LOGIN")

	// Create a new request
	req := client.R().
		SetFormDataFromValues(formData).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("Cookie", "spacetrack_csrf_cookie="+spacecraftCookie+";"+"chocolatechip="+chocolatechip)

	resp, _ := req.Post(SPACE_TRACK_LOGIN_URL)

	if resp.StatusCode() != 200 {
		return LoggedInCredentials{}, fmt.Errorf("error making request to log in: %s, %s", resp.Status(), resp.String())
	}
	return LoggedInCredentials{credentials.Identity, credentials.Password, spacecraftCookie, chocolatechip}, nil
}

func FetchData(client *resty.Client, tleFilepath string, loggedInCredentials LoggedInCredentials) ([]TLE, error) {
	// if file exists, read from file
	if _, err := os.Stat(tleFilepath); err == nil {
		log.Info("Reading data from file")
		data, err := readDataFromFile(tleFilepath)
		if err != nil {
			return nil, err
		}
		return data, nil

	}
	log.Info("Fetching data from spacetrack")
	req, err := client.R().
		SetHeader("Cookie", "spacetrack_csrf_cookie="+loggedInCredentials.spacecraftCookie+";"+"chocolatechip="+loggedInCredentials.chocolatechip).
		Get(SPACE_TRACK_API)

	if err != nil {
		log.Error("Error making request:", err)
	}
	if req.StatusCode() != 200 {
		log.Error("Error making request:", req.Status())
		return nil, fmt.Errorf("error making request: %s, %s", req.Status(), req.String())
	}

	// unmarshal the response body
	var data []TLE
	err = json.Unmarshal(req.Body(), &data)
	if err != nil {
		return nil, err
	}

	// save the data to a file
	err = saveDataToFile(tleFilepath, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
