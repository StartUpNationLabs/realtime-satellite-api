package spacetrack

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

import (
	log "github.com/sirupsen/logrus"
)

// Credentials  is a struct that holds the username and password of the user
type Credentials struct {
	Identity string
	Password string
}
type LoggedInCredentials struct {
	identity         string
	password         string
	spacecraftCookie string
	chocolatechip    string
}

func Login(credentials Credentials) (LoggedInCredentials, error) {
	var req *http.Request
	var resp *http.Response
	// Create all the mandatory request to login and load the data
	client := &http.Client{}

	req = createRequest("GET", SPACE_TRACK_LOGIN_URL, nil)

	resp = sendRequest(*req, *client)

	// Close the body at the end of the method
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error("Error closing the body:", err)
		}
	}(resp.Body)

	spacecraftCookie := ""
	chocolatechip := ""

	for _, cookie := range resp.Cookies() {
		if cookie.Name == "spacetrack_csrf_cookie" {
			spacecraftCookie = cookie.Value
		} else if cookie.Name == "chocolatechip" {
			chocolatechip = cookie.Value
		}
	}

	if chocolatechip == "" || spacecraftCookie == "" {
		return LoggedInCredentials{}, fmt.Errorf("the cookie aren't fectched successfully :/")
	}

	// Print the response status code
	log.Info("Response Status:", resp.Status)

	// Create form data
	formData := url.Values{}
	formData.Set("identity", credentials.Identity)
	formData.Set("password", credentials.Password)
	formData.Set("spacetrack_csrf_token", spacecraftCookie)
	formData.Set("btnLogin", "LOGIN")

	// Create a new request
	req = createRequest("POST", SPACE_TRACK_LOGIN_URL, formData)

	// Set the content type header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "spacetrack_csrf_cookie="+spacecraftCookie+";"+"chocolatechip="+chocolatechip)

	// Make the request
	resp = sendRequest(*req, *client)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error("Error closing the body:", err)
		}
	}(resp.Body)

	// Print the response status code
	log.Info("Response Status:", resp.Status)
	return LoggedInCredentials{credentials.Identity, credentials.Password, spacecraftCookie, chocolatechip}, nil
}

func FetchData(tleFilepath string, loggedInCredentials LoggedInCredentials) {
	log.Fatalf("Not implemented yet")
}
