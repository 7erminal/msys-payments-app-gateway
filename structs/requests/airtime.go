package requests

type BuyAirtimeFormulatedRequest struct {
	RequestId    int64   `json:"request_id"`
	Amount       float64 `json:"amount"`
	Network      string  `json:"network"`
	Destination  string  `json:"destination"`
	SourceSystem string  `json:"source_system"`
	PhoneNumber  string  `json:"phone_number"`
}

type BuyAirtimeRequest struct {
	Amount      float64 `json:"amount" valid:"required~Amount is required"`
	Network     string  `json:"network" valid:"required~Network is required"`
	Destination string  `json:"destination" valid:"required~Destination is required"`
}

type BuyAirtimeAPIRequest struct {
	Amount      float64 `json:"amount" valid:"required~Amount is required"`
	Destination string  `json:"destination" valid:"required~Destination is required"`
}
