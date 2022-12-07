package orangemoney

import "net/http"

type clientConfig struct {
	httpClient *http.Client
	username   string
	password   string
	baseURL    string
	authToken  string
}

func defaultClientConfig() *clientConfig {
	return &clientConfig{
		httpClient: http.DefaultClient,
		username:   "",
		password:   "",
		authToken:  "",
		baseURL:    "https://api-s1.orange.cm",
	}
}
