package orangemoney

import (
	"context"
	"encoding/json"
	"net/http"
)

// merchantPaymentService is the API client for the `/mp` endpoint
type merchantPaymentService service

// Init allows a consumer to get a PayToken which uniquely identity the transaction.
func (service *merchantPaymentService) Init(ctx context.Context) (*OrangeResponse[PayToken], *Response, error) {
	err := service.client.refreshToken(ctx)
	if err != nil {
		return nil, nil, err
	}

	request, err := service.client.newRequest(ctx, http.MethodPost, "/omcoreapis/1.0.2/mp/init", nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	token := new(OrangeResponse[PayToken])
	if err = json.Unmarshal(*response.Body, token); err != nil {
		return nil, response, err
	}

	return token, response, nil
}

// Push the initiated transaction to the customer's mobile phone.
func (service *merchantPaymentService) Push(ctx context.Context, payToken *string) (*OrangeResponse[MerchantPaymentTransaction], *Response, error) {
	err := service.client.refreshToken(ctx)
	if err != nil {
		return nil, nil, err
	}

	request, err := service.client.newRequest(ctx, http.MethodGet, "/omcoreapis/1.0.2/mp/push/"+*payToken, nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	transaction := new(OrangeResponse[MerchantPaymentTransaction])
	if err = json.Unmarshal(*response.Body, transaction); err != nil {
		return nil, response, err
	}

	return transaction, response, nil
}

// TransactionStatus returns the status of an initiated transaction
func (service *merchantPaymentService) TransactionStatus(ctx context.Context, payToken *string) (*OrangeResponse[MerchantPaymentTransaction], *Response, error) {
	err := service.client.refreshToken(ctx)
	if err != nil {
		return nil, nil, err
	}

	request, err := service.client.newRequest(ctx, http.MethodGet, "/omcoreapis/1.0.2/mp/paymentstatus/"+*payToken, nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	transaction := new(OrangeResponse[MerchantPaymentTransaction])
	if err = json.Unmarshal(*response.Body, transaction); err != nil {
		return nil, response, err
	}

	return transaction, response, nil
}

// Pay executes an initiated transaction
func (service *merchantPaymentService) Pay(ctx context.Context, params *MerchantPaymentPayPrams) (*OrangeResponse[MerchantPaymentTransaction], *Response, error) {
	err := service.client.refreshToken(ctx)
	if err != nil {
		return nil, nil, err
	}

	request, err := service.client.newRequest(ctx, http.MethodPost, "/omcoreapis/1.0.2/mp/pay", params)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	transaction := new(OrangeResponse[MerchantPaymentTransaction])
	if err = json.Unmarshal(*response.Body, transaction); err != nil {
		return nil, response, err
	}

	return transaction, response, nil
}
