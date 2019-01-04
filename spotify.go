package main

import (
	"fmt"
	"net/http"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
	"io/ioutil"
	"../spotify/apikey"
)

const htmlIndex = "<html><body><a href='/SpotifyLogin'>Log in with Spotify</a></body></html>"

var (
    spotifyOauthConfig = &oauth2.Config{
        RedirectURL:    "http://localhost:3000/SpotifyCallback",
        ClientID:    constants.ClientID,
        ClientSecret: constants.ClientSecret,
        Scopes:       []string{"user-read-private",
            "user-read-email"},
        Endpoint:     spotify.Endpoint,
    }
    oauthStateString = "state"
)

func main() {
	http.HandleFunc("/", handleMain)
    http.HandleFunc("/SpotifyLogin", handleSpotifyLogin)
    http.HandleFunc("/SpotifyCallback", handleSpotifyCallback)
    fmt.Println(http.ListenAndServe(":3000", nil))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlIndex)
}

func handleSpotifyLogin(w http.ResponseWriter, r *http.Request) {
	// AuthCodeURL --> Exchange --> Client --> Get requests
	url := spotifyOauthConfig.AuthCodeURL(oauthStateString)
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleSpotifyCallback(w http.ResponseWriter, r *http.Request) {
	// See the state
	state := r.FormValue("state")
    if state != oauthStateString {
        fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }
	// Check whether obtained a code for token
    code := r.FormValue("code")
    token, err := spotifyOauthConfig.Exchange(oauth2.NoContext, code)
    if err != nil {
        fmt.Println("Code exchange failed with '%s'\n", err)
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
	}
	// Use token
	client := spotifyOauthConfig.Client(oauth2.NoContext, token)
	
	response, err := client.Get("https://api.spotify.com/v1/me")

    defer response.Body.Close()
    contents, err := ioutil.ReadAll(response.Body)
    fmt.Fprintf(w, "Content: %s\n", contents)
}