package orangemoney

import (
	"context"
	"net/http"
	"testing"

	"github.com/NdoleStudio/orangemoney-go/internal/helpers"
	"github.com/NdoleStudio/orangemoney-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestMerchantPaymentService_Init(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]http.Request, 0)
	responses := [][]byte{stubs.TokenResponse(), stubs.MerchantPaymentInitResponse()}
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, responses, &requests)
	client := New(
		WithBaseURL(server.URL),
		WithUsername(testUsername),
		WithPassword(testPassword),
		WithAuthToken(testAuthToken),
	)

	// Act
	payToken, response, err := client.MerchantPayment.Init(context.Background())

	// Assert
	assert.Nil(t, err)

	assert.GreaterOrEqual(t, len(requests), 2)
	request := requests[len(requests)-1]

	assert.Equal(t, "/omcoreapis/1.0.2/mp/init", request.URL.Path)
	assert.Equal(t, testAuthToken, request.Header.Get("X-AUTH-TOKEN"))
	assert.Equal(t, "Bearer 19077204-9d0a-31fa-85cf-xxxxxxxxxx", request.Header.Get("Authorization"))

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	assert.Equal(t, &PayTokenResponse{
		Message: "Payment request successfully initiated",
		Data: struct {
			PayToken string `json:"payToken"`
		}{
			PayToken: "MP22120771FEB7B21FD2381C3786",
		},
	}, payToken)

	// Teardown
	server.Close()
}
