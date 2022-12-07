package orangemoney

// OrangeResponse represents an orange payment response
type OrangeResponse[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}
