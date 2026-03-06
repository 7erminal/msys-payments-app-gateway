package requests

type DataBundlesListRequest struct {
	NetworkId          string
	DestinationAccount string
	PhoneNumber        string
}

type DataBundlesListFormulatedRequest struct {
	NetworkId          string
	DestinationAccount string
	PhoneNumber        string
	SourceSystem       string
	ClientId           string
}

type BuyDataBundleFormulatedRequest struct {
	TransactionId string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
	Network       string  `json:"network"`
	Destination   string  `json:"destination"`
	BundleId      string  `json:"bundle_id"`
	SourceSystem  string  `json:"source_system"`
	PhoneNumber   string  `json:"phone_number"`
	ClientId      string  `json:"client_id"`
}

type GetBundlesAPIRequest struct {
	RequestId   string `json:"request_id"`
	Destination string `json:"destination"`
}

type BuyDataBundleAPIRequest struct {
	Destination string  `json:"destination"`
	Amount      float64 `json:"amount"`
	BundleId    string  `json:"bundle_id"`
	Network     string  `json:"network"`
	ClientId    string  `json:"client_id"`
}

// type GetClientsRequest struct {
// 	Destination string `json:"destination"`
// }
