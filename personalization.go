package main

import (
	"fmt"
	"encoding/json"
	"log"
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
	ImageUrl string
	Name string
	Followers int
	Popularity int
	Genres []string
}

type TopSongs struct {
	Items []struct {
		Album struct {
			AlbumType string `json:"album_type"`
			Artists   []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"artists"`
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href   string `json:"href"`
			ID     string `json:"id"`
			Images []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name                 string `json:"name"`
			ReleaseDate          string `json:"release_date"`
			ReleaseDatePrecision string `json:"release_date_precision"`
			Type                 string `json:"type"`
			URI                  string `json:"uri"`
		} `json:"album"`
		Artists []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"artists"`
		DiscNumber  int  `json:"disc_number"`
		DurationMs  int  `json:"duration_ms"`
		Explicit    bool `json:"explicit"`
		ExternalIds struct {
			Isrc string `json:"isrc"`
		} `json:"external_ids"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href        string `json:"href"`
		ID          string `json:"id"`
		IsLocal     bool   `json:"is_local"`
		IsPlayable  bool   `json:"is_playable"`
		Name        string `json:"name"`
		Popularity  int    `json:"popularity"`
		PreviewURL  string `json:"preview_url"`
		TrackNumber int    `json:"track_number"`
		Type        string `json:"type"`
		URI         string `json:"uri"`
	} `json:"items"`
	Total    int         `json:"total"`
	Limit    int         `json:"limit"`
	Offset   int         `json:"offset"`
	Href     string      `json:"href"`
	Previous interface{} `json:"previous"`
	Next     string      `json:"next"`
}

type Song struct {
	Id string
	Uri string
	SongUrl string
	Name string
	AlbumName string
	AlbumArtUrl string
	// TODO: Add multi artist
	Artist string
}

func (u *User) GetDisplayPicture() (string) {
	return u.Images[0].URL
}

// timeSpan - long_term: years, medium_term: 6 months, short_term: 4 weeks
func (c *Client) GetUserTopArtists(timeSpan string) (*[]Artist, error) {
	// spotify get
	var apiURL string

	switch timeSpan {
		case "long_term":
			apiURL = fmt.Sprintf("%sme/top/artists?time_range=long_term&limit=5", c.baseUrl)
		case "medium_term":
			apiURL = fmt.Sprintf("%sme/top/artists?time_range=medium_term&limit=5", c.baseUrl)
		case "short_term":
			apiURL = fmt.Sprintf("%sme/top/artists?time_range=short_term&limit=5", c.baseUrl)
		default:
			apiURL = fmt.Sprintf("%sme/top/artists?time_range=long_term&limit=5", c.baseUrl)
	}

	var ta TopArtists

	// Get api request
	response, err := c.connection.Get(apiURL)
	if err != nil {
		log.Println("Failed to request")
		return nil, err
	}

	defer response.Body.Close()
	
    if err := json.NewDecoder(response.Body).Decode(&ta); err != nil {
		log.Println("Failed to convert JSON")
		log.Println(err)
		return nil, err
	}

	artists := make([]Artist, 0)

	for _, item := range ta.Items {
		artists = append(artists, Artist{Id: item.ID, Uri: item.URI, ProfileUrl: item.ExternalUrls.Spotify, ImageUrl: item.Images[0].URL,
			 Name: item.Name, Followers: item.Followers.Total, Popularity: item.Popularity, Genres: item.Genres})
	}

	return &artists, nil
}

// TODO: Add return value
func (c *Client) GetUserTopSongs(timeSpan string) (*[]Song, error) {
	// spotify get
	var apiURL string

	switch timeSpan {
		case "long_term":
			apiURL = fmt.Sprintf("%sme/top/tracks?time_range=long_term&limit=10", c.baseUrl)
		case "medium_term":
			log.Println("Medium term")
			apiURL = fmt.Sprintf("%sme/top/tracks?time_range=medium_term&limit=10", c.baseUrl)
		case "short_term":
			apiURL = fmt.Sprintf("%sme/top/tracks?time_range=short_term&limit=10", c.baseUrl)
		default:
			apiURL = fmt.Sprintf("%sme/top/tracks?time_range=long_term&limit=10", c.baseUrl)
	}

	var ts TopSongs

	// Get api request
	response, err := c.connection.Get(apiURL)
	if err != nil {
		log.Println("Failed to request")
		return nil, err
	}

	defer response.Body.Close()
	
    if err := json.NewDecoder(response.Body).Decode(&ts); err != nil {
		log.Println("Failed to convert JSON")
		log.Println(err)
		return nil, err
	}

	songs := make([]Song, 0)

	for _, item := range ts.Items {
		songs = append(songs, Song {Id: item.ID, Uri: item.URI, SongUrl: item.Href, Name: item.Name,
			AlbumName: item.Album.Name, AlbumArtUrl: item.Album.Images[0].URL, Artist: item.Artists[0].Name})
	}

	return &songs, nil
}

// TODO: Add return value
func (c *Client) GetUserFavouriteGenres() {
	// Get artists genres, find top count
}

// TODO: Add return value
func (c *Client) GetRecommendedArtists() {
	// spotify get --> Take artist IDs, and use https://developer.spotify.com/console/get-artist-related-artists/
}

// TODO: Add return value
func (c *Client) GetRecommendedSongs() {
	// Get top artists --> seed --> Get top songs
}

// TODO: Add return value
func (c *Client) CreatePersonalizedPlaylist() {
	// Get recommended songs --> Use post to create playlist
}