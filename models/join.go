package models

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// struct for device info
type deviceInfo struct {
	ID          string `json:"id"`
	RegID       string `json:"regId"`
	RegID2      string `json:"regId2"`
	UserAccount string `json:"userAccount"`
	DeviceID    string `json:"deviceId"`
	DeviceName  string `json:"deviceName"`
	DeviceType  int    `json:"deviceType"`
	APILevel    int    `json:"apiLevel"`
	Model       string `json:"model"`
	HasTasker   bool   `json:"hasTasker"`
}

// struct for base response
type baseResponse struct {
	Success       string       `json:"success"`
	UserAuthError string       `json:"userAuthError"`
	Records       []deviceInfo `json:"records"`
}

// Client struct for the main http client
type Client struct {
	httpClient *http.Client
	BaseURL    string
	APIKey     string
}

// makes a GET request to the provided URL
func (c *Client) makeGetRequest(url string) ([]byte, error) {
	fmt.Println("Making the get req to " + url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// set the api key in the request header
	req.Header.Set("X-Api-Key", c.APIKey)
	resp, err := doReq(req, c.httpClient)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func doReq(req *http.Request, client *http.Client) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
