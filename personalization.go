package main

import (
	"fmt"
	"encoding/json"
	"log"
	"sort"
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

type SpotifyArtist struct {
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

type Artists struct {
	Artists []struct {
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
	} `json:"artists"`
}

type SeedRecommendations struct {
	Tracks []struct {
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
			AvailableMarkets []string `json:"available_markets"`
			ExternalUrls     struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href   string `json:"href"`
			ID     string `json:"id"`
			Images []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name string `json:"name"`
			Type string `json:"type"`
			URI  string `json:"uri"`
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
		AvailableMarkets []string `json:"available_markets"`
		DiscNumber       int      `json:"disc_number"`
		DurationMs       int      `json:"duration_ms"`
		Explicit         bool     `json:"explicit"`
		ExternalIds      struct {
			Isrc string `json:"isrc"`
		} `json:"external_ids"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href        string `json:"href"`
		ID          string `json:"id"`
		Name        string `json:"name"`
		Popularity  int    `json:"popularity"`
		PreviewURL  string `json:"preview_url"`
		TrackNumber int    `json:"track_number"`
		Type        string `json:"type"`
		URI         string `json:"uri"`
	} `json:"tracks"`
	Seeds []struct {
		InitialPoolSize    int    `json:"initialPoolSize"`
		AfterFilteringSize int    `json:"afterFilteringSize"`
		AfterRelinkingSize int    `json:"afterRelinkingSize"`
		ID                 string `json:"id"`
		Type               string `json:"type"`
		Href               string `json:"href"`
	} `json:"seeds"`
}

func (u *User) GetDisplayPicture() (string) {
	return u.Images[0].URL
}

// timeSpan - long_term: years, medium_term: 6 months, short_term: 4 weeks
// limit 1-50
func (c *Client) GetUserTopArtists(timeSpan string, limit int) (*[]Artist, error) {
	// spotify get
	var apiURL string

	switch timeSpan {
		case "long_term":
			apiURL = fmt.Sprintf("%sme/top/artists?time_range=long_term&limit=%d", c.baseUrl, limit)
		case "medium_term":
			apiURL = fmt.Sprintf("%sme/top/artists?time_range=medium_term&limit=%d", c.baseUrl, limit)
		case "short_term":
			apiURL = fmt.Sprintf("%sme/top/artists?time_range=short_term&limit=%d", c.baseUrl, limit)
		default:
			apiURL = fmt.Sprintf("%sme/top/artists?time_range=long_term&limit=%d", c.baseUrl, limit)
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

// timeSpan - long_term: years, medium_term: 6 months, short_term: 4 weeks
// limit 1-50
func (c *Client) GetUserTopSongs(timeSpan string, limit int) (*[]Song, error) {
	// spotify get
	var apiURL string

	switch timeSpan {
		case "long_term":
			apiURL = fmt.Sprintf("%sme/top/tracks?time_range=long_term&limit=%d", c.baseUrl, limit)
		case "medium_term":
			log.Println("Medium term")
			apiURL = fmt.Sprintf("%sme/top/tracks?time_range=medium_term&limit=%d", c.baseUrl, limit)
		case "short_term":
			apiURL = fmt.Sprintf("%sme/top/tracks?time_range=short_term&limit=%d", c.baseUrl, limit)
		default:
			apiURL = fmt.Sprintf("%sme/top/tracks?time_range=long_term&limit=%d", c.baseUrl, limit)
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
func (c *Client) GetUserFavouriteGenres(artists []Artist) (*[]string, error) {
	// Get artists genres, find top count
	//nil
	var m map[string]int
	// allocates and initializes
	m = make(map[string]int)

	// Get a count for each genre
	for _, item := range artists {
		for _, genre := range item.Genres {
				m[genre]++
		}
	}

	type GenreCount struct {
		GenreName string
		Count int
	}

	var genreArray []GenreCount
	for gn, c := range m {
		genreArray = append(genreArray, GenreCount{gn, c})
	}

	// Sorting the slice by value
	sort.Slice(genreArray, func(i, j int) bool {
		return genreArray[i].Count > genreArray[j].Count
	})

	top := make([]string, 0)

	// Get top 5 if there is enough
	if len(genreArray) >= 5 {
		top := make([]string, 0)

		for _, genre := range genreArray[:5] {
			top = append(top, genre.GenreName)
		}
		return &top, nil
	}

	for _, genre := range genreArray {
		top = append(top, genre.GenreName)
	}
	
	// Getting info for each

	return &top, nil
}

// TODO: Add return value
func (c *Client) GetRecommendedArtists(artists []Artist) (*[]Artist, error) {
	// spotify get --> Take artist IDs, and use https://developer.spotify.com/console/get-artist-related-artists/
	var apiURL string
	
	//nil
	var m map[string]int
	// allocates and initializes
	m = make(map[string]int)

	type ArtistCount struct {
		ArtistID string
		Count int
	}

	// Get recommended top songs from each top artist, find the top count EXCEPT the top artists

	for _, artist := range artists {
		apiURL = fmt.Sprintf("%sartists/%s/related-artists", c.baseUrl, artist.Id)

		var recommended Artists

		// Get api request
		response, err := c.connection.Get(apiURL)
		if err != nil {
			log.Println("Failed to request")
			return nil, err
		}

		defer response.Body.Close()
		
		if err := json.NewDecoder(response.Body).Decode(&recommended); err != nil {
			log.Println("Failed to convert JSON")
			log.Println(err)
			return nil, err
		}

		// Counting the recommended for each artist recommendation
		for _, artist := range recommended.Artists {
			m[artist.ID]++
		}
	}

	var artistArray []ArtistCount
	for aid, c := range m {
		artistArray = append(artistArray, ArtistCount{aid, c})
	}

	// Sorting the slice by value
	sort.Slice(artistArray, func(i, j int) bool {
		return artistArray[i].Count > artistArray[j].Count
	})

	top := make([]string, 0)
	var unique bool
	var counter int

	for _, recArtist := range artistArray {
		for _, artist := range artists {
			if recArtist.ArtistID == artist.Id {
				unique = false
				continue
			}
			unique = true
		}

		// Add the artist into top array
		if unique {
			top = append(top, recArtist.ArtistID)
			counter++
		}

		// Only get 5
		if counter == 5 {
			break
		}
	}

	var recommendArtists Artists

	apiURL2 := fmt.Sprintf("%sartists?ids=%s,%s,%s,%s,%s", c.baseUrl, top[0], top[1], top[2], top[3], top[4])

	// Get info of artists through multiple artists endpoint
	response, err := c.connection.Get(apiURL2)
	if err != nil {
		log.Println("Failed to request")
		return nil, err
	}

	defer response.Body.Close()
	
	if err := json.NewDecoder(response.Body).Decode(&recommendArtists); err != nil {
		log.Println("Failed to convert JSON")
		log.Println(err)
		return nil, err
	}

	// Get the top 5 recommended artists information
	finalzedArtists := make([]Artist, 0)

	for _, item := range recommendArtists.Artists {
		finalzedArtists = append(finalzedArtists, Artist{Id: item.ID, Uri: item.URI, ProfileUrl: item.ExternalUrls.Spotify, ImageUrl: item.Images[0].URL,
			 Name: item.Name, Followers: item.Followers.Total, Popularity: item.Popularity, Genres: item.Genres})
	}
	

	return &finalzedArtists, nil
}

func (c *Client) GetRecommendedSongs(recommendedArtists []Artist) (*[]Song, error) {
	var recommendations SeedRecommendations
	
	// TODO: Seed genre, artists, dyanamic 
	apiURL2 :=  fmt.Sprintf("%srecommendations?limit=10&seed_artists=%s,%s,%s,%s,%s", c.baseUrl, recommendedArtists[0].Id, recommendedArtists[1].Id, recommendedArtists[2].Id, recommendedArtists[3].Id, recommendedArtists[4].Id)

	// Get info of artists through multiple artists endpoint
	response, err := c.connection.Get(apiURL2)
	if err != nil {
		log.Println("Failed to request")
		return nil, err
	}

	defer response.Body.Close()
	
	if err := json.NewDecoder(response.Body).Decode(&recommendations); err != nil {
		log.Println("Failed to convert JSON")
		log.Println(err)
		return nil, err
	}

	songs := make([]Song, 0)

	for _, item := range recommendations.Tracks {
		songs = append(songs, Song {Id: item.ID, Uri: item.URI, SongUrl: item.Href, Name: item.Name,
			AlbumName: item.Album.Name, AlbumArtUrl: item.Album.Images[0].URL, Artist: item.Artists[0].Name})
	}

	return &songs, nil
}

// TODO: Add return value
func (c *Client) CreatePersonalizedPlaylist() {
	// Get recommended songs --> Use post to create playlist
}