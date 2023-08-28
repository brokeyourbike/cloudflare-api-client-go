package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const defaultBaseURL = "https://api.cloudflare.com/client/v4"

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client interface {
	FetchZeroTrustUsers(ctx context.Context, accountID string, page int) (data FetchZeroTrustUsersResponse, err error)
}

type client struct {
	httpClient HttpClient
	baseUrl    string
	token      string
}

func NewClient(httpClient HttpClient, opts ...ClientOption) *client {
	c := &client{
		httpClient: httpClient,
		baseUrl:    defaultBaseURL,
	}

	for _, o := range opts {
		o.Apply(c)
	}

	return c
}

// FetchZeroTrustUsers returns the list of zero trust users for an account.
//
// API Reference: https://developers.cloudflare.com/api/operations/zero-trust-users-get-users
func (c *client) FetchZeroTrustUsers(ctx context.Context, accountID string, page int) (data FetchZeroTrustUsersResponse, err error) {
	url := fmt.Sprintf("%s/accounts/%s/access/users", c.baseUrl, accountID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return data, fmt.Errorf("cannot create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	q := req.URL.Query()
	q.Add("page", strconv.Itoa(page))
	req.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return data, fmt.Errorf("cannot send request: %w", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return data, fmt.Errorf("failed to decode response: %w", err)
	}

	return data, nil
}
