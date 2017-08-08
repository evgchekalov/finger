package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var ClientID = os.Getenv("CLIENT_ID")
var RedirectURL = os.Getenv("REDIRECT_URL")
var ClientSecret = os.Getenv("CLIENT_SECRET")

func showDialog(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf(
		"https://www.facebook.com/v2.10/dialog/oauth?client_id=%s&redirect_uri=%s",
		ClientID, RedirectURL)
	fmt.Println(url)
	http.Redirect(w, r, url, 302)
}

func getToken(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if len(code) == 0 {
		http.Error(w, "no code", 400)
		return
	}

	url := fmt.Sprintf(
		"https://graph.facebook.com/v2.10/oauth/access_token?client_id=%s&redirect_uri=%s&client_secret=%s&code=%s",
		ClientID, RedirectURL, ClientSecret, code)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	defer resp.Body.Close()

	var i interface{}
	json.NewDecoder(resp.Body).Decode(&i)
	m := i.(map[string]interface{})
	t, ok := m["access_token"]
	if !ok {
		http.Error(w, "no access_token", 400)
		return
	}
	fmt.Fprintf(w, t.(string))
}

func hc(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/hc", hc)

	http.HandleFunc("/fb-dialog", showDialog)
	http.HandleFunc("/fb-token", getToken)
	http.ListenAndServe(":8080", nil)
}
