package helpers

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
)

// MakeTestServer creates an api server for testing
func MakeTestServer(responseCode int, body []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(responseCode)
		_, err := res.Write(body)
		if err != nil {
			panic(err)
		}
	}))
}

// MakeRequestCapturingTestServer creates an api server that captures the request object
func MakeRequestCapturingTestServer(responseCodes []int, responses [][]byte, requests *[]http.Request) *httptest.Server {
	index := 0
	return httptest.NewServer(http.HandlerFunc(func(responseWriter http.ResponseWriter, req *http.Request) {
		clonedRequest := req.Clone(context.Background())

		// clone body
		body, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		req.Body = io.NopCloser(bytes.NewReader(body))
		clonedRequest.Body = io.NopCloser(bytes.NewReader(body))

		*requests = append(*requests, *clonedRequest)

		responseWriter.WriteHeader(responseCodes[index])
		_, err = responseWriter.Write(responses[index])
		index++
		if err != nil {
			panic(err)
		}
	}))
}
