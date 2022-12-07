package orangemoney

import (
	"net/http"
	"strings"
)

// Option is options for constructing a client
type Option interface {
	apply(config *clientConfig)
}

type clientOptionFunc func(config *clientConfig)

func (fn clientOptionFunc) apply(config *clientConfig) {
	fn(config)
}

// WithHTTPClient sets the underlying HTTP client used for API requests.
// By default, http.DefaultClient is used.
func WithHTTPClient(httpClient *http.Client) Option {
	return clientOptionFunc(func(config *clientConfig) {
		if httpClient != nil {
			config.httpClient = httpClient
		}
	})
}

// WithBaseURL set's the base url for the flutterwave API
func WithBaseURL(baseURL string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		if baseURL != "" {
			config.baseURL = strings.TrimRight(baseURL, "/")
		}
	})
}

// WithUsername sets the Orange API Username used to fetch the access token
func WithUsername(username string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.username = username
	})
}

// WithPassword sets the Orange API password used to fetch the access token
func WithPassword(password string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.password = password
	})
}

// WithAuthToken sets the X-AUTH-TOKEN used as a header of API requests
func WithAuthToken(authToken string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.authToken = authToken
	})
}
