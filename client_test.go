package cloudflare

import (
	"bytes"
	"context"
	_ "embed"
	"io"
	"net/http"
	"testing"

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

func TestListAccessUsers_One(t *testing.T) {
	mockHttpClient := NewMockHttpClient(t)
	client := NewClient(mockHttpClient, WithToken("token123"), WithBaseURL("https://example.com"))

	ctx := context.Background()

	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewReader(successOneResponse)),
	}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil).Once()

	got, err := client.ListZeroTrustUsers(ctx, "account456")
	assert.NoError(t, err)
	assert.Len(t, got, 1)
}

func TestListAccessUsers_Multiple(t *testing.T) {
	mockHttpClient := NewMockHttpClient(t)
	client := NewClient(mockHttpClient, WithToken("token123"))

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

	got, err := client.ListZeroTrustUsers(ctx, "account456")
	require.NoError(t, err)
	require.Len(t, got, 2)
	assert.Equal(t, got[0].Email, "john@doe.com")
	assert.Equal(t, got[1].Email, "jane@doe.com")
}
