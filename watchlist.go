package tdameritrade

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// NewWatchlist is a watchlist to be created by a user.
type NewWatchlist struct {
	Name           string          `json:"name"`
	WatchlistItems []WatchlistItem `json:"watchlistItems"`
}

// WatchlistItem is a security to be added to a NewWatchlist.
type WatchlistItem struct {
	Quantity      float64             `json:"quantity"`
	AveragePrice  float64             `json:"averagePrice"`
	Commission    float64             `json:"commission"`
	PurchasedDate string              `json:"purchasedDate"`
	Instrument    WatchlistInstrument `json:"instrument"`
}

// WatchlistInstrument is the specific information about the security being added to the watchlist.
type WatchlistInstrument struct {
	Symbol    string `json:"symbol"`
	AssetType string `json:"assetType"`
}

// UpdateWatchlist is a watchlist used to update an existing watchlist.
type UpdateWatchlist struct {
	Name           string                `json:"name"`
	WatchlistID    string                `json:"watchlistId"`
	WatchlistItems []UpdateWatchlistItem `json:"watchlistItems"`
}

// UpdateWatchlistItem is an item in the user's existing watchlist.
type UpdateWatchlistItem struct {
	SequenceID    int                 `json:"sequenceId"`
	Quantity      float64             `json:"quantity"`
	AveragePrice  float64             `json:"averagePrice"`
	Commission    float64             `json:"commission"`
	PurchasedDate string              `json:"purchasedDate,omitempty"`
	Instrument    WatchlistInstrument `json:"instrument"`
}

// StoredWatchlist is an existing watchlist in a user's account.
type StoredWatchlist struct {
	Name           string                `json:"name"`
	WatchlistID    string                `json:"watchlistId"`
	AccountID      string                `json:"accountId"`
	Status         string                `json:"status"`
	WatchlistItems []StoredWatchlistItem `json:"watchlistItems"`
}

// StoredWatchlistItem is an item in the user's existing watchlist.
type StoredWatchlistItem struct {
	SequenceID    int                       `json:"sequenceId"`
	Quantity      float64                   `json:"quantity"`
	AveragePrice  float64                   `json:"averagePrice"`
	Commission    float64                   `json:"commission"`
	PurchasedDate string                    `json:"purchasedDate"`
	Instrument    StoredWatchlistInstrument `json:"instrument"`
	Status        string                    `json:"status"`
}

// StoredWatchlistInstrument is an investment instrument in a user's watchlist.
type StoredWatchlistInstrument struct {
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
	AssetType   string `json:"assetType"`
}

// WatchlistService allows CRUD operations on watchlists in a user's account.
// See https://developer.tdameritrade.com/watchlist/apis.
type WatchlistService struct {
	client *Client
}

// CreateWatchlist adds a new watchlist to a user's account
// See https://developer.tdameritrade.com/watchlist/apis/post/accounts/%7BaccountId%7D/watchlists-0
func (s *WatchlistService) CreateWatchlist(ctx context.Context, accountID string, newWatchlist *NewWatchlist) (*Response, error) {
	if accountID == "" {
		return nil, fmt.Errorf("accountID cannot be empty")
	}

	u := fmt.Sprintf("accounts/%s/watchlists", accountID)
	req, err := s.client.NewRequest("POST", u, newWatchlist)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DeleteWatchlist removes a watchlist from a user's account
// See https://developer.tdameritrade.com/watchlist/apis/delete/accounts/%7BaccountId%7D/watchlists/%7BwatchlistId%7D-0
func (s *WatchlistService) DeleteWatchlist(ctx context.Context, accountID, watchlistID string) (*Response, error) {
	if accountID == "" {
		return nil, fmt.Errorf("accountID cannot be empty")
	}

	if watchlistID == "" {
		return nil, fmt.Errorf("watchlistID cannot be empty")
	}

	u := fmt.Sprintf("accounts/%s/watchlists/%s", accountID, watchlistID)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// GetWatchlist returns a single watchlist in a user's account
// See https://developer.tdameritrade.com/watchlist/apis/get/accounts/%7BaccountId%7D/watchlists/%7BwatchlistId%7D-0
func (s *WatchlistService) GetWatchlist(ctx context.Context, accountID, watchlistID string) (*StoredWatchlist, *Response, error) {
	if accountID == "" {
		return nil, nil, fmt.Errorf("accountID cannot be empty")
	}

	if watchlistID == "" {
		return nil, nil, fmt.Errorf("watchlistID cannot be empty")
	}

	u := fmt.Sprintf("accounts/%s/watchlists/%s", accountID, watchlistID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	storedWatchlist := new(StoredWatchlist)
	resp, err := s.client.Do(ctx, req, storedWatchlist)
	return storedWatchlist, resp, err
}

// GetAllWatchlists returns all watchlists for all of a user's linked accounts.
// See https://developer.tdameritrade.com/watchlist/apis/get/accounts/watchlists-0
func (s *WatchlistService) GetAllWatchlists(ctx context.Context) (*[]StoredWatchlist, *Response, error) {
	u := "accounts/watchlists"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	buf := new(strings.Builder)
	resp, err := s.client.Do(ctx, req, buf)
	if err != nil {
		return nil, nil, err
	}

	watchlists := new([]StoredWatchlist)
	err = json.Unmarshal([]byte(buf.String()), watchlists)
	return watchlists, resp, err
}

// GetAllWatchlistsForAccount returns all watchlists for a single user account.
// See https://developer.tdameritrade.com/watchlist/apis/get/accounts/%7BaccountId%7D/watchlists-0
func (s *WatchlistService) GetAllWatchlistsForAccount(ctx context.Context, accountID string) (*[]StoredWatchlist, *Response, error) {
	if accountID == "" {
		return nil, nil, fmt.Errorf("accountID cannot be empty")
	}

	u := fmt.Sprintf("accounts/%s/watchlists", accountID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	watchlists := new([]StoredWatchlist)
	resp, err := s.client.Do(ctx, req, watchlists)
	return watchlists, resp, err
}

// ReplaceWatchlist replaces a watchlist in an account with a new watchlist.
// It does not verify that symbols are valid.
// See https://developer.tdameritrade.com/watchlist/apis/put/accounts/%7BaccountId%7D/watchlists/%7BwatchlistId%7D-0
func (s *WatchlistService) ReplaceWatchlist(ctx context.Context, accountID, watchlistID string, newWatchlist *NewWatchlist) (*Response, error) {
	if accountID == "" {
		return nil, fmt.Errorf("accountID cannot be empty")
	}

	if watchlistID == "" {
		return nil, fmt.Errorf("watchlistID cannot be empty")
	}

	u := fmt.Sprintf("accounts/%s/watchlists/%s", accountID, watchlistID)
	req, err := s.client.NewRequest("PUT", u, newWatchlist)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// UpdateWatchlist partially updates watchlist for a specific account.
// Callers can:
//  - change the watchlist's name
//  - add to the beginning/end of a watchlist
//  - update or delete items in a watchlist
// This method does not verify that the symbol or asset type are valid.
// See https://developer.tdameritrade.com/watchlist/apis/patch/accounts/%7BaccountId%7D/watchlists/%7BwatchlistId%7D-0
func (s *WatchlistService) UpdateWatchlist(ctx context.Context, accountID string, updateWatchlist *UpdateWatchlist) (*Response, error) {
	if accountID == "" {
		return nil, fmt.Errorf("accountID cannot be empty")
	}

	watchlistID := updateWatchlist.WatchlistID
	if watchlistID == "" {
		return nil, fmt.Errorf("watchlistID cannot be empty")
	}

	u := fmt.Sprintf("accounts/%s/watchlists/%s", accountID, watchlistID)
	req, err := s.client.NewRequest("PATCH", u, updateWatchlist)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
