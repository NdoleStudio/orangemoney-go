package orangemoney

// MerchantPaymentPayPrams are the parameters for executing a payment transaction
type MerchantPaymentPayPrams struct {
	SubscriberMSISDN  string `json:"subscriberMsisdn"`
	ChannelUserMSISDN string `json:"channelUserMsisdn"`
	Amount            string `json:"amount"`
	Description       string `json:"description"`
	OrderID           string `json:"orderId"`
	Pin               string `json:"pin"`
	PayToken          string `json:"payToken"`
	NotificationURL   string `json:"notifUrl"`
}

// MerchantPaymentTransaction represents a payment request sent to a subscriber
type MerchantPaymentTransaction struct {
	ID                        int     `json:"id"`
	CreatedTime               string  `json:"createtime"`
	SubscriberMSISDN          string  `json:"subscriberMsisdn"`
	Amount                    int     `json:"amount"`
	PayToken                  string  `json:"payToken"`
	TransactionID             string  `json:"txnid"`
	TransactionMode           string  `json:"txnmode"`
	InitTransactionMessage    string  `json:"inittxnmessage"`
	InitTransactionStatus     string  `json:"inittxnstatus"`
	ConfirmTransactionStatus  *string `json:"confirmtxnstatus"`
	ConfirmTransactionMessage *string `json:"confirmtxnmessage"`
	Status                    string  `json:"status"`
	NotificationURL           string  `json:"notifUrl"`
	Description               string  `json:"description"`
	ChannelUserMSISDN         string  `json:"channelUserMsisdn"`
}
