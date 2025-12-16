package responses

import (
	"msys_payment_app_gateway/models"
	"time"
)

type TransactionApiResponse struct {
	StatusCode int
	StatusDesc string
	Result     *Bil_transactions
}

type Bil_transactions struct {
	TransactionId           int64            `orm:"auto"`
	TransactionRefNumber    string           `orm:"size(255);unique"`
	Service                 *models.Services `orm:"rel(fk)"`
	BillerCode              string           `orm:"size(255)"`
	TransactionBy           *Customers       `orm:"rel(fk);column(transaction_by)"`
	Amount                  float64
	TransactingCurrency     string `orm:"size(255)"`
	SourceChannel           string `orm:"size(255)"`
	Source                  string `orm:"size(255)"`
	Destination             string `orm:"size(255)"`
	Package                 string `orm:"size(255)"`
	Charge                  float64
	Commission              float64
	ExternalReferenceNumber string               `orm:"size(255)"`
	Status                  *models.Status_codes `orm:"rel(fk)"`
	ExtraDetails1           string               `orm:"size(255)"`
	ExtraDetails2           string               `orm:"size(255)"`
	ExtraDetails3           string               `orm:"size(255)"`
	DateCreated             time.Time            `orm:"type(datetime)"`
	DateModified            time.Time            `orm:"type(datetime)"`
	CreatedBy               int
	ModifiedBy              int
	Active                  int
}

type LogTransactionResponse struct {
	StatusCode int
	Result     *Bil_transactions
	StatusDesc string
}

type TransferApiResponseData struct {
	ClientRefernce string  `json:"clientReference"`
	Amount         float64 `json:"amount"`
	Charges        float64 `json:"charges"`
	Description    string  `json:"description"`
	RecipientName  string  `json:"recipientName"`
	Meta           string  `json:"meta"`
}

type Trx_transactions struct {
	TransactionId          string
	Amount                 float64
	TotalDebitAmount       float64
	Charge                 float64
	Commission             float64
	SenderAccountNumber    string
	RecipientAccountNumber string
	TransferCode           string
	Status                 *Status
	ResponseCode           string
	ResponseMessage        string
	Description            string
	DateCreated            time.Time
	DateModified           time.Time
	CreatedBy              int
	ModifiedBy             int
	Active                 int
}

type TransferApiResponseDTO struct {
	StatusCode int               `json:"success"`
	StatusDesc string            `json:"statusDesc"`
	Result     *Trx_transactions `json:"result"`
}
