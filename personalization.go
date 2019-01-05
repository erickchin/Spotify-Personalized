package main

import (
	"fmt"
	"encoding/json"
)

type TopArtists struct {
	Items []struct {
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Followers struct {
			Href  interface{} `json:"href"`
			Total int         `json:"total"`
		} `json:"followers"`
		Genres []string `json:"genres"`
		Href   string   `json:"href"`
		ID     string   `json:"id"`
		Images []struct {
			Height int    `json:"height"`
			URL    string `json:"url"`
			Width  int    `json:"width"`
		} `json:"images"`
		Name       string `json:"name"`
		Popularity int    `json:"popularity"`
		Type       string `json:"type"`
		URI        string `json:"uri"`
	} `json:"items"`
	Total    int         `json:"total"`
	Limit    int         `json:"limit"`
	Offset   int         `json:"offset"`
	Href     string      `json:"href"`
	Previous interface{} `json:"previous"`
	Next     string      `json:"next"`
}

type Artist struct {
	Id string
	Uri string
	ProfileUrl string
	ImagesUrl string
	Name string
	Followers int
	Popularity int
	Genres []string
}

func (u *User) GetDisplayPicture() (string) {
	return u.Images[0].URL
}

// TODO: Add return value
// timeSpan - long_term: years, medium_term: 6 months, short_term: 4 weeks
func (c *Client) GetUserTopArtists(timeSpan string) (*[]Artist, error) {
	// spotify get
	var apiURL string

	switch timeSpan {
		case "long_term":
			apiURL = fmt.Sprintf("%sme/top/artists?time_range=long_term&limit=5", c.baseUrl)
		case "medium_term":
			apiURL = fmt.Sprintf("%sme/top/artists?time_range=long_term&limit=5", c.baseUrl)
		case "short_term":
			apiURL = fmt.Sprintf("%sme/top/artists?time_range=long_term&limit=5", c.baseUrl)
		default:
			apiURL = fmt.Sprintf("%sme/top/artists?time_range=long_term&limit=5", c.baseUrl)
	}

	var ta TopArtists

	// Get api request
	response, err := c.connection.Get(apiURL)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

    if err := json.NewDecoder(response.Body).Decode(&ta); err != nil {
		return nil, err
	}

	var artists []Artist

	/*for _, item := range ta.Items {

	}*/

	return &artists, nil
}

// TODO: Add return value
func (c *Client) GetUserTopSongs() {
	// spotify get
}

// TODO: Add return value
func (c *Client) GetUserFavouriteGenres() {
	// spotify get
}

// TODO: Add return value
func (c *Client) GetRecommendedArtists() {
	// spotify get
}

// TODO: Add return value
func (c *Client) GetRecommendedSongs() {
	// Get top artists --> seed --> Get top songs
}

// TODO: Add return value
func (c *Client) CreatePersonalizedPlaylist() {
	// Get recommended songs --> Use post to create playlist
}