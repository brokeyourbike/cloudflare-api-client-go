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
