package client

import (
    "net/http"
)

type Client struct {
    httpClient *http.Client
    BaseURL    string
}

func NewClient(baseURL string) *Client {
    return &Client{
        httpClient: &http.Client{},
        BaseURL:    baseURL,
    }
}

func (c *Client) MakeRequest(endpoint string, method string) (*http.Response, error) {
    // Construct the full URL
    url := c.BaseURL + endpoint

    // Create a new request
    req, err := http.NewRequest(method, url, nil)
    if err != nil {
        return nil, err
    }

    // Send the request
    return c.httpClient.Do(req)
}
