package responses

type AgentResponse struct {
	AgentCode   string
	AgentName   string
	PhoneNumber string
	Email       string
}

type AgentTransactionsResponse struct {
	Success    bool
	Result     *[]TxnResp
	StatusDesc string
}
