package controllers

import (
	"encoding/json"
	"errors"
	apifunctions "msys_payment_app_gateway/controllers/api_functions"
	"msys_payment_app_gateway/models"
	"msys_payment_app_gateway/structs/requests"
	"msys_payment_app_gateway/structs/responses"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"

	beego "github.com/beego/beego/v2/server/web"
)

// Api_requestsController operations for Api_requests
type Api_requestsController struct {
	beego.Controller
}

// URLMapping ...
func (c *Api_requestsController) URLMapping() {
	c.Mapping("GetBundles", c.GetBundles)
	c.Mapping("GetCorporatives", c.GetCorporatives)
	c.Mapping("BuyDataBundle", c.BuyDataBundle)
	c.Mapping("BuyAirtime", c.BuyAirtime)
	c.Mapping("AccountQuery", c.AccountQuery)
	c.Mapping("PayDSTV", c.PayDSTV)
	c.Mapping("PayGOTV", c.PayGOTV)
	c.Mapping("PayECG", c.PayECG)
	c.Mapping("GetCustomerAccounts", c.GetCustomerAccounts)
	c.Mapping("ValidateCustomer", c.ValidateCustomer)
	c.Mapping("NameInquiry", c.NameInquiry)
	c.Mapping("AccountBalance", c.AccountBalance)
	c.Mapping("ResetPin", c.ResetPin)
	c.Mapping("GetCustomerDetails", c.GetCustomerDetails)
}

// GetCorporatives ...
// @Title Get Corporatives
// @Description Get Corporatives Available
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	Network		header 	string true		"header for network"
// @Param	body		body 	requests.GetBundlesAPIRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /get-corporatives [post]
func (c *Api_requestsController) GetCorporatives() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	network := c.Ctx.Input.Header("Network")

	var req requests.GetBundlesAPIRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	destinationPhoneNumber := req.Destination

	logs.Info("GetBundles called with PhoneNumber: %s, SourceSystem: %s, Network: %s, DestinationPhoneNumber: %s", phoneNumber, sourceSystem, network, destinationPhoneNumber)

	reqBody := c.Ctx.Input.RequestBody
	reqHeaders := c.Ctx.Request.Header

	requestMap := map[string]interface{}{
		"headers": reqHeaders,
		"body":    string(reqBody),
	}

	reqText, err := json.Marshal(requestMap)
	if err != nil {
		logs.Error("Error marshalling request input: ", err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	var v models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		RequestType:  "Get Clients",
		PhoneNumber:  phoneNumber,
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)

		logs.Info("Formatted request for Corporatives: ")
		resp := apifunctions.GetCorporatives(&c.Controller)
		logs.Info("Response from Get corporatives API: ", resp)

		var response responses.CorporativeResponse = responses.CorporativeResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong",
			Result:        nil,
		}

		if resp.StatusCode != 200 {
			response = responses.CorporativeResponse{
				StatusCode:    false,
				StatusMessage: resp.StatusMessage,
				Result:        resp.Result,
			}
		} else {
			responseText, err := json.Marshal(response.Result)
			if err != nil {
				logs.Error("Error marshalling response result: ", err)
				responseText = []byte("[]")
			}
			v.RequestResponse = string(responseText)
			v.DateModified = time.Now()
			v.ResponseDate = time.Now()
			if err := models.UpdateApi_requestsById(&v); err != nil {
				logs.Error("Error updating API request with response: ", err)
			} else {
				logs.Info("API request updated with response successfully: ", v)
			}
			response = responses.CorporativeResponse{
				StatusCode:    true,
				StatusMessage: "Corporatives fetched successfully",
				Result:        resp.Result,
			}
		}

		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = response

	} else {
		var response responses.CorporativeResponse = responses.CorporativeResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong:: " + err.Error(),
			Result:        nil,
		}

		c.Data["json"] = response
	}
	c.ServeJSON()
}

// GetCustomerDetails ...
// @Title Get Customer Details
// @Description Get Customer Details
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /get-customer-details [post]
func (c *Api_requestsController) GetCustomerDetails() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	// sourceSystem := c.Ctx.Input.Header("SourceSystem")
	cust := c.Ctx.Input.GetData("customer")

	logs.Info("Customer details: %s", cust)
	customerData, ok := cust.(*responses.Customer)
	if !ok {
		logs.Error("Error asserting customer data")
		c.Data["json"] = "Invalid customer data"
		c.ServeJSON()
		return
	}

	reqBody := c.Ctx.Input.RequestBody
	reqHeaders := c.Ctx.Request.Header

	requestMap := map[string]interface{}{
		"headers": reqHeaders,
		"body":    string(reqBody),
	}

	reqText, err := json.Marshal(requestMap)
	if err != nil {
		logs.Error("Error marshalling request input: ", err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	var v models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		RequestType:  "Get Customer details",
		PhoneNumber:  phoneNumber,
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)

		var response responses.CustomerGatewayResponseDTO = responses.CustomerGatewayResponseDTO{
			StatusCode:    false,
			StatusMessage: "Customer fetch failed",
			Result:        nil,
		}

		var fields []string
		var sortby []string
		var order []string
		var query = make(map[string]string)
		var limit int64 = 10
		var offset int64

		customerNumberSearch := "CustomerNumber:" + customerData.CustomerNumber

		if v := customerNumberSearch; v != "" {
			for _, cond := range strings.Split(v, ",") {
				kv := strings.SplitN(cond, ":", 2)
				if len(kv) != 2 {
					c.Data["json"] = errors.New("Error: invalid query key/value pair")
					c.ServeJSON()
					return
				}
				k, v := kv[0], kv[1]
				query[k] = v
			}
		}

		logs.Debug("Query for customer corporatives is ", query)

		if customerCorps, err := models.GetAllCustomer_corporatives(query, fields, sortby, order, offset, limit); err == nil {

			logs.Debug("Returned customer corporatives data is ", customerCorps)
			var customerCorpsDTO []responses.CustomerCorporativesResponseDTO
			for _, v := range customerCorps {
				logs.Debug("Processing customer corporative: ", v)
				var corpDTO responses.CustomerCorporativesResponseDTO
				corpBytes, err := json.Marshal(v)
				if err != nil {
					logs.Error("Error marshalling customer corporative data: ", err)
					continue
				}
				if err := json.Unmarshal(corpBytes, &corpDTO); err != nil {
					logs.Error("Error unmarshalling customer corporative data: ", err)
					continue
				}
				customerCorpsDTO = append(customerCorpsDTO, corpDTO)
			}

			// Log customer corporatives data as readable JSON
			corpsJSON, err := json.MarshalIndent(customerCorpsDTO, "", "  ")
			if err != nil {
				logs.Error("Error marshalling customer corporatives to JSON: ", err)
			} else {
				logs.Debug("Formatted customer corporatives data is: %s", string(corpsJSON))
			}

			customerResp := responses.CustomerGateway{
				CustomerId:           customerData.CustomerId,
				FullName:             customerData.FullName,
				ImagePath:            customerData.ImagePath,
				Email:                customerData.Email,
				PhoneNumber:          customerData.PhoneNumber,
				Location:             customerData.Location,
				IdentificationType:   customerData.IdentificationType,
				IdentificationNumber: customerData.IdentificationNumber,
				DateCreated:          customerData.DateCreated,
				Status:               customerData.Active,
				CustomerCorporatives: &customerCorpsDTO,
			}

			logs.Info("Formatted request for customer: ")

			response = responses.CustomerGatewayResponseDTO{
				StatusCode:    true,
				StatusMessage: "Customer fetched successfully",
				Result:        &customerResp,
			}
		} else {
			logs.Error("Error fetching customer corporatives: ", err)
			response = responses.CustomerGatewayResponseDTO{
				StatusCode:    false,
				StatusMessage: "Something went wrong:: " + err.Error(),
				Result:        nil,
			}
			c.Data["json"] = response
		}

		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = response

	} else {
		var response responses.CustomerGatewayResponseDTO = responses.CustomerGatewayResponseDTO{
			StatusCode:    false,
			StatusMessage: "Something went wrong:: " + err.Error(),
			Result:        nil,
		}

		c.Data["json"] = response
	}
	c.ServeJSON()
}

// GetCustomerAccounts ...
// @Title Get Customer Accounts
// @Description Get customer accounts
// @Param	Authorization		header 	string true		"header for User"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	body		body 	requests.NumberExistsApiRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /get-customer-accounts [post]
func (c *Api_requestsController) GetCustomerAccounts() {
	// Extract headers
	// phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")

	var req requests.NumberExistsApiRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	phoneNumber := req.MobileNumber

	logs.Info("Get customer accounts called with PhoneNumber: %s, SourceSystem: %s", phoneNumber, sourceSystem)
	reqBody := c.Ctx.Input.RequestBody
	reqHeaders := c.Ctx.Request.Header

	requestMap := map[string]interface{}{
		"headers": reqHeaders,
		"body":    string(reqBody),
	}

	reqText, err := json.Marshal(requestMap)
	if err != nil {
		logs.Error("Error marshalling request input: ", err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	var v models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		PhoneNumber:  phoneNumber,
		RequestType:  "List Customer Accounts",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)

		var response responses.CustomerAccountsResponse = responses.CustomerAccountsResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong",
			Result:        nil,
		}

		var clientId int64
		clientId, err = strconv.ParseInt(req.ClientId, 10, 64)
		if err != nil {
			logs.Error("Error converting ClientId to int64: ", err)
			response = responses.CustomerAccountsResponse{
				StatusCode:    false,
				StatusMessage: "Invalid ClientId",
				Result:        nil,
			}
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = response
		}

		if client, err := models.GetClientsById(clientId); err != nil {

			clientCorpId := client.ClientCorpId

			listAccountsRequest := requests.NumberExistsApiRequest{
				MobileNumber: phoneNumber,
				ClientId:     clientCorpId,
			}

			logs.Info("Formatted request for customer accounts: ", listAccountsRequest)
			resp := apifunctions.ListCustomerAccounts(&c.Controller, listAccountsRequest)
			logs.Info("Response from customer accounts API: ", resp)

			if resp.Data.StatusCode != 200 {
				response = responses.CustomerAccountsResponse{
					StatusCode:    false,
					StatusMessage: resp.Data.StatusMessage,
					Result:        nil,
				}
			} else {
				responseText, err := json.Marshal(response.Result)
				if err != nil {
					logs.Error("Error marshalling response result: ", err)
					responseText = []byte("[]")
				}
				v.RequestResponse = string(responseText)
				v.DateModified = time.Now()
				v.ResponseDate = time.Now()
				if err := models.UpdateApi_requestsById(&v); err != nil {
					logs.Error("Error updating API request with response: ", err)
				} else {
					logs.Info("API request updated with response successfully: ", v)
				}
				response = responses.CustomerAccountsResponse{
					StatusCode:    true,
					StatusMessage: "Accounts fetched successfully",
					Result:        resp.Data.Result,
				}
			}
		}

		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = response

	} else {
		var response responses.CustomerAccountsResponse = responses.CustomerAccountsResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong:: " + err.Error(),
			Result:        nil,
		}

		c.Data["json"] = response
	}
	c.ServeJSON()
}

// NameInquiry ...
// @Title Name Inquiry
// @Description Name Inquiry with number
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	body		body 	requests.NumberExistsApiRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /name-inquiry [post]
func (c *Api_requestsController) NameInquiry() {
	// Extract headers
	// phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")

	var req requests.NumberExistsApiRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	// phoneNumber := req.MobileNumber

	logs.Info("Login called with PhoneNumber: %s, SourceSystem: %s", phoneNumber, sourceSystem)
	reqBody := c.Ctx.Input.RequestBody
	reqHeaders := c.Ctx.Request.Header

	requestMap := map[string]interface{}{
		"headers": reqHeaders,
		"body":    string(reqBody),
	}

	reqText, err := json.Marshal(requestMap)
	if err != nil {
		logs.Error("Error marshalling request input: ", err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	var v models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		PhoneNumber:  phoneNumber,
		RequestType:  "Name Inquiry",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)
		nameInquiryRequest := requests.NumberExistsApiRequest{
			MobileNumber: req.MobileNumber,
			ClientId:     req.ClientId,
		}

		logs.Info("Formatted request for name inquiry: ", nameInquiryRequest)
		resp := apifunctions.NameInquiry(&c.Controller, nameInquiryRequest)
		logs.Info("Response from name inquiry API: ", resp)

		var response responses.NameInquiryResponse = responses.NameInquiryResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong",
			Result:        "",
		}

		if resp.Data.StatusCode != 200 {
			response = responses.NameInquiryResponse{
				StatusCode:    false,
				StatusMessage: resp.Data.StatusMessage,
				Result:        "",
			}
		} else {
			responseText, err := json.Marshal(response.Result)
			if err != nil {
				logs.Error("Error marshalling response result: ", err)
				responseText = []byte("[]")
			}
			v.RequestResponse = string(responseText)
			v.DateModified = time.Now()
			v.ResponseDate = time.Now()
			if err := models.UpdateApi_requestsById(&v); err != nil {
				logs.Error("Error updating API request with response: ", err)
			} else {
				logs.Info("API request updated with response successfully: ", v)
			}
			response = responses.NameInquiryResponse{
				StatusCode:    true,
				StatusMessage: "Name inquiry successful",
				Result:        resp.Data.Result,
			}
		}

		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = response

	} else {
		var response responses.NameInquiryResponse = responses.NameInquiryResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong:: " + err.Error(),
			Result:        "",
		}

		c.Data["json"] = response
	}
	c.ServeJSON()
}

// AccountBalance ...
// @Title Account Balance
// @Description Account Balance
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	body		body 	requests.NumberExistsApiRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /account-balance [post]
func (c *Api_requestsController) AccountBalance() {
	// Extract headers
	// phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")

	var req requests.AccountBalanceApiRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	logs.Info("Account Balance called with AccountNumber: %s, SourceSystem: %s", req.AccountNumber, sourceSystem)
	reqBody := c.Ctx.Input.RequestBody
	reqHeaders := c.Ctx.Request.Header

	requestMap := map[string]interface{}{
		"headers": reqHeaders,
		"body":    string(reqBody),
	}

	reqText, err := json.Marshal(requestMap)
	if err != nil {
		logs.Error("Error marshalling request input: ", err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	var v models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		PhoneNumber:  phoneNumber,
		RequestType:  "Account Balance",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)
		accountBalanceRequest := requests.AccountBalanceApiRequest{
			AccountNumber: req.AccountNumber,
			ClientId:      req.ClientId,
		}

		logs.Info("Formatted request for account balance: ", accountBalanceRequest)
		resp := apifunctions.GetAccountBalance(&c.Controller, accountBalanceRequest)
		logs.Info("Response from account balance API: ", resp)

		var response responses.AccountBalanceResponse = responses.AccountBalanceResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong",
			Result:        nil,
		}

		if resp.Data.StatusCode != 200 {
			response = responses.AccountBalanceResponse{
				StatusCode:    false,
				StatusMessage: resp.Data.StatusMessage,
				Result:        nil,
			}
		} else {
			responseText, err := json.Marshal(response.Result)
			if err != nil {
				logs.Error("Error marshalling response result: ", err)
				responseText = []byte("[]")
			}
			v.RequestResponse = string(responseText)
			v.DateModified = time.Now()
			v.ResponseDate = time.Now()
			if err := models.UpdateApi_requestsById(&v); err != nil {
				logs.Error("Error updating API request with response: ", err)
			} else {
				logs.Info("API request updated with response successfully: ", v)
			}
			response = responses.AccountBalanceResponse{
				StatusCode:    true,
				StatusMessage: "Account balance fetched succeefully",
				Result:        resp.Data.Result,
			}
		}

		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = response

	} else {
		var response responses.AccountBalanceResponse = responses.AccountBalanceResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong:: " + err.Error(),
			Result:        nil,
		}

		c.Data["json"] = response
	}
	c.ServeJSON()
}

// ResetPin ...
// @Title Reset Pin
// @Description Reset Customer Pin
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	body		body 	requests.NumberExistsApiRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /reset-pin [post]
func (c *Api_requestsController) ResetPin() {
	// Extract headers
	// phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")

	var req requests.ResetPinRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	logs.Info("Reset Pin called with PhoneNumber: %s, SourceSystem: %s", phoneNumber, sourceSystem)
	reqBody := c.Ctx.Input.RequestBody
	reqHeaders := c.Ctx.Request.Header

	requestMap := map[string]interface{}{
		"headers": reqHeaders,
		"body":    string(reqBody),
	}

	reqText, err := json.Marshal(requestMap)
	if err != nil {
		logs.Error("Error marshalling request input: ", err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	var v models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		PhoneNumber:  phoneNumber,
		RequestType:  "Reset Pin",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)
		resetPinRequest := requests.ResetPinApiRequest{
			Number:      req.Number,
			OldPassword: req.OldPassword,
			NewPassword: req.NewPassword,
			ClientId:    req.ClientId,
		}

		logs.Info("Formatted request for pin reset: ", resetPinRequest)
		resp := apifunctions.ResetPin(&c.Controller, resetPinRequest)
		logs.Info("Response from pin reset API: ", resp)

		var response responses.ResetPinResponse = responses.ResetPinResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong",
			Result:        "",
		}

		if resp.Data.StatusCode != 200 {
			response = responses.ResetPinResponse{
				StatusCode:    false,
				StatusMessage: resp.Data.StatusMessage,
				Result:        "",
			}
		} else {
			responseText, err := json.Marshal(response.Result)
			if err != nil {
				logs.Error("Error marshalling response result: ", err)
				responseText = []byte("[]")
			}
			v.RequestResponse = string(responseText)
			v.DateModified = time.Now()
			v.ResponseDate = time.Now()
			if err := models.UpdateApi_requestsById(&v); err != nil {
				logs.Error("Error updating API request with response: ", err)
			} else {
				logs.Info("API request updated with response successfully: ", v)
			}
			response = responses.ResetPinResponse{
				StatusCode:    true,
				StatusMessage: "Login successful",
				Result:        resp.Data.Client,
			}
		}

		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = response

	} else {
		var response responses.ResetPinResponse = responses.ResetPinResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong:: " + err.Error(),
			Result:        "",
		}

		c.Data["json"] = response
	}
	c.ServeJSON()
}

// GetBundles ...
// @Title Get Bundles
// @Description Get Data Bundles Available
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	Network		header 	string true		"header for network"
// @Param	body		body 	requests.GetBundlesAPIRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /get-data-bundles [post]
func (c *Api_requestsController) GetBundles() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	network := c.Ctx.Input.Header("Network")

	var req requests.GetBundlesAPIRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	destinationPhoneNumber := req.Destination

	logs.Info("GetBundles called with PhoneNumber: %s, SourceSystem: %s, Network: %s, DestinationPhoneNumber: %s", phoneNumber, sourceSystem, network, destinationPhoneNumber)

	reqBody := c.Ctx.Input.RequestBody
	reqHeaders := c.Ctx.Request.Header

	requestMap := map[string]interface{}{
		"headers": reqHeaders,
		"body":    string(reqBody),
	}

	reqText, err := json.Marshal(requestMap)
	if err != nil {
		logs.Error("Error marshalling request input: ", err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	var v models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		PhoneNumber:  phoneNumber,
		RequestType:  "Get Bundles",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)
		getBundlesRequest := requests.DataBundlesListFormulatedRequest{
			NetworkId:          network,
			DestinationAccount: destinationPhoneNumber,
			PhoneNumber:        phoneNumber,
			SourceSystem:       sourceSystem,
		}

		logs.Info("Formatted request for GetBundles: ", getBundlesRequest)
		resp := apifunctions.GetBundles(&c.Controller, getBundlesRequest)
		logs.Info("Response from GetBundles API: ", resp)

		var response responses.DataBundlesListAPIResponse = responses.DataBundlesListAPIResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong",
			Result:        resp.Result,
		}

		if !resp.StatusCode {
			response = responses.DataBundlesListAPIResponse{
				StatusCode:    false,
				StatusMessage: resp.StatusMessage,
				Result:        resp.Result,
			}
		} else {
			responseText, err := json.Marshal(response.Result)
			if err != nil {
				logs.Error("Error marshalling response result: ", err)
				responseText = []byte("[]")
			}
			v.RequestResponse = string(responseText)
			v.DateModified = time.Now()
			v.ResponseDate = time.Now()
			if err := models.UpdateApi_requestsById(&v); err != nil {
				logs.Error("Error updating API request with response: ", err)
			} else {
				logs.Info("API request updated with response successfully: ", v)
			}
			response = responses.DataBundlesListAPIResponse{
				StatusCode:    true,
				StatusMessage: "Bundles retrieved successfully",
				Result:        resp.Result,
			}
		}

		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = response

	} else {
		var response responses.DataBundlesListAPIResponse = responses.DataBundlesListAPIResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong:: " + err.Error(),
			Result:        nil,
		}

		c.Data["json"] = response
	}
	c.ServeJSON()
}

// BuyBundle ...
// @Title Buy Data Bundle
// @Description Buy Data Bundle Available
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	Network		header 	string true		"header for network"
// @Param	body		body 	requests.BuyDataBundleAPIRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /buy-data-bundle [post]
func (c *Api_requestsController) BuyDataBundle() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	network := c.Ctx.Input.Header("Network")

	var req requests.BuyDataBundleAPIRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	destinationPhoneNumber := req.Destination

	logs.Info("GetBundles called with PhoneNumber: %s, SourceSystem: %s, Network: %s, DestinationPhoneNumber: %s", phoneNumber, sourceSystem, network, destinationPhoneNumber)

	reqBody := c.Ctx.Input.RequestBody
	reqHeaders := c.Ctx.Request.Header

	requestMap := map[string]interface{}{
		"headers": reqHeaders,
		"body":    string(reqBody),
	}

	reqText, err := json.Marshal(requestMap)
	if err != nil {
		logs.Error("Error marshalling request input: ", err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	var v models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		PhoneNumber:  phoneNumber,
		RequestType:  "Buy Data Bundle",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)
		buyBundleRequest := requests.BuyDataBundleFormulatedRequest{
			RequestId:    v.Id,
			Amount:       req.Amount,
			Network:      network,
			Destination:  destinationPhoneNumber,
			BundleId:     req.BundleId,
			SourceSystem: sourceSystem,
			PhoneNumber:  phoneNumber,
		}

		logs.Info("Formatted request for Buy Bundle: ", buyBundleRequest)
		resp := apifunctions.BuyDataBundle(&c.Controller, buyBundleRequest)
		logs.Info("Response from Buy Bundle API: ", resp)

		var response responses.BuyDataBundleAPIResponse = responses.BuyDataBundleAPIResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong",
			Result:        resp.Result,
		}

		if !resp.StatusCode {
			response = responses.BuyDataBundleAPIResponse{
				StatusCode:    false,
				StatusMessage: resp.StatusMessage,
				Result:        resp.Result,
			}
		} else {
			responseText, err := json.Marshal(response.Result)
			if err != nil {
				logs.Error("Error marshalling response result: ", err)
				responseText = []byte("[]")
			}
			v.RequestResponse = string(responseText)
			v.DateModified = time.Now()
			v.ResponseDate = time.Now()
			if err := models.UpdateApi_requestsById(&v); err != nil {
				logs.Error("Error updating API request with response: ", err)
			} else {
				logs.Info("API request updated with response successfully: ", v)
			}
			response = responses.BuyDataBundleAPIResponse{
				StatusCode:    true,
				StatusMessage: "Data bundle purchase is being processed",
				Result:        resp.Result,
			}
		}

		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = response

	} else {
		var response responses.BuyDataBundleAPIResponse = responses.BuyDataBundleAPIResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong:: " + err.Error(),
			Result:        nil,
		}

		c.Data["json"] = response
	}
	c.ServeJSON()
}

// BuyAirtime ...
// @Title Buy Airtime
// @Description Buy Airtime
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	Network		header 	string true		"header for network"
// @Param	body		body 	requests.BuyAirtimeAPIRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /buy-airtime [post]
func (c *Api_requestsController) BuyAirtime() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	network := c.Ctx.Input.Header("Network")

	var req requests.BuyAirtimeAPIRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	destinationPhoneNumber := req.Destination

	logs.Info("GetBundles called with PhoneNumber: %s, SourceSystem: %s, Network: %s, DestinationPhoneNumber: %s", phoneNumber, sourceSystem, network, destinationPhoneNumber)

	reqBody := c.Ctx.Input.RequestBody
	reqHeaders := c.Ctx.Request.Header

	requestMap := map[string]interface{}{
		"headers": reqHeaders,
		"body":    string(reqBody),
	}

	reqText, err := json.Marshal(requestMap)
	if err != nil {
		logs.Error("Error marshalling request input: ", err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	var v models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		PhoneNumber:  phoneNumber,
		RequestType:  "Buy Airtime",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)
		buyAirtimeRequest := requests.BuyAirtimeFormulatedRequest{
			RequestId:    v.Id,
			Amount:       req.Amount,
			Network:      network,
			Destination:  destinationPhoneNumber,
			SourceSystem: sourceSystem,
			PhoneNumber:  phoneNumber,
		}

		logs.Info("Formatted request for Buy Airtime: ", buyAirtimeRequest)
		resp := apifunctions.BuyAirtime(&c.Controller, buyAirtimeRequest)
		logs.Info("Response from Buy Airtime API: ", resp)

		var response responses.BuyAirtimeAPIResponse = responses.BuyAirtimeAPIResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong",
			Result:        resp.Result,
		}

		if !resp.StatusCode {
			response = responses.BuyAirtimeAPIResponse{
				StatusCode:    false,
				StatusMessage: resp.StatusMessage,
				Result:        resp.Result,
			}
		} else {
			responseText, err := json.Marshal(response.Result)
			if err != nil {
				logs.Error("Error marshalling response result: ", err)
				responseText = []byte("[]")
			}
			v.RequestResponse = string(responseText)
			v.DateModified = time.Now()
			v.ResponseDate = time.Now()
			if err := models.UpdateApi_requestsById(&v); err != nil {
				logs.Error("Error updating API request with response: ", err)
			} else {
				logs.Info("API request updated with response successfully: ", v)
			}
			response = responses.BuyAirtimeAPIResponse{
				StatusCode:    true,
				StatusMessage: "Airtime purchase is being processed",
				Result:        resp.Result,
			}
		}

		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = response

	} else {
		var response responses.BuyAirtimeAPIResponse = responses.BuyAirtimeAPIResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong:: " + err.Error(),
			Result:        nil,
		}

		c.Data["json"] = response
	}
	c.ServeJSON()
}

// AccountQuery ...
// @Title Account Query
// @Description Account Query
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	BillerCode		header 	string true		"header for network"
// @Param	body		body 	requests.DSTVAccountQueryApiRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /account-query [post]
func (c *Api_requestsController) AccountQuery() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	billerCode := c.Ctx.Input.Header("BillerCode")

	var req requests.BillPaymentAccountQueryRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	accountNumber := req.AccountNumber

	logs.Info("AccountQuery called with PhoneNumber: %s, SourceSystem: %s, Network: %s, AccountNumber: %s", phoneNumber, sourceSystem, accountNumber)
	reqBody := c.Ctx.Input.RequestBody
	reqHeaders := c.Ctx.Request.Header

	requestMap := map[string]interface{}{
		"headers": reqHeaders,
		"body":    string(reqBody),
	}

	reqText, err := json.Marshal(requestMap)
	if err != nil {
		logs.Error("Error marshalling request input: ", err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	var v models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		RequestType:  "Account Query",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)
		accountQueryRequest := requests.BillPaymentAccountQueryApiRequest{
			AccountNumber: accountNumber,
			SourceSystem:  sourceSystem,
			PhoneNumber:   phoneNumber,
			BillerCode:    billerCode,
		}

		logs.Info("Formatted request for account query ", accountQueryRequest)
		resp := apifunctions.AccountQuery(&c.Controller, accountQueryRequest)
		logs.Info("Response from Account query API: ", resp)

		var response responses.AccountQueryAPIResponse = responses.AccountQueryAPIResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong",
			Result:        resp.Result,
		}

		if !resp.StatusCode {
			response = responses.AccountQueryAPIResponse{
				StatusCode:    false,
				StatusMessage: resp.StatusMessage,
				Result:        resp.Result,
			}
		} else {
			responseText, err := json.Marshal(response.Result)
			if err != nil {
				logs.Error("Error marshalling response result: ", err)
				responseText = []byte("[]")
			}
			v.RequestResponse = string(responseText)
			v.DateModified = time.Now()
			v.ResponseDate = time.Now()
			if err := models.UpdateApi_requestsById(&v); err != nil {
				logs.Error("Error updating API request with response: ", err)
			} else {
				logs.Info("API request updated with response successfully: ", v)
			}
			response = responses.AccountQueryAPIResponse{
				StatusCode:    true,
				StatusMessage: "Accounts queried successfully",
				Result:        resp.Result,
			}
		}

		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = response

	} else {
		var response responses.AccountQueryAPIResponse = responses.AccountQueryAPIResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong:: " + err.Error(),
			Result:        nil,
		}

		c.Data["json"] = response
	}
	c.ServeJSON()
}

// PayDSTV ...
// @Title Pay DSTV
// @Description Pay DSTV
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	body		body 	requests.DSTVPaymentApiRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /pay-dstv [post]
func (c *Api_requestsController) PayDSTV() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	// network := c.Ctx.Input.Header("Network")

	var req requests.DSTVPaymentApiRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	destinationAccount := req.DestinationAccount

	logs.Info("PayDSTV called with PhoneNumber: %s, SourceSystem: %s, DestinationAccount: %s", phoneNumber, sourceSystem, destinationAccount)
	reqBody := c.Ctx.Input.RequestBody
	reqHeaders := c.Ctx.Request.Header

	requestMap := map[string]interface{}{
		"headers": reqHeaders,
		"body":    string(reqBody),
	}

	reqText, err := json.Marshal(requestMap)
	if err != nil {
		logs.Error("Error marshalling request input: ", err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	var v models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		PhoneNumber:  phoneNumber,
		RequestType:  "Buy DSTV",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)
		payDSTVRequest := requests.DSTVPaymentRequest{
			RequestId:          v.Id,
			Amount:             req.Amount,
			DestinationAccount: destinationAccount,
			PackageType:        req.PackageType,
			SourceSystem:       sourceSystem,
			PhoneNumber:        phoneNumber,
		}

		logs.Info("Formatted request for Buy Bundle: ", payDSTVRequest)
		resp := apifunctions.PayDSTVBill(&c.Controller, payDSTVRequest)
		logs.Info("Response from Buy Bundle API: ", resp)

		var response responses.DSTVBillPaymentApiResponse = responses.DSTVBillPaymentApiResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong",
			Result:        resp.Result,
		}

		if !resp.StatusCode {
			response = responses.DSTVBillPaymentApiResponse{
				StatusCode:    false,
				StatusMessage: resp.StatusMessage,
				Result:        resp.Result,
			}
		} else {
			responseText, err := json.Marshal(response.Result)
			if err != nil {
				logs.Error("Error marshalling response result: ", err)
				responseText = []byte("[]")
			}
			v.RequestResponse = string(responseText)
			v.DateModified = time.Now()
			v.ResponseDate = time.Now()
			if err := models.UpdateApi_requestsById(&v); err != nil {
				logs.Error("Error updating API request with response: ", err)
			} else {
				logs.Info("API request updated with response successfully: ", v)
			}
			response = responses.DSTVBillPaymentApiResponse{
				StatusCode:    true,
				StatusMessage: "DSTV payment successful",
				Result:        resp.Result,
			}
		}

		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = response

	} else {
		var response responses.DSTVBillPaymentApiResponse = responses.DSTVBillPaymentApiResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong:: " + err.Error(),
			Result:        nil,
		}

		c.Data["json"] = response
	}
	c.ServeJSON()
}

// PayGOTV ...
// @Title Pay GOTV
// @Description Pay GOTV
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	body		body 	requests.BuyDataBundleAPIRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /pay-gotv [post]
func (c *Api_requestsController) PayGOTV() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	// network := c.Ctx.Input.Header("Network")

	var req requests.GoTvPaymentApiRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	destinationAccount := req.DestinationAccount

	logs.Info("PayGOTV called with PhoneNumber: %s, SourceSystem: %s, DestinationAccount: %s", phoneNumber, sourceSystem, destinationAccount)
	reqBody := c.Ctx.Input.RequestBody
	reqHeaders := c.Ctx.Request.Header

	requestMap := map[string]interface{}{
		"headers": reqHeaders,
		"body":    string(reqBody),
	}

	reqText, err := json.Marshal(requestMap)
	if err != nil {
		logs.Error("Error marshalling request input: ", err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	var v models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		RequestType:  "Buy GOTV",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)
		payDSTVRequest := requests.GoTvPaymentApiRequest{
			RequestId:          v.Id,
			Amount:             req.Amount,
			DestinationAccount: destinationAccount,
			PackageType:        req.PackageType,
			SourceSystem:       sourceSystem,
			PhoneNumber:        phoneNumber,
		}

		logs.Info("Formatted request for Buy Bundle: ", payDSTVRequest)
		resp := apifunctions.PayGoTvBill(&c.Controller, payDSTVRequest)
		logs.Info("Response from GOTV API: ", resp)

		var response responses.GoTvBillPaymentResponse = responses.GoTvBillPaymentResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong",
			Result:        resp.Result,
		}

		if !resp.StatusCode {
			response = responses.GoTvBillPaymentResponse{
				StatusCode:    false,
				StatusMessage: resp.StatusMessage,
				Result:        resp.Result,
			}
		} else {
			responseText, err := json.Marshal(response.Result)
			if err != nil {
				logs.Error("Error marshalling response result: ", err)
				responseText = []byte("[]")
			}
			v.RequestResponse = string(responseText)
			v.DateModified = time.Now()
			v.ResponseDate = time.Now()
			if err := models.UpdateApi_requestsById(&v); err != nil {
				logs.Error("Error updating API request with response: ", err)
			} else {
				logs.Info("API request updated with response successfully: ", v)
			}
			response = responses.GoTvBillPaymentResponse{
				StatusCode:    true,
				StatusMessage: "Go TV payment successful",
				Result:        resp.Result,
			}
		}

		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = response

	} else {
		var response responses.DSTVBillPaymentApiResponse = responses.DSTVBillPaymentApiResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong:: " + err.Error(),
			Result:        nil,
		}

		c.Data["json"] = response
	}
	c.ServeJSON()
}

// PayECG ...
// @Title Pay ECG
// @Description Pay ECG
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	body		body 	requests.ECGPaymentRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /pay-ecg [post]
func (c *Api_requestsController) PayECG() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	// network := c.Ctx.Input.Header("Network")

	var req requests.ECGPaymentRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	destinationAccount := req.DestinationAccount

	logs.Info("Pay ECG called with PhoneNumber: %s, SourceSystem: %s, DestinationAccount: %s", phoneNumber, sourceSystem, destinationAccount)
	reqBody := c.Ctx.Input.RequestBody
	reqHeaders := c.Ctx.Request.Header

	requestMap := map[string]interface{}{
		"headers": reqHeaders,
		"body":    string(reqBody),
	}

	reqText, err := json.Marshal(requestMap)
	if err != nil {
		logs.Error("Error marshalling request input: ", err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	var v models.Api_requests = models.Api_requests{
		PhoneNumber:  phoneNumber,
		Request:      string(reqText),
		RequestType:  "Pay ECG",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)
		payDSTVRequest := requests.ECGPaymentApiRequest{
			RequestId:          v.Id,
			Amount:             req.Amount,
			DestinationAccount: destinationAccount,
			PackageType:        req.PackageType,
			SourceSystem:       sourceSystem,
			PhoneNumber:        phoneNumber,
		}

		logs.Info("Formatted request for Buy Bundle: ", payDSTVRequest)
		resp := apifunctions.PayECGBill(&c.Controller, payDSTVRequest)
		logs.Info("Response from Buy Bundle API: ", resp)

		var response responses.ECGBillPaymentApiResponse = responses.ECGBillPaymentApiResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong",
			Result:        resp.Result,
		}

		if !resp.StatusCode {
			response = responses.ECGBillPaymentApiResponse{
				StatusCode:    false,
				StatusMessage: resp.StatusMessage,
				Result:        resp.Result,
			}
		} else {
			responseText, err := json.Marshal(response.Result)
			if err != nil {
				logs.Error("Error marshalling response result: ", err)
				responseText = []byte("[]")
			}
			v.RequestResponse = string(responseText)
			v.DateModified = time.Now()
			v.ResponseDate = time.Now()
			if err := models.UpdateApi_requestsById(&v); err != nil {
				logs.Error("Error updating API request with response: ", err)
			} else {
				logs.Info("API request updated with response successfully: ", v)
			}
			response = responses.ECGBillPaymentApiResponse{
				StatusCode:    true,
				StatusMessage: "Payment is being processed",
				Result:        resp.Result,
			}
		}

		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = response

	} else {
		var response responses.ECGBillPaymentApiResponse = responses.ECGBillPaymentApiResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong:: " + err.Error(),
			Result:        nil,
		}

		c.Data["json"] = response
	}
	c.ServeJSON()
}

// PayWater ...
// @Title Pay Water bill
// @Description Pay Water Bill
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	body		body 	requests.GhanaWaterPaymentRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /pay-water [post]
func (c *Api_requestsController) PayWaterBill() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	// network := c.Ctx.Input.Header("Network")

	var req requests.GhanaWaterPaymentRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	destinationAccount := req.DestinationAccount

	logs.Info("Pay water called with PhoneNumber: %s, SourceSystem: %s, DestinationAccount: %s", phoneNumber, sourceSystem, destinationAccount)
	reqBody := c.Ctx.Input.RequestBody
	reqHeaders := c.Ctx.Request.Header

	requestMap := map[string]interface{}{
		"headers": reqHeaders,
		"body":    string(reqBody),
	}

	reqText, err := json.Marshal(requestMap)
	if err != nil {
		logs.Error("Error marshalling request input: ", err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	var v models.Api_requests = models.Api_requests{
		PhoneNumber:  phoneNumber,
		Request:      string(reqText),
		RequestType:  "Buy Water",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)
		payDSTVRequest := requests.GhanaWaterPaymentApiRequest{
			RequestId:          v.Id,
			Amount:             req.Amount,
			DestinationAccount: destinationAccount,
			PackageType:        req.PackageType,
			SourceSystem:       sourceSystem,
			PhoneNumber:        phoneNumber,
		}

		logs.Info("Formatted request for Buy Bundle: ", payDSTVRequest)
		resp := apifunctions.PayGhanaWaterBill(&c.Controller, payDSTVRequest)
		logs.Info("Response from Buy Bundle API: ", resp)

		var response responses.GhanaWaterBillPaymentApiResponse = responses.GhanaWaterBillPaymentApiResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong",
			Result:        resp.Result,
		}

		if !resp.StatusCode {
			response = responses.GhanaWaterBillPaymentApiResponse{
				StatusCode:    false,
				StatusMessage: resp.StatusMessage,
				Result:        resp.Result,
			}
		} else {
			responseText, err := json.Marshal(response.Result)
			if err != nil {
				logs.Error("Error marshalling response result: ", err)
				responseText = []byte("[]")
			}
			v.RequestResponse = string(responseText)
			v.DateModified = time.Now()
			v.ResponseDate = time.Now()
			if err := models.UpdateApi_requestsById(&v); err != nil {
				logs.Error("Error updating API request with response: ", err)
			} else {
				logs.Info("API request updated with response successfully: ", v)
			}
			response = responses.GhanaWaterBillPaymentApiResponse{
				StatusCode:    true,
				StatusMessage: "Payment is being processed",
				Result:        resp.Result,
			}
		}

		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = response

	} else {
		var response responses.GhanaWaterBillPaymentApiResponse = responses.GhanaWaterBillPaymentApiResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong:: " + err.Error(),
			Result:        nil,
		}

		c.Data["json"] = response
	}
	c.ServeJSON()
}

// PayStartimes ...
// @Title Pay Startimes bill
// @Description Pay Startimes Bill
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	body		body 	requests.StartimesPaymentRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /pay-startimes [post]
func (c *Api_requestsController) PayStartimesTvBill() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	// network := c.Ctx.Input.Header("Network")

	var req requests.StartimesPaymentRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	destinationAccount := req.DestinationAccount

	logs.Info("Pay startimes called with PhoneNumber: %s, SourceSystem: %s, DestinationAccount: %s", phoneNumber, sourceSystem, destinationAccount)
	reqBody := c.Ctx.Input.RequestBody
	reqHeaders := c.Ctx.Request.Header

	requestMap := map[string]interface{}{
		"headers": reqHeaders,
		"body":    string(reqBody),
	}

	reqText, err := json.Marshal(requestMap)
	if err != nil {
		logs.Error("Error marshalling request input: ", err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	var v models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		PhoneNumber:  phoneNumber,
		RequestType:  "Buy Startimes",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)
		payDSTVRequest := requests.StartimesPaymentApiRequest{
			RequestId:          v.Id,
			Amount:             req.Amount,
			DestinationAccount: destinationAccount,
			PackageType:        req.PackageType,
			SourceSystem:       sourceSystem,
			PhoneNumber:        phoneNumber,
		}

		logs.Info("Formatted request for Buy Bundle: ", payDSTVRequest)
		resp := apifunctions.PayStartimesBill(&c.Controller, payDSTVRequest)
		logs.Info("Response from Buy Bundle API: ", resp)

		var response responses.StartimesBillPaymentApiResponse = responses.StartimesBillPaymentApiResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong",
			Result:        resp.Result,
		}

		if !resp.StatusCode {
			response = responses.StartimesBillPaymentApiResponse{
				StatusCode:    false,
				StatusMessage: resp.StatusMessage,
				Result:        resp.Result,
			}
		} else {
			responseText, err := json.Marshal(response.Result)
			if err != nil {
				logs.Error("Error marshalling response result: ", err)
				responseText = []byte("[]")
			}
			v.RequestResponse = string(responseText)
			v.DateModified = time.Now()
			v.ResponseDate = time.Now()
			if err := models.UpdateApi_requestsById(&v); err != nil {
				logs.Error("Error updating API request with response: ", err)
			} else {
				logs.Info("API request updated with response successfully: ", v)
			}
			response = responses.StartimesBillPaymentApiResponse{
				StatusCode:    true,
				StatusMessage: "Payment is being processed",
				Result:        resp.Result,
			}
		}

		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = response

	} else {
		var response responses.StartimesBillPaymentApiResponse = responses.StartimesBillPaymentApiResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong:: " + err.Error(),
			Result:        nil,
		}

		c.Data["json"] = response
	}
	c.ServeJSON()
}

// ValidateCustomer ...
// @Title Validate Customer
// @Description Validate Customer
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /validate-customer [post]
func (c *Api_requestsController) ValidateCustomer() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	// network := c.Ctx.Input.Header("Network")

	logs.Info("ValidateCustomer called with PhoneNumber: %s, SourceSystem: %s", phoneNumber, sourceSystem)
	reqBody := c.Ctx.Input.RequestBody
	reqHeaders := c.Ctx.Request.Header

	requestMap := map[string]interface{}{
		"headers": reqHeaders,
		"body":    string(reqBody),
	}

	reqText, err := json.Marshal(requestMap)
	if err != nil {
		logs.Error("Error marshalling request input: ", err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	var v models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		PhoneNumber:  phoneNumber,
		RequestType:  "Validate customer",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		proceed := true
		logs.Info("API request logged successfully: ", v)
		mobileNumberReq := requests.MobileNumberRequest{
			MobileNumber: phoneNumber,
		}

		logs.Info("Formatted request for Validate Customer: ", mobileNumberReq)
		resp := apifunctions.GetCustomer(&c.Controller, mobileNumberReq)
		logs.Info("Response from Buy Bundle API: ", resp)

		responseText, err := json.Marshal(resp)
		if err != nil {
			logs.Error("Error marshalling response result: ", err)
			responseText = []byte("[]")
		}
		v.RequestResponse = string(responseText)
		v.DateModified = time.Now()
		v.ResponseDate = time.Now()
		if err := models.UpdateApi_requestsById(&v); err != nil {
			logs.Error("Error updating API request with response: ", err)
		} else {
			logs.Info("API request updated with response successfully: ", v)
		}

		var response responses.CustomerGatewayResponseDTO = responses.CustomerGatewayResponseDTO{
			StatusCode:    false,
			StatusMessage: "Something went wrong",
			Result:        nil,
		}

		if resp.StatusCode != 200 {
			proceed = false
			response = responses.CustomerGatewayResponseDTO{
				StatusCode:    false,
				StatusMessage: resp.StatusDesc,
				Result:        nil,
			}

			// Add customer
			addCustomerReq := requests.AddCustomer{
				Name:         resp.Customer.FullName,
				Email:        resp.Customer.Email,
				PhoneNumber:  resp.Customer.PhoneNumber,
				Location:     resp.Customer.Location,
				IdType:       "",
				IdNumber:     "",
				ImagePath:    "",
				AddedBy:      "1",
				CustomerType: "Temporary",
			}

			addCust := apifunctions.AddCustomer(&c.Controller, addCustomerReq)
			logs.Info("Response from Add Customer API: ", addCust)
			if addCust.StatusCode == 200 {
				logs.Error("Customer added successfully: ", addCust.StatusDesc)
			} else {
				logs.Info("Failed to add customer: ", addCust.StatusDesc)
			}
		}

		if proceed {

			customer := responses.CustomerGateway{
				CustomerId:           resp.Customer.CustomerId,
				FullName:             resp.Customer.FullName,
				Email:                resp.Customer.Email,
				PhoneNumber:          resp.Customer.PhoneNumber,
				Location:             resp.Customer.Location,
				IdentificationType:   resp.Customer.IdentificationType,
				IdentificationNumber: resp.Customer.IdentificationNumber,
				DateCreated:          resp.Customer.DateCreated,
				Status:               resp.Customer.Active,
			}

			response = responses.CustomerGatewayResponseDTO{
				StatusCode:    true,
				StatusMessage: "Customer validated successfully",
				Result:        &customer,
			}
		}

		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = response

	} else {
		var response responses.CustomerGatewayResponseDTO = responses.CustomerGatewayResponseDTO{
			StatusCode:    false,
			StatusMessage: "Something went wrong:: " + err.Error(),
			Result:        nil,
		}

		c.Data["json"] = response
	}
	c.ServeJSON()
}
