package orangemoney

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

type service struct {
	client *Client
}

// Client is the campay API client.
// Do not instantiate this client with Client{}. Use the New method instead.
type Client struct {
	httpClient *http.Client
	common     service

	baseURL   string
	username  string
	password  string
	authToken string

	accessToken         string
	tokenExpirationTime int64
	mutex               sync.Mutex

	MerchantPayment *merchantPaymentService
}

// New creates and returns a new campay.Client from a slice of campay.ClientOption.
func New(options ...Option) *Client {
	config := defaultClientConfig()

	for _, option := range options {
		option.apply(config)
	}

	client := &Client{
		httpClient: config.httpClient,
		baseURL:    config.baseURL,
		username:   config.username,
		password:   config.password,
		authToken:  config.authToken,
		mutex:      sync.Mutex{},
	}

	client.common.client = client
	client.MerchantPayment = (*merchantPaymentService)(&client.common)
	return client
}

// AccessToken fetches the access token used to authenticate api requests.
func (client *Client) AccessToken(ctx context.Context) (*AccessTokenResponse, *Response, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, client.baseURL+"/token", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, nil, err
	}

	request.SetBasicAuth(client.username, client.password)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.do(request)
	if err != nil {
		return nil, resp, err
	}

	var token AccessTokenResponse
	if err = json.Unmarshal(*resp.Body, &token); err != nil {
		return nil, resp, err
	}

	return &token, resp, nil
}

// refreshToken refreshes the authentication AccessTokenResponse
func (client *Client) refreshToken(ctx context.Context) error {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	if client.tokenExpirationTime > time.Now().UTC().Unix() {
		return nil
	}

	client.accessToken = ""

	token, _, err := client.AccessToken(ctx)
	if err != nil {
		return err
	}

	client.accessToken = token.AccessToken
	client.tokenExpirationTime = time.Now().UTC().Unix() + token.ExpiresIn - 100 // Give extra 100 second buffer

	return nil
}

// newRequest creates an API request. A relative URL can be provided in uri,
// in which case it is resolved relative to the BaseURL of the Client.
// URI's should always be specified without a preceding slash.
func (client *Client) newRequest(ctx context.Context, method, uri string, body any) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, client.baseURL+uri, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-AUTH-TOKEN", client.authToken)
	req.Header.Set("Authorization", "Bearer "+client.accessToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return req, nil
}

// do carries out an HTTP request and returns a Response
func (client *Client) do(req *http.Request) (*Response, error) {
	if req == nil {
		return nil, fmt.Errorf("%T cannot be nil", req)
	}

	httpResponse, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = httpResponse.Body.Close() }()

	resp, err := client.newResponse(httpResponse)
	if err != nil {
		return resp, err
	}

	_, err = io.Copy(io.Discard, httpResponse.Body)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// newResponse converts an *http.Response to *Response
func (client *Client) newResponse(httpResponse *http.Response) (*Response, error) {
	if httpResponse == nil {
		return nil, fmt.Errorf("%T cannot be nil", httpResponse)
	}

	resp := new(Response)
	resp.HTTPResponse = httpResponse

	buf, err := io.ReadAll(resp.HTTPResponse.Body)
	if err != nil {
		return nil, err
	}
	resp.Body = &buf

	return resp, resp.Error()
}
