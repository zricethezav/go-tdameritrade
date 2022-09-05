package tdameritrade

import (
	"context"
	"fmt"
)

// InstrumentService handles communication with the marketdata related methods of
// the TDAmeritrade API.
//
// TDAmeritrade API docs: https://developer.tdameritrade.com/instruments/apis
type InstrumentService struct {
	client *Client
}

type Instruments map[string]*InstrumentInfo

type InstrumentInfo struct {
	Cusip       string `json:"cusip,omitempty"`
	Symbol      string `json:"symbol"`
	Description string `json:"description,omitempty"`
	Type        string `json:"assetType"` //"'NOT_APPLICABLE' or 'OPEN_END_NON_TAXABLE' or 'OPEN_END_TAXABLE' or 'NO_LOAD_NON_TAXABLE' or 'NO_LOAD_TAXABLE'"
	Exchange    string `json:"exchange"`
}

func (s *InstrumentService) GetInstrument(ctx context.Context, cusip string) (*Instruments, *Response, error) {
	if cusip == "" {
		return nil, nil, fmt.Errorf("no cusip present")
	}
	u := fmt.Sprintf("instruments/%s", cusip)
	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}

	instruments := new(Instruments)

	resp, err := s.client.Do(ctx, req, instruments)
	if err != nil {
		return nil, resp, err
	}
	return nil, nil, nil
}

func (s *InstrumentService) SearchInstruments(ctx context.Context, symbol, projection string) (*Instruments, *Response, error) {
	u := "instruments"
	if symbol == "" {
		return nil, nil, fmt.Errorf("no symbol present")
	}
	if projection == "" {
		projection = "symbol-search"
	}
	u = fmt.Sprintf("%s?symbol=%s&projection=%s", u, symbol, projection)

	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}

	instruments := new(Instruments)

	resp, err := s.client.Do(ctx, req, instruments)
	if err != nil {
		return nil, resp, err
	}

	return instruments, resp, nil
}
