package responses

type LoanData struct {
	LoanId   int
	LoanType string
	LoanBal  float64
}

type ListLoanApiResponse struct {
	Result     *LoanData `json:"Result,omitempty"`
	Client     string    `json:"Client,omitempty"`
	StatusCode int       `json:"StatusCode,omitempty"`
	StatusDesc string    `json:"StatusDesc,omitempty"`
}

type ListLoanResponse struct {
	Result     *LoanData `json:"Result,omitempty"`
	StatusCode bool      `json:"StatusCode,omitempty"`
	StatusDesc string    `json:"StatusDesc,omitempty"`
}

type ListLoansApiResponse struct {
	Result     *[]LoanData `json:"Result,omitempty"`
	Client     string      `json:"Client,omitempty"`
	StatusCode int         `json:"StatusCode,omitempty"`
	StatusDesc string      `json:"StatusDesc,omitempty"`
}

type ListLoansResponse struct {
	Result     *[]LoanData `json:"Result,omitempty"`
	StatusCode bool        `json:"StatusCode,omitempty"`
	StatusDesc string      `json:"StatusDesc,omitempty"`
}

type RepayLoanApiResponse struct {
	Result     string `json:"Result,omitempty"`
	Client     string `json:"Client,omitempty"`
	StatusCode int    `json:"StatusCode,omitempty"`
	StatusDesc string `json:"StatusDesc,omitempty"`
}

type RepayLoanResponse struct {
	Result        DepositData `json:"Result,omitempty"`
	Success       bool        `json:"StatusCode,omitempty"`
	StatusMessage string      `json:"StatusDesc,omitempty"`
}
