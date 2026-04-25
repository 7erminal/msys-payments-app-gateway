package requests

type AccountLoansRequest struct {
	AccountNumber string `json:"accountNumber"`
}

type LoanRepaymentApiRequest struct {
	AccountNumber string `json:"accountNumber"`
	Amount        string `json:"amount"`
	MobileNumber  string `json:"mobileNumber"`
	LoanId        string `json:"loanId"`
	ClientId      string `json:"clientId"`
	PaymentMode   string `json:"paymentMode"`
}

type LoanRepaymentRequest struct {
	CustomerName   string
	CustomerNumber string
	AccountNumber  string `json:"accountNumber"`
	Amount         string `json:"amount"`
	PaymentMethod  string `json:"paymentMethod"`
	LoanId         string `json:"loanId"`
	ClientId       string `json:"clientId"`
}
