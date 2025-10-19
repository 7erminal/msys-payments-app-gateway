package responses

type PaymentRequestResponse struct {
	Success       bool
	StatusMessage string
	Result        interface{}
}
