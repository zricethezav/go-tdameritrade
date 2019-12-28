package tdameritrade

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	baseURL = "https://api.tdameritrade.com/v1/"
)

// A Client manages communication with the TD-Ameritrade API.
type Client struct {
	client *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests. Defaults to the public TD-Ameritrade API, but can be
	// set to any endpoint. This allows for more manageable testing.
	BaseURL *url.URL

	// services used for talking to different parts of the tdameritrade api
	PriceHistory *PriceHistoryService
	Account      *AccountsService
	MarketHours  *MarketHoursService
	Quotes       *QuotesService
	Instrument   *InstrumentService
}

type Response struct {
	*http.Response

	// TODO add additional items if needed
}

// NewClient returns a new TD-Ameritrade API client. If a nil httpClient is
// provided, a new http.Client will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	b, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{client: httpClient, BaseURL: b}
	c.PriceHistory = &PriceHistoryService{client: c}
	c.Account = &AccountsService{client: c}
	c.MarketHours = &MarketHoursService{client: c}
	c.Quotes = &QuotesService{client: c}
	c.Instrument = &InstrumentService{client: c}

	return c, nil
}

func (c *Client) UpdateBaseURL(baseURL string) error {
	b, err := url.Parse(baseURL)
	if err != nil {
		return err
	}
	c.BaseURL = b
	return nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	if ctx == nil {
		return nil, errors.New("context must be non-nil")
	}

	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			return nil, err
		}
	}

	defer resp.Body.Close()

	if err := checkResponse(resp); err != nil {
		return nil, err
	}

	response := newResponse(resp)

	// write to v for that good shit
	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, _ = io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return response, err
}

func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errMsg, _ := ioutil.ReadAll(r.Body)
	return errors.New(string(errMsg))
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	fmt.Println(buf)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}
