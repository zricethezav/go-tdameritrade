package tdameritrade

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Accounts []*Account

type Account struct {
	SecuritiesAccount `json:"securitiesAccount"`
}

type _Instrument Instrument

type Instrument struct {
	AssetType string `json:"assetType"`
	Data      interface{}
}

type OptionDeliverable struct {
	Symbol           string  `json:"symbol"`
	DeliverableUnits float64 `json:"deliverableUnits"`
	CurrencyType     string  `json:"currencyType"`
	AssetType        string  `json:"assetType"`
}

type OptionA struct {
	Cusip              string               `json:"cusip,omitempty"`
	Symbol             string               `json:"symbol"`
	Description        string               `json:"description,omitempty"`
	Type               string               `json:"type"`
	PutCall            string               `json:"putCall"`
	UnderlyingSymbol   string               `json:"underlyingSymbol"`
	OptionMultiplier   float64              `json:"optionMultiplier"`
	OptionDeliverables []*OptionDeliverable `json:"optionDeliverables"`
}

type MutualFund struct {
	Cusip       string `json:"cusip,omitempty"`
	Symbol      string `json:"symbol"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type"` //"'NOT_APPLICABLE' or 'OPEN_END_NON_TAXABLE' or 'OPEN_END_TAXABLE' or 'NO_LOAD_NON_TAXABLE' or 'NO_LOAD_TAXABLE'"
}

type CashEquivalent struct {
	Cusip       string `json:"cusip,omitempty"`
	Symbol      string `json:"symbol"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type"` //"'SAVINGS' or 'MONEY_MARKET_FUND'"
}

type Equity struct {
	Cusip       string `json:"cusip,omitempty"`
	Symbol      string `json:"symbol"`
	Description string `json:"description,omitempty"`
}

type FixedIncome struct {
	Cusip        string  `json:"cusip"`
	Symbol       string  `json:"symbol"`
	Description  string  `json:"description"`
	MaturityDate string  `json:"maturityDate"`
	VariableRate float64 `json:"variableRate"`
	Factor       float64 `json:"factor"`
}

type SecuritiesAccount struct {
	Type                    string  `json:"type"`
	AccountID               string  `json:"accountId"`
	RoundTrips              float64 `json:"roundTrips"`
	IsDayTrader             bool    `json:"isDayTrader"`
	IsClosingOnlyRestricted bool    `json:"isClosingOnlyRestricted"`
	Positions               []struct {
		ShortQuantity                  float64    `json:"shortQuantity"`
		AveragePrice                   float64    `json:"averagePrice"`
		CurrentDayProfitLoss           float64    `json:"currentDayProfitLoss"`
		CurrentDayProfitLossPercentage float64    `json:"currentDayProfitLossPercentage"`
		LongQuantity                   float64    `json:"longQuantity"`
		SettledLongQuantity            float64    `json:"settledLongQuantity"`
		SettledShortQuantity           float64    `json:"settledShortQuantity"`
		AgedQuantity                   float64    `json:"agedQuantity"`
		Instrument                     Instrument `json:"instrument"`
		MarketValue                    float64    `json:"marketValue"`
	} `json:"positions"`
	OrderStrategies []struct {
		Session    string `json:"session"`
		Duration   string `json:"duration"`
		OrderType  string `json:"orderType"`
		CancelTime struct {
			Date        string `json:"date"`
			ShortFormat bool   `json:"shortFormat"`
		} `json:"cancelTime"`
		ComplexOrderStrategyType string  `json:"complexOrderStrategyType"`
		Quantity                 float64 `json:"quantity"`
		FilledQuantity           float64 `json:"filledQuantity"`
		RemainingQuantity        float64 `json:"remainingQuantity"`
		RequestedDestination     string  `json:"requestedDestination"`
		DestinationLinkName      string  `json:"destinationLinkName"`
		ReleaseTime              string  `json:"releaseTime"`
		StopPrice                float64 `json:"stopPrice"`
		StopPriceLinkBasis       string  `json:"stopPriceLinkBasis"`
		StopPriceLinkType        string  `json:"stopPriceLinkType"`
		StopPriceOffset          float64 `json:"stopPriceOffset"`
		StopType                 string  `json:"stopType"`
		PriceLinkBasis           string  `json:"priceLinkBasis"`
		PriceLinkType            string  `json:"priceLinkType"`
		Price                    float64 `json:"price"`
		TaxLotMethod             string  `json:"taxLotMethod"`
		OrderLegCollection       []struct {
			OrderLegType   string  `json:"orderLegType"`
			LegID          int64   `json:"legId"`
			Instrument     string  `json:"instrument"`
			Instruction    string  `json:"instruction"`
			PositionEffect string  `json:"positionEffect"`
			Quantity       float64 `json:"quantity"`
			QuantityType   string  `json:"quantityType"`
		} `json:"orderLegCollection"`
		ActivationPrice          float64  `json:"activationPrice"`
		SpecialInstruction       string   `json:"specialInstruction"`
		OrderStrategyType        string   `json:"orderStrategyType"`
		OrderID                  int64    `json:"orderId"`
		Cancelable               bool     `json:"cancelable"`
		Editable                 bool     `json:"editable"`
		Status                   string   `json:"status"`
		EnteredTime              string   `json:"enteredTime"`
		CloseTime                string   `json:"closeTime"`
		Tag                      string   `json:"tag"`
		AccountID                int64    `json:"accountId"`
		OrderActivityCollection  []string `json:"orderActivityCollection"`
		ReplacingOrderCollection []struct {
		} `json:"replacingOrderCollection"`
		ChildOrderStrategies []struct {
		} `json:"childOrderStrategies"`
		StatusDescription string `json:"statusDescription"`
	} `json:"orderStrategies"`
	InitialBalances   Balance `json:"initialBalances"`
	CurrentBalances   Balance `json:"currentBalances"`
	ProjectedBalances Balance `json:"projectedBalances"`
}

type Balance struct {
	AccruedInterest              float64 `json:"accruedInterest"`
	CashBalance                  float64 `json:"cashBalance"`
	CashReceipts                 float64 `json:"cashReceipts"`
	LongOptionMarketValue        float64 `json:"longOptionMarketValue"`
	LiquidationValue             float64 `json:"liquidationValue"`
	LongMarketValue              float64 `json:"longMarketValue"`
	MoneyMarketFund              float64 `json:"moneyMarketFund"`
	Savings                      float64 `json:"savings"`
	ShortMarketValue             float64 `json:"shortMarketValue"`
	PendingDeposits              float64 `json:"pendingDeposits"`
	CashAvailableForTrading      float64 `json:"cashAvailableForTrading"`
	CashAvailableForWithdrawal   float64 `json:"cashAvailableForWithdrawal"`
	CashCall                     float64 `json:"cashCall"`
	LongNonMarginableMarketValue float64 `json:"longNonMarginableMarketValue"`
	TotalCash                    float64 `json:"totalCash"`
	ShortOptionMarketValue       float64 `json:"shortOptionMarketValue"`
	MutualFundValue              float64 `json:"mutualFundValue"`
	BondValue                    float64 `json:"bondValue"`
	CashDebitCallValue           float64 `json:"cashDebitCallValue"`
	UnsettledCash                float64 `json:"unsettledCash"`
}

type OrderLegCollection struct {
	OrderLegType   string     `json:"orderLegType,omitempty"`
	LegID          int        `json:"legId,omitempty"`
	Instrument     Instrument `json:"instrument"`
	Instruction    string     `json:"instruction"`
	PositionEffect string     `json:"positionEffect,omitempty"`
	Quantity       float64    `json:"quantity"`
	QuantityType   string     `json:"quantityType,omitempty"`
}

type CancelTime struct {
	Date        string `json:"date,omitempty"`
	ShortFormat bool   `json:"shortFormat,omitempty"`
}

type Orders []Order

type Order struct {
	Session                  string                `json:"session"`
	Duration                 string                `json:"duration"`
	OrderType                string                `json:"orderType"`
	CancelTime               *CancelTime           `json:"cancelTime,omitempty"`
	ComplexOrderStrategyType string                `json:"complexOrderStrategyType,omitempty"`
	Quantity                 float64               `json:"quantity,omitempty"`
	FilledQuantity           float64               `json:"filledQuantity,omitempty"`
	RemainingQuantity        float64               `json:"remainingQuantity,omitempty"`
	RequestedDestination     string                `json:"requestedDestination,omitempty"`
	DestinationLinkName      string                `json:"destinationLinkName,omitempty"`
	ReleaseTime              string                `json:"releaseTime,omitempty"`
	StopPrice                float64               `json:"stopPrice,omitempty"`
	StopPriceLinkBasis       string                `json:"stopPriceLinkBasis,omitempty"`
	StopPriceLinkType        string                `json:"stopPriceLinkType,omitempty"`
	StopPriceOffset          float64               `json:"stopPriceOffset,omitempty"`
	StopType                 string                `json:"stopType,omitempty"`
	PriceLinkBasis           string                `json:"priceLinkBasis,omitempty"`
	PriceLinkType            string                `json:"priceLinkType,omitempty"`
	Price                    float64               `json:"price,omitempty"`
	TaxLotMethod             string                `json:"taxLotMethod,omitempty"`
	OrderLegCollection       []*OrderLegCollection `json:"orderLegCollection"`
	ActivationPrice          float64               `json:"activationPrice,omitempty"`
	SpecialInstruction       string                `json:"specialInstruction,omitempty"`
	OrderStrategyType        string                `json:"orderStrategyType"`
	OrderID                  int64                 `json:"orderId,omitempty"`
	Cancelable               bool                  `json:"cancelable,omitempty"`
	Editable                 bool                  `json:"editable,omitempty"`
	Status                   string                `json:"status,omitempty"`
	EnteredTime              string                `json:"enteredTime,omitempty"`
	CloseTime                string                `json:"closeTime,omitempty"`
	Tag                      string                `json:"tag,omitempty"`
	AccountID                float64               `json:"accountId,omitempty"`
	OrderActivityCollection  []*Execution          `json:"orderActivityCollection,omitempty"`
	ReplacingOrderCollection []*Order              `json:"replacingOrderCollection,omitempty"`
	ChildOrderStrategies     []*Order              `json:"childOrderStrategies,omitempty"`
	StatusDescription        string                `json:"statusDescription,omitempty"`
}

type ExecutionLeg struct {
	LegID             int64   `json:"legId"`
	Quantity          float64 `json:"quantity"`
	MismarkedQuantity float64 `json:"mismarkedQuantity"`
	Price             float64 `json:"price"`
	Time              string  `json:"time"`
}

type Execution struct {
	ActivityType           string          `json:"activityType"`  //"'EXECUTION' or 'ORDER_ACTION'",
	ExecutionType          string          `json:"executionType"` //"'FILL'",
	Quantity               float64         `json:"quantity"`
	OrderRemainingQuantity float64         `json:"orderRemainingQuantity"`
	ExecutionLegs          []*ExecutionLeg `json:"executionLegs"`
}

// AccountsService handles communication with the account related methods of
// the TDAmeritrade API.
//
// TDAmeritrade API docs: https://developer.tdameritrade.com/account-access/apis
type AccountsService struct {
	client *Client
}

type AccountOptions struct {
	Position bool
	Orders   bool
}

type OrderParams struct {
	MaxResults int
	From       time.Time
	To         time.Time
	Status     string
}

func (i *Instrument) UnmarshalJSON(bs []byte) (err error) {
	instrument := _Instrument{}

	err = json.Unmarshal(bs, &instrument)
	if err != nil {
		return err
	}

	switch instrument.AssetType {
	case "EQUITY":
		instrument.Data = &Equity{}
	case "OPTION":
		instrument.Data = &OptionA{}
	case "MUTUAL_FUND":
		instrument.Data = &MutualFund{}
	case "CASH_EQUIVALENT":
		instrument.Data = &CashEquivalent{}
	case "FIXED_INCOME":
		instrument.Data = &FixedIncome{}
	default:
		return fmt.Errorf("unsupported type %s", instrument.AssetType)
	}
	err = json.Unmarshal(bs, instrument.Data)
	*i = Instrument(instrument)

	return err
}

func (i *Instrument) MarshalJSON() ([]byte, error) {
	switch data := i.Data.(type) {
	case *Equity:
		return json.Marshal(&struct {
			AssetType string `json:"assetType"`
			*Equity
		}{
			AssetType: i.AssetType,
			Equity:    data,
		})
	case *OptionA:
		return json.Marshal(&struct {
			AssetType string `json:"assetType"`
			*OptionA
		}{
			AssetType: i.AssetType,
			OptionA:   data,
		})
	case *MutualFund:
		return json.Marshal(&struct {
			AssetType string `json:"assetType"`
			*MutualFund
		}{
			AssetType:  i.AssetType,
			MutualFund: data,
		})
	case *CashEquivalent:
		return json.Marshal(&struct {
			AssetType string `json:"assetType"`
			*CashEquivalent
		}{
			AssetType:      i.AssetType,
			CashEquivalent: data,
		})
	case *FixedIncome:
		return json.Marshal(&struct {
			AssetType string `json:"assetType"`
			*FixedIncome
		}{
			AssetType:   i.AssetType,
			FixedIncome: data,
		})
	default:
		return nil, fmt.Errorf("unexpected type %T: %v", data, data)
	}
}

func (s *AccountsService) GetAccounts(ctx context.Context, opts *AccountOptions) (*Accounts, *Response, error) {
	u := "accounts"
	if opts != nil {
		if opts.Position {
			u = fmt.Sprintf("%s?fields=%s", u, "positions")
		}
		if opts.Orders {
			u = fmt.Sprintf("%s,%s", u, "orders")
		}
	}
	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err

	}
	accounts := new(Accounts)
	resp, err := s.client.Do(ctx, req, accounts)
	if err != nil {
		return nil, resp, err
	}
	return accounts, resp, err
}

func (s *AccountsService) GetAccount(ctx context.Context, accountID string, opts *AccountOptions) (*Account, *Response, error) {
	u := fmt.Sprintf("accounts/%s", accountID)
	if opts != nil {
		if opts.Position {
			u = fmt.Sprintf("%s?fields=%s", u, "positions")
		}
		if opts.Orders {
			u = fmt.Sprintf("%s,%s", u, "orders")
		}
	}
	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err

	}
	account := new(Account)
	resp, err := s.client.Do(ctx, req, account)
	if err != nil {
		return nil, resp, err
	}
	return account, resp, err
}

func (s *AccountsService) PlaceOrder(ctx context.Context, accountID string, order *Order) (*Response, error) {
	u := fmt.Sprintf("accounts/%s/orders", accountID)
	if order == nil {
		return nil, fmt.Errorf("order is nil")
	}

	req, err := s.client.NewRequest("POST", u, order)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return s.client.Do(ctx, req, nil)
}

func (s *AccountsService) CancelOrder(ctx context.Context, accountID, orderID string) (*Response, error) {
	u := fmt.Sprintf("accounts/%s/orders/%s", accountID, orderID)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

func (s *AccountsService) ReplaceOrder(ctx context.Context, accountID string, orderID string, order *Order) (*Response, error) {
	u := fmt.Sprintf("accounts/%s/orders/%s", accountID, orderID)
	if order == nil {
		return nil, fmt.Errorf("order is nil")
	}

	req, err := s.client.NewRequest("PUT", u, order)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

func (s *AccountsService) GetOrder(ctx context.Context, accountID, orderID string) (*Response, error) {
	u := fmt.Sprintf("accounts/%s/orders/%s", accountID, orderID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

func (s *AccountsService) GetOrderByPath(ctx context.Context, accountID string, orderParams *OrderParams) (*Orders, *Response, error) {
	u := fmt.Sprintf("accounts/%s/orders", accountID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	orders := new(Orders)

	resp, err := s.client.Do(ctx, req, orders)
	if err != nil {
		return nil, resp, err
	}

	return orders, resp, nil
}

func (s *AccountsService) GetOrderByQuery(ctx context.Context, accountID string, orderParams *OrderParams) (*Orders, *Response, error) {
	u := fmt.Sprintf("accounts/%s/orders", accountID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	orders := new(Orders)

	resp, err := s.client.Do(ctx, req, orders)
	if err != nil {
		return nil, resp, err
	}

	return orders, resp, nil
}

func (s *AccountsService) CreateSavedOrder(ctx context.Context, accountID string, order *Order) (*Response, error) {
	u := fmt.Sprintf("accounts/%s/savedorders", accountID)
	if order == nil {
		return nil, fmt.Errorf("order is nil")
	}

	req, err := s.client.NewRequest("POST", u, order)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return s.client.Do(ctx, req, nil)
}

func (s *AccountsService) DeleteSavedOrder(ctx context.Context, accountID, savedOrderID string) (*Response, error) {
	u := fmt.Sprintf("accounts/%s/savedorders/%s", accountID, savedOrderID)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

func (s *AccountsService) GetSavedOrder(ctx context.Context, accountID, savedOrderID string, orderParams *OrderParams) (*Response, error) {
	u := fmt.Sprintf("accounts/%s/savedorders/%s", accountID, savedOrderID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

func (s *AccountsService) ReplaceSavedOrder(ctx context.Context, accountID, savedOrderID string, order *Order) (*Response, error) {
	u := fmt.Sprintf("accounts/%s/savedorders/%s", accountID, savedOrderID)
	if order == nil {
		return nil, fmt.Errorf("order is nil")
	}

	req, err := s.client.NewRequest("PUT", u, order)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return s.client.Do(ctx, req, nil)
}
