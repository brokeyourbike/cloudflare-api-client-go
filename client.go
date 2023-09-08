package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"golang.org/x/sync/errgroup"
)

const defaultBaseURL = "https://api.cloudflare.com/client/v4"

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client interface {
	ListZeroTrustUsers(ctx context.Context) ([]ZeroTrustUser, error)
}

var _ Client = (*client)(nil)

type client struct {
	httpClient HttpClient
	baseUrl    string
	token      string
	accountID  string
}

// ClientOption is a function that configures a Client.
type ClientOption func(*client)

func NewClient(token, accountID string, opts ...ClientOption) *client {
	c := &client{
		httpClient: http.DefaultClient,
		baseUrl:    defaultBaseURL,
		token:      token,
		accountID:  accountID,
	}

	for _, option := range opts {
		option(c)
	}

	return c
}

// ListZeroTrustUsers returns the list of zero trust users for an account, automatically handling the pagination.
//
// API Reference: https://developers.cloudflare.com/api/operations/zero-trust-users-get-users
func (c *client) ListZeroTrustUsers(ctx context.Context) ([]ZeroTrustUser, error) {
	r, err := c.fetchZeroTrustUsers(ctx, c.accountID, 1)
	if err != nil {
		return []ZeroTrustUser{}, fmt.Errorf("cannot fetch users: %w", err)
	}

	if r.ResultInfo.TotalCount == r.ResultInfo.Count {
		return r.Result, nil
	}

	pages := r.ResultInfo.TotalCount / r.ResultInfo.PerPage
	users := r.Result

	mu := sync.Mutex{}
	g, ctx := errgroup.WithContext(ctx)

	for page := 2; page <= pages; page++ {
		page := page
		g.Go(func() error {
			r, err := c.fetchZeroTrustUsers(ctx, c.accountID, page)
			if err != nil {
				return err
			}

			mu.Lock()
			defer mu.Unlock()

			users = append(users, r.Result...)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return users, err
	}

	return users, nil
}

func (c *client) fetchZeroTrustUsers(ctx context.Context, accountID string, page int) (data FetchZeroTrustUsersResponse, err error) {
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
