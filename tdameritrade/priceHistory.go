package tdameritrade

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	"time"
)

var (
	validPeriodTypes = []string{"day", "month", "year", "ytd"}
	validFrequencyTypes = []string{"minute", "daily", "weekly", "monthly"}
)

const (
	defaultPeriodType = "day"
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
	PeriodType            string    `url:"periodType"`
	Period                int       `url:"period"`
	FrequencyType         string    `url:"frequencyType"`
	Frequency             int       `url:"frequency"`
	EndDate               time.Time `url:"endDate"`
	StartDate             time.Time `url:"startDate"`
	NeedExtendedHoursData *bool      `url:"needExtendedHoursData"`
}

type PriceHistory struct {
	Candles []struct {
		Close    int `json:"close"`
		Datetime int `json:"datetime"`
		High     int `json:"high"`
		Low      int `json:"low"`
		Open     int `json:"open"`
		Volume   int `json:"volume"`
	} `json:"candles"`
	Empty  bool   `json:"empty"`
	Symbol string `json:"symbol"`
}

// PriceHistory get the price history for a symbol
// TDAmeritrade API Docs: https://developer.tdameritrade.com/price-history/apis/get/marketdata/%7Bsymbol%7D/pricehistory
func (s *PriceHistoryService) PriceHistory(ctx context.Context, opts PriceHistoryOptions) (*PriceHistory, *Response, error) {
	if err := opts.validate(); err != nil {
		return nil, nil, err
	}
	q, err := query.Values(opts)
	if err != nil {
		return nil, nil, err
	}
	// u.RawQuery = q.Encode()



	return nil, nil, nil
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
