package cloudflare

type ClientOption interface {
	Apply(*client)
}

type withToken struct {
	token string
}

func (w withToken) Apply(c *client) {
	c.token = w.token
}

func WithToken(token string) ClientOption {
	return withToken{token: token}
}

type withBaseURL struct {
	baseUrl string
}

func (w withBaseURL) Apply(c *client) {
	c.baseUrl = w.baseUrl
}

func WithBaseURL(baseUrl string) ClientOption {
	return withBaseURL{baseUrl: baseUrl}
}

type withAccountID struct {
	accountID string
}

func (w withAccountID) Apply(c *client) {
	c.accountID = w.accountID
}

func WithAccountID(accountID string) ClientOption {
	return withAccountID{accountID: accountID}
}
