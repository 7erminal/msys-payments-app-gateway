package requests

type DSTVAccountQueryApiRequest struct {
	AccountNumber string `json:"account_number" valid:"required~Account number is required"`
}

type BillPaymentAccountQueryRequest struct {
	AccountNumber string `json:"account_number" valid:"required~Account number is required"`
	BillerCode    string `json:"biller_code" valid:"required~Biller code is required"`
	ClientId      string `json:"client_id"`
}

type AccountQueryRequest struct {
	AccountNumber string
	PhoneNumber   string
	SourceSystem  string
}

type BillPaymentAccountQueryApiRequest struct {
	AccountNumber string
	PhoneNumber   string
	SourceSystem  string
	BillerCode    string
}

type DSTVAccountQueryRequest struct {
	AccountNumber string
	PhoneNumber   string
	SourceSystem  string
}

type DSTVPaymentRequest struct {
	TransactionId      string
	DestinationAccount string
	Amount             float64
	PackageType        string
	PhoneNumber        string
	SourceSystem       string
	ClientId           string
}

type DSTVPaymentApiRequest struct {
	DestinationAccount string
	Amount             float64
	PackageType        string
	ClientId           string
}

type ECGPaymentRequest struct {
	DestinationAccount string
	Amount             float64
	PackageType        string
	ClientId           string
}

type ECGPaymentApiRequest struct {
	TransactionId      string
	DestinationAccount string
	Amount             float64
	PackageType        string
	PhoneNumber        string
	SourceSystem       string
	ClientId           string
}

type GhanaWaterPaymentRequest struct {
	DestinationAccount string
	Amount             float64
	PackageType        string
	ClientId           string
	CustomerName       string
	CustomerEmail      string
}

type GhanaWaterPaymentFuncRequest struct {
	TransactionId      string
	DestinationAccount string
	Amount             float64
	PackageType        string
	PhoneNumber        string
	SourceSystem       string
	Name               string
	Email              string
	ClientId           string
}

type GhanaWaterExtraData struct {
	Bundle    string `json:"bundle"`
	Email     string `json:"Email"`
	SessionId string `json:"SessionId"`
}

type GhanaWaterPaymentApiRequest struct {
	TransactionId      string
	DestinationAccount string
	Amount             float64
	PackageType        string
	PhoneNumber        string
	SourceSystem       string
	ExtraData          GhanaWaterExtraData
	ClientId           string
}

type GoTvPaymentRequest struct {
	DestinationAccount string
	Amount             float64
	PackageType        string
	ClientId           string
}

type GoTvPaymentRequest1 struct {
	RequestId          int64
	DestinationAccount string
	Amount             float64
	PackageType        string
	PhoneNumber        string
	SourceSystem       string
	ClientId           string
}

type GoTvPaymentApiRequest struct {
	TransactionId      string
	DestinationAccount string
	Amount             float64
	PackageType        string
	PhoneNumber        string
	SourceSystem       string
	ClientId           string
}

type StartimesPaymentRequest struct {
	DestinationAccount string
	Amount             float64
	PackageType        string
	PhoneNumber        string
	SourceSystem       string
	ClientId           string
}

type StartimesPaymentApiRequest struct {
	TransactionId      string
	DestinationAccount string
	Amount             float64
	PackageType        string
	PhoneNumber        string
	SourceSystem       string
	ClientId           string
}
