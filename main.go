package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func createRequest(requestType, uri string, formData url.Values) (req *http.Request) {
	var err error

	if formData != nil {
		req, err = http.NewRequest(requestType, uri, strings.NewReader(formData.Encode()))
	} else {
		req, err = http.NewRequest(requestType, uri, nil)
	}

	if err != nil {
		fmt.Println("Error with the creation of the request:", err)
		os.Exit(1)
	}

	return req
}

func sendRequest(req http.Request, client http.Client) (resp *http.Response) {
	var err error

	resp, err = client.Do(&req)

	if err != nil {
		fmt.Println("Error making request:", err)
		os.Exit(1)
	}

	return resp
}

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
	var req *http.Request
	var resp *http.Response
	var username, password string
	var spacecraftCookie, chocolatechip string

	username, password = loadEnv()
	// Create all the mandatory request to login and load the data
	client := &http.Client{}

	req = createRequest("GET", "https://www.space-track.org/auth/login", nil)

	resp = sendRequest(*req, *client)

	// Close the body at the end of the method
	defer resp.Body.Close()

	for _, cookie := range resp.Cookies() {
		if cookie.Name == "spacetrack_csrf_cookie" {
			spacecraftCookie = cookie.Value
		} else if cookie.Name == "chocolatechip" {
			chocolatechip = cookie.Value
		}
	}

	if chocolatechip == "" || spacecraftCookie == "" {
		fmt.Println("The cookie aren't fectched successfully :/")
		return
	}

	fmt.Println("cookie_spacecraft_csr :", spacecraftCookie, ", chocolatechip :", chocolatechip)

	// Create form data
	formData := url.Values{}
	formData.Set("identity", username)
	formData.Set("password", password)
	formData.Set("spacetrack_csrf_token", spacecraftCookie)
	formData.Set("btnLogin", "LOGIN")

	// Create a new request
	req = createRequest("POST", "https://www.space-track.org/auth/login", formData)

	// Set the content type header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "spacetrack_csrf_cookie="+spacecraftCookie+";"+"chocolatechip="+chocolatechip)

	// Make the request
	resp = sendRequest(*req, *client)

	defer resp.Body.Close()

	// Print the response status code
	fmt.Println("Response Status:", resp.Status)
}
