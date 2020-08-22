package models

import "net/http"

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
