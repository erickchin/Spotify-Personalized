package main

import (
	"fmt"
	"encoding/json"
)

type User struct {
	DisplayName  string `json:"display_name"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Followers struct {
		Href  interface{} `json:"href"`
		Total int         `json:"total"`
	} `json:"followers"`
	Href   string `json:"href"`
	ID     string `json:"id"`
	Images []struct {
		Height interface{} `json:"height"`
		URL    string      `json:"url"`
		Width  interface{} `json:"width"`
	} `json:"images"`
	Type string `json:"type"`
	URI  string `json:"uri"`
}

func (c *Client) GetPersonalInfo() (*User, error) {
	apiURL := fmt.Sprintf("%s/me", c.baseUrl)

	var u User

	response, err := c.connection.Get(apiURL)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
    json.NewDecoder(response.Body).Decode(u)

	return &u, nil
}