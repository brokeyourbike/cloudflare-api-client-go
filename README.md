# cloudflare-api-client-go

[![Go Reference](https://pkg.go.dev/badge/github.com/brokeyourbike/cloudflare-api-client-go.svg)](https://pkg.go.dev/github.com/brokeyourbike/cloudflare-api-client-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/brokeyourbike/cloudflare-api-client-go)](https://goreportcard.com/report/github.com/brokeyourbike/cloudflare-api-client-go)
[![Maintainability](https://api.codeclimate.com/v1/badges/67ee4ef0f7416cbda159/maintainability)](https://codeclimate.com/github/brokeyourbike/cloudflare-api-client-go/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/67ee4ef0f7416cbda159/test_coverage)](https://codeclimate.com/github/brokeyourbike/cloudflare-api-client-go/test_coverage)

A couple of useful functions for Cloudflare API which are currently not supported by the [official client](https://github.com/cloudflare/cloudflare-api-client-go)

## Installation

```bash
go get github.com/brokeyourbike/cloudflare-api-client-go
```

## Usage

```go
client := cloudflare.NewClient(
	httpClient,
	cloudflare.WithToken("securetoken"),
	cloudflare.WithAccountID("00000000-0000-0000-0000-000000000000"),
)

users, err := client.ListZeroTrustUsers(context.TODO())
require.NoError(t, err)
```

## Authors
- [Ivan Stasiuk](https://github.com/brokeyourbike) | [Twitter](https://twitter.com/brokeyourbike) | [LinkedIn](https://www.linkedin.com/in/brokeyourbike) | [stasi.uk](https://stasi.uk)

## License
[BSD-3-Clause License](https://github.com/brokeyourbike/cloudflare-api-client-go/blob/main/LICENSE)
