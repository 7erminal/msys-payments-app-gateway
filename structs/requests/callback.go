package requests

type CallbackMeta struct {
	Commission string
}

type CallbackData struct {
	AmountDebited         float64
	TransactionId         *string
	ClientReference       *string
	Description           *string
	ExternalTransactionId *string
	Amount                float64
	Charges               float64
	Meta                  *CallbackMeta
	RecipientName         *string
}

type CallbackFormulateRequest struct {
	ResponseCode string
	Data         CallbackData
}

type PaymentCallbackData struct {
	AmountCharged         float64
	TransactionId         *string
	ClientReference       *string
	Description           *string
	ExternalTransactionId *string
	Amount                float64
	Charges               float64
	AmountAfterCharges    float64
	PaymentDate           *string
	OrderId               *string
	Status                string
}

type PaymentCallbackFormulateRequest struct {
	ResponseCode string
	Data         PaymentCallbackData
	Message      string
}

type CallbackAPIRequest struct {
	ResponseCode string
	Data         CallbackData
}

type TransactionStatusRequest struct {
	TransactionID string
}

type TransactionStatusApiRequest struct {
	TransactionID string
}

type ReceiveMoneyCallbackAPIRequest struct {
	ResponseCode string
	Data         CallbackData
	Message      string
}
