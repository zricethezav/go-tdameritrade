package tdameritrade

import (
	"context"
)

// InstrumentService handles communication with the marketdata related methods of
// the TDAmeritrade API.
//
// TDAmeritrade API docs: https://developer.tdameritrade.com/instruments/apis
type InstrumentService struct {
	client *Client
}

func (s *MarketHoursService) GetInstrument(ctx context.Context, cusip string) (*Instrument, *Response, error) {
	// TODO
	return nil, nil, nil
}

func (s *MarketHoursService) SearchInstruments(ctx context.Context, symbol, projection string) (*Instrument, *Response, error) {
	// TODO
	return nil, nil, nil
}
