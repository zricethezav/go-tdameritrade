package main

import (
	"context"
	"fmt"
	"github.com/zricethezav/go-tdameritrade"
	"golang.org/x/oauth2"
	"log"
	"os"
	"time"
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

	hours, _, err := c.MarketHours.GetMarketHours(ctx, "Equity", time.Now())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", (*hours)["equity"]["EQ"])

	hours, _, err = c.MarketHours.GetMarketHoursMulti(ctx, "EQUITY,OPTION", time.Now())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", (*hours)["option"]["EQO"])
}

