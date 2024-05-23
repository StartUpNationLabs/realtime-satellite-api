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

type Client struct {
	client              *resty.Client
	loggedInCredentials LoggedInCredentials
	configuration       Conf
}

func New() (Client, error) {
	c := Client{}
	c.client = resty.New()
	c.configuration = confFromEnv()

	return c, nil
}

func (c *Client) Login() error {
	log.Info("Logging in")
	loginPage, _ := c.client.R().
		Get(c.configuration.LoginUrl)

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
		return fmt.Errorf("error getting cookies")
	}

	// Create form data
	formData := url.Values{}
	formData.Set("identity", c.configuration.username)
	formData.Set("password", c.configuration.password)
	formData.Set("spacetrack_csrf_token", spacecraftCookie)
	formData.Set("btnLogin", "LOGIN")

	// Create a new request
	req := c.client.R().
		SetFormDataFromValues(formData).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("Cookie", "spacetrack_csrf_cookie="+spacecraftCookie+";"+"chocolatechip="+chocolatechip)

	resp, _ := req.Post(c.configuration.LoginUrl)

	if resp.StatusCode() != 200 {
		return fmt.Errorf("error making request: %s, %s", resp.Status(), resp.String())
	}
	c.loggedInCredentials = LoggedInCredentials{spacecraftCookie, chocolatechip}
	return nil
}

func markData(data []TLE) []TLE {
	for i := range data {
		data[i].Group = []string{"Space-Track"}
	}
	return data
}

func (c *Client) FetchData() ([]TLE, error) {
	// if file exists, read from file
	if _, err := os.Stat(c.configuration.tleFile); err == nil {
		log.Info("Reading data from file")
		data, err := readDataFromFile(c.configuration.tleFile)
		data = markData(data)
		if err != nil {
			return nil, err
		}
		return data, nil

	}
	log.Info("Fetching data from spacetrack")
	req, err := c.client.R().
		SetHeader("Cookie", "spacetrack_csrf_cookie="+c.loggedInCredentials.spacecraftCookie+";"+"chocolatechip="+c.loggedInCredentials.chocolatechip).
		Get(c.configuration.FetchAllUrl)

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
	data = markData(data)
	if err != nil {
		return nil, err
	}

	// save the data to a file
	err = saveDataToFile(c.configuration.tleFile, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
