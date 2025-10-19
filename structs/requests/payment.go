package requests

type PaymentApiRequestDTO struct {
	InitiatedBy     int64
	Amount          float32
	Service         int64
	Sender          int64
	Reciever        int64
	PaymentMethod   int64
	TransactionId   int64
	PaymentProofUrl string
	ReferenceNumber string
	CallThirdParty  bool
}

type RequestMoneyApiRequestDTO struct {
	InitiatedBy     int64
	Amount          float64
	Service         string
	Sender          int64
	Reciever        int64
	PhoneNumber     string
	CustomerName    string
	CustomerMsisdn  string
	CustomerEmail   string
	SenderAccount   string
	ReceiverAccount string
	Currency        string
	PaymentMethod   string
	TransactionId   string
	PaymentProofUrl string
	ReferenceNumber string
	CallThirdParty  bool
	Operator        string
	Network         string
	ServiceNetwork  string
}

type MakePaymentApiRequestDTO struct {
	InitiatedBy     int64
	Amount          float64
	Service         string
	Sender          int64
	Reciever        int64
	SenderAccount   string
	ReceiverAccount string
	Currency        string
	PaymentMethod   string
	TransactionId   string
	PaymentProofUrl string
	ReferenceNumber string
	CallThirdParty  bool
	Operator        string
	Network         string
}

type MomoPaymentRequestDTO struct {
	PaymentId          string
	CustomerName       string
	CustomerMsisdn     string
	CustomerEmail      string
	Operator           string
	Amount             float64
	PrimaryCallbackUrl string
	Description        string
	ClientReference    string
}

type MomoPaymentApiRequestDTO struct {
	PaymentId          string
	CustomerName       string
	CustomerMsisdn     string
	CustomerEmail      string
	Operator           string
	Amount             float64
	PrimaryCallbackUrl string
	Description        string
	ClientReference    string
	Channel            string
}
