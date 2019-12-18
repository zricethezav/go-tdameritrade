package tdameritrade

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
)
const (
	baseURL="https://api.tdameritrade.com/v1"
)

// A Client manages communication with the TD-Ameritrade API.
type Client struct {
	client *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests. Defaults to the public TD-Ameritrade API, but can be
	// set to any endpoint. This allows for more manageable testing.
	BaseURL *url.URL
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

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
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

	if err := checkResponse(resp); err != nil  {
		return nil, err
	}


	// write to v for that good shit
	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, _ = io.Copy(w, resp.Body)
		} else {
			// is not an io writer
		}
	}

	return resp, nil
}

func checkResponse(resp *http.Response) error {
	return nil
}
