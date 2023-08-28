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
)

//go:embed testdata/success.json
var successResponse []byte

func TestFetchAccessUsers(t *testing.T) {
	mockHttpClient := NewMockHttpClient(t)
	client := NewClient(mockHttpClient, WithToken("token123"))

	ctx := context.Background()

	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewReader(successResponse)),
	}
	mockHttpClient.On("Do", mock.AnythingOfType("*http.Request")).Return(resp, nil)

	got, err := client.FetchZeroTrustUsers(ctx, "account456", 1)
	assert.NoError(t, err)
	assert.Len(t, got.Result, 1)
}
