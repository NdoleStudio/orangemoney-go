package orangemoney

// PayTokenResponse is the response when initializing the fetch pay token
type PayTokenResponse struct {
	Message string `json:"message"`
	Data    struct {
		PayToken string `json:"payToken"`
	} `json:"data"`
}
