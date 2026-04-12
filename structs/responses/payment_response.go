package responses

type PaymentRequestResponse struct {
	Success       bool
	StatusMessage string
	Result        interface{}
}

type DepositData struct {
	TransactionReference string
	Amount               float64
	Currency             string
}

type DepositResponse struct {
	Success       bool
	StatusMessage string
	Result        *DepositData
}
