package main

import (
    "net/http"
    "html/template"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
    "../spotify/constants"
    "log"
)

const htmlIndex = "<html><body><a href='/SpotifyLogin'>Log in with Spotify</a></body></html>"

var (
    spotifyOauthConfig = &oauth2.Config {
        RedirectURL:    "http://localhost:8080/SpotifyCallback",
        ClientID:    c.ClientID,
        ClientSecret: c.ClientSecret,
        Scopes:       []string{"user-read-private",
            "user-read-email", "user-top-read"},
        Endpoint:     spotify.Endpoint,
    }
    oauthStateString = "state"
)

var welcome *template.Template
var index *template.Template

type Client struct  {
    connection *http.Client
    baseUrl string
}

type Data struct {
    Profile User
    //{{(index .Artist 0)}} {{(index .Artist 1)}} access arrays
    TopArtists []Artist
    TopSongs []Song  
    TopGenre []string
    RecommendedArtists []Artist
    RecommendedSongs []Song
}

func init() {
    welcome = template.Must(template.ParseGlob("./welcome.html"))
    index = template.Must(template.ParseGlob("./index.html"))
}

func main() {
	/*http.HandleFunc("/", handleMain)
    http.HandleFunc("/SpotifyLogin", handleSpotifyLogin)
    http.HandleFunc("/SpotifyCallback", handleSpotifyCallback)*/
    
    mux := http.NewServeMux()
    mux.HandleFunc("/", handleMain)
    mux.HandleFunc("/SpotifyLogin", handleSpotifyLogin)
    mux.HandleFunc("/SpotifyCallback", handleSpotifyCallback)
    
    fileServer := http.FileServer(http.Dir("./img"))


    mux.Handle("/img/", http.StripPrefix("/img", fileServer))
    log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
    //fmt.Println(http.ListenAndServe(":3000", nil))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintf(w, htmlIndex)
    welcome.ExecuteTemplate(w, "welcome.html", nil)
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
        log.Printf("Invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }
	// Check whether obtained a code for token
    code := r.FormValue("code")
    token, err := spotifyOauthConfig.Exchange(oauth2.NoContext, code)
    if err != nil {
        log.Println("Code exchange failed with '%s'\n", err)
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
	}
    // Use token
    userClient := spotifyOauthConfig.Client(oauth2.NoContext, token)

	client := Client {userClient, "https://api.spotify.com/v1/"}
    
    userProfile, err := client.GetPersonalInfo()
    topArtists5, err := client.GetUserTopArtists("medium_term", 5)
    topArtists20, err := client.GetUserTopArtists("medium_term", 20)
    topSongs, err := client.GetUserTopSongs("medium_term", 20)
    topGenres, err := client.GetUserFavouriteGenres(*topArtists5)
    recommendedArtists, err := client.GetRecommendedArtists(*topArtists20)
    recommendedSongs, err := client.GetRecommendedSongs(*recommendedArtists)
    log.Println((*topSongs)[0].SongUrl)
    log.Println((*topSongs)[0].SongUrl)
    log.Println((*topSongs)[0].SongUrl)

    /*
    user, err := client.GetUserTopArtists("medium_term", 20)
    recommended, err := client.GetRecommendedArtists(*user)
    songs, err := client.GetRecommendedSongs(*recommended)
    log.Println((songs))*/
    
    n := Data {*userProfile, *topArtists5, *topSongs, *topGenres, *recommendedArtists, *recommendedSongs}
    index.Execute(w, n)
    index.ExecuteTemplate(w, "index.html", nil)

}