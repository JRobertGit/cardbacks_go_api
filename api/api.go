package api

import (
	"context"
	"encoding/json"
	"fmt"
	"go-api/config"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/oauth2"
)

func GetAuthenticatedClient(ctx context.Context, config config.ExternalAPI) (*http.Client, error) {
	req, err := generateOAuthReq(ctx, config)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	token := extractToken(resp.Body)
	return oauth2.NewClient(ctx, oauth2.StaticTokenSource(token)), nil
}

func generateOAuthReq(ctx context.Context, config config.ExternalAPI) (*http.Request, error) {
	v := url.Values{}
	v.Set("grant_type", "client_credentials")
	params := strings.NewReader(v.Encode())
	req, err := http.NewRequestWithContext(ctx, "POST", config.TokenURL, params)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(config.ClientID, config.Secret)

	return req, nil
}

func extractToken(responseBody io.Reader) *oauth2.Token {
	body, _ := ioutil.ReadAll(responseBody)
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)

	var duration time.Duration
	duration, _ = time.ParseDuration(fmt.Sprint(result["expires_in"]) + "s")

	token := oauth2.Token{
		AccessToken: result["access_token"].(string),
		TokenType:   result["token_type"].(string),
		Expiry:      time.Now().Local().Add(duration),
	}

	return &token
}
