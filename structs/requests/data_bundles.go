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
}

type BuyDataBundleFormulatedRequest struct {
	RequestId    int64   `json:"request_id"`
	Amount       float64 `json:"amount"`
	Network      string  `json:"network"`
	Destination  string  `json:"destination"`
	BundleId     string  `json:"bundle_id"`
	SourceSystem string  `json:"source_system"`
	PhoneNumber  string  `json:"phone_number"`
}

type GetBundlesAPIRequest struct {
	Destination string `json:"destination"`
}

type BuyDataBundleAPIRequest struct {
	Destination string  `json:"destination"`
	Amount      float64 `json:"amount"`
	BundleId    string  `json:"bundle_id"`
}

// type GetClientsRequest struct {
// 	Destination string `json:"destination"`
// }
