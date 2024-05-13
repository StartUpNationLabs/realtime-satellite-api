package spacetrack

import (
	"fmt"
	"net/http"
	"net/url"
)

func Login(username, password string) (spacecraftCookie, chocolatechip string) {
	var req *http.Request
	var resp *http.Response
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
	return spacecraftCookie, chocolatechip
}

func FetchData(tleFilepath, spacecraftCookie, chocolatechip string) {
	fmt.Println("Not implemented yet !")
}
