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
}

type LoanRepaymentRequest struct {
	AccountNumber string `json:"accountNumber"`
	Amount        string `json:"amount"`
	MobileNumber  string `json:"mobileNumber"`
	LoanId        string `json:"loanId"`
	ClientId      string `json:"clientId"`
}
