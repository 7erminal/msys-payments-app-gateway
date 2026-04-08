package responses

type DebitAccountV2Response struct {
	StatusCode int    `json:"StatusCode"`
	StatusDesc string `json:"StatusDesc"`
	Result     string `json:"Result,omitempty"`
	Client     string `json:"Client,omitempty"`
}

type CreditAccountV2Response struct {
	StatusCode int    `json:"StatusCode"`
	StatusDesc string `json:"StatusDesc"`
	Result     string `json:"Result,omitempty"`
	Client     string `json:"Client,omitempty"`
}

type CloseAccountResponse struct {
	StatusCode bool   `json:"StatusCode"`
	StatusDesc string `json:"StatusDesc"`
	Result     string `json:"Result,omitempty"`
}

type CloseAccountApiResponse struct {
	StatusCode int    `json:"StatusCode"`
	StatusDesc string `json:"StatusDesc"`
	Result     string `json:"Result,omitempty"`
	Client     string `json:"Client,omitempty"`
}

type AccountDetailsData struct {
	AccountName   string
	AccountNumber string
	Product       string
}

type AccountDetailsApiResponse struct {
	StatusCode int
	StatusDesc string
	Result     []*AccountDetailsData
}
