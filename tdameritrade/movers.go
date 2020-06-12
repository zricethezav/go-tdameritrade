package tdameritrade

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
)

const (
	defaultChangeType    = "percent"
	defaultDirectionType = "up"
)

var (
	ChangeTypes    = []string{"value", "percent"}
	DirectionTypes = []string{"up", "down"}
)

type MoverService struct {
	client *Client
}

type MoverOptions struct {
	Direction  string `url:"direction"`
	ChangeType string `url:"change"`
}

type Mover struct {
	Change      float64 `json:"change"`
	Description string  `json:"description"`
	Direction   string  `json:"direction"`
	Last        float64 `json:"last"`
	TotalVolume float64 `json:"totalVolume"`
	Symbol      string  `json:"symbol"`
}

func (s *MoverService) Mover(ctx context.Context, symbol string, opts *MoverOptions) (*[]Mover, *Response, error) {
	u := fmt.Sprintf("marketdata/%s/movers", symbol)
	if opts != nil {
		if err := opts.validate(); err != nil {
			return nil, nil, err
		}

	} else {
		opts = &MoverOptions{
			Direction:  defaultDirectionType,
			ChangeType: defaultChangeType,
		}
	}
	q, err := query.Values(opts)
	if err != nil {
		return nil, nil, err
	}
	u = fmt.Sprintf("%s?%s", u, q.Encode())

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	movers := new([]Mover)
	resp, err := s.client.Do(ctx, req, movers)
	if err != nil {
		return nil, resp, err
	}

	return movers, resp, nil
}

func (opts *MoverOptions) validate() error {
	if opts.ChangeType != "" {
		if !contains(opts.ChangeType, ChangeTypes) {
			return fmt.Errorf("invalid changeType, must have the value of one of the following %v", ChangeTypes)
		}
	} else {
		opts.ChangeType = defaultChangeType
	}

	if opts.Direction != "" {
		if !contains(opts.Direction, DirectionTypes) {
			return fmt.Errorf("invalid direction, must have the value of one of the following %v", DirectionTypes)
		}
	} else {
		opts.ChangeType = defaultDirectionType
	}

	return nil
}
