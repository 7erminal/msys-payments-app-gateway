package requests

type AgentRequest struct {
	AgentCode   string
	AgentName   string
	PhoneNumber string
	Email       string
	AddedBy     string
}

type AgentTransactionsRequest struct {
	AgentCode string
	FromDate  string
	ToDate    string
}
