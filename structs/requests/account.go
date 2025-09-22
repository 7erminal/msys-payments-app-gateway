package requests

type DebitAccountRequestV2 struct {
	AccountNumber string
	Amount        string
	Reference     string
	Channel       string
}

type CreditAccountRequestV2 struct {
	AccountNumber string
	Amount        string
	Reference     string
	Channel       string
}

type DebitAccountRequest struct {
	Amount     float64
	ModifiedBy int
	Reason     string
	AccountId  int64
}

type CreditAccountRequest struct {
	Amount     float64
	ModifiedBy int
	Reason     string
	AccountId  int64
}
