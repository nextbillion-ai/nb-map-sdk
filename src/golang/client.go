package client

import (
    "net/http"
)

type Client struct {
    // client properties like base URL, http client instances, headers, and more.
	BaseURL    string
    HTTPClient *http.Client
}

func NewClient() *Client {
    return &Client{
        // initializing client properties

		BaseURL:    baseURL,
        HTTPClient: &http.Client{},
    }
}

func (c *Client) MakeRequest(endpoint string) (*http.Response, error) {
    req, err := http.NewRequest("GET", c.BaseURL+endpoint, nil)
    if err != nil {
        return nil, err
    }

    return c.HTTPClient.Do(req)
}
