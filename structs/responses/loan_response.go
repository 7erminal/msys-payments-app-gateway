package responses

type LoanData struct {
	LoanDate        string
	LoanDescription string
	LoanAmount      float64
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
	Result     string `json:"Result,omitempty"`
	StatusCode bool   `json:"StatusCode,omitempty"`
	StatusDesc string `json:"StatusDesc,omitempty"`
}
