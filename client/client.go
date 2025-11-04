package client

import "fmt"

type Client struct {
	BaseUrl string
}

func New(baseUrl string) (*Client, error) {
	if baseUrl == "" {
		return nil, fmt.Errorf("BaseUrl must not be empty")
	}
	return &Client{BaseUrl: baseUrl}, nil
}
