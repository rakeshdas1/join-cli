package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
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

// JoinAPIClient struct for the main http client
type JoinAPIClient struct {
	httpClient *http.Client
	BaseURL    string
	APIKey     string
}

// makes a GET request to the provided URL
func (c *JoinAPIClient) makeGetRequest(url string) ([]byte, error) {
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
func (c *JoinAPIClient) constructAPIRequestURL(APIPath string) (apiReqURL string) {
	u, _ := url.Parse(c.BaseURL)
	u.Path = "_ah/api/registration/v1"
	u.Path = path.Join(u.Path, url.QueryEscape(APIPath))
	q := u.Query()
	q.Set("apikey", c.APIKey)
	u.RawQuery = q.Encode()
	return u.String()
}
func (c *JoinAPIClient) GetAllDevices() (*baseResponse, error) {
	var APIListAllDevicesPath = "listDevices"

	// container for API response
	var returnedResponse *baseResponse

	reqURL := c.constructAPIRequestURL(APIListAllDevicesPath)

	resp, err := c.makeGetRequest(reqURL)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(resp), &returnedResponse)
	return returnedResponse, nil
}
func (c *JoinAPIClient) NewHTTPClient() error {
	if c.httpClient == nil {
		fmt.Printf("Http Client does not exist, creating a new one..\n")
		c.httpClient = http.DefaultClient
	}
	return nil
}
