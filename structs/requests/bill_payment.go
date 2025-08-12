package requests

type DSTVAccountQueryApiRequest struct {
	AccountNumber string `json:"account_number" valid:"required~Account number is required"`
}

type BillPaymentAccountQueryRequest struct {
	AccountNumber string `json:"account_number" valid:"required~Account number is required"`
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
	RequestId          int64
	DestinationAccount string
	Amount             float64
	PackageType        string
	PhoneNumber        string
	SourceSystem       string
}

type DSTVPaymentApiRequest struct {
	DestinationAccount string  `json:"destination_account" valid:"required~Destination account is required"`
	Amount             float64 `json:"amount" valid:"required~Amount is required"`
	PackageType        string  `json:"package_type" valid:"required~Package type is required"`
}

type ECGPaymentRequest struct {
	DestinationAccount string  `json:"destination_account" valid:"required~Destination account is required"`
	Amount             float64 `json:"amount" valid:"required~Amount is required"`
	PackageType        string  `json:"package_type" valid:"required~Package type is required"`
}

type ECGPaymentApiRequest struct {
	RequestId          int64
	DestinationAccount string
	Amount             float64
	PackageType        string
	PhoneNumber        string
	SourceSystem       string
}

type GhanaWaterPaymentRequest struct {
	DestinationAccount string
	Amount             float64
	PackageType        string
	PhoneNumber        string
	SourceSystem       string
}

type GhanaWaterPaymentApiRequest struct {
	RequestId          int64
	DestinationAccount string
	Amount             float64
	PackageType        string
	PhoneNumber        string
	SourceSystem       string
}

type GoTvPaymentRequest struct {
	DestinationAccount string
	Amount             float64
	PackageType        string
	PhoneNumber        string
	SourceSystem       string
}

type GoTvPaymentApiRequest struct {
	RequestId          int64
	DestinationAccount string
	Amount             float64
	PackageType        string
	PhoneNumber        string
	SourceSystem       string
}

type StartimesPaymentRequest struct {
	DestinationAccount string
	Amount             float64
	PackageType        string
	PhoneNumber        string
	SourceSystem       string
}

type StartimesPaymentApiRequest struct {
	RequestId          int64
	DestinationAccount string
	Amount             float64
	PackageType        string
	PhoneNumber        string
	SourceSystem       string
}
