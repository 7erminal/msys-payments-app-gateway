package responses

type AccountQueryData struct {
	Display string
	Value   string
	Amount  float64
}

type AccountQueryResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        *[]AccountQueryData
}

type AccountQueryAPIResponse struct {
	StatusCode    bool                `json:"status_code"`
	StatusMessage string              `json:"status_message"`
	Result        *[]AccountQueryData `json:"result"`
}

type DSTVBillPaymentDataResponse struct {
	Description   string
	Amount        float64
	TransactionId string
}

type DSTVBillPaymentResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        *DSTVBillPaymentDataResponse
}

type DSTVBillPaymentApiResponse struct {
	StatusCode    bool                         `json:"status_code"`
	StatusMessage string                       `json:"status_message"`
	Result        *DSTVBillPaymentDataResponse `json:"result"`
}

type GoTvBillPaymentDataResponse struct {
	Description   string
	Amount        float64
	TransactionId string
}

type GoTvBillPaymentResponse struct {
	StatusCode    bool                         `json:"status_code"`
	StatusMessage string                       `json:"status_message"`
	Result        *GoTvBillPaymentDataResponse `json:"result"`
}

type GoTvBillPaymentApiResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        *GoTvBillPaymentDataResponse
}

type GhanaWaterBillPaymentDataResponse struct {
	Description   string
	Amount        float64
	TransactionId string
}

type GhanaWaterBillPaymentResponse struct {
	StatusCode    bool                               `json:"status_code"`
	StatusMessage string                             `json:"status_message"`
	Result        *GhanaWaterBillPaymentDataResponse `json:"result"`
}

type GhanaWaterBillPaymentApiResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        *GhanaWaterBillPaymentDataResponse
}

type StartimesBillPaymentDataResponse struct {
	Description   string
	Amount        float64
	TransactionId string
}

type StartimesBillPaymentResponse struct {
	StatusCode    bool                              `json:"status_code"`
	StatusMessage string                            `json:"status_message"`
	Result        *StartimesBillPaymentDataResponse `json:"result"`
}

type StartimesBillPaymentApiResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        *StartimesBillPaymentDataResponse
}

type ECGBillPaymentDataResponse struct {
	Description   string
	Amount        float64
	TransactionId string
}

type ECGBillPaymentResponse struct {
	StatusCode    bool                        `json:"status_code"`
	StatusMessage string                      `json:"status_message"`
	Result        *ECGBillPaymentDataResponse `json:"result"`
}

type ECGBillPaymentApiResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        *ECGBillPaymentDataResponse
}

type BilTransactionsApiData struct {
	TransactionId           int64
	TransactionRefNumber    string
	Service                 string
	Request                 int64
	TransactionBy           string
	Amount                  float64
	TransactingCurrency     string
	SourceChannel           string
	Source                  string
	Destination             string
	Charge                  float64
	BillerName              string
	NetworkName             string
	Commission              float64
	ExternalReferenceNumber string
	Status                  string
	DateCreated             string
	DateModified            string
	CreatedBy               int
	ModifiedBy              int
	Active                  int
}

type BilTransactionsData struct {
	TransactionId           int64
	TransactionRefNumber    string
	Service                 string
	TransactionBy           string
	Amount                  float64
	TransactingCurrency     string
	Source                  string
	Destination             string
	Charge                  float64
	BillerName              string
	NetworkName             string
	Commission              float64
	ExternalReferenceNumber string
	Status                  string
	TransactionDate         string
}

type BilTransactionsApiResponse struct {
	StatusCode    string
	StatusMessage string
	Result        []*BilTransactionsApiData
}

type BilTransactionsResponse struct {
	Success    bool
	StatusDesc string
	Result     []*BilTransactionsData
}

type BilTransactionResponse struct {
	Success    bool
	StatusDesc string
	Result     *BilTransactionsData
}

type BilTransactionApiResponse struct {
	StatusCode    string
	StatusMessage string
	Result        *BilTransactionsApiData
}
