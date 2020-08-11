package tdameritrade

import (
	"context"
	"fmt"
	"time"
)

// MarketHoursService handles communication with the marketdata related methods of
// the TDAmeritrade API.
//
// TDAmeritrade API docs: https://developer.tdameritrade.com/market-hours/apis
type MarketHoursService struct {
	client *Client
}

type MarketHours map[string]map[string]*Hours

type Period struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type SessionHours struct {
	PreMarket     []Period `json:"preMarket"`
	RegularMarket []Period `json:"regularMarket"`
	PostMarket    []Period `json:"postMarket"`
}

type Hours struct {
	Category     string       `json:"category"`
	Date         string       `json:"date"`
	Exchange     string       `json:"exchange"`
	IsOpen       bool         `json:"isOpen"`
	MarketType   string       `json:"marketType"`
	Product      string       `json:"product"`
	ProductName  string       `json:"productName"`
	SessionHours SessionHours `json:"sessionHours"`
}

func (s *MarketHoursService) GetMarketHoursMulti(ctx context.Context, markets string, date time.Time) (*MarketHours, *Response, error) {
	u := fmt.Sprintf("marketdata/hours")
	if markets == "" {
		return nil, nil, fmt.Errorf("no markets present")
	}
	u = fmt.Sprintf("%s?markets=%s", u, markets)
	if !date.IsZero() {
		u = fmt.Sprintf("%s&date=%s", u, date.Format("2006-01-02"))
	}

	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}

	hours := new(MarketHours)

	resp, err := s.client.Do(ctx, req, hours)
	if err != nil {
		return nil, resp, err
	}

	return hours, resp, nil
}

func (s *MarketHoursService) GetMarketHours(ctx context.Context, market string, date time.Time) (*MarketHours, *Response, error) {
	u := fmt.Sprintf("marketdata/%s/hours", market)

	if !date.IsZero() {
		u = fmt.Sprintf("%s?date=%s", u, date.Format("2006-01-02"))
	}

	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}

	hours := new(MarketHours)

	resp, err := s.client.Do(ctx, req, hours)
	if err != nil {
		return nil, resp, err
	}

	return hours, resp, nil
}
