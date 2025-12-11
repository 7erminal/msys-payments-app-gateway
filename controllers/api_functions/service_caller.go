package apifunctions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"msys_payment_app_gateway/api"
	"msys_payment_app_gateway/structs/requests"
	"msys_payment_app_gateway/structs/responses"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func GetCorporatives(c *beego.Controller) (resp responses.CorporativeApiResponse) {
	host, _ := beego.AppConfig.String("clientBaseUrl")

	logs.Info("Getting corporatives")
	request := api.NewRequest(
		host,
		"/v2/clients",
		api.GET)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}

	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CorporativeApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func Login(c *beego.Controller, req requests.LoginApiRequest) (resp responses.LoginApiResponse) {
	host, _ := beego.AppConfig.String("authenticationBaseUrl")

	logs.Info("Verify pin ", req.PhoneNumber, " for ", req.Password)

	request := api.NewRequest(
		host,
		"/v1/auth/validate-customer-credentials/token",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["clientId"] = req.ClientId

	request.InterfaceParams["Username"] = req.PhoneNumber
	request.InterfaceParams["Password"] = req.Password

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.LoginApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func Register(c *beego.Controller, req requests.AddCustomer) (resp responses.CustomerResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Sending name ", req.Name)

	// Get date
	// now := time.Now()
	// y, m, d := now.Date()
	// d_str := strconv.Itoa(d)
	// m_str := strconv.Itoa(int(m))
	// if len(d_str) < 2 {
	// 	d_str = "0" + d_str
	// }
	// if len(m_str) < 2 {
	// 	m_str = "0" + m_str
	// }
	// dob := strconv.Itoa(y) + "/" + m_str + "/" + d_str

	request := api.NewRequest(
		host,
		"/v1/customers/add-customer",
		api.POST)
	request.Params["Name"] = req.Name
	request.Params["Email"] = req.Email
	request.Params["IdType"] = req.IdType
	request.Params["PhoneNumber"] = req.PhoneNumber
	request.Params["IdNumber"] = req.IdNumber
	request.Params["Dob"] = req.Dob
	request.Params["AddedBy"] = req.AddedBy
	request.Params["Location"] = req.Location
	request.Params["Branch"] = req.Branch
	request.Params["Status"] = req.Status
	if req.ImagePath != "" {
		request.FileField["CustomerImage"] = req.ImagePath
	}
	request.Params["Category"] = req.CustomerType
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CustomerResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	// logs.Info("Response received ", c.Data["json"])
	// logs.Info("Access token ", data["access_token"])
	// logs.Info("Expires in ", data["expires_in"])
	// logs.Info("Scope is ", data["scope"])
	// logs.Info("Token Type is ", data["token_type"])
	// logs.Info("Response received ", c.Data["json"])
	// logs.Info("Access token ", data.Access_token)
	// logs.Info("Expires in ", data.Expires_in)
	// logs.Info("Scope is ", data.Scope)
	// logs.Info("Token Type is ", data.Token_type)

	return data
}

func OpenAccount(c *beego.Controller, req requests.OpenAccountApiRequest) (resp responses.RegisterApiResponse) {
	host, _ := beego.AppConfig.String("clientBaseUrl")

	logs.Info("Registering user ", req.MobileNumber, " with name ", req.FirstName, " ", req.LastName)
	request := api.NewRequest(
		host,
		"/v2/api/2/register-customer",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["clientId"] = req.ClientId

	request.InterfaceParams["FirstName"] = req.FirstName
	request.InterfaceParams["LastName"] = req.LastName
	request.InterfaceParams["Gender"] = req.Gender
	request.InterfaceParams["MobileNumber"] = req.MobileNumber
	request.InterfaceParams["ChargeAmount"] = req.ChargeAmount
	request.InterfaceParams["Source"] = req.Source

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.RegisterApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func AddAccount(c *beego.Controller, req requests.AddCustomerAccountApiRequest) (resp responses.CustomerAccountApiResponse) {
	host, _ := beego.AppConfig.String("accountBaseUrl")

	logs.Info("Adding account for user ", req.AccountNumber, " with alias ", req.AccountAlias)

	request := api.NewRequest(
		host,
		"/v1/customer-accounts/add-account",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}

	request.InterfaceParams["account_number"] = req.AccountNumber
	request.InterfaceParams["account_alias"] = req.AccountAlias
	request.InterfaceParams["created_by"] = req.CreatedBy
	request.InterfaceParams["active"] = req.Active

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CustomerAccountApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func VerifyCustomer(c *beego.Controller, req requests.VerifyCustomerApiRequest) (resp responses.VerifyCustomerApiResponse) {
	host, _ := beego.AppConfig.String("clientBaseUrl")

	logs.Info("Registering user ", req.MobileNumber, " with name ", req.FirstName, " ", req.LastName, " username ", req.Username, " email ", req.Email, " dob ", req.Dob)
	request := api.NewRequest(
		host,
		"/v2/api/verify-customer",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["clientId"] = req.ClientId

	request.InterfaceParams["FirstName"] = req.FirstName
	request.InterfaceParams["LastName"] = req.LastName
	request.InterfaceParams["Gender"] = req.Gender
	request.InterfaceParams["MobileNumber"] = req.MobileNumber
	request.InterfaceParams["Username"] = req.Username
	request.InterfaceParams["Email"] = req.Email
	request.InterfaceParams["Dob"] = req.Dob

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.VerifyCustomerApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func FetchApprovedAccounts(c *beego.Controller, clientId string) (resp responses.CustAccountsApiResponse) {
	host, _ := beego.AppConfig.String("clientBaseUrl")

	request := api.NewRequest(
		host,
		"/v2/api/fetch-approved-customers",
		api.GET)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["clientId"] = clientId

	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CustAccountsApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func ActivateVerifiedCustomer(c *beego.Controller, req requests.ActivateVerifiedCustomerApiRequest) (resp responses.VerifyApprovedCustomerApiResponse) {
	host, _ := beego.AppConfig.String("clientBaseUrl")

	logs.Info("Activating verified customer ", req.MobileNumber, " with username ", req.Username)
	request := api.NewRequest(
		host,
		"/v2/api/activate-verified-customers",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["clientId"] = req.ClientId

	request.InterfaceParams["MobileNumber"] = req.MobileNumber
	request.InterfaceParams["Username"] = req.Username

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.VerifyApprovedCustomerApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func NumberExists(c *beego.Controller, req requests.NumberExistsApiRequest) (resp responses.BuyAirtimeResponse) {
	host, _ := beego.AppConfig.String("clientBaseUrl")

	logs.Info("Checking if number exists ", req.MobileNumber)
	request := api.NewRequest(
		host,
		"/v2/api/existing-number/"+req.MobileNumber,
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["clientId"] = req.ClientId

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.BuyAirtimeResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func GetCustomerAccounts(c *beego.Controller, customerId string) (resp responses.CustomerApprovalAccountsResponse) {
	host, _ := beego.AppConfig.String("accountBaseUrl")

	logs.Info("Getting customer accounts for customer ID ", customerId)
	request := api.NewRequest(
		host,
		"/v1/customer-accounts/customer/"+customerId,
		api.GET)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}

	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CustomerApprovalAccountsResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func GetCustomerAccount(c *beego.Controller, accountNumber string) (resp responses.CreateCustomerAccountApiResponse) {
	host, _ := beego.AppConfig.String("accountBaseUrl")

	logs.Info("Getting customer accounts for account number::: ", accountNumber)
	request := api.NewRequest(
		host,
		"/v1/customer-accounts/account/"+accountNumber,
		api.GET)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}

	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CreateCustomerAccountApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func AddCustomerAccount(c *beego.Controller, req requests.CreateCustomerAccountApiRequest) (resp responses.CreateCustomerAccountApiResponse) {
	host, _ := beego.AppConfig.String("accountBaseUrl")

	logs.Info("Creating customer account for account number ", req.AccountNumber)
	request := api.NewRequest(
		host,
		"/v1/customer-accounts/add-account",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.InterfaceParams["account_number"] = req.AccountNumber
	request.InterfaceParams["account_alias"] = req.AccountAlias
	request.InterfaceParams["active"] = req.Active
	request.InterfaceParams["created_by"] = "1"
	request.InterfaceParams["customer_id"] = req.CustomerId
	request.InterfaceParams["account_type"] = req.AccountType
	request.InterfaceParams["reference"] = req.Reference

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CreateCustomerAccountApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func GetCustomerAccountStatement(c *beego.Controller, accountNumber string, clientId string) (resp responses.CustomerAccountStatementApiResponseDTO) {
	host, _ := beego.AppConfig.String("clientBaseUrl")

	logs.Info("Getting customer account statement for account number::: ", accountNumber)
	request := api.NewRequest(
		host,
		"/v2/api/v2/account-statement",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["clientId"] = clientId

	request.InterfaceParams["AccountNumber"] = accountNumber
	request.InterfaceParams["FromDate"] = "2023-01-01"
	request.InterfaceParams["ToDate"] = fmt.Sprintf("%d-%02d-%02d", 2024, 12, 31)

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CustomerAccountStatementApiResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func GetCustomerAccountHistory(c *beego.Controller, accountNumber string) (resp responses.CustomersAccountHistoryApiResponseDTO) {
	host, _ := beego.AppConfig.String("accountBaseUrl")

	logs.Info("Getting customer accounts for account number::: ", accountNumber)
	request := api.NewRequest(
		host,
		"/v1/customer-accounts/account-history/"+accountNumber,
		api.GET)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}

	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CustomersAccountHistoryApiResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func LogTransaction(c *beego.Controller, req requests.LogTransactionApiRequest) (resp responses.TransactionApiResponse) {
	host, _ := beego.AppConfig.String("transactionBaseUrl")

	logs.Info("Logging transaction for account number ", req.SourceAccountNumber)

	request := api.NewRequest(
		host,
		"/v2/transactions",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["sourceSystem"] = "MSYS_PAYMENT_APP_GATEWAY"

	type ExtraData struct {
		ExtraDetails1 string
		ExtraDetails2 string
		ExtraDetails3 string
	}

	extraData := ExtraData{
		ExtraDetails1: req.ExtraDetails1,
		ExtraDetails2: req.ExtraDetails2,
		ExtraDetails3: req.ExtraDetails3,
	}

	request.InterfaceParams["RequestId"] = req.RequestId
	request.InterfaceParams["Amount"] = req.Amount
	request.InterfaceParams["ServiceCode"] = req.ServiceCode
	request.InterfaceParams["Source"] = req.SourceAccountNumber
	request.InterfaceParams["PhoneNumber"] = req.PhoneNumber
	request.InterfaceParams["Destination"] = req.DestinationAccountNumber
	request.InterfaceParams["ExtraData"] = extraData
	request.InterfaceParams["BillerCode"] = req.TransactionReference
	request.InterfaceParams["ClientReference"] = req.ExternalReferenceNumber
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.TransactionApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func GetBilTransaction(c *beego.Controller, id string) (resp responses.LogTransactionResponse) {
	host, _ := beego.AppConfig.String("transactionBaseUrl")

	logs.Info("Getting transaction for ", id)
	request := api.NewRequest(
		host,
		"/v2/transactions/"+id,
		api.GET)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	// request.HeaderField["clientId"] = req.ClientId

	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.LogTransactionResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func GetBilTransactionWithTransactionRef(c *beego.Controller, id string) (resp responses.LogTransactionResponse) {
	host, _ := beego.AppConfig.String("transactionBaseUrl")

	logs.Info("Getting transaction for ", id)
	request := api.NewRequest(
		host,
		"/v2/transactions/ref/"+id,
		api.GET)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	// request.HeaderField["clientId"] = req.ClientId

	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.LogTransactionResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func GetCustomerAirtimeBundleBillPayHistory(c *beego.Controller, query string) (resp responses.BilTransactionsApiResponse) {
	host, _ := beego.AppConfig.String("airtimeBaseUrl")

	logs.Info("Getting airtime and bundle history::: ", query)
	request := api.NewRequest(
		host,
		"/v1/requests/bil-transactions",
		api.GET)
	request.Params["query"] = query
	// request.Params = {"UserId": strconv.Itoa(int(userid))}

	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.BilTransactionsApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func GetCustomerBillPaymentHistory(c *beego.Controller, query string) (resp responses.BilTransactionsApiResponse) {
	host, _ := beego.AppConfig.String("billpaymentBaseUrl")

	logs.Info("Getting bill payment transactions::: ", query)
	request := api.NewRequest(
		host,
		"/v1/bill-payment/bil-transactions",
		api.GET)
	request.Params["query"] = query
	// request.Params = {"UserId": strconv.Itoa(int(userid))}

	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.BilTransactionsApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func ReportAccountAnomaly(c *beego.Controller, req requests.CustomerAccountAnomaliesRequest) (resp responses.CreateCustomerAccountApiResponse) {
	host, _ := beego.AppConfig.String("accountBaseUrl")

	logs.Info("Creating customer account for account number ", req.AccountNumber)
	request := api.NewRequest(
		host,
		"/v1/customer-account-anomalies/",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.InterfaceParams["AccountNumber"] = req.AccountNumber
	request.InterfaceParams["Amount"] = req.Amount
	request.InterfaceParams["Desc"] = req.Desc
	request.InterfaceParams["Balance"] = req.Balance
	request.InterfaceParams["CheckedBalance"] = req.CheckedBalance
	request.InterfaceParams["CreatedBy"] = 1
	request.InterfaceParams["ModifiedBy"] = 1
	request.InterfaceParams["Active"] = req.Active
	//

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CreateCustomerAccountApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func GetAccountBalance(c *beego.Controller, req requests.AccountBalanceApiRequest) (resp responses.AccountBalanceApiResponse) {
	host, _ := beego.AppConfig.String("clientBaseUrl")

	logs.Info("Getting account balance for ", req.AccountNumber)
	request := api.NewRequest(
		host,
		"/v2/api/v2/account-balance",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["clientId"] = req.ClientId

	request.InterfaceParams["AccountNumber"] = req.AccountNumber

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.AccountBalanceApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func UpdateAccountBalance(c *beego.Controller, req requests.UpdateAccountBalanceApiRequest) (resp responses.CreateCustomerAccountApiResponse) {
	host, _ := beego.AppConfig.String("accountBaseUrl")

	logs.Info("Getting account balance for ", req.AccountNumber, " to ", req.Balance)

	accountIdStr := fmt.Sprintf("%d", req.AccountId)
	request := api.NewRequest(
		host,
		"/v1/customer-accounts/update-balance/"+accountIdStr,
		api.PUT)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	// request.HeaderField["clientId"] = req.ClientId

	request.InterfaceParams["balance"] = req.Balance
	request.InterfaceParams["modified_by"] = req.ModifiedBy
	request.InterfaceParams["reason"] = req.Reason

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CreateCustomerAccountApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func DebitAccountPro(c *beego.Controller, req requests.DebitAccountRequestV2) (resp responses.DebitAccountV2Response) {
	host, _ := beego.AppConfig.String("clientBaseUrl")

	logs.Info("Creating customer account for account number ", req.AccountNumber)
	logs.Info("Clien ID ", req.ClientId)
	request := api.NewRequest(
		host,
		"/v2/api/v2/debit-account",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["clientId"] = req.ClientId

	request.InterfaceParams["AccountNumber"] = req.AccountNumber
	request.InterfaceParams["Amount"] = req.Amount
	request.InterfaceParams["Reference"] = req.Reference
	request.InterfaceParams["Channel"] = req.Channel

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.DebitAccountV2Response
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func DebitAccount(c *beego.Controller, req requests.DebitAccountRequest) (resp responses.CreateCustomerAccountApiResponse) {
	host, _ := beego.AppConfig.String("accountBaseUrl")

	logs.Info("Creating customer account for account number ", req.AccountId)
	logs.Info("Amount to debit being sent to accounts api ", req.Amount, " and ", fmt.Sprintf("%v", req.Amount))
	request := api.NewRequest(
		host,
		"/v1/customer-accounts/debit-account/"+fmt.Sprintf("%d", req.AccountId),
		api.PUT)
	// request.Params["username"] = username
	request.Params["amount"] = fmt.Sprintf("%v", req.Amount)
	request.Params["modified_by"] = fmt.Sprintf("%v", req.ModifiedBy)
	request.Params["reason"] = req.Reason

	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CreateCustomerAccountApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func CreditAccountPro(c *beego.Controller, req requests.CreditAccountRequestV2) (resp responses.CreditAccountV2Response) {
	host, _ := beego.AppConfig.String("clientBaseUrl")

	logs.Info("Creating customer account for account number ", req.AccountNumber)
	request := api.NewRequest(
		host,
		"/v2/api/v2/credit-account",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["clientId"] = req.ClientId

	request.InterfaceParams["AccountNumber"] = req.AccountNumber
	request.InterfaceParams["Amount"] = req.Amount
	request.InterfaceParams["Reference"] = req.Reference
	request.InterfaceParams["Channel"] = req.Channel

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CreditAccountV2Response
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func CreditAccount(c *beego.Controller, req requests.CreditAccountRequest) (resp responses.CreateCustomerAccountApiResponse) {
	host, _ := beego.AppConfig.String("accountBaseUrl")

	logs.Info("Creating customer account for account number ", req.AccountId)
	request := api.NewRequest(
		host,
		"/v1/customer-accounts/credit-account/"+fmt.Sprintf("%d", req.AccountId),
		api.PUT)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.Params["amount"] = fmt.Sprintf("%v", req.Amount)
	request.Params["modified_by"] = fmt.Sprintf("%v", req.ModifiedBy)
	request.Params["reason"] = req.Reason

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CreateCustomerAccountApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func ListCustomerAccounts(c *beego.Controller, req requests.NumberExistsApiRequest) (resp responses.ListCustAccountsApiResponse) {
	host, _ := beego.AppConfig.String("clientBaseUrl")

	logs.Info("Listing customer accounts for number ", req.MobileNumber, " for client ", req.ClientId)
	request := api.NewRequest(
		host,
		"/v2/api/v2/list-cust-accounts",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["clientId"] = req.ClientId

	request.InterfaceParams["Number"] = req.MobileNumber

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.ListCustAccountsApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func CustomerAccountStatement(c *beego.Controller, req requests.AccountBalanceApiRequest) (resp responses.ListCustAccountsApiResponse) {
	host, _ := beego.AppConfig.String("clientBaseUrl")

	logs.Info("Getting statement ", req.AccountNumber, " for client ", req.ClientId)
	request := api.NewRequest(
		host,
		"/v2/api/v2/account-statment",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["clientId"] = req.ClientId

	request.InterfaceParams["AccountNumber"] = req.AccountNumber

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.ListCustAccountsApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func GetAccountInfo(c *beego.Controller, req requests.ClientApiRequest) (resp responses.BuyAirtimeResponse) {
	host, _ := beego.AppConfig.String("clientBaseUrl")

	logs.Info("Getting account info for client ", req.ClientId)
	request := api.NewRequest(
		host,
		"/v2/api/get-contact-info/",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["clientId"] = req.ClientId

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.BuyAirtimeResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func NameInquiry(c *beego.Controller, req requests.NumberExistsApiRequest) (resp responses.NameInquiryApiResponse) {
	host, _ := beego.AppConfig.String("clientBaseUrl")

	logs.Info("Listing customer accounts for client ", req.ClientId)
	request := api.NewRequest(
		host,
		"/v2/api/name-inquiry/"+req.MobileNumber,
		api.GET)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["clientId"] = req.ClientId

	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.NameInquiryApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func NameInquiryViaMobileMoney(c *beego.Controller, req requests.NameInquiryApiRequestDTO) (resp responses.NameInquiryApiResponseDTO) {
	host, _ := beego.AppConfig.String("paymentBaseUrl")

	logs.Info("Name inquiry ", req.CustomerMsisdn)
	request := api.NewRequest(
		host,
		"/v1/payments/name-inquiry/",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.InterfaceParams["CustomerMsisdn"] = req.CustomerMsisdn
	request.InterfaceParams["Channel"] = req.Channel

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.NameInquiryApiResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func GetPaymentMethods(c *beego.Controller) (resp responses.PaymentMethodsApiResponseDTO) {
	host, _ := beego.AppConfig.String("paymentBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/payment-methods/",
		api.GET)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.PaymentMethodsApiResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func MakePayment(c *beego.Controller, req requests.MakePaymentApiRequestDTO) (resp responses.PaymentApiResponseDTO) {
	host, _ := beego.AppConfig.String("paymentBaseUrl")

	logs.Info("Requesting Money ", req.Amount, " from ", req.InitiatedBy)
	request := api.NewRequest(
		host,
		"/v1/payments/",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.InterfaceParams["InitiatedBy"] = req.InitiatedBy
	request.InterfaceParams["Amount"] = req.Amount
	request.InterfaceParams["Service"] = req.Service
	request.InterfaceParams["Sender"] = req.Sender
	request.InterfaceParams["PaymentMethod"] = req.PaymentMethod
	request.InterfaceParams["SenderAccount"] = req.SenderAccount
	request.InterfaceParams["ReceiverAccount"] = req.ReceiverAccount
	request.InterfaceParams["Reciever"] = req.Reciever
	request.InterfaceParams["TransactionId"] = req.TransactionId
	request.InterfaceParams["PaymentProofUrl"] = req.PaymentProofUrl
	request.InterfaceParams["ReferenceNumber"] = req.ReferenceNumber
	request.InterfaceParams["CallThirdParty"] = req.CallThirdParty
	request.InterfaceParams["Operator"] = req.Operator
	request.InterfaceParams["Network"] = req.Network
	request.InterfaceParams["ServiceNetwork"] = req.ServiceNetwork
	request.InterfaceParams["ServicePackage"] = req.ServicePackage

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.PaymentApiResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func RequestMoneyViaMobileMoney(c *beego.Controller, req requests.MomoPaymentApiRequestDTO) (resp responses.PaymentApiResponseDTO) {
	host, _ := beego.AppConfig.String("paymentBaseUrl")

	logs.Info("Requesting Money ", req.Amount, " from ", req.CustomerMsisdn)
	request := api.NewRequest(
		host,
		"/v1/request-money/momo",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.InterfaceParams["CustomerMsisdn"] = req.CustomerMsisdn
	request.InterfaceParams["CustomerEmail"] = req.CustomerEmail
	request.InterfaceParams["CustomerName"] = req.CustomerName
	request.InterfaceParams["Amount"] = req.Amount
	request.InterfaceParams["PrimaryCallbackUrl"] = req.PrimaryCallbackUrl
	request.InterfaceParams["Description"] = req.Description
	request.InterfaceParams["ClientReference"] = req.ClientReference
	request.InterfaceParams["Operator"] = req.Operator
	request.InterfaceParams["Channel"] = req.Channel
	request.InterfaceParams["PaymentId"] = req.PaymentId

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.PaymentApiResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func SendMoneyViaMobileMoney(c *beego.Controller, req requests.MomoPaymentApiRequestDTO) (resp responses.PaymentApiResponseDTO) {
	host, _ := beego.AppConfig.String("paymentBaseUrl")

	logs.Info("Requesting Money ", req.Amount, " from ", req.CustomerMsisdn)
	request := api.NewRequest(
		host,
		"/v1/payments/momo",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.InterfaceParams["CustomerMsisdn"] = req.CustomerMsisdn
	request.InterfaceParams["CustomerEmail"] = req.CustomerEmail
	request.InterfaceParams["CustomerName"] = req.CustomerName
	request.InterfaceParams["Amount"] = req.Amount
	request.InterfaceParams["PrimaryCallbackUrl"] = req.PrimaryCallbackUrl
	request.InterfaceParams["Description"] = req.Description
	request.InterfaceParams["ClientReference"] = req.ClientReference
	request.InterfaceParams["Operator"] = req.Operator
	request.InterfaceParams["Channel"] = req.Channel
	request.InterfaceParams["PaymentId"] = req.PaymentId

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.PaymentApiResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func ResetPin(c *beego.Controller, req requests.ResetPinApiRequest) (resp responses.ResetPinApiResponse) {
	host, _ := beego.AppConfig.String("clientBaseUrl")

	logs.Info("Resetting pin for number ", req.Number)
	request := api.NewRequest(
		host,
		"/v2/api/reset-pin",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["clientId"] = req.ClientId
	request.InterfaceParams["Number"] = req.Number
	request.InterfaceParams["OldPassword"] = req.OldPassword
	request.InterfaceParams["NewPassword"] = req.NewPassword

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.ResetPinApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func GetBundles(c *beego.Controller, req requests.DataBundlesListFormulatedRequest) (resp responses.DataBundlesListResponse) {
	host, _ := beego.AppConfig.String("airtimeBaseUrl")

	logs.Info("Getting data bundles ", req.NetworkId, " for ", req.DestinationAccount)

	request := api.NewRequest(
		host,
		"/v1/requests/bundles/"+req.NetworkId+"/"+req.DestinationAccount,
		api.GET)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["PhoneNumber"] = req.PhoneNumber
	request.HeaderField["SourceSystem"] = req.SourceSystem
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.DataBundlesListResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func RequestMoney(c *beego.Controller, req requests.PaymentApiRequestDTO) (resp responses.BuyAirtimeResponse) {
	host, _ := beego.AppConfig.String("paymentBaseUrl")

	logs.Info("Requesting Money ", req.Amount, " from ", req.PaymentMethod)

	request := api.NewRequest(
		host,
		"/v1/payments/request",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}

	request.InterfaceParams["InitiatedBy"] = req.InitiatedBy
	request.InterfaceParams["Amount"] = req.Amount
	request.InterfaceParams["Service"] = req.Service
	request.InterfaceParams["Sender"] = req.Sender
	request.InterfaceParams["PaymentMethod"] = req.PaymentMethod
	request.InterfaceParams["Reciever"] = req.Reciever
	request.InterfaceParams["TransactionId"] = req.TransactionId
	request.InterfaceParams["PaymentProofUrl"] = req.PaymentProofUrl
	request.InterfaceParams["ReferenceNumber"] = req.ReferenceNumber
	request.InterfaceParams["CallThirdParty"] = req.CallThirdParty
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.BuyAirtimeResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func NameInquiryPhoneNumber(c *beego.Controller, req requests.BuyAirtimeFormulatedRequest) (resp responses.BuyAirtimeResponse) {
	host, _ := beego.AppConfig.String("paymentBaseUrl")

	logs.Info("Buying Airtime ", req.Amount, " for ", req.Destination)

	request := api.NewRequest(
		host,
		"/v1/requests/buy-airtime",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["PhoneNumber"] = req.PhoneNumber
	request.HeaderField["SourceSystem"] = req.SourceSystem

	request.InterfaceParams["destination"] = req.Destination
	request.InterfaceParams["amount"] = req.Amount
	request.InterfaceParams["network"] = req.Network
	request.InterfaceParams["request_id"] = req.TransactionId
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.BuyAirtimeResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func BuyAirtime(c *beego.Controller, req requests.BuyAirtimeFormulatedRequest) (resp responses.BuyAirtimeResponse) {
	host, _ := beego.AppConfig.String("airtimeBaseUrl")

	logs.Info("Buying Airtime ", req.Amount, " for ", req.Destination)

	request := api.NewRequest(
		host,
		"/v1/requests/buy-airtime",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["PhoneNumber"] = req.PhoneNumber
	request.HeaderField["SourceSystem"] = req.SourceSystem

	request.InterfaceParams["destination"] = req.Destination
	request.InterfaceParams["amount"] = req.Amount
	request.InterfaceParams["network"] = req.Network
	request.InterfaceParams["transaction_id"] = req.TransactionId
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.BuyAirtimeResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func BuyDataBundle(c *beego.Controller, req requests.BuyDataBundleFormulatedRequest) (resp responses.BuyDataBundleResponse) {
	host, _ := beego.AppConfig.String("airtimeBaseUrl")

	logs.Info("Buying Bundle ", req.Amount, " for ", req.Destination)

	request := api.NewRequest(
		host,
		"/v1/requests/buy-bundle",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["PhoneNumber"] = req.PhoneNumber
	request.HeaderField["SourceSystem"] = req.SourceSystem

	request.InterfaceParams["destination"] = req.Destination
	request.InterfaceParams["amount"] = req.Amount
	request.InterfaceParams["network"] = req.Network
	request.InterfaceParams["bundle_id"] = req.BundleId
	request.InterfaceParams["transaction_id"] = req.TransactionId
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	// logs.Info("Raw response received is ", res)
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.BuyDataBundleResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func Callback(c *beego.Controller, req requests.CallbackFormulateRequest) (resp responses.CallbackResponse) {
	host, _ := beego.AppConfig.String("transactionBaseUrl")

	logs.Info("Sending callback ", req.ResponseCode, " for ", req.Data.TransactionId)
	request := api.NewRequest(
		host,
		"/v2/callback/process",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}

	request.InterfaceParams["ResponseCode"] = req.ResponseCode
	request.InterfaceParams["Data"] = req.Data

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	// logs.Info("Raw response received is ", res)
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CallbackResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func ReceivePaymentCallback(c *beego.Controller, req requests.PaymentCallbackData) (resp responses.CallbackAPIResponse) {
	host, _ := beego.AppConfig.String("paymentBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/callback/process",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}

	request.InterfaceParams["AmountCharged"] = req.AmountCharged
	request.InterfaceParams["ClientReference"] = req.ClientReference
	request.InterfaceParams["TransactionId"] = req.TransactionId
	request.InterfaceParams["Description"] = req.Description
	request.InterfaceParams["ExternalTransactionId"] = req.ExternalTransactionId
	request.InterfaceParams["Amount"] = req.Amount
	request.InterfaceParams["Charges"] = req.Charges
	request.InterfaceParams["AmountAfterCharges"] = req.AmountAfterCharges
	request.InterfaceParams["PaymentDate"] = req.PaymentDate
	request.InterfaceParams["OrderId"] = req.OrderId
	request.InterfaceParams["Status"] = req.Status

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	// logs.Info("Raw response received is ", res)
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CallbackAPIResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func PayDSTVBill(c *beego.Controller, req requests.DSTVPaymentRequest) (resp responses.DSTVBillPaymentApiResponse) {
	host, _ := beego.AppConfig.String("billpaymentBaseUrl")

	logs.Info("Paying DSTV ", req.Amount, " for ", req.DestinationAccount)

	request := api.NewRequest(
		host,
		"/v1/bill-payment/pay-dstv-bill",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["PhoneNumber"] = req.PhoneNumber
	request.HeaderField["SourceSystem"] = req.SourceSystem

	request.InterfaceParams["DestinationAccount"] = req.DestinationAccount
	request.InterfaceParams["Amount"] = req.Amount
	request.InterfaceParams["PackageType"] = req.PackageType
	request.InterfaceParams["TransactionId"] = req.TransactionId
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	// logs.Info("Raw response received is ", res)
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.DSTVBillPaymentApiResponse
	json.Unmarshal(read, &data)
	if err := json.Unmarshal(read, &data); err != nil {
		logs.Info("Failed to unmarshal response: %v", err)
		logs.Info("Raw response: %s", string(read))
	} else {
		logs.Info("Resp is %+v", data)
	}
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func DSTVAccountQuery(c *beego.Controller, req requests.DSTVAccountQueryRequest) (resp responses.AccountQueryResponse) {
	host, _ := beego.AppConfig.String("billpaymentBaseUrl")

	logs.Info("DSTV query for account ", req.AccountNumber)

	request := api.NewRequest(
		host,
		"/v1/bill-payment/dstv-account-query/"+req.AccountNumber,
		api.GET)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["PhoneNumber"] = req.PhoneNumber
	request.HeaderField["SourceSystem"] = req.SourceSystem

	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	// logs.Info("Raw response received is ", res)
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.AccountQueryResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func AccountQuery(c *beego.Controller, req requests.BillPaymentAccountQueryApiRequest) (resp responses.AccountQueryResponse) {
	host, _ := beego.AppConfig.String("billpaymentBaseUrl")

	logs.Info("Account query for ", req.AccountNumber, " with BillerCode ", req.BillerCode)

	request := api.NewRequest(
		host,
		"/v1/bill-payment/account-query/"+req.BillerCode+"/"+req.AccountNumber,
		api.GET)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["PhoneNumber"] = req.PhoneNumber
	request.HeaderField["SourceSystem"] = req.SourceSystem

	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	// logs.Info("Raw response received is ", res)
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.AccountQueryResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func GhanaWaterAccountQuery(c *beego.Controller, req requests.BillPaymentAccountQueryApiRequest) (resp responses.AccountQueryResponse) {
	host, _ := beego.AppConfig.String("billpaymentBaseUrl")

	logs.Info("Account query for ", req.AccountNumber, " with BillerCode ", req.BillerCode)

	request := api.NewRequest(
		host,
		"/v1/bill-payment/ghana-water-account-query/"+req.AccountNumber+"/"+req.PhoneNumber,
		api.GET)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["PhoneNumber"] = req.PhoneNumber
	request.HeaderField["SourceSystem"] = req.SourceSystem

	// request.Params[""] = req.AccountNumber

	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	// logs.Info("Raw response received is ", res)
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.AccountQueryResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func PayECGBill(c *beego.Controller, req requests.ECGPaymentApiRequest) (resp responses.ECGBillPaymentApiResponse) {
	host, _ := beego.AppConfig.String("billpaymentBaseUrl")

	logs.Info("Paying ECG ", req.Amount, " for ", req.DestinationAccount)

	request := api.NewRequest(
		host,
		"/v1/bill-payment/pay-ecg-bill",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["PhoneNumber"] = req.PhoneNumber
	request.HeaderField["SourceSystem"] = req.SourceSystem

	request.InterfaceParams["DestinationAccount"] = req.DestinationAccount
	request.InterfaceParams["Amount"] = req.Amount
	request.InterfaceParams["PackageType"] = req.PackageType
	request.InterfaceParams["TransactionId"] = req.TransactionId
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	// logs.Info("Raw response received is ", res)
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.ECGBillPaymentApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func PayGhanaWaterBill(c *beego.Controller, req requests.GhanaWaterPaymentApiRequest) (resp responses.GhanaWaterBillPaymentApiResponse) {
	host, _ := beego.AppConfig.String("billpaymentBaseUrl")

	logs.Info("Paying water bill ", req.Amount, " for ", req.DestinationAccount)

	request := api.NewRequest(
		host,
		"/v1/bill-payment/pay-water-bill",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["PhoneNumber"] = req.PhoneNumber
	request.HeaderField["SourceSystem"] = req.SourceSystem

	request.InterfaceParams["DestinationAccount"] = req.DestinationAccount
	request.InterfaceParams["Amount"] = req.Amount
	request.InterfaceParams["Bundle"] = req.PackageType
	request.InterfaceParams["TransactionId"] = req.TransactionId
	request.InterfaceParams["SessionId"] = req.ExtraData.SessionId
	request.InterfaceParams["Email"] = req.ExtraData.Email
	request.InterfaceParams["PhoneNumber"] = req.ExtraData.Bundle
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	// logs.Info("Raw response received is ", res)
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.GhanaWaterBillPaymentApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func PayGoTvBill(c *beego.Controller, req requests.GoTvPaymentApiRequest) (resp responses.GoTvBillPaymentApiResponse) {
	host, _ := beego.AppConfig.String("billpaymentBaseUrl")

	logs.Info("Buying GoTV ", req.Amount, " for ", req.DestinationAccount)

	request := api.NewRequest(
		host,
		"/v1/bill-payment/pay-gotv-bill",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["PhoneNumber"] = req.PhoneNumber
	request.HeaderField["SourceSystem"] = req.SourceSystem

	request.InterfaceParams["DestinationAccount"] = req.DestinationAccount
	request.InterfaceParams["Amount"] = req.Amount
	request.InterfaceParams["PackageType"] = req.PackageType
	request.InterfaceParams["TransactionId"] = req.TransactionId
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	// logs.Info("Raw response received is ", res)
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.GoTvBillPaymentApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func PayStartimesBill(c *beego.Controller, req requests.StartimesPaymentApiRequest) (resp responses.StartimesBillPaymentApiResponse) {
	host, _ := beego.AppConfig.String("billpaymentBaseUrl")

	logs.Info("Buying Bundle ", req.Amount, " for ", req.DestinationAccount)

	request := api.NewRequest(
		host,
		"/v1/bill-payment/pay-startimes-bill",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.HeaderField["PhoneNumber"] = req.PhoneNumber
	request.HeaderField["SourceSystem"] = req.SourceSystem

	request.InterfaceParams["DestinationAccount"] = req.DestinationAccount
	request.InterfaceParams["Amount"] = req.Amount
	request.InterfaceParams["PackageType"] = req.PackageType
	request.InterfaceParams["TransactionId"] = req.TransactionId
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	// logs.Info("Raw response received is ", res)
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.StartimesBillPaymentApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func CheckTransactionStatus(c *beego.Controller, req requests.TransactionStatusApiRequest) (resp responses.TransactionStatusApiResponse) {
	host, _ := beego.AppConfig.String("billpaymentBaseUrl")

	logs.Info("Checking transaction status for ", req.TransactionID)
	request := api.NewRequest(
		host,
		"/v1/callback/transaction-status-check",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}

	request.InterfaceParams["TransactionID"] = req.TransactionID
	request.InterfaceParams["ThirdParthTransactionID"] = ""
	request.InterfaceParams["NetworkTransactionID"] = ""

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	// logs.Info("Raw response received is ", res)
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.TransactionStatusApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func GetCustomer(c *beego.Controller, req requests.MobileNumberRequest) (resp responses.CustomerResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/customers/phone-number/"+req.MobileNumber,
		api.GET)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}

	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	// logs.Info("Raw response received is ", res)
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CustomerResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func GetCustomerByUsername(c *beego.Controller, req requests.UsernameRequest) (resp responses.CustomerResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/customers/username/"+req.Username,
		api.GET)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}

	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	// logs.Info("Raw response received is ", res)
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CustomerResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}
