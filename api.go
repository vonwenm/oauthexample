package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"code.google.com/p/gcfg"

	"github.com/codegangsta/martini"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

type Config struct {
	Oauth struct {
		ClientID     string
		ClientSecret string
	}
}

var (
	cfg   Config
	conf  *oauth2.Config
	state = "secret_state"
)

func init() {
	err := gcfg.ReadFileInto(&cfg, "conf.gcfg")
	if err != nil {
		log.Fatalf("Trouble reading config: %s", err)
	}
	conf = &oauth2.Config{
		ClientID:     cfg.Oauth.ClientID,
		ClientSecret: cfg.Oauth.ClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/plus.login", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
		RedirectURL: "https://localhost:8001/oauth2callback",
	}
}

func GetAuthCodeUrl() string {
	url := conf.AuthCodeURL(state, oauth2.AccessTypeOnline)
	return url
}

func GetUserInfo(w http.ResponseWriter, r *http.Request, params martini.Params) string {
	token, err := GetToken(r, params)
	if err != nil {
		return err.Error()
	}

	info, err := RequestUserInfo(token.AccessToken)
	if err != nil {
		return err.Error()
	}
	return info
}

func RequestUserInfo(token string) (string, error) {
	reqUrl := "https://www.googleapis.com/oauth2/v1/userinfo?alt=json"
	req := fmt.Sprintf("%s&access_token=%s", reqUrl, token)

	response, err := http.Get(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	info, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(info), nil
}

func GetToken(r *http.Request, params martini.Params) (*oauth2.Token, error) {
	qs := r.URL.Query()

	if qs.Get("state") != state {
		return nil, errors.New("State returned from auth server does not match!")
	}

	code := qs.Get("code")
	if code == "" {
		return nil, errors.New("Unable to acquire code!")
	}

	token, err := conf.Exchange(context.TODO(), code)
	if err != nil {
		return nil, err
	}

	return token, nil
}
