package tdameritrade

import (
	"context"
	"fmt"
	"strings"
)

type Preferences struct {
	ExpressTrading                   bool   `json:"expressTrading"`
	DirectOptionsRouting             bool   `json:"directOptionsRouting"`
	DirectEquityRouting              bool   `json:"directEquityRouting"`
	DefaultEquityOrderLegInstruction string `json:"defaultEquityOrderLegInstruction"`
	DefaultEquityOrderType           string `json:"defaultEquityOrderType"`
	DefaultEquityOrderPriceLinkType  string `json:"defaultEquityOrderPriceLinkType"`
	DefaultEquityOrderDuration       string `json:"defaultEquityOrderDuration"`
	DefaultEquityOrderMarketSession  string `json:"defaultEquityOrderMarketSession"`
	DefaultEquityQuantity            int    `json:"defaultEquityQuantity"`
	MutualFundTaxLotMethod           string `json:"mutualFundTaxLotMethod"`
	OptionTaxLotMethod               string `json:"optionTaxLotMethod"`
	EquityTaxLotMethod               string `json:"equityTaxLotMethod"`
	DefaultAdvancedToolLaunch        string `json:"defaultAdvancedToolLaunch"`
	AuthTokenTimeout                 string `json:"authTokenTimeout"`
}

type StreamerSubscriptionKeys struct {
	Keys []KeyEntry `json:"keys"`
}

type KeyEntry struct {
	Key string `json:"key"`
}

type UserPrincipal struct {
	AuthToken                string                   `json:"authToken"`
	UserID                   string                   `json:"userId"`
	UserCdDomainID           string                   `json:"userCdDomainId"`
	PrimaryAccountID         string                   `json:"primaryAccountId"`
	LastLoginTime            string                   `json:"lastLoginTime"`
	TokenExpirationTime      string                   `json:"tokenExpirationTime"`
	LoginTime                string                   `json:"loginTime"`
	AccessLevel              string                   `json:"accessLevel"`
	StalePassword            bool                     `json:"stalePassword"`
	StreamerInfo             StreamerInfo             `json:"streamerInfo"`
	ProfessionalStatus       string                   `json:"professionalStatus"`
	Quotes                   QuoteDelays              `json:"quotes"`
	StreamerSubscriptionKeys StreamerSubscriptionKeys `json:"streamerSubscriptionKeys"`
	Accounts                 []UserAccountInfo        `json:"accounts"`
}

type UserAccountInfo struct {
	AccountID         string         `json:"accountId"`
	Description       string         `json:"description"`
	DisplayName       string         `json:"displayName"`
	AccountCdDomainID string         `json:"accountCdDomainId"`
	Company           string         `json:"company"`
	Segment           string         `json:"segment"`
	SurrogateIds      string         `json:"surrogateIds"`
	Preferences       Preferences    `json:"preferences"`
	ACL               string         `json:"acl"`
	Authorizations    Authorizations `json:"authorizations"`
}

type Authorizations struct {
	Apex               bool   `json:"apex"`
	LevelTwoQuotes     bool   `json:"levelTwoQuotes"`
	StockTrading       bool   `json:"stockTrading"`
	MarginTrading      bool   `json:"marginTrading"`
	StreamingNews      bool   `json:"streamingNews"`
	OptionTradingLevel string `json:"optionTradingLevel"`
	StreamerAccess     bool   `json:"streamerAccess"`
	AdvancedMargin     bool   `json:"advancedMargin"`
	ScottradeAccount   bool   `json:"scottradeAccount"`
}

type StreamerInfo struct {
	StreamerBinaryURL string `json:"streamerBinaryUrl"`
	StreamerSocketURL string `json:"streamerSocketUrl"`
	Token             string `json:"token"`
	TokenTimestamp    string `json:"tokenTimestamp"`
	UserGroup         string `json:"userGroup"`
	AccessLevel       string `json:"accessLevel"`
	ACL               string `json:"acl"`
	AppID             string `json:"appId"`
}

type QuoteDelays struct {
	IsNyseDelayed   bool `json:"isNyseDelayed"`
	IsNasdaqDelayed bool `json:"isNasdaqDelayed"`
	IsOpraDelayed   bool `json:"isOpraDelayed"`
	IsAmexDelayed   bool `json:"isAmexDelayed"`
	IsCmeDelayed    bool `json:"isCmeDelayed"`
	IsIceDelayed    bool `json:"isIceDelayed"`
	IsForexDelayed  bool `json:"isForexDelayed"`
}

// UserService exposes operations on a user's preferences.
// See https://developer.tdameritrade.com/user-principal/apis.
type UserService struct {
	client *Client
}

// GetPreferences returns Preferences for a specific account.
func (s *UserService) GetPreferences(ctx context.Context, accountID string) (*Preferences, *Response, error) {
	u := fmt.Sprintf("accounts/%s/preferences", accountID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	preferences := new(Preferences)
	resp, err := s.client.Do(ctx, req, preferences)
	if err != nil {
		return nil, resp, err
	}

	return preferences, resp, err
}

// GetStreamerSubscriptionKeys returns Subscription Keys for provided accounts or default accounts.
func (s *UserService) GetStreamerSubscriptionKeys(ctx context.Context, accountIDs ...string) (*StreamerSubscriptionKeys, *Response, error) {
	u := fmt.Sprintf("userprincipals/streamersubscriptionkeys?accountIds=%s", strings.Join(accountIDs, ","))
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	streamerSubscriptionKeys := new(StreamerSubscriptionKeys)
	resp, err := s.client.Do(ctx, req, streamerSubscriptionKeys)
	if err != nil {
		return nil, resp, err
	}

	return streamerSubscriptionKeys, resp, err
}

// GetUserPrincipals returns User Principal details.
func (s *UserService) GetUserPrincipals(ctx context.Context, fields ...string) (*UserPrincipal, *Response, error) {
	u := "userpricipals"
	if len(fields) > 0 {
		u = fmt.Sprintf("%s?fields=%s", u, strings.Join(fields, ","))
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	userPrincipal := new(UserPrincipal)
	resp, err := s.client.Do(ctx, req, userPrincipal)
	if err != nil {
		return nil, resp, err
	}

	return userPrincipal, resp, err
}

// UpdatePreferences updates Preferences for a specific account.
// Please note that the directOptionsRouting and directEquityRouting values cannot be modified via this operation
func (s *UserService) UpdatePreferences(ctx context.Context, accountID string, newPreferences *Preferences) (*Response, error) {
	if newPreferences == nil {
		return nil, fmt.Errorf("newPreferences is nil")
	}

	u := fmt.Sprintf("accounts/%s/preferences", accountID)
	req, err := s.client.NewRequest("PUT", u, newPreferences)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
