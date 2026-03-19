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
	CorpId                   string
	ExtraDetails1            string
	ExtraDetails2            string
	ExtraDetails3            string
	Reference                string
	ClientID                 string
	PhoneNumber              string
	TransactionPackage       string
	ExternalReferenceNumber  string
	CreatedBy                string
	Source                   string
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
	CorpId                   string
	ExtraDetails1            string
	ExtraDetails2            string
	ExtraDetails3            string
	Reference                string
	ClientID                 string
	PhoneNumber              string
	TransactionPackage       string
	ExternalReferenceNumber  string
	CreatedBy                string
}

type TransferApiRequest struct {
	RequestId              string
	Amount                 float64
	Charge                 float64
	Commission             float64
	TotalDebitAmount       float64
	SenderAccountNumber    string
	RecipientAccountNumber string
	TransferCode           string
	Description            string
	RecipientName          string
	Status                 string
	ServiceCode            string
	CreatedBy              string
}

type TransferCommissionApiRequest struct {
	RequestId              string
	TransactionId          string
	Amount                 float64
	Charge                 float64
	Commission             float64
	TotalDebitAmount       float64
	SenderAccountNumber    string
	RecipientAccountNumber string
	TransferCode           string
	Description            string
	RecipientName          string
	Status                 string
}
