package cloudflare_test

import (
	"bytes"
	"context"
	_ "embed"
	"io"
	"net/http"
	"testing"

	"github.com/brokeyourbike/cloudflare-api-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/success-one.json
var successOneResponse []byte

//go:embed testdata/success-multiple-page1.json
var successMultiplePage1Response []byte

//go:embed testdata/success-multiple-page2.json
var successMultiplePage2Response []byte

func TestListAccessUsers_RequestErr(t *testing.T) {
	mockHttpClient := cloudflare.NewMockHttpClient(t)
	client := cloudflare.NewClient("token123", "zone1", "account456", cloudflare.WithHTTPClient(mockHttpClient))

	_, err := client.ListZeroTrustUsers(nil) //lint:ignore SA1012 testing failure
	require.Error(t, err)
	require.Contains(t, err.Error(), "cannot create request")
}

func TestListAccessUsers_One(t *testing.T) {
	mockHttpClient := cloudflare.NewMockHttpClient(t)
	client := cloudflare.NewClient("token123", "zone1", "account456", cloudflare.WithHTTPClient(mockHttpClient), cloudflare.WithBaseURL("https://c.om"))

	ctx := context.Background()

	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewReader(successOneResponse)),
	}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.ListZeroTrustUsers(ctx)
	assert.NoError(t, err)
	assert.Len(t, got, 1)
}

func TestListAccessUsers_Multiple(t *testing.T) {
	mockHttpClient := cloudflare.NewMockHttpClient(t)
	client := cloudflare.NewClient("token123", "zone1", "account456", cloudflare.WithHTTPClient(mockHttpClient))

	ctx := context.Background()

	resp1 := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewReader(successMultiplePage1Response)),
	}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Twice().Return(resp1, nil).Once()

	resp2 := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewReader(successMultiplePage2Response)),
	}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp2, nil).Once()

	got, err := client.ListZeroTrustUsers(ctx)
	require.NoError(t, err)
	require.Len(t, got, 2)
	assert.Equal(t, got[0].Email, "john@doe.com")
	assert.Equal(t, got[1].Email, "jane@doe.com")
}

func TestPurge(t *testing.T) {
	mockHttpClient := cloudflare.NewMockHttpClient(t)
	client := cloudflare.NewClient("token123", "zone1", "account456", cloudflare.WithHTTPClient(mockHttpClient), cloudflare.WithBaseURL("https://c.om"))

	ctx := context.Background()

	resp := &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(nil))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	err := client.Purge(ctx)
	assert.NoError(t, err)
}

func TestPurge_Err(t *testing.T) {
	mockHttpClient := cloudflare.NewMockHttpClient(t)
	client := cloudflare.NewClient("token123", "zone1", "account456", cloudflare.WithHTTPClient(mockHttpClient), cloudflare.WithBaseURL("https://c.om"))

	ctx := context.Background()

	resp := &http.Response{StatusCode: 400, Body: io.NopCloser(bytes.NewReader(nil))}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	err := client.Purge(ctx)
	assert.ErrorContains(t, err, "unexpected status code")
}
