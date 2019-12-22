package tdameritrade

import (
	"context"
	"encoding/json"
	"fmt"
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
		AccountID                int64    `json:"accountId, string"`
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
