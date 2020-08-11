package tdameritrade

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

type testStore struct{}

func (s *testStore) StoreToken(token *oauth2.Token, w http.ResponseWriter, req *http.Request) error {
	return nil
}

func (s *testStore) GetToken(req *http.Request) (*oauth2.Token, error) {
	return &oauth2.Token{
		AccessToken:  "ACCESSTOKEN",
		RefreshToken: "REFRESHTOKEN",
		Expiry:       time.Now().AddDate(0, 0, 1),
	}, nil
}

func (s *testStore) StoreState(state string, w http.ResponseWriter, req *http.Request) error {
	return nil
}

func (s *testStore) GetState(req *http.Request) (string, error) {
	return "state", nil
}

type testEmptyStore struct{}

func (s *testEmptyStore) StoreToken(token *oauth2.Token, w http.ResponseWriter, req *http.Request) error {
	return nil
}

func (s *testEmptyStore) GetToken(req *http.Request) (*oauth2.Token, error) {
	return &oauth2.Token{
		AccessToken:  "ACCESSTOKEN",
		RefreshToken: "REFRESHTOKEN",
		Expiry:       time.Now().AddDate(0, 0, 1),
	}, nil
}

func (s *testEmptyStore) StoreState(state string, w http.ResponseWriter, req *http.Request) error {
	return nil
}

func (s *testEmptyStore) GetState(req *http.Request) (string, error) {
	return "", nil
}

func TestStartOAuth2ReturnsCorrectURL(t *testing.T) {
	authenticator := &Authenticator{
		Store: &testEmptyStore{},
		OAuth2: oauth2.Config{
			ClientID: "CLIENTID",
			Endpoint: oauth2.Endpoint{
				TokenURL: "https://api.tdameritrade.com/v1/oauth2/token",
				AuthURL:  "https://auth.tdameritrade.com/auth",
			},
			RedirectURL: "https://localhost:8080/callback",
		},
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/callback?code=code&state=state", nil)
	u, err := authenticator.StartOAuth2Flow(w, req)
	if err != nil {
		t.Fatalf(err.Error())
	}

	expectedURL, _ := url.Parse("https://auth.tdameritrade.com/auth?client_id=CLIENTID&redirect_uri=https%3A%2F%2Flocalhost%3A8080%2Fcallback&response_type=code")
	actualURL, _ := url.Parse(u)

	// State query param is different each invocation.
	// This cannot be changed by design.
	modifiedQuery := actualURL.Query()
	modifiedQuery.Del("state")
	actualURL.RawQuery = modifiedQuery.Encode()
	if actualURL.String() != expectedURL.String() {
		t.Fatalf("invalid auth code URL. expected: '%v', got: '%v'", expectedURL, actualURL)
	}
}

func TestFinishOAuth2RejectsEmptyState(t *testing.T) {
	authenticator := &Authenticator{
		Store: &testEmptyStore{},
		OAuth2: oauth2.Config{
			ClientID: "CLIENTID",
			Endpoint: oauth2.Endpoint{
				TokenURL: "https://api.tdameritrade.com/v1/oauth2/token",
				AuthURL:  "https://auth.tdameritrade.com/auth",
			},
			RedirectURL: "https://localhost:8080/callback",
		},
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/callback?code=code&state=state", nil)
	ctx := context.Background()
	client, err := authenticator.FinishOAuth2Flow(ctx, w, req)
	if err == nil {
		t.Fatalf("empty state not rejected.")
	}

	if err.Error() != "missing state in request from TD Ameritrade" {
		t.Fatalf("unexpected error: %v", err)
	}

	if client != nil {
		t.Fatalf("client returned despite empty state.")
	}
}

func TestFinishOAuth2RejectsInvalidState(t *testing.T) {
	authenticator := &Authenticator{
		Store: &testStore{},
		OAuth2: oauth2.Config{
			ClientID: "CLIENTID",
			Endpoint: oauth2.Endpoint{
				TokenURL: "https://api.tdameritrade.com/v1/oauth2/token",
				AuthURL:  "https://auth.tdameritrade.com/auth",
			},
			RedirectURL: "https://localhost:8080/callback",
		},
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/callback?code=code&state=invalid", nil)
	ctx := context.Background()
	client, err := authenticator.FinishOAuth2Flow(ctx, w, req)
	if err == nil {
		t.Fatalf("invalid state not rejected.")
	}

	if err.Error() != "invalid state. expected: 'state', got 'invalid'" {
		t.Fatalf("unexpected error: %v", err)
	}

	if client != nil {
		t.Fatalf("client returned despite invalid state.")
	}
}

func TestFinishOAuth2RejectsEmptyCode(t *testing.T) {
	authenticator := &Authenticator{
		Store: &testStore{},
		OAuth2: oauth2.Config{
			ClientID: "CLIENTID",
			Endpoint: oauth2.Endpoint{
				TokenURL: "https://api.tdameritrade.com/v1/oauth2/token",
				AuthURL:  "https://auth.tdameritrade.com/auth",
			},
			RedirectURL: "https://localhost:8080/callback",
		},
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/callback?code=&state=state", nil)
	ctx := context.Background()
	client, err := authenticator.FinishOAuth2Flow(ctx, w, req)
	if err == nil {
		t.Fatalf("empty code not rejected.")
	}

	if err.Error() != "missing code in request from TD Ameritrade" {
		t.Fatalf("unexpected error: %v", err)
	}

	if client != nil {
		t.Fatalf("client returned despite empty code.")
	}
}
