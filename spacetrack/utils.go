package spacetrack

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
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
