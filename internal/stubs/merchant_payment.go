package stubs

// MerchantPaymentInitResponse is the response when initializing a merchant payment transaction
func MerchantPaymentInitResponse() []byte {
	return []byte(`
{
   "message":"Payment request successfully initiated",
   "data":{
      "payToken":"MP22120771FEB7B21FD2381C3786"
   }
}
`)
}

// MerchantPaymentPayResponseWithInsufficientFunds is the response when the user has insufficient funds
func MerchantPaymentPayResponseWithInsufficientFunds() []byte {
	return []byte(`
{
   "message":"60019 :: Le solde du compte du payeur est insuffisant",
   "data":{
      "id":48462449,
      "createtime":"1670442106",
      "subscriberMsisdn":"69XXXXXXX",
      "amount":100,
      "payToken":"MP22120771FEB7B21FD2381C3786",
      "txnid":null,
      "txnmode":"12345",
      "inittxnmessage":"Le solde du compte du payeur est insuffisant",
      "inittxnstatus":"60019",
      "confirmtxnstatus":null,
      "confirmtxnmessage":null,
      "status":"FAILED",
      "notifUrl":"https://example.com/payment-notification",
      "description":"Payment Description",
      "channelUserMsisdn":"69XXXXXXX"
   }
}
`)
}

// MerchantPaymentPayResponse is the response after executing a payment
func MerchantPaymentPayResponse() []byte {
	return []byte(`
{
   "message":"Merchant payment successfully initiated",
   "data":{
      "id":48463325,
      "createtime":"1670442691",
      "subscriberMsisdn":"69XXXXXXX",
      "amount":100,
      "payToken":"MP22120771FEB7B21FD2381C3786",
      "txnid":"MP221207.2051.B56929",
      "txnmode":"12345",
      "inittxnmessage":"Paiement e la clientele done.The devrez confirmer le paiement en saisissant son code PIN et vous recevrez alors un SMS. Merci dutiliser des services Orange Money.",
      "inittxnstatus":"200",
      "confirmtxnstatus":null,
      "confirmtxnmessage":null,
      "status":"PENDING",
      "notifUrl":"https://example.com/payment-notification",
      "description":"Payment Description",
      "channelUserMsisdn":"69XXXXXXX"
   }
}
`)
}

// MerchantPaymentPushResponse is the response after sending a push notification
func MerchantPaymentPushResponse() []byte {
	return []byte(`
{
   "message":"Push sent to customer",
   "data":{
      "id":48463325,
      "createtime":"1670442691",
      "subscriberMsisdn":"69XXXXXXX",
      "amount":100,
      "payToken":"MP22120771FEB7B21FD2381C3786",
      "txnid":"MP221207.2051.B56929",
      "txnmode":"12345",
      "inittxnmessage":"Paiement e la clientele done.The devrez confirmer le paiement en saisissant son code PIN et vous recevrez alors un SMS. Merci dutiliser des services Orange Money.",
      "inittxnstatus":"200",
      "confirmtxnstatus":null,
      "confirmtxnmessage":null,
      "status":"PENDING",
      "notifUrl":"https://example.com/payment-notification",
      "description":"Payment Description",
      "channelUserMsisdn":"69XXXXXXX"
   }
}
`)
}

// MerchantPaymentTransactionStatusResponse is the transaction status response for a confirmed payment
func MerchantPaymentTransactionStatusResponse() []byte {
	return []byte(`
{
   "message":"Transaction retrieved successfully",
   "data":{
      "id":48463325,
      "createtime":"1670442691",
      "subscriberMsisdn":"69XXXXXXX",
      "amount":100,
      "payToken":"MP22120771FEB7B21FD2381C3786",
      "txnid":"MP221207.2051.B56929",
      "txnmode":"12345",
      "inittxnmessage":"Paiement e la clientele done.The devrez confirmer le paiement en saisissant son code PIN et vous recevrez alors un SMS. Merci dutiliser des services Orange Money.",
      "inittxnstatus":"200",
      "confirmtxnstatus":"200",
      "confirmtxnmessage":"Successful Payment of COMPANY_NAME from 69XXXXXXX CUSTOMER_NAME. Transaction ID:MP221207.2051.B56929, Amount:100, New balance:1103.5.",
      "status":"SUCCESSFULL",
      "notifUrl":"https://example.com/payment-notification",
      "description":"Payment Description",
      "channelUserMsisdn":"69XXXXXXX"
   }
}
`)
}

// MerchantPaymentTransactionStatusResponseWithExpired is the transaction status response for an expired payment
func MerchantPaymentTransactionStatusResponseWithExpired() []byte {
	return []byte(`
{
   "message":"Transaction retrieved successfully",
   "data":{
      "id":48445436,
      "createtime":"1670436067",
      "subscriberMsisdn":null,
      "amount":null,
      "payToken":"MP22120771FEB7B21FD2381C3786",
      "txnid":null,
      "txnmode":null,
      "inittxnmessage":null,
      "inittxnstatus":null,
      "confirmtxnstatus":null,
      "confirmtxnmessage":null,
      "status":"EXPIRED",
      "notifUrl":null,
      "description":null,
      "channelUserMsisdn":null
   }
}
`)
}
