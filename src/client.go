package client

import (
    "bytes"
    "encoding/json"
    "net/http"
)

type Client struct {
    BaseURL string
    APIKey  string
}

func NewClient(baseURL, apiKey string) *Client {
    return &Client{
        BaseURL: baseURL,
        APIKey:  apiKey,
    }
}

func (c *Client) MakeRequest(endpoint, method string, data interface{}) (*http.Response, error) {
    // Convert the request data to JSON
    jsonData, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }

    // Create the full URL
    url := c.BaseURL + endpoint

    // Create a new request
    req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
    if err != nil {
        return nil, err
    }

    // Set headers
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+c.APIKey)

    // Make the request
    client := &http.Client{}
    return client.Do(req)
}
