package responses

type CallbackResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        *PaymentResponse
}

type CallbackAPIResponse struct {
	StatusCode    int
	StatusMessage string
	Result        *PaymentResponse
}

type TransferCallbackResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        *Trx_transactions
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
