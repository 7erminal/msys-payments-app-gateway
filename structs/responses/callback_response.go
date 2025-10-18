package responses

type CallbackResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        *PaymentResponse
}

type CallbackAPIResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        *PaymentResponse
}

type TransactionStatusApiResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        string
}

type TransactionStatusResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        string
}
