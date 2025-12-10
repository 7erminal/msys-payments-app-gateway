package requests

type LogTransactionApiRequest struct {
	RequestId                string
	SourceAccountNumber      string
	DestinationAccountNumber string
	Amount                   float64
	Charge                   float64
	ServiceCode              string
	TransactionReference     string
	TransactionType          string
	StatusCode               string
	ExtraDetails1            string
	ExtraDetails2            string
	ExtraDetails3            string
	Reference                string
	ClientID                 string
	PhoneNumber              string
	TransactionPackage       string
	ExternalReferenceNumber  string
}

type LogTransactionRequest struct {
	RequestId                string
	SourceAccountNumber      string
	DestinationAccountNumber string
	Amount                   float64
	Charge                   float64
	ServiceCode              string
	TransactionReference     string
	TransactionType          string
	StatusCode               string
	ExtraDetails1            string
	ExtraDetails2            string
	ExtraDetails3            string
	Reference                string
	ClientID                 string
	PhoneNumber              string
	TransactionPackage       string
	ExternalReferenceNumber  string
}
