package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	var req *http.Request
	var resp *http.Response
	var err error
	var username, password string
	var spacecraftCookie, chocolatechip string

	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	//Load .env file variables
	username = os.Getenv("identity")
	password = os.Getenv("password")

	// Create all the mandatory request to login and load the data
	client := &http.Client{}
	req, err = http.NewRequest("GET", "https://www.space-track.org/auth/login", nil)

	if err != nil {
		fmt.Println("Error with the creation of the authentification request:", err)
		return
	}

	resp, err = client.Do(req)

	if err != nil {
		fmt.Println("Error with the first authentification request:", err)
		return
	}

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
	req, err = http.NewRequest("POST", "https://www.space-track.org/auth/login", strings.NewReader(formData.Encode()))

	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the content type header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "spacetrack_csrf_cookie="+spacecraftCookie+";"+"chocolatechip="+chocolatechip)

	// Make the request
	resp, err = client.Do(req)

	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Print the response status code
	fmt.Println("Response Status:", resp.Status)
}
