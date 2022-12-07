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
