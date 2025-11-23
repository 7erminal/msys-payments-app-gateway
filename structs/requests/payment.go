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
	ServicePackage  string
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
	ServiceNetwork  string
	ServicePackage  string
}

type PaymentRequestApiRequestDTO struct {
	Amount          float64
	Service         string
	PaymentMethod   string
	SenderAccount   string
	ReceiverAccount string
	ClientId        string
	Network         string
	ServiceNetwork  string
	ServicePackage  string
	MobileNumber    string
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

type DepositAPIRequest struct {
	Amount      float64 `json:"amount" valid:"required~Amount is required"`
	Destination string  `json:"destination" valid:"required~Destination is required"`
	Source      string  `json:"source" valid:"required~Source is required"`
	ClientId    string  `json:"client_id" valid:"required~Client ID is required"`
}

type WithdrawalAPIRequest struct {
	Amount      float64 `json:"amount" valid:"required~Amount is required"`
	Destination string  `json:"destination" valid:"required~Destination is required"`
	Source      string  `json:"source" valid:"required~Source is required"`
	ClientId    string  `json:"client_id" valid:"required~Client ID is required"`
}
