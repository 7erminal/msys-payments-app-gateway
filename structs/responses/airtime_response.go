package responses

type AirtimeResponseResult struct {
	TransactionID     string  `json:"transaction_id"`
	PhoneNumber       string  `json:"phone_number"`
	Amount            float64 `json:"amount"`
	Network           string  `json:"network"`
	Destination       string  `json:"destination"`
	TransactionStatus string  `json:"transaction_status"`
	TransactionDate   string  `json:"transaction_date"`
}

type BuyAirtimeResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        *AirtimeResponseResult
}

type BuyAirtimeAPIResponse struct {
	StatusCode    bool                   `json:"status_code"`
	StatusMessage string                 `json:"status_message"`
	Result        *AirtimeResponseResult `json:"result"`
}
