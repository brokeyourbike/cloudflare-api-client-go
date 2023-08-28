package cloudflare

import (
	"strings"
)

// WithHTTPClient sets the HTTP client for the Cloudflare API client.
func WithHTTPClient(c HttpClient) ClientOption {
	return func(target *client) {
		target.httpClient = c
	}
}

// WithBaseURL sets the base URL for the Cloudflare API client.
func WithBaseURL(baseUrl string) ClientOption {
	return func(target *client) {
		target.baseUrl = strings.TrimSuffix(baseUrl, "/")
	}
}
