package responses

import (
	"time"
)

type LoginDataResponse struct {
	PhoneNumber  string `json:"phone_number"`
	SourceSystem string `json:"source_system"`
	ClientId     string `json:"client_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type LoginDataResponseResult struct {
	StatusCode    int
	StatusMessage string
	Result        bool
	Client        string
}

type LoginApiResponse struct {
	StatusCode int
	Value      string
	StatusDesc string
}

type LoginAccountApiResponse struct {
	Data LoginDataResponseResult `json:"data"`
}

type LoginResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        string
}

type CustAccountsData struct {
	Username     string
	MobileNumber string
}

type CustCorpAccountsData struct {
	Product       string
	AccountNumber string
}

type CustAccountsDataResponseResult struct {
	StatusCode    int
	StatusMessage string
	Result        *[]CustAccountsData
	Client        string
}

type CustAccountsApiResponse struct {
	StatusCode int
	StatusDesc string
	Result     *[]CustAccountsData
}

type ListCustAccountsApiResponse struct {
	StatusCode int
	StatusDesc string
	Result     *[]CustCorpAccountsData
}

type CustomerAccountApiResponse struct {
	StatusCode    string                      `json:"statusCode"`
	StatusMessage string                      `json:"statusMessage"`
	Result        *CustomerAccountResponseObj `json:"result,omitempty"`
}

type Identification_types struct {
	IdentificationTypeId int64  `orm:"auto"`
	Name                 string `orm:"size(100)"`
	Code                 string `orm:"size(100)"`
	Active               int
}

type Customer_categories struct {
	CustomerCategoryId int64
	Category           string
	Description        string
	DateCreated        time.Time
	DateModified       time.Time
	CreatedBy          int
	ModifiedBy         int
	Active             int
}

type Customer struct {
	CustomerId           int64
	CustomerNumber       string
	FullName             string
	Email                string
	PhoneNumber          string
	Gender               string
	Location             string
	IdentificationType   *Identification_types
	IdentificationNumber string
	CustomerCategory     *Customer_categories
	Nickname             string
	Dob                  time.Time
	DateCreated          time.Time
	DateModified         time.Time
	CreatedBy            int
	ModifiedBy           int
	Active               int
	LastTxnDate          time.Time
	ImagePath            string
}

type CustomerResponseDTO struct {
	StatusCode int
	Customer   *Customer
	StatusDesc string
}

// type UserResponseDTO struct {
// 	StatusCode int
// 	User   *UserGateway
// 	StatusDesc string
// }

type CustomerResponseDTO2 struct {
	StatusCode int
	Result     *Customer
	StatusDesc string
}

type Clients struct {
	Id           int64
	ClientCorpId string
	ClientName   string
	ClientCode   string
	ClientUrl    string
	DateCreated  time.Time
	DateModified time.Time
	CreatedBy    int
	ModifiedBy   int
	Active       int
}

type CustomerCorporativesResponseDTO struct {
	Id             int64
	CustomerNumber string
	CorpId         *Clients
	IsDefault      int
	IsActive       int
	DateCreated    time.Time
	DateModified   time.Time
	CreatedBy      int
	ModifiedBy     int
}

type CustomerGateway struct {
	CustomerId           int64
	FullName             string
	ImagePath            string
	Email                string
	PhoneNumber          string
	Gender               string
	Location             string
	IdentificationType   *Identification_types
	IdentificationNumber string
	DateCreated          time.Time
	Status               int
	LastDeal             time.Time
	Branch               *string
	CustomerCorporatives *[]CustomerCorporativesResponseDTO
}

type CustomerGatewayResponseDTO struct {
	StatusCode    bool
	Result        *CustomerGateway
	StatusMessage string
}

type RegisterDataResponseResult struct {
	StatusCode int
	StatusDesc string
	Result     bool
	Client     string
}

type RegisterApiResponse struct {
	Data RegisterDataResponseResult `json:"data"`
}

type RegisterResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        *CustomerGateway
}

type RegisterAccountResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        *bool
}

type VerifyCustomerApiResponse struct {
	StatusCode    int
	StatusMessage string
	Result        string
}

type VerifyCustomerResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        *CustomerGateway
}

type FetchApprovedCustomersApiResponse struct {
	StatusCode int
	StatusDesc string
	Result     *CustomerGateway
}

type FetchApprovedCustomersResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        *CustomerGateway
}

type VerifyApprovedCustomerApiResponse struct {
	StatusCode int
	StatusDesc string
	Result     *CustomerGateway
}

type VerifyApprovedCustomerResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        *CustomerGateway
}

type CustomerAccountsDataResponseResult struct {
	StatusCode    int
	StatusMessage string
	Result        bool
	Client        string
}

type CustomerAccountsApiResponse struct {
	Data CustomerAccountsDataResponseResult `json:"data"`
}

type CustomerAccountResponse struct {
	CustomerAccountId int64
	AccountNumber     string
	AccountAlias      string
	AccountType       string
	Reference         string
	ClientId          string
	Balance           float64
	SharesBalance     float64
	LoanBalance       float64
	CLearBalance      float64
	FrozenAmount      float64
	BalanceBefore     float64
	DateCreated       string
	Active            int
}

type CustomerAccountsResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        *[]CustomerAccountResponse
}

type CustomerCorporativeAccountsResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        *[]CustomerAccountsApiResponse
}

type AccountInfoDataResponseResult struct {
	StatusCode    int
	StatusMessage string
	Result        bool
	Client        string
}

type AccountInfoApiResponse struct {
	Data AccountInfoDataResponseResult `json:"data"`
}

type AccountInfoResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        string
}

type AccountBalanceData struct {
	AccountStatus    string
	AvailableBalance *float64
	ClearBalance     *float64
	LoanBalance      *float64
	SharesBalance    *float64
}

type AccountBalanceDataResp struct {
	AccountNumber    string
	AccountStatus    string
	AvailableBalance *float64
	ClearBalance     *float64
	LoanBalance      *float64
	SharesBalance    *float64
}

type AccountBalanceDataResponseResult struct {
	StatusCode    int
	StatusMessage string
	Result        *AccountBalanceData
	Client        string
}

type AccountBalanceApiResponse struct {
	StatusCode int                 `json:"StatusCode"`
	StatusDesc string              `json:"StatusDesc"`
	Result     *AccountBalanceData `json:"Result,omitempty"`
	Client     string              `json:"Client,omitempty"`
}

type AccountBalanceResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        *AccountBalanceDataResp
}

type CustAccountBalanceResponse struct {
	StatusCode      bool
	StatusMessage   string
	Result          *AccountBalanceData
	CustomerAccount *CustomerAccountResponseObj
}

type AccountActivityResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        string
}

type CorporativeData struct {
	Id           int64
	ClientName   string
	ClientCode   string
	ClientUrl    string
	ClientCorpId int64
	DateCreated  string
	DateModified string
	Active       int
}

type CorporativeDataResponseResult struct {
	StatusCode    int
	StatusMessage string
	Result        *CorporativeData
	Client        string
}

type CorporativeApiResponse struct {
	StatusCode    int
	StatusMessage string
	Result        *[]CorporativeData
}

type CorporativeResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        *[]CorporativeData
}

type ResetPinDataResponseResult struct {
	StatusCode    int
	StatusMessage string
	Result        bool
	Client        string
}

type ResetPinApiResponse struct {
	Data ResetPinDataResponseResult `json:"data"`
}

type ResetPinResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        string
}

type NameInquiryDataResponseResult struct {
	StatusCode    int
	StatusMessage string
	Result        string
	Client        string
}

type NameInquiryApiResponse struct {
	Data NameInquiryDataResponseResult `json:"data"`
}

type NameInquiryResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        string
}

type NameInquiryApiResponseData struct {
	IsRegistered string
	Name         string
	Status       string
	Profile      string
}

type NameInquiryApiResponseDTO struct {
	Success    bool
	Result     *NameInquiryApiResponseData
	StatusDesc string
}

type TokenDestructureResponseDTO struct {
	Email    string
	RoleId   string
	InviteBy string
}

type InviteDecodeResponseDTO struct {
	StatusCode int
	Value      *TokenDestructureResponseDTO
	StatusDesc string
}

type InviteResponseDTO struct {
	Success    bool
	Result     *TokenDestructureResponseDTO
	StatusDesc string
}

type StringResponseDTO struct {
	Success    bool
	Result     *string
	StatusDesc string
}

type StringResponseCodeDTO struct {
	StatusCode int
	Result     *string
	StatusDesc string
}

type StringOriResponseDTO struct {
	StatusCode int
	Value      string
	StatusDesc string
}

type TokenResponseDTO struct {
	AccessToken  string
	RefreshToken string
	TokenType    string
	ExpiresIn    int64
}

type LoginDataResponseDTO struct {
	UserType string
	Token    *TokenResponseDTO
}

type LoginResponseDTO struct {
	StatusCode int
	Result     string
	StatusDesc string
}

type CustomerLoginApiResponseDTO struct {
	StatusCode int
	Result     *LoginDataResponseDTO
	StatusDesc string
}

type CustomerLoginDataResponseDTO struct {
	UserType string
	Token    *TokenResponseDTO
}

type CustomerLoginResponse struct {
	StatusCode    bool
	StatusMessage string
	Result        *CustomerLoginDataResponseDTO
}

type UserOriResponseDTO struct {
	StatusCode int
	User       *UsersOri
	StatusDesc string
}

type UserResponseDTO struct {
	StatusCode int
	User       *Users
	StatusDesc string
}

type UserGatewayResponseDTO struct {
	Success    bool
	Result     *UserGateway
	StatusDesc string
}

type UsersData struct {
	Data  *[]UserGateway
	Count int
}

type UsersGatewayResponseDTO struct {
	Success    bool
	Result     *UsersData
	StatusDesc string
}

type UsersOriResponseDTO struct {
	StatusCode int
	Users      *[]UsersOri
	StatusDesc string
}

type Role struct {
	RoleId int64
	Role   string
}

type UserExtraDetails struct {
	// CustomerId int64
	// User       int64
	// Branch *BranchResp
	// Shop             *Shops
	// CustomerCategory *Customer_categories
	// Nickname         string
	DateCreated time.Time
	// DateModified     time.Time
	// CreatedBy        int
	// ModifiedBy       int
	// Active           int
}

type UsersOri struct {
	UserId        int64
	UserType      int
	ImagePath     string
	UserDetails   *UserExtraDetails
	FullName      string
	Username      string
	Password      string
	Email         string
	PhoneNumber   string
	Gender        string
	Dob           time.Time
	Address       string
	IdType        string
	IdNumber      string
	Role          *Role
	MaritalStatus string
	Active        int
	IsVerified    bool
	DateCreated   time.Time
	DateModified  time.Time
	CreatedBy     int
	ModifiedBy    int
}

type Users struct {
	UserId        int64 `orm:"auto"`
	UserType      int
	ImagePath     string
	Customer      *UserExtraDetails `orm:"rel(fk);column(customer_id)"`
	FullName      string            `orm:"size(255)"`
	Username      string            `orm:"size(255)"`
	Password      string            `orm:"size(255)"`
	Email         string            `orm:"size(255)"`
	PhoneNumber   string            `orm:"size(255)"`
	Gender        string            `orm:"size(10)"`
	Dob           time.Time         `orm:"type(datetime)"`
	Address       string            `orm:"size(255)"`
	IdType        string            `orm:"size(5)"`
	IdNumber      string            `orm:"size(100)"`
	Role          *Role
	MaritalStatus string `orm:"size(255);omitempty"`
	Active        int
	IsVerified    bool
	DateCreated   time.Time `orm:"type(datetime)"`
	DateModified  time.Time `orm:"type(datetime)"`
	CreatedBy     int
	ModifiedBy    int
	// Branch        *BranchResp
}

type UserGateway struct {
	UserId int64 `orm:"auto"`
	// UserType    int
	FullName    string `orm:"size(255)"`
	Username    string `orm:"size(255)"`
	Email       string `orm:"size(255)"`
	PhoneNumber string `orm:"size(255)"`
	ImagePath   string
	Customer    *UserExtraDetails
	// Gender         string    `orm:"size(10)"`
	// Dob            time.Time `orm:"type(datetime)"`
	// Address        string    `orm:"size(255)"`
	// IdType         string    `orm:"size(5)"`
	// IdNumber       string    `orm:"size(100)"`
	Status         string
	IsVerified     bool
	Role           *Role
	DateRegistered time.Time `orm:"type(datetime)"`
	// DateModified time.Time `orm:"type(datetime)"`
	// CreatedBy    int
	// ModifiedBy   int
}

type UserTokens struct {
	Token      string
	ExpiryDate time.Time
}

type UserInvitesOri struct {
	UserInviteId    int64
	InvitedBy       *UsersOri
	InvitationToken *UserTokens
	Email           string
	Role            *Role
	Status          string
	DateCreated     time.Time
}

type UserInvites struct {
	UserInviteId int64
	// InvitedBy    *UserGateway
	// InvitationToken *UserTokens
	Email       string
	Role        string
	Status      string
	DateCreated time.Time
}

type UserInvitesResponseDTO struct {
	StatusCode  int
	UserInvites *[]UserInvitesOri
	StatusDesc  string
}

type UserInvitesResponse struct {
	Success    bool
	Result     *[]UserInvites
	StatusDesc string
}

type UserInviteResponseDTO struct {
	StatusCode int
	UserInvite *UserInvitesOri
	StatusDesc string
}

type UserInviteResponse struct {
	Success    bool
	Result     *UserInvites
	StatusDesc string
}

type Customers struct {
	CustomerId           int64     `orm:"auto"`
	FullName             string    `orm:"column(full_name);size(255)"`
	ImagePath            string    `orm:"column(image_path);size(255)"`
	Email                string    `orm:"column(email);size(255);null"`
	PhoneNumber          string    `orm:"column(phone_number);size(255);null"`
	Location             string    `orm:"column(location);size(255);null"`
	IdentificationNumber string    `orm:"column(identification_number);size(255);null"`
	Nickname             string    `orm:"size(100);omitempty;null"`
	Dob                  time.Time `orm:"column(dob);type(datetime)"`
	DateCreated          time.Time `orm:"type(datetime)"`
	DateModified         time.Time `orm:"type(datetime)"`
	CreatedBy            int
	ModifiedBy           int
	Active               int
	User                 *Users `orm:"rel(fk);omitempty;null"`
}

type CustomersResponseDTO struct {
	StatusCode int
	Customers  *[]Customer
	StatusDesc string
}

type CustomersData struct {
	Data  *[]CustomerGateway
	Count int
}

type CustomersGatewayResponseDTO struct {
	Success    bool
	Result     *CustomersData
	StatusDesc string
}

type IDTypeResponse struct {
	IdentificationTypeId int64
	Name                 string
	Code                 string
}

type IDTypeResponseDTO struct {
	StatusCode int
	IdTypes    *[]IDTypeResponse
	StatusDesc string
}

type IDTypesGatewayResponseDTO struct {
	Success    bool
	Result     *[]IDTypeResponse
	StatusDesc string
}

type CustomerEmergencyContact struct {
	CustomerEmergencyContactId int64
	Name                       string
	Contact                    string
}

type CustomerGuarantor struct {
	CustomerGuarantorId int64
	Name                string
	Contact             string
}

type CustomerEmergencyContactGateway struct {
	CustomerEmergencyContactId int64
	Name                       string
	PhoneNumber                string
}

type CustomerGuarantorGateway struct {
	CustomerGuarantorId int64
	Name                string
	PhoneNumber         string
}

type CustomerEmergencyContactResponseDTO struct {
	StatusCode               int
	CustomerEmergencyContact *CustomerEmergencyContact
	StatusDesc               string
}

type CustomerGuarantorResponseDTO struct {
	StatusCode        int
	CustomerGuarantor *CustomerGuarantor
	StatusDesc        string
}

type CustomerAccountResponseObj struct {
	CustomerAccountId int64   `json:"customer_account_id"`
	AccountNumber     string  `json:"account_number"`
	AccountAlias      string  `json:"account_alias"`
	AccountType       string  `json:"account_type,omitempty"`
	Reference         string  `json:"reference,omitempty"`
	Balance           float64 `json:"balance"`
	FrozenAmount      float64 `json:"frozen_amount"`
	BalanceBefore     float64 `json:"balance_before"`
	DateCreated       string  `json:"date_created"`
	Active            int     `json:"active"`
}

type CreateCustomerAccountApiResponse struct {
	StatusCode    string                      `json:"statusCode"`
	StatusMessage string                      `json:"statusMessage"`
	Result        *CustomerAccountResponseObj `json:"result,omitempty"`
}

type CustomerApprovalAccountsResponse struct {
	StatusCode    string                        `json:"statusCode"`
	StatusMessage string                        `json:"statusMessage"`
	Result        []*CustomerAccountResponseObj `json:"result,omitempty"`
}

type CustomerAccountHistoryData struct {
	CustomerAccountHistoryId int64
	CustomerAccount          string
	DebitAmount              float64
	CreditAmount             float64
	TransactionDate          string
	Reference                string
	CreatedBy                int
	ModifiedBy               int
}

type CustomersAccountHistoryResponseDTO struct {
	Success    bool
	Result     []*CustomerAccountHistoryData
	StatusDesc string
}

type CustomersAccountHistoryApiResponseDTO struct {
	StatusCode    string
	Result        []*CustomerAccountHistoryData
	StatusMessage string
}

type CustomerAccountStatementApiData struct {
	TransactionDate        string
	TransactionDescription string
	Debit                  float64
	Credit                 float64
}

type CustomerAccountStatementData struct {
	Account                string
	TransactionDate        string
	TransactionDescription string
	Debit                  float64
	Credit                 float64
	TransactionType        string
}

type CustomerAccountStatementApiResponseDTO struct {
	StatusCode    int
	Result        []*CustomerAccountStatementApiData
	StatusMessage string
}

type CustomersAccountStatementResponseDTO struct {
	Success    bool
	Result     []*CustomerAccountStatementData
	StatusDesc string
}

type Status struct {
	StatusId     int64     `orm:"auto"`
	Status       string    `orm:"size(128)"`
	StatusCode   string    `orm:"size(128)"`
	DateCreated  time.Time `orm:"type(datetime)"`
	DateModified time.Time `orm:"type(datetime)"`
	CreatedBy    int
	ModifiedBy   int
	Active       int
}

type Payments struct {
	PaymentId       int64 `orm:"auto"`
	InitiatedBy     int64
	Sender          *Customers `orm:"rel(fk);column(sender)"`
	Reciever        *Users     `orm:"rel(fk);column(reciever)"`
	Amount          float64
	Commission      float64
	Charge          float64
	Narration       string  `orm:"size(255)"`
	PaymentProof    string  `orm:"null"`
	Status          *Status `orm:"rel(fk);column(status)"`
	PaymentAccount  int
	ReferenceNumber string
	DateCreated     time.Time `orm:"type(datetime)"`
	DateModified    time.Time `orm:"type(datetime)"`
	CreatedBy       int64
	ModifiedBy      int64
	Active          int
}
