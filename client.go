package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HostURL - Default Render API URL
const HostURL string = "https://api.render.com/v1"

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	APIKey     string
	Request    *http.Request
}

func NewClient(apiKey, host *string) (*Client, error) {
	client := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
	}

	if host != nil {
		client.HostURL = *host
	}

	if apiKey == nil || *apiKey == "" {
		return nil, fmt.Errorf("API Key is required")
	}

	client.APIKey = *apiKey

	return &client, nil
}

func (c *Client) doRequest(method, url string, data interface{}, jsonSchema interface{}) error {

	buf := new(bytes.Buffer)
	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return err
		}
		buf = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.APIKey)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent && res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusAccepted {
		return fmt.Errorf("status code: %d, details: %s", res.StatusCode, body)
	}

	if jsonSchema != nil {
		err = json.Unmarshal(body, &jsonSchema)
		if err != nil {
			return err
		}
	}

	return nil
}
