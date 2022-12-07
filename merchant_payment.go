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
