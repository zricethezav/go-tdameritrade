package tdameritrade

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	"time"
)

var (
	validPeriodTypes    = []string{"day", "month", "year", "ytd"}
	validFrequencyTypes = []string{"minute", "daily", "weekly", "monthly"}
)

const (
	defaultPeriodType    = "day"
	defaultFrequencyType = "minute"
)

// PriceHistoryService handles communication with the marketdata related methods of
// the TDAmeritrade API.
//
// TDAmeritrade API docs: https://developer.tdameritrade.com/price-history/apis
type PriceHistoryService struct {
	client *Client
}

// PriceHistoryOptions is parsed and translated to query options in the https request
type PriceHistoryOptions struct {
	PeriodType            string    `url:"periodType,omitempty"`
	Period                int       `url:"period,omitempty"`
	FrequencyType         string    `url:"frequencyType,omitempty"`
	Frequency             int       `url:"frequency,omitempty"`
	EndDate               time.Time `url:"endDate,omitempty"`
	StartDate             time.Time `url:"startDate,omitempty"`
	NeedExtendedHoursData *bool     `url:"needExtendedHoursData,omitempty"`
}

type PriceHistory struct {
	Candles []struct {
		Close    float64 `json:"close"`
		Datetime int     `json:"datetime"`
		High     float64 `json:"high"`
		Low      float64 `json:"low"`
		Open     float64 `json:"open"`
		Volume   float64 `json:"volume"`
	} `json:"candles"`
	Empty  bool   `json:"empty"`
	Symbol string `json:"symbol"`
}

// PriceHistory get the price history for a symbol
// TDAmeritrade API Docs: https://developer.tdameritrade.com/price-history/apis/get/marketdata/%7Bsymbol%7D/pricehistory
func (s *PriceHistoryService) PriceHistory(ctx context.Context, symbol string, opts *PriceHistoryOptions) (*PriceHistory, *Response, error) {
	u := fmt.Sprintf("marketdata/%s/pricehistory", symbol)
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, nil, err
		}
		q, err := query.Values(opts)
		if err != nil {
			return nil, nil, err
		}
		u = fmt.Sprintf("%s?%s", u, q.Encode())
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	priceHistory := new(PriceHistory)
	resp, err := s.client.Do(ctx, req, priceHistory)
	if err != nil {
		return nil, resp, err
	}
	if priceHistory.Empty {
		return priceHistory, resp, fmt.Errorf("no data, check time period and/or ticker %s", symbol)
	}
	return priceHistory, resp, nil
}

func (opts *PriceHistoryOptions) validate() error {
	if opts.PeriodType != "" {
		if !contains(opts.PeriodType, validPeriodTypes) {
			return fmt.Errorf("invalid periodType, must have the value of one of the following %v", validPeriodTypes)
		}
	} else {
		opts.PeriodType = defaultPeriodType
	}

	if opts.FrequencyType != "" {
		if !contains(opts.FrequencyType, validFrequencyTypes) {
			return fmt.Errorf("invalid frequencyType, must have the value of one of the following %v", validFrequencyTypes)
		}
	} else {
		opts.PeriodType = defaultFrequencyType
	}

	return nil
}

func contains(s string, lst []string) bool {
	for _, e := range lst {
		if e == s {
			return true
		}
	}
	return false
}
