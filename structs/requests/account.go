package requests

type DebitAccountRequestV2 struct {
	AccountNumber string
	Amount        string
	Reference     string
	Channel       string
	ClientId      string
}

type CreditAccountRequestV2 struct {
	AccountNumber string
	Amount        string
	Reference     string
	Channel       string
	ClientId      string
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

type CloseAccountApiRequest struct {
	AccountNumber string `json:"accountNumber"`
	ClientId      string `json:"clientId"`
}

type CloseAccountRequest struct {
	AccountNumber string `json:"accountNumber"`
	ClientId      string `json:"clientId"`
}

type GetCorporativesRequest struct {
	RequestId string `json:"request_id"`
}

type AccountActivityRequest struct {
	AccountNumber string
	ClientId      string
	Reference     string
	Amount        float64
	ActivityType  string
	ActivityBy    string
	MobileNumber  string
	PaymentMethod string
}
