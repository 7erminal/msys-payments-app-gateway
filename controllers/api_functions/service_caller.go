package apifunctions

import (
	"bytes"
	"encoding/json"
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

	logs.Info("Raw response received is ", res)
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

	logs.Info("Raw response received is ", res)
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

	logs.Info("Raw response received is ", res)
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

func OpenAccount(c *beego.Controller, req requests.RegisterApiRequest) (resp responses.RegisterApiResponse) {
	host, _ := beego.AppConfig.String("clientBaseUrl")

	logs.Info("Registering user ", req.MobileNumber, " with name ", req.FirstName, " ", req.LastName)
	request := api.NewRequest(
		host,
		"/v2/api/register-customer",
		api.POST)
	// request.Params["username"] = username
	// request.Params = {"UserId": strconv.Itoa(int(userid))}

	request.InterfaceParams["FirstName"] = req.FirstName
	request.InterfaceParams["LastName"] = req.LastName
	request.InterfaceParams["Gender"] = req.Gender
	request.InterfaceParams["MobileNumber"] = req.MobileNumber

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

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.RegisterApiResponse
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

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.BuyAirtimeResponse
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
		"/v2/api/account-balance",
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

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.AccountBalanceApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func ListCustomerAccounts(c *beego.Controller, req requests.NumberExistsApiRequest) (resp responses.CustAccountsApiResponse) {
	host, _ := beego.AppConfig.String("clientBaseUrl")

	logs.Info("Listing customer accounts for number ", req.MobileNumber)
	request := api.NewRequest(
		host,
		"/v2/api/list-cust-accounts",
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

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.CustAccountsApiResponse
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

	logs.Info("Raw response received is ", res)
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

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.NameInquiryApiResponse
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

	logs.Info("Raw response received is ", res)
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

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.DataBundlesListResponse
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
	request.InterfaceParams["request_id"] = req.RequestId
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

	logs.Info("Raw response received is ", res)
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
	request.InterfaceParams["request_id"] = req.RequestId
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
	host, _ := beego.AppConfig.String("airtimeBaseUrl")

	logs.Info("Sending callback ", req.ResponseCode, " for ", req.Data.TransactionId)
	request := api.NewRequest(
		host,
		"/v1/callback/process",
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

func PayDSTVBill(c *beego.Controller, req requests.DSTVPaymentRequest) (resp responses.DSTVBillPaymentResponse) {
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
	request.InterfaceParams["RequestId"] = req.RequestId
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
	var data responses.DSTVBillPaymentResponse
	json.Unmarshal(read, &data)
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
	request.InterfaceParams["RequestId"] = req.RequestId
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
	request.InterfaceParams["PackageType"] = req.PackageType
	request.InterfaceParams["RequestId"] = req.RequestId
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
	request.InterfaceParams["RequestId"] = req.RequestId
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
	request.InterfaceParams["RequestId"] = req.RequestId
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
