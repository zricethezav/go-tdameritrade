package main

import (
	"context"
	"fmt"
	"github.com/zricethezav/go-tdameritrade"
	"golang.org/x/oauth2"
	"log"
	"os"
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
	if refreshToken == "" {
		log.Fatal("No Account ID present")
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

	// more examples here: https://developer.tdameritrade.com/content/place-order-samples
	resp, err := c.Account.PlaceOrder(ctx, accountID, &tdameritrade.Order{
		Session: "NORMAL",
		Duration: "DAY",
		OrderType: "MARKET",
		OrderStrategyType: "SINGLE",
		OrderLegCollection: []*tdameritrade.OrderLegCollection{
			{
				Instruction: "Sell",
				Quantity: 2,
				Instrument: tdameritrade.Instrument{
					AssetType: "EQUITY",
					Data: &tdameritrade.Equity{
						Symbol: "XYZ",
					},
				},

			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode)
}

