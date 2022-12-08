package orangemoney

import (
	"context"
	"net/http"
	"testing"

	"github.com/NdoleStudio/orangemoney-go/internal/helpers"
	"github.com/NdoleStudio/orangemoney-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

const (
	testAuthToken = "auth-token-1"
	testUsername  = "username-2"
	testPassword  = "password-3"
)

func TestClient_Token(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]http.Request, 0)
	server := helpers.MakeRequestCapturingTestServer([]int{http.StatusOK}, [][]byte{stubs.TokenResponse()}, &requests)

	client := New(
		WithBaseURL(server.URL),
		WithUsername(testUsername),
		WithPassword(testPassword),
	)

	// Act
	accessToken, response, err := client.AccessToken(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, 1, len(requests))

	request := requests[0]
	actualUsername, actualPassword, ok := request.BasicAuth()
	assert.True(t, ok)

	assert.Equal(t, testUsername, actualUsername)
	assert.Equal(t, testPassword, actualPassword)
	assert.Equal(t, "/token", request.URL.Path)
	assert.Equal(t, "application/x-www-form-urlencoded", request.Header.Get("Content-Type"))

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, int64(2496), accessToken.ExpiresIn)
	assert.Equal(t, "Bearer", accessToken.TokenType)
	assert.Equal(t, "19077204-9d0a-31fa-85cf-xxxxxxxxxx", accessToken.AccessToken)

	// Teardown
	server.Close()
}
