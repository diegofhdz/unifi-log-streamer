package unifi

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Client struct {
	client *http.Client
}

func parseUrl(baseURL string, queryParams []string) string {
	if queryParams == nil {
		return baseURL
	}
	paramsString := strings.Join(queryParams, "&")
	return baseURL + "?" + paramsString
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	httpClient2 := *httpClient
	c := &Client{client: &httpClient2}
	c.initialize()
	return c
}

func (c *Client) ListHosts(pageSize int, nextToken string) (string, error) {
	apiURL := "https://api.ui.com/v1/hosts"
	fmt.Printf("%s\n", apiURL)

	params := []string{}
	params = append(params, fmt.Sprintf("pageSize=%s", strconv.Itoa(pageSize)))
	params = append(params, fmt.Sprintf("nextToken=%s", nextToken))

	parsedUrl := parseUrl(apiURL, params)
	fmt.Printf("%s\n", parsedUrl)

	return parsedUrl, nil
}
