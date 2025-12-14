package responses

import (
	"msys_payment_app_gateway/models"
	"time"
)

type PaymentResponse struct {
	PaymentId       string
	TransactionId   string
	Sender          string
	Reciever        string
	Amount          float64
	Commission      float64
	Charge          float64
	OtherCharge     float64
	PaymentAmount   float64
	PaymentMethod   *models.Payment_methods
	PaymentProof    string
	Status          *Status
	Service         string
	SenderAccount   string
	ReceiverAccount string
	ReferenceNumber string
	ServiceNetwork  string
	ServicePackage  string
	DateCreated     time.Time `orm:"type(datetime)"`
	DateModified    time.Time `orm:"type(datetime)"`
	ProcessedDate   time.Time `orm:"type(datetime);null"`

	Active          int
	CallbackUrl     string
	ClientReference string
	PaymentHistory  *[]PaymentHistoryResponse
}

type PaymentHistoryResponse struct {
	PaymentHistoryId int64
	PaymentId        int64
	Status           string
	Service          string
	Narration        string
	Reference        string
	DateCreated      time.Time
	DateModified     time.Time
	CreatedBy        int64
	ModifiedBy       int64
	Active           int
}

type PaymentApiResponseDTO struct {
	StatusCode int
	Payment    *PaymentResponse
	StatusDesc string
}

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
	StatusCode    bool
	StatusMessage string
	Result        *[]AccountQueryData
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
	StatusCode    bool
	StatusMessage string
	Result        *DSTVBillPaymentDataResponse
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
