package responses

import "msys_payment_app_gateway/models"

type DataBundleData struct {
	Display string
	Value   string
	Amount  float64
}

type DataBundlesListResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        []DataBundleData
}

type DataBundlesListAPIResponse struct {
	StatusCode    bool             `json:"status_code"`
	StatusMessage string           `json:"message"`
	Result        []DataBundleData `json:"result"`
}

type BuyDataBundleResponseResult struct {
	TransactionID     string  `json:"transaction_id"`
	PhoneNumber       string  `json:"phone_number"`
	Amount            float64 `json:"amount"`
	Network           string  `json:"network"`
	Destination       string  `json:"destination"`
	TransactionStatus string  `json:"transaction_status"`
	TransactionDate   string  `json:"transaction_date"`
}
type BuyDataBundleResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        *BuyDataBundleResponseResult
}

type BuyDataBundleAPIResponse struct {
	StatusCode    bool                         `json:"status_code"`
	StatusMessage string                       `json:"status_message"`
	Result        *BuyDataBundleResponseResult `json:"result"`
}

type Payment_methods struct {
	PaymentMethodId int64
	PaymentMethod   string
}

type PaymentMethodsApiResponseDTO struct {
	StatusCode     int
	PaymentMethods *[]models.Payment_methods
	StatusDesc     string
}

type PaymentMethodsResponseDTO struct {
	Success    bool
	Result     []*Payment_methods
	StatusDesc string
}
