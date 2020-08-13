package main

import (
	"context"
	"log"
	"os"

	"github.com/zricethezav/go-tdameritrade/tdameritrade"
	"golang.org/x/oauth2"
)

func main() {
	// pass an http client with auth
	token := os.Getenv("TDAMERITRADE_CLIENT_ID")
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}
	refreshToken := os.Getenv("TDAMERITRADE_REFRESH_TOKEN")
	if refreshToken == "" {
		log.Fatal("Unauthorized: No refresh token present")
	}

	accountID := os.Getenv("TDAMERITRADE_ACCOUNT_ID")

	conf := oauth2.Config{
		ClientID: token,
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://api.tdameritrade.com/v1/oauth2/token",
		},
		RedirectURL: "https://localhost",
	}

	tkn := &oauth2.Token{
		RefreshToken: refreshToken,
	}

	ctx := context.Background()
	tc := conf.Client(ctx, tkn)

	c, err := tdameritrade.NewClient(tc)
	if err != nil {
		log.Fatal(err)
	}

	preferences, resp, err := c.User.GetPreferences(ctx, accountID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Preferences: [Status: %d] %+v", resp.StatusCode, *preferences)

	userPrincipals, resp, err := c.User.GetUserPrincipals(ctx, "streamerSubscriptionKeys", "streamerConnectionInfo")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("User Principals: [Status: %d] %+v", resp.StatusCode, *userPrincipals)

	streamerSubscriptionKeys, resp, err := c.User.GetStreamerSubscriptionKeys(ctx, accountID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Streamer Subscription Keys: [Status: %d] %+v", resp.StatusCode, *streamerSubscriptionKeys)
}
