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

type ExtraData struct {
	ExtraData1 string
	ExtraData2 string
	ExtraData3 string
}

type UserTransactionRequestDTO struct {
	SourceChannel            string
	SourceAccountNumber      string
	PhoneNumber              string
	Amount                   float64
	DestinationAccountNumber string
	ClientReference          string
	Package                  string
	ServiceCode              string
	RequestId                string
	ExtraData                ExtraData
	CreatedBy                string
	Status                   string
}

type SendDepositRequest struct {
	Amount        float64
	AccountNumber string
	MobileNumber  string
	PaymentMethod string
	ClientId      string
}
