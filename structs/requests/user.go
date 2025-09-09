package requests

type LoginApiRequest struct {
	PhoneNumber  string
	Password     string
	SourceSystem string
	ClientId     string
}

type LoginRequest struct {
	PhoneNumber string
	Password    string
	ClientId    string `validate:"omitempty"`
}

type RegisterRequest struct {
	ClientId     int64  `validate:"required"`
	FirstName    string `validate:"required"`
	LastName     string `validate:"required"`
	Email        string
	Gender       string `validate:"required"`
	MobileNumber string `validate:"required"`
	Password     string `validate:"required"`
	Dob          string `validate:"required"`
	Username     string `validate:"required"`
}

type RegisterApiRequest struct {
	FirstName    string
	LastName     string
	Gender       string
	MobileNumber string
	ClientId     string
	Password     string
	Dob          string
	Username     string
}

type OpenAccountRequest struct {
	ClientId     int64  `validate:"required"`
	FirstName    string `validate:"required"`
	LastName     string `validate:"required"`
	Gender       string `validate:"required"`
	MobileNumber string `validate:"required"`
}

type OpenAccountApiRequest struct {
	FirstName    string
	LastName     string
	Gender       string
	MobileNumber string
	ClientId     string
}

type VerifyCustomerApiRequest struct {
	Username     string
	FirstName    string
	LastName     string
	Email        string
	Dob          string
	Gender       string
	MobileNumber string
	ClientId     string
}

type ActivateVerifiedCustomerRequest struct {
	MobileNumber string
	Username     string
}

type ActivateVerifiedCustomerApiRequest struct {
	MobileNumber string
	Username     string
	ClientId     string
}

type NumberExistsApiRequest struct {
	MobileNumber string
	ClientId     string
}

type AccountBalanceApiRequest struct {
	AccountNumber string
	ClientId      string
}

type ClientApiRequest struct {
	ClientId string
}

type MobileNumberRequest struct {
	MobileNumber string
}

type UsernameRequest struct {
	Username string
}

type AddCustomer struct {
	Email        string
	Name         string
	PhoneNumber  string
	Location     string
	IdType       string
	IdNumber     string
	ImagePath    string
	AddedBy      string
	CustomerType string
	Branch       string
	Dob          string
	Status       string
}

type ResetPinRequest struct {
	Number      string
	OldPassword string
	NewPassword string
	ClientId    string
}

type ResetPinApiRequest struct {
	Number      string
	OldPassword string
	NewPassword string
	ClientId    string
}

type SignIn struct {
	Email    string
	Password string
}

type ChangePassword struct {
	OldPassword string
	NewPassword string
}

type EmailReq struct {
	Email string
}

type EmailTokenReq struct {
	Email    string
	Password string
	Token    string
}

type EmailPasswordReq struct {
	Email    string
	Password string
}

type UpdateCustomer struct {
	Email            string
	Name             string
	PhoneNumber      string
	Location         string
	IdType           string
	IdNumber         string
	ImagePath        string
	EmergencyContact []EditCustomerEmergencyContact
	Guarantor        []EditCustomerGuarantor
}

type AddCustomerEmergencyContact struct {
	Name        string
	PhoneNumber string
}

type AddCustomerGuarantor struct {
	Name        string
	PhoneNumber string
}

type EditCustomerEmergencyContact struct {
	CustomerEmergencyContactId int64
	Name                       string
	PhoneNumber                string
}

type EditCustomerGuarantor struct {
	CustomerGuarantorId int64
	Name                string
	PhoneNumber         string
}

type Registration struct {
	Email       string
	FirstName   string
	LastName    string
	ImagePath   string
	PhoneNumber string
	Password    string
	Token       string
}

type RegisterUser struct {
	Email       string
	Name        string
	Gender      string
	PhoneNumber string
	Password    string
	Dob         string
	RoleId      string
	Branch      string
}

type AddCustomerCredential struct {
	CustomerId int64
	Username   string
	Password   string
	Pin        string
}

type CreateCustomerAccountApiRequest struct {
	AccountNumber string
	AccountAlias  string
	Balance       float64
	CreatedBy     int
	Active        int
}
