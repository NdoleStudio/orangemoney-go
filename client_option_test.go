package orangemoney

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithHTTPClient(t *testing.T) {
	t.Run("httpClient is not set when the httpClient is nil", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		config := defaultClientConfig()

		// Act
		WithHTTPClient(nil).apply(config)

		// Assert
		assert.NotNil(t, config.httpClient)
	})

	t.Run("httpClient is set when the httpClient is not nil", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		config := defaultClientConfig()
		newClient := &http.Client{Timeout: 300}

		// Act
		WithHTTPClient(newClient).apply(config)

		// Assert
		assert.NotNil(t, config.httpClient)
		assert.Equal(t, newClient.Timeout, config.httpClient.Timeout)
	})
}

func TestWithBaseURL(t *testing.T) {
	t.Run("baseURL is set successfully", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		baseURL := "https://example.com"
		config := defaultClientConfig()

		// Act
		WithBaseURL(baseURL).apply(config)

		// Assert
		assert.Equal(t, baseURL, config.baseURL)
	})

	t.Run("tailing / is trimmed from baseURL", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		baseURL := "https://example.com/"
		config := defaultClientConfig()

		// Act
		WithBaseURL(baseURL).apply(config)

		// Assert
		assert.Equal(t, "https://example.com", config.baseURL)
	})
}

func TestWithUsername(t *testing.T) {
	t.Run("username is set successfully", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		username := "username-1"
		config := defaultClientConfig()

		// Act
		WithUsername(username).apply(config)

		// Assert
		assert.Equal(t, username, config.username)
	})
}

func TestWithPassword(t *testing.T) {
	t.Run("password is set successfully", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		password := "password-1"
		config := defaultClientConfig()

		// Act
		WithPassword(password).apply(config)

		// Assert
		assert.Equal(t, password, config.password)
	})
}

func TestWithAuthToken(t *testing.T) {
	t.Run("authToken is set successfully", func(t *testing.T) {
		// Setup
		t.Parallel()

		// Arrange
		authToken := "token-1"
		config := defaultClientConfig()

		// Act
		WithAuthToken(authToken).apply(config)

		// Assert
		assert.Equal(t, authToken, config.authToken)
	})
}
