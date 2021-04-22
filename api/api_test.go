package api

import (
	"bytes"
	"context"
	"go-api/config"
	"testing"

	"golang.org/x/oauth2"
)

var resBody []byte = []byte{123, 34, 97, 99, 99, 101, 115, 115, 95, 116,
	111, 107, 101, 110, 34, 58, 34, 85, 83, 51, 108, 66, 106, 75, 74, 65,
	115, 68, 74, 97, 51, 52, 122, 112, 89, 74, 52, 49, 100, 78, 97, 68, 79,
	53, 72, 87, 80, 103, 52, 98, 100, 34, 44, 34, 116, 111, 107, 101, 110,
	95, 116, 121, 112, 101, 34, 58, 34, 98, 101, 97, 114, 101, 114, 34, 44,
	34, 101, 120, 112, 105, 114, 101, 115, 95, 105, 110, 34, 58, 56, 54, 51,
	57, 57, 125}

var token oauth2.Token = oauth2.Token{
	AccessToken: "US3lBjKJAsDJa34zpYJ41dNaDO5HWPg4bd",
	TokenType:   "bearer",
}

var con config.ExternalAPI = config.ExternalAPI{
	ClientID: "fd08d53532b6465da5424211bbfcd537",
	Secret:   "TNOTekALFK7PVwlEG8zQBOrS1bzjN3Wa",
	TokenURL: "https://us.battle.net/oauth/token",
	BaseURL:  "https://us.api.blizzard.com/hearthstone/cardbacks",
}

func TestGenerateOAuthReq(t *testing.T) {
	t.Run("", func(t *testing.T) {
		req, _ := generateOAuthReq(context.Background(), con)
		if req.Method != "POST" {
			t.Errorf("actual %v, expected %v", req.Method, "POST")
		}
		if req.Host != "us.battle.net" {
			t.Errorf("actual %v, expected %v", req.Host, "us.battle.net")
		}
		if req.URL.Path != "/oauth/token" {
			t.Errorf("actual %v, expected %v", req.URL.Path, "/oauth/token")
		}
		user, pass, _ := req.BasicAuth()
		if user != con.ClientID {
			t.Errorf("actual %v, expected %v", user, con.ClientID)
		}
		if pass != con.Secret {
			t.Errorf("actual %v, expected %v", pass, con.Secret)
		}
	})
}

func TestExtractToken(t *testing.T) {
	b := bytes.NewBuffer(resBody)

	t.Run("should extract access token from response body", func(t *testing.T) {
		expected := token
		actual := extractToken(b)
		if actual.AccessToken != expected.AccessToken ||
			actual.TokenType != expected.TokenType {
			t.Errorf("actual %v, expected %v", actual, expected)
		}
	})
}
