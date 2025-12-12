package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	apifunctions "msys_payment_app_gateway/controllers/api_functions"
	"msys_payment_app_gateway/controllers/helpers"
	"msys_payment_app_gateway/controllers/services"
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
	c.Mapping("RegisterAccount", c.RegisterAccount)
	c.Mapping("GetBilTransactions", c.GetBilTransactions)
	c.Mapping("GetCustomerAccountStatement", c.GetCustomerAccountStatement)
	c.Mapping("Deposit", c.Deposit)
	c.Mapping("Withdrawal", c.Withdrawal)
	// c.Mapping("TransferFunds", c.TransferFunds)
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
				Gender:               customerData.Gender,
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

			if customerData.Active == 2 {
				logs.Info("Customer is pending activation, checking accounts and profile completion")
				// Fetch customer corporatives

				logs.Info("Before update, Customer was created on ", customerData.DateCreated)
				helpers.CheckProfileCompletion(&c.Controller, customerData)

				logs.Info("Is customer completed? ", customerData.Active)
				logs.Info("Date customer was created is ", customerData.DateCreated)

				if customerData.Active == 2 && customerData.DateCreated != (time.Time{}) && time.Since(customerData.DateCreated) > 20*time.Minute {
					logs.Info("Date created is ", customerData.DateCreated)
					logs.Info("Time since customer was created is more than 20 minutes::: ", time.Since(customerData.DateCreated).Minutes(), " minutes. Its is over ", 20*time.Minute, " minutes")
					// Do your logic here
					logs.Info("Customer was created over 20 minutes ago")

					status := 1

					updatedcustmer := requests.UpdateCustomer{
						Name:        customerData.FullName,
						Email:       customerData.Email,
						PhoneNumber: customerData.PhoneNumber,
						IdNumber:    customerData.IdentificationNumber,
						Location:    customerData.Location,
						Branch:      1, // Assuming default branch
						UserId:      1, // Assuming system user
						CustomerId:  customerData.CustomerId,
						Status:      status,
					}

					logs.Info("About to update customer status to active")

					updateCustomerResp := apifunctions.UpdateCustomer(&c.Controller, updatedcustmer)
					logs.Info("Update customer response: ", updateCustomerResp)
					logs.Info("After update, Customer was created on ", customerData.DateCreated)
					if updateCustomerResp.StatusCode != 200 {
						logs.Error("Error updating customer status: ", updateCustomerResp.StatusDesc)
					} else {
						logs.Info("Customer status updated successfully: ", updateCustomerResp.Customer)
					}
				}

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
	logs.Info("Final response to be sent: ", c.Data["json"])
	c.ServeJSON()
}

// GetCustomerAccounts ...
// @Title Get Customer Accounts
// @Description Get customer accounts
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	body		body 	requests.NumberExistsApiRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /get-customer-accounts [post]
func (c *Api_requestsController) GetCustomerAccounts() {
	// Extract headers
	// phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")

	cust := c.Ctx.Input.GetData("customer")

	logs.Info("Customer details: %s", cust)
	customerData, ok := cust.(*responses.Customer)
	if !ok {
		logs.Error("Error asserting customer data")
		c.Data["json"] = "Invalid customer data"
		c.ServeJSON()
		return
	}

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
		RequestType:  "Get Customer Accounts",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}

	var response responses.CustomerAccountsResponse = responses.CustomerAccountsResponse{
		StatusCode:    false,
		StatusMessage: "Something went wrong",
		Result:        nil,
	}

	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)

		if customerData != nil {
			logs.Info("Customer exists: ", customerData.CustomerNumber)
			var fields []string
			var sortby []string
			var order []string
			var query = make(map[string]string)
			var limit int64 = 10
			var offset int64

			logs.Info("Customer status is ", customerData.Active)
			switch customerData.Active {
			case 1, 2:
				logs.Info("Customer is active or pending activation, fetching accounts")
				customerNumberSearch := "CustomerNumber:" + customerData.CustomerNumber

				// query: k:v,k:v
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

				if customerCorporatives, err := models.GetAllCustomer_corporatives(query, fields, sortby, order, offset, limit); err == nil {
					for _, v := range customerCorporatives {
						logs.Debug("Returned customer corporative data is ", v)
						// Convert to DTO
						var corpDTO models.Customer_corporatives
						corpBytes, err := json.Marshal(v)
						if err != nil {
							logs.Error("Error marshalling customer corporative data: ", err)
							continue
						}
						if err := json.Unmarshal(corpBytes, &corpDTO); err != nil {
							logs.Error("Error unmarshalling customer corporative data: ", err)
						}
						logs.Info("Customer corporative DTO: ", corpDTO)

						if corpDTO.IsActive == 0 {
							logs.Info("Customer corporative is inactive, fetching accounts skipped")
							helpers.FetchCustomerAccounts(&c.Controller, customerData, corpDTO, phoneNumber)
						}
					}

					resp := apifunctions.GetCustomerAccounts(&c.Controller, strconv.FormatInt(customerData.CustomerId, 10))
					logs.Info("Response from customer accounts API: ", resp)

					if resp.StatusCode != "200" {
						response = responses.CustomerAccountsResponse{
							StatusCode:    false,
							StatusMessage: resp.StatusMessage,
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

						corps := apifunctions.GetCorporatives(&c.Controller)
						logs.Info("Response from Get corporatives API: ", resp)

						reference := ""
						if corps.StatusCode != 200 {
							logs.Error("Error fetching corporatives: ", corps.StatusMessage)
						} else {
							logs.Info("Corporatives fetched successfully: ", corps.Result)

							if corps.Result == nil || len(*corps.Result) == 0 {
								logs.Error("No corporatives found")
							}

							custAccounts := make([]responses.CustomerAccountResponse, 0)
							for _, acc := range resp.Result {
								logs.Info("Account reference: ", acc.Reference)
								// Map each account to the response object
								reference = acc.Reference

								// Find the matching corporate for the customer
								for _, corp := range *corps.Result {
									corpIdStr := strings.TrimSpace(strconv.FormatInt(corp.Id, 10))
									if corpIdStr == acc.Reference {
										reference = corp.ClientCode
										break
									}
								}

								logs.Info("Account reference mapped to corporate code: ", reference)

								sharesBalance := 0.0
								loanBalance := 0.0
								clearBalance := 0.0
								// Convert reference to int64 for client lookup
								clientRef := strings.TrimSpace(acc.Reference)
								clientID, parseErr := strconv.ParseInt(clientRef, 10, 64)
								if parseErr != nil {
									logs.Error("failed to convert acc.Reference to int64: ", parseErr)
									response = responses.CustomerAccountsResponse{
										StatusCode:    false,
										StatusMessage: "Invalid account reference: " + parseErr.Error(),
										Result:        nil,
									}
									continue
								}

								if client, err := models.GetClientsById(clientID); err != nil {
									logs.Error("Error getting client by ID: ", err)
									response = responses.CustomerAccountsResponse{
										StatusCode:    false,
										StatusMessage: "Error getting client by ID: " + err.Error(),
										Result:        nil,
									}
								} else {
									// Convert client
									checkAndUpdateAccountBalance := requests.UpdateAccountBalanceApiRequest{
										AccountNumber: acc.AccountNumber,
										AccountId:     acc.CustomerAccountId,
										Balance:       acc.Balance,
										ModifiedBy:    int(customerData.CustomerId),
										Reason:        "Fetch latest balance",
										ClientId:      client.ClientCode,
									}

									balanceResp := helpers.UpdateAccountBalance(&c.Controller, checkAndUpdateAccountBalance)
									logs.Info("Account balance update response: ", balanceResp)

									if !balanceResp.StatusCode {
										logs.Error("Error fetching and updating account balance: ", balanceResp.StatusMessage)
									}

									if balanceResp.Result.SharesBalance != nil {
										sharesBalance = *balanceResp.Result.SharesBalance
									}
									if balanceResp.Result.LoanBalance != nil {
										loanBalance = *balanceResp.Result.LoanBalance
									}
									if balanceResp.Result.ClearBalance != nil {
										clearBalance = *balanceResp.Result.ClearBalance
									}

									// keep using client (from GetClientsById) as needed below
									_ = client
								}

								custAccount := responses.CustomerAccountResponse{
									CustomerAccountId: acc.CustomerAccountId,
									AccountNumber:     acc.AccountNumber,
									AccountAlias:      acc.AccountAlias,
									AccountType:       acc.AccountType,
									Reference:         reference,
									ClientId:          acc.Reference,
									Balance:           acc.Balance,
									SharesBalance:     sharesBalance,
									LoanBalance:       loanBalance,
									CLearBalance:      clearBalance,
									FrozenAmount:      acc.FrozenAmount,
									BalanceBefore:     acc.BalanceBefore,
									DateCreated:       acc.DateCreated,
									Active:            acc.Active,
								}
								custAccounts = append(custAccounts, custAccount)
							}
							response = responses.CustomerAccountsResponse{
								StatusCode:    true,
								StatusMessage: "Accounts fetched successfully",
								Result:        &custAccounts,
							}
						}
					}
				}
			default:
				response = responses.CustomerAccountsResponse{
					StatusCode:    false,
					StatusMessage: "Customer is not active",
					Result:        nil,
				}
			}

			// go helpers.CheckAccountsStatus(&c.Controller, customerData)
		} else {
			response = responses.CustomerAccountsResponse{
				StatusCode:    false,
				StatusMessage: "Invalid customer data",
				Result:        nil,
			}
		}

		c.Ctx.Output.SetStatus(200)
	} else {
		response = responses.CustomerAccountsResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong:: " + err.Error(),
			Result:        nil,
		}
	}

	c.Data["json"] = response
	c.ServeJSON()
}

// RegisterAccount ...
// @Title Register Account
// @Description Register customer
// @Param	Authorization		header 	string true		"header for User"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	Network		header 	string true		"header for network"
// @Param	body		body 	requests.OpenAccountRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /register-account [post]
func (c *Api_requestsController) RegisterAccount() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	network := c.Ctx.Input.Header("Network")

	cust := c.Ctx.Input.GetData("customer")

	logs.Info("Customer details: %s", cust)
	customerData, ok := cust.(*responses.Customer)
	if !ok {
		logs.Error("Error asserting customer data")
		c.Data["json"] = "Invalid customer data"
		c.ServeJSON()
		return
	}

	var req requests.OpenAccountRequest
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		logs.Error("Error unmarshalling request body: ", err)
		c.Data["json"] = "Invalid request body"
		c.ServeJSON()
		return
	}

	logs.Info("Register account called with SourceSystem: %s", sourceSystem)
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

	var response responses.RegisterAccountResponse
	var v models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		PhoneNumber:  phoneNumber,
		RequestType:  "Register Account",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)

		logs.Info("Customer exists: ", customerData.CustomerNumber)

		if client, err := models.GetClientsById(req.ClientId); err != nil {
			logs.Error("Error getting client by ID: ", err)
			response = responses.RegisterAccountResponse{
				StatusCode:    false,
				StatusMessage: "Error getting client by ID: " + err.Error(),
				Result:        nil,
			}
		} else {
			// logs.Info("Mobile Number: ", req.MobileNumber)
			// logs.Info("First Name: ", req.FirstName)
			// logs.Info("Last Name: ", req.LastName)
			// registerAccountRequest := requests.OpenAccountApiRequest{
			// 	FirstName:    req.FirstName,
			// 	LastName:     req.LastName,
			// 	Gender:       req.Gender,
			// 	MobileNumber: req.MobileNumber,
			// 	ClientId:     client.ClientCorpId,
			// }

			// logs.Info("Formatted request for Register account: ", registerAccountRequest)
			// resp := apifunctions.OpenAccount(&c.Controller, registerAccountRequest)
			// logs.Info("Response from Register account API: ", resp)

			accountOpeningFee := 1.0

			req := requests.PaymentRequestApiRequestDTO{
				ClientId:        client.ClientCode,
				Amount:          accountOpeningFee,
				PaymentMethod:   "MOBILEMONEY",
				Service:         "ACCOUNT OPENING",
				SenderAccount:   phoneNumber,
				ReceiverAccount: customerData.PhoneNumber,
				Network:         network,
				ServiceNetwork:  client.ClientCode,
				ServicePackage:  strconv.FormatFloat(accountOpeningFee, 'f', -1, 64),
				MobileNumber:    phoneNumber,
			}
			//

			resp, err := helpers.RequestPaymentMain(&c.Controller, req)
			if err != nil {
				logs.Error("Error processing account opening fee payment: ", err)
				response = responses.RegisterAccountResponse{
					StatusCode:    false,
					StatusMessage: "Error processing account opening fee payment: " + err.Error(),
					Result:        nil,
				}
			} else {
				logs.Info("Account opening fee payment response: ", resp)
				if !resp.Success {
					response = responses.RegisterAccountResponse{
						StatusCode:    false,
						StatusMessage: "Account opening fee payment failed: " + resp.StatusMessage,
						Result:        nil,
					}
				} else {
					// Payment successful, proceed with account registration
					result_ := true
					response = responses.RegisterAccountResponse{
						StatusCode:    true,
						StatusMessage: "Account opening is being processed",
						Result:        &result_,
					}
				}
			}

		}

		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = response

	} else {
		var response responses.RegisterAccountResponse = responses.RegisterAccountResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong:: " + err.Error(),
			Result:        nil,
		}

		c.Data["json"] = response
	}
	c.ServeJSON()
}

// GetCustomerAccountHistory ...
// @Title Get Customer Account History
// @Description Get customer account history
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	body		body 	requests.AccountNumberRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /get-customer-account-history [post]
func (c *Api_requestsController) GetCustomerAccountHistory() {
	// Extract headers
	// phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")

	cust := c.Ctx.Input.GetData("customer")

	logs.Info("Customer details: %s", cust)
	// customerData, ok := cust.(*responses.Customer)
	// if !ok {
	// 	logs.Error("Error asserting customer data")
	// 	c.Data["json"] = "Invalid customer data"
	// 	c.ServeJSON()
	// 	return
	// }

	var req requests.AccountNumberRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	accountNumber := req.AccountNumber

	phoneNumber := c.Ctx.Input.Header("PhoneNumber")

	logs.Info("Get customer account history called with Account Number: %s, SourceSystem: %s", accountNumber, sourceSystem)
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
		RequestType:  "Get Customer Account History",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}

	var response responses.CustomersAccountHistoryResponseDTO = responses.CustomersAccountHistoryResponseDTO{
		Success:    false,
		StatusDesc: "Something went wrong",
		Result:     nil,
	}

	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)

		resp := apifunctions.GetCustomerAccountHistory(&c.Controller, accountNumber)
		logs.Info("Response from customer account history API: ", resp)

		if resp.StatusCode != "200" {
			response = responses.CustomersAccountHistoryResponseDTO{
				Success:    false,
				StatusDesc: resp.StatusMessage,
				Result:     nil,
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

			accountHistory := make([]*responses.CustomerAccountHistoryData, 0)

			for _, ah := range resp.Result {
				// Map each account history to the response object

				accHistory := &responses.CustomerAccountHistoryData{
					CustomerAccountHistoryId: ah.CustomerAccountHistoryId,
					CustomerAccount:          ah.CustomerAccount,
					DebitAmount:              ah.DebitAmount,
					CreditAmount:             ah.CreditAmount,
					TransactionDate:          ah.TransactionDate,
					CreatedBy:                ah.CreatedBy,
					ModifiedBy:               ah.ModifiedBy,
				}
				accountHistory = append(accountHistory, accHistory)
			}

			response = responses.CustomersAccountHistoryResponseDTO{
				Success:    true,
				StatusDesc: "Account history fetched successfully",
				Result:     accountHistory,
			}
		}

		c.Ctx.Output.SetStatus(200)
	} else {
		response = responses.CustomersAccountHistoryResponseDTO{
			Success:    false,
			StatusDesc: "Something went wrong:: " + err.Error(),
			Result:     nil,
		}
	}

	c.Data["json"] = response
	c.ServeJSON()
}

// GetBilTransactions ...
// @Title Get Biller Transaction History
// @Description Get customer biller transaction history
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /get-biller-transaction-history [post]
func (c *Api_requestsController) GetBilTransactions() {
	// Extract headers
	// phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")

	cust := c.Ctx.Input.GetData("customer")

	logs.Info("Customer details: %s", cust)
	// customerData, ok := cust.(*responses.Customer)
	// if !ok {
	// 	logs.Error("Error asserting customer data")
	// 	c.Data["json"] = "Invalid customer data"
	// 	c.ServeJSON()
	// 	return
	// }

	phoneNumber := c.Ctx.Input.Header("PhoneNumber")

	logs.Info("Get customer account biller transaction history called with Account Number: %s, SourceSystem: %s", phoneNumber, sourceSystem)
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
		RequestType:  "Get Customer Account History",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}

	var response responses.BilTransactionsResponse = responses.BilTransactionsResponse{
		Success:    false,
		StatusDesc: "Something went wrong",
		Result:     nil,
	}

	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)

		query := "BilTransactionId__TransactionBy__PhoneNumber:" + phoneNumber

		resp := apifunctions.GetCustomerBillPaymentHistory(&c.Controller, query)
		logs.Info("Response from customer bill payment history API: ", resp)

		if resp.StatusCode != "200" {
			response = responses.BilTransactionsResponse{
				Success:    false,
				StatusDesc: resp.StatusMessage,
				Result:     nil,
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

			billPaymentHist := make([]*responses.BilTransactionsData, 0)

			for _, ah := range resp.Result {
				// Map each account history to the response object

				billPayment := &responses.BilTransactionsData{
					TransactionId:           ah.TransactionId,
					TransactionBy:           ah.TransactionBy,
					TransactionDate:         ah.DateCreated,
					Amount:                  ah.Amount,
					TransactionRefNumber:    ah.TransactionRefNumber,
					Status:                  ah.Status,
					TransactingCurrency:     ah.TransactingCurrency,
					Source:                  ah.Source,
					Destination:             ah.Destination,
					Charge:                  ah.Charge,
					BillerName:              ah.BillerName,
					NetworkName:             ah.NetworkName,
					ExternalReferenceNumber: ah.ExternalReferenceNumber,
					Service:                 ah.Service,
				}
				billPaymentHist = append(billPaymentHist, billPayment)
			}

			response = responses.BilTransactionsResponse{
				Success:    true,
				StatusDesc: "Account history fetched successfully",
				Result:     billPaymentHist,
			}
		}

		c.Ctx.Output.SetStatus(200)
	} else {
		response = responses.BilTransactionsResponse{
			Success:    false,
			StatusDesc: "Something went wrong:: " + err.Error(),
			Result:     nil,
		}
	}

	c.Data["json"] = response
	c.ServeJSON()
}

// GetCustomerAccountHistory ...
// @Title Get Customer Account History
// @Description Get customer account history
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	body		body 	requests.AccountNumberRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /get-customer-account-statement [post]
func (c *Api_requestsController) GetCustomerAccountStatement() {
	// Extract headers
	// phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")

	cust := c.Ctx.Input.GetData("customer")

	logs.Info("Customer details: %s", cust)
	// customerData, ok := cust.(*responses.Customer)
	// if !ok {
	// 	logs.Error("Error asserting customer data")
	// 	c.Data["json"] = "Invalid customer data"
	// 	c.ServeJSON()
	// 	return
	// }

	var req requests.AccountNumberRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	accountNumber := req.AccountNumber

	phoneNumber := c.Ctx.Input.Header("PhoneNumber")

	logs.Info("Get customer account statement called with Account Number: %s, SourceSystem: %s", accountNumber, sourceSystem)
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
		RequestType:  "Get Customer Account Statement",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}

	var response responses.CustomersAccountStatementResponseDTO = responses.CustomersAccountStatementResponseDTO{
		Success:    false,
		StatusDesc: "Something went wrong",
		Result:     nil,
	}

	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)

		clientId, err := strconv.ParseInt(req.ClientId, 10, 64)

		if err != nil {
			logs.Error("Error parsing client ID: ", err)
			response = responses.CustomersAccountStatementResponseDTO{
				Success:    false,
				StatusDesc: "Error parsing client ID: " + err.Error(),
				Result:     nil,
			}
		}

		if client, err := models.GetClientsById(clientId); err != nil {
			logs.Error("Error getting client by ID: ", err)
			response = responses.CustomersAccountStatementResponseDTO{
				Success:    false,
				StatusDesc: "Error getting client by ID: " + err.Error(),
				Result:     nil,
			}

			c.Data["json"] = response
			c.ServeJSON()
		} else {

			resp := apifunctions.GetCustomerAccountStatement(&c.Controller, accountNumber, client.ClientCode)
			logs.Info("Response from customer account history API: ", resp)

			if resp.StatusCode != 200 {
				response = responses.CustomersAccountStatementResponseDTO{
					Success:    false,
					StatusDesc: resp.StatusMessage,
					Result:     nil,
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

				accountStatement := make([]*responses.CustomerAccountStatementData, 0)

				for _, ah := range resp.Result {
					// Map each account history to the response object

					var dateStr string
					// TransactionDate is a string; try RFC3339, then unix timestamp, else keep original string
					t := ah.TransactionDate
					dateStr = t
					// Try common parse layouts, then fall back to extracting the date portion
					if parsed, err := time.Parse(time.RFC3339, t); err == nil {
						dateStr = parsed.Format("Jan 2 2006")
					} else if parsed, err := time.Parse("Jan 2 2006 03:04:05PM", t); err == nil {
						dateStr = parsed.Format("Jan 2 2006")
					} else if parsed, err := time.Parse("Jan 2 2006 15:04:05", t); err == nil {
						dateStr = parsed.Format("Jan 2 2006")
					} else if i, err := strconv.ParseInt(strings.TrimSpace(t), 10, 64); err == nil {
						dateStr = time.Unix(i, 0).Format("Jan 2 2006")
					} else {
						parts := strings.Fields(t)
						if len(parts) >= 3 {
							// e.g. "Jun 17 2025 12:00:00AM" -> "Jun 17 2025"
							dateStr = strings.Join(parts[:3], " ")
						} else {
							dateStr = t
						}
					}

					accStatement := &responses.CustomerAccountStatementData{
						Account:                req.AccountNumber,
						TransactionDate:        dateStr,
						TransactionDescription: ah.TransactionDescription,
						Credit:                 ah.Credit,
						Debit:                  ah.Debit,
						// Balance:           0, // Balance not provided in the original response
						TransactionType: func() string {
							if ah.Debit > 0 {
								return "Debit"
							} else if ah.Credit > 0 {
								return "Credit"
							}
							return "N/A"
						}(),
					}
					accountStatement = append(accountStatement, accStatement)
				}

				response = responses.CustomersAccountStatementResponseDTO{
					Success:    true,
					StatusDesc: "Account history fetched successfully",
					Result:     accountStatement,
				}
			}
		}

		c.Ctx.Output.SetStatus(200)
	} else {
		response = responses.CustomersAccountStatementResponseDTO{
			Success:    false,
			StatusDesc: "Something went wrong:: " + err.Error(),
			Result:     nil,
		}
	}

	c.Data["json"] = response
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

		if accountResp := apifunctions.GetCustomerAccount(&c.Controller, req.AccountNumber); accountResp.StatusCode != "200" {
			logs.Error("Error fetching account details: ", accountResp.StatusMessage)
			var response responses.AccountBalanceResponse = responses.AccountBalanceResponse{
				StatusCode:    false,
				StatusMessage: "Error fetching account details: " + accountResp.StatusMessage,
				Result:        nil,
			}
			c.Ctx.Output.SetStatus(200)
			c.Data["json"] = response
			c.ServeJSON()
			return
		}
		logs.Info("Formatted request for account balance: ", accountBalanceRequest)
		resp := apifunctions.GetAccountBalance(&c.Controller, accountBalanceRequest)
		logs.Info("Response from account balance API: ", resp)

		var response responses.AccountBalanceResponse = responses.AccountBalanceResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong",
			Result:        nil,
		}

		if resp.StatusCode != 200 {
			response = responses.AccountBalanceResponse{
				StatusCode:    false,
				StatusMessage: resp.StatusDesc,
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

			accBal := responses.AccountBalanceDataResp{
				AccountNumber:    req.AccountNumber,
				AccountStatus:    resp.Result.AccountStatus,
				AvailableBalance: resp.Result.AvailableBalance,
				ClearBalance:     resp.Result.ClearBalance,
				LoanBalance:      resp.Result.LoanBalance,
				SharesBalance:    resp.Result.SharesBalance,
			}
			response = responses.AccountBalanceResponse{
				StatusCode:    true,
				StatusMessage: "Account balance fetched succeefully",
				Result:        &accBal,
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

// ChangePassword ...
// @Title Change Password
// @Description change password
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	body		body 	models.Auth_requests	true		"body for Auth_requests content"
// @Success 201 {object} models.Auth_requests
// @Failure 403 body is empty
// @router /change-password [post]
func (c *Auth_requestsController) ChangePassword() {
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")

	cust := c.Ctx.Input.GetData("customer")

	logs.Info("Customer details: %s", cust)
	customerData, ok := cust.(*responses.Customer)
	if !ok {
		logs.Error("Error asserting customer data")
		c.Data["json"] = "Invalid customer data"
		c.ServeJSON()
		return
	}

	logs.Info("Change Password called with PhoneNumber: %s, SourceSystem: %s", phoneNumber, sourceSystem)

	var req requests.ChangePasswordRequest
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		logs.Error("Error unmarshalling request body: ", err)
		c.Data["json"] = "Invalid request body"
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
		PhoneNumber:  phoneNumber,
		RequestType:  "Change Password",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)
		changePasswordRequest := requests.ChangePassword{
			OldPassword: req.OldPassword,
			NewPassword: req.NewPassword,
		}

		logs.Info("Formatted request for change password: ", changePasswordRequest)
		resp := apifunctions.ChangeCustomerPassword(&c.Controller, strconv.FormatInt(customerData.CustomerId, 10), changePasswordRequest)

		logs.Info("Response from change password API: ", resp)

		var response responses.StringResponseDTO = responses.StringResponseDTO{
			Success:    false,
			StatusDesc: "Something went wrong",
			Result:     nil,
		}
		if resp.StatusCode != 200 {
			response = responses.StringResponseDTO{
				Success:    false,
				StatusDesc: resp.StatusDesc,
				Result:     nil,
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
			response = responses.StringResponseDTO{
				Success:    true,
				StatusDesc: "Password changed successfully",
				Result:     &resp.Value,
			}
		}
	}

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
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /get-payment-methods [get]
func (c *Api_requestsController) GetPaymentMethods() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	network := c.Ctx.Input.Header("Network")

	logs.Info("GetBundles called with PhoneNumber: %s, SourceSystem: %s, Network: %s", phoneNumber, sourceSystem, network)

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
		RequestType:  "Get Payment Methods",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)

		resp := apifunctions.GetPaymentMethods(&c.Controller)
		logs.Info("Response from Get payment methods API: ", resp)

		isSuccess := false
		message := "Failed to retrieve payment methods"

		var response responses.PaymentMethodsResponseDTO = responses.PaymentMethodsResponseDTO{
			Success:    isSuccess,
			StatusDesc: message,
			Result:     nil,
		}

		if resp.StatusCode != 200 {
			response = responses.PaymentMethodsResponseDTO{
				Success:    isSuccess,
				StatusDesc: resp.StatusDesc,
				Result:     nil,
			}
		} else {
			isSuccess = true
			message = "Payment Methods retrieved successfully"
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

			paymentMethods := make([]*responses.Payment_methods, 0)

			for _, pm := range *resp.PaymentMethods {
				// Map each payment method to the response object

				payMethod := &responses.Payment_methods{
					PaymentMethodId: pm.PaymentMethodId,
					PaymentMethod:   pm.PaymentMethod,
					Networks:        pm.Networks,
				}
				paymentMethods = append(paymentMethods, payMethod)
			}
			response = responses.PaymentMethodsResponseDTO{
				Success:    isSuccess,
				StatusDesc: message,
				Result:     paymentMethods,
			}
		}

		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = response

	} else {
		var response responses.PaymentMethodsResponseDTO = responses.PaymentMethodsResponseDTO{
			Success:    false,
			StatusDesc: "Something went wrong:: " + err.Error(),
			Result:     nil,
		}

		c.Data["json"] = response
	}
	c.ServeJSON()
}

// Deposit ...
// @Title Deposit
// @Description Deposit to account
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	AccountNumber		header 	string true		"header for Customer's account number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	Network		header 	string true		"header for network"
// @Param	body		body 	requests.DepositAPIRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /deposit [post]
func (c *Api_requestsController) Deposit() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	accountNumber := c.Ctx.Input.Header("AccountNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	network := c.Ctx.Input.Header("Network")

	var req requests.DepositAPIRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	destinationPhoneNumber := req.Destination

	logs.Info("Deposit called with PhoneNumber: %s, SourceSystem: %s, Network: %s, DestinationPhoneNumber: %s", phoneNumber, sourceSystem, network, destinationPhoneNumber)

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

	isSuccess := false
	message := "Deposit failed"

	var response responses.PaymentRequestResponse = responses.PaymentRequestResponse{
		Success:       isSuccess,
		StatusMessage: message,
		Result:        nil,
	}
	var v models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		PhoneNumber:  phoneNumber,
		RequestType:  "Deposit",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)

		requestIdStr := fmt.Sprintf("%d", v.Id)
		amountString := strconv.FormatFloat(req.Amount, 'f', -1, 64)
		transactionLog := requests.LogTransactionRequest{
			RequestId:                requestIdStr,
			SourceAccountNumber:      accountNumber,
			DestinationAccountNumber: req.Destination,
			Amount:                   req.Amount,
			Charge:                   0.0,
			TransactionType:          "DEPOSIT",
			ServiceCode:              "DEPOSIT",
			TransactionReference:     "DEPOSIT",
			StatusCode:               "PENDING",
			ExtraDetails1:            amountString,
			ExtraDetails2:            strconv.FormatFloat(req.Amount, 'f', -1, 64),
			ExtraDetails3:            network,
			Reference:                amountString,
			ClientID:                 req.ClientId,
			PhoneNumber:              phoneNumber,
			TransactionPackage:       amountString,
			ExternalReferenceNumber:  "",
		}

		if txn, err := helpers.LogTransaction(&c.Controller, transactionLog); err != nil {
			logs.Error("Error logging transaction: ", err)
			isSuccess = false
			message = "Error logging transaction: " + err.Error()

		} else {

			req := requests.PaymentRequestApiRequestDTO{
				ClientId:        req.ClientId,
				Amount:          req.Amount,
				PaymentMethod:   "MOBILEMONEY",
				Service:         "DEPOSIT",
				SenderAccount:   accountNumber,
				ReceiverAccount: destinationPhoneNumber,
				Network:         network,
				ServiceNetwork:  req.ClientId,
				ServicePackage:  strconv.FormatFloat(req.Amount, 'f', -1, 64),
				MobileNumber:    phoneNumber,
				TransactionId:   txn.Result.TransactionRefNumber,
			}
			//

			resp, err := helpers.RequestPaymentMain(&c.Controller, req)

			logs.Info("Response from Deposit API: ", resp)

			if err != nil {
				message = err.Error()
			} else {
				if resp.Success {
					accountCheckResp := helpers.LogAccountActivity(&c.Controller, accountNumber, "Deposit", req.Amount, req.ClientId, "credit")

					isSuccess = true
					message = "Deposit successful"
					responseText, err := json.Marshal(resp)

					if !accountCheckResp.StatusCode {
						logs.Error("Error logging account activity for deposit: ", accountCheckResp.StatusMessage)
						message = "Deposit failed: " + accountCheckResp.StatusMessage
					}

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
				} else {
					message = "Deposit failed: " + resp.StatusMessage
				}
			}
		}

		c.Ctx.Output.SetStatus(200)

	} else {
		logs.Error("Error logging API request: ", err)
	}

	response = responses.PaymentRequestResponse{
		Success:       isSuccess,
		StatusMessage: message,
		Result:        nil,
	}

	c.Data["json"] = response

	c.ServeJSON()
}

// Withdrawal ...
// @Title Withdrawal
// @Description Withdrawal from account
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	AccountNumber		header 	string true		"header for Customer's account number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	Network		header 	string true		"header for network"
// @Param	body		body 	requests.WithdrawalAPIRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /withdrawal [post]
func (c *Api_requestsController) Withdrawal() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	accountNumber := c.Ctx.Input.Header("AccountNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	network := c.Ctx.Input.Header("Network")

	var req requests.WithdrawalAPIRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	destinationPhoneNumber := req.Destination

	logs.Info("Withdrawal called with PhoneNumber: %s, SourceSystem: %s, Network: %s, DestinationPhoneNumber: %s", phoneNumber, sourceSystem, network, destinationPhoneNumber)

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

	isSuccess := false
	message := "Withdrawal failed"
	var v models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		PhoneNumber:  phoneNumber,
		RequestType:  "Withdrawal",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)

		requestIdStr := fmt.Sprintf("%d", v.Id)
		amountString := strconv.FormatFloat(req.Amount, 'f', -1, 64)
		transactionLog := requests.LogTransactionRequest{
			RequestId:                requestIdStr,
			SourceAccountNumber:      accountNumber,
			DestinationAccountNumber: req.Destination,
			Amount:                   req.Amount,
			Charge:                   0.0,
			TransactionType:          "WITHDRAWAL",
			ServiceCode:              "WITHDRAWAL",
			TransactionReference:     "WITHDRAWAL",
			StatusCode:               "PENDING",
			ExtraDetails1:            amountString,
			ExtraDetails2:            strconv.FormatFloat(req.Amount, 'f', -1, 64),
			ExtraDetails3:            network,
			Reference:                amountString,
			ClientID:                 req.ClientId,
			PhoneNumber:              phoneNumber,
			TransactionPackage:       amountString,
			ExternalReferenceNumber:  "",
		}

		if txn, err := helpers.LogTransaction(&c.Controller, transactionLog); err != nil {
			logs.Error("Error logging transaction: ", err)
			isSuccess = false
			message = "Error logging transaction: " + err.Error()

		} else {

			accountCheckResp := helpers.LogAccountActivity(&c.Controller, accountNumber, "Withdrawal", req.Amount, req.ClientId, "debit")

			if !accountCheckResp.StatusCode {
				message = "Withdrawal failed: " + accountCheckResp.StatusMessage
				c.Ctx.Output.SetStatus(200)

				var response responses.PaymentRequestResponse = responses.PaymentRequestResponse{
					Success:       isSuccess,
					StatusMessage: message,
					Result:        nil,
				}

				c.Ctx.Output.SetStatus(200)
				c.Data["json"] = response
			} else {
				req := requests.PaymentRequestApiRequestDTO{
					ClientId:        req.ClientId,
					Amount:          req.Amount,
					PaymentMethod:   "MOBILEMONEY",
					Service:         "WITHDRAWAL",
					SenderAccount:   accountNumber,
					ReceiverAccount: destinationPhoneNumber,
					Network:         network,
					ServiceNetwork:  req.ClientId,
					ServicePackage:  strconv.FormatFloat(req.Amount, 'f', -1, 64),
					MobileNumber:    phoneNumber,
					TransactionId:   txn.Result.TransactionRefNumber,
				}
				//

				resp, err := helpers.MakePaymentMain(&c.Controller, req)

				logs.Info("Response from Withdrawal API: ", resp)

				if err != nil {
					message = err.Error()
				} else {
					if resp.Success {
						isSuccess = true
						message = "Withdrawal successful"
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
					} else {
						message = "Withdrawal failed: " + resp.StatusMessage
					}
				}

				c.Ctx.Output.SetStatus(200)
				c.Data["json"] = responses.PaymentRequestResponse{
					Success:       isSuccess,
					StatusMessage: message,
					Result:        resp,
				}
			}
		}

	} else {
		logs.Error("Error logging API request: ", err)

		var response responses.PaymentRequestResponse = responses.PaymentRequestResponse{
			Success:       isSuccess,
			StatusMessage: message,
			Result:        nil,
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
// @Param	AccountNumber		header 	string true		"header for Customer's account number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	Network		header 	string true		"header for network"
// @Param	body		body 	requests.BuyDataBundleAPIRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /buy-data-bundle [post]
func (c *Api_requestsController) BuyDataBundle() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	accountNumber := c.Ctx.Input.Header("AccountNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	network := c.Ctx.Input.Header("Network")

	var req requests.BuyDataBundleAPIRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	destinationPhoneNumber := req.Destination

	logs.Info("GetBundles called with PhoneNumber: %s, SourceSystem: %s, Network: %s, DestinationPhoneNumber: %s", phoneNumber, sourceSystem, network, destinationPhoneNumber)

	logs.Info("Amount to be debited: %f", req.Amount)
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

	isSuccess := false
	message := "Data bundle purchase failed"
	result := responses.BuyDataBundleResponseResult{}

	var response responses.BuyDataBundleResponse = responses.BuyDataBundleResponse{
		StatusCode:    isSuccess,
		StatusMessage: message,
		Result:        nil,
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

		requestIdStr := fmt.Sprintf("%d", v.Id)
		transactionLog := requests.LogTransactionRequest{
			RequestId:                requestIdStr,
			SourceAccountNumber:      accountNumber,
			DestinationAccountNumber: req.Destination,
			Amount:                   req.Amount,
			Charge:                   0.0,
			TransactionType:          "DATA_BUNDLE",
			ServiceCode:              "BILL_PAYMENT",
			TransactionReference:     "DATA_BUNDLE",
			StatusCode:               "PENDING",
			ExtraDetails1:            req.BundleId,
			ExtraDetails2:            strconv.FormatFloat(req.Amount, 'f', -1, 64),
			ExtraDetails3:            req.Network,
			Reference:                req.BundleId,
			ClientID:                 req.ClientId,
			PhoneNumber:              phoneNumber,
			TransactionPackage:       req.BundleId,
			ExternalReferenceNumber:  "",
		}

		if txn, err := helpers.LogTransaction(&c.Controller, transactionLog); err != nil {
			logs.Error("Error logging transaction: ", err)
			isSuccess = false
			message = "Error logging transaction: " + err.Error()

		} else {
			if accountNumber != "" {
				accountResp := apifunctions.GetCustomerAccount(&c.Controller, accountNumber)
				proceed := false
				if accountResp.StatusCode == "200" {
					accountCheckResp := helpers.LogAccountActivity(&c.Controller, accountNumber, "Data Bundle Purchase", req.Amount, req.ClientId, "debit")

					if !accountCheckResp.StatusCode {
						isSuccess = false
						message = accountCheckResp.StatusMessage
						// result = nil
						proceed = false
					} else {
						logs.Info("Account activity logged successfully for account number: ", accountNumber)

						// Log payment request
						makePaymentRequest := requests.PaymentRequestApiRequestDTO{
							ClientId:        req.ClientId,
							Amount:          req.Amount,
							PaymentMethod:   "ACCOUNT",
							Service:         "DATA_BUNDLE",
							SenderAccount:   accountNumber,
							ReceiverAccount: destinationPhoneNumber,
							Network:         network,
							ServiceNetwork:  req.Network,
							ServicePackage:  req.BundleId,
							MobileNumber:    phoneNumber,
							TransactionId:   txn.Result.TransactionRefNumber,
						}

						helpers.MakePaymentMain(&c.Controller, makePaymentRequest)

						proceed = true

						isSuccess = true
						message = accountCheckResp.StatusMessage
						// result = nil
					}
				} else {
					logs.Error("Error fetching account details for account number: ", accountNumber)
					logs.Info("Register Customer")

					req := requests.PaymentRequestApiRequestDTO{
						ClientId:        req.ClientId,
						Amount:          req.Amount,
						PaymentMethod:   "MOBILEMONEY",
						Service:         "DATA_BUNDLE",
						SenderAccount:   accountNumber,
						ReceiverAccount: destinationPhoneNumber,
						Network:         network,
						ServiceNetwork:  req.Network,
						ServicePackage:  req.BundleId,
						MobileNumber:    accountNumber,
						TransactionId:   txn.Result.TransactionRefNumber,
					}
					//

					resp, err := helpers.RequestPaymentMain(&c.Controller, req)
					if err != nil {
						logs.Error("Error requesting payment: ", err)
						message = "Error requesting payment: " + err.Error()
						isSuccess = false
					} else {
						logs.Info("Payment requested successfully: ", resp)
						if !resp.Success {
							isSuccess = false
							message = resp.StatusMessage
						} else {
							isSuccess = true
							message = "Data bundle purchase is being processed"
						}
					}
				}

				// txnString := fmt.Sprintf("%d", txn.Result.TransactionId)
				buyBundleRequest := requests.BuyDataBundleFormulatedRequest{
					TransactionId: txn.Result.TransactionRefNumber,
					Amount:        req.Amount,
					Network:       network,
					Destination:   destinationPhoneNumber,
					BundleId:      req.BundleId,
					SourceSystem:  sourceSystem,
					PhoneNumber:   phoneNumber,
				}

				if proceed {
					logs.Info("Formatted request for Buy Bundle: ", buyBundleRequest)
					resp := services.BuyDataBundle(&c.Controller, buyBundleRequest)
					logs.Info("Response from Buy Bundle API: ", resp)

					var response responses.BuyDataBundleResponse = responses.BuyDataBundleResponse{
						StatusCode:    false,
						StatusMessage: "Something went wrong",
						Result:        resp.Result,
					}

					if !resp.StatusCode {
						response = responses.BuyDataBundleResponse{
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

						// if accountNumber != "" {
						// 	helpers.LogAccountActivity(&c.Controller, accountNumber, "Data Bundle Purchase", strconv.FormatFloat(req.Amount, 'f', -1, 64), req.ClientId, "debit")
						// }
						response = responses.BuyDataBundleResponse{
							StatusCode:    true,
							StatusMessage: "Data bundle purchase is being processed",
							Result:        resp.Result,
						}
					}
				}
			} else {

				logs.Error("Account number is required for data bundle purchase")
				isSuccess = false
				message = "Account number is required for data bundle purchase"
			}
		}

		c.Ctx.Output.SetStatus(200)

	} else {
		var response responses.BuyDataBundleAPIResponse = responses.BuyDataBundleAPIResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong:: " + err.Error(),
			Result:        nil,
		}

		c.Data["json"] = response
	}

	// Log c.Data["json"] as JSON for debugging
	if respData, err := json.MarshalIndent(c.Data["json"], "", "  "); err != nil {
		logs.Error("Error marshalling c.Data[\"json\"] for logging: ", err)
	} else {
		logs.Info("Response JSON: %s", string(respData))
	}

	response = responses.BuyDataBundleResponse{
		StatusCode:    isSuccess,
		StatusMessage: message,
		Result:        &result,
	}
	c.Data["json"] = response

	c.ServeJSON()
}

// BuyAirtime ...
// @Title Buy Airtime
// @Description Buy Airtime
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	AccountNumber		header 	string true		"header for Customer's account number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	Network		header 	string true		"header for network"
// @Param	body		body 	requests.BuyAirtimeAPIRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /buy-airtime [post]
func (c *Api_requestsController) BuyAirtime() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	accountNumber := c.Ctx.Input.Header("AccountNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	network := c.Ctx.Input.Header("Network")

	var req requests.BuyAirtimeAPIRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	destinationPhoneNumber := req.Destination

	logs.Info("GetAirtime called with PhoneNumber: %s, SourceSystem: %s, Network: %s, DestinationPhoneNumber: %s", phoneNumber, sourceSystem, network, destinationPhoneNumber)

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

	isSuccess := false
	message := "Data bundle purchase failed"
	result := responses.AirtimeResponseResult{}

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

		requestIdStr := fmt.Sprintf("%d", v.Id)
		amountString := strconv.FormatFloat(req.Amount, 'f', -1, 64)
		transactionLog := requests.LogTransactionRequest{
			RequestId:                requestIdStr,
			SourceAccountNumber:      accountNumber,
			DestinationAccountNumber: req.Destination,
			Amount:                   req.Amount,
			Charge:                   0.0,
			TransactionType:          "AIRTIME",
			ServiceCode:              "BILL_PAYMENT",
			TransactionReference:     "AIRTIME",
			StatusCode:               "PENDING",
			ExtraDetails1:            amountString,
			ExtraDetails2:            strconv.FormatFloat(req.Amount, 'f', -1, 64),
			ExtraDetails3:            req.Network,
			Reference:                amountString,
			ClientID:                 req.ClientId,
			PhoneNumber:              phoneNumber,
			TransactionPackage:       amountString,
			ExternalReferenceNumber:  "",
		}

		if txn, err := helpers.LogTransaction(&c.Controller, transactionLog); err != nil {
			logs.Error("Error logging transaction: ", err)
			isSuccess = false
			message = "Error logging transaction: " + err.Error()

		} else {
			// txnString := fmt.Sprintf("%d", txn.Result.TransactionId)
			logs.Info("Transaction ID returned is ", txn.Result.TransactionId)
			buyAirtimeRequest := requests.BuyAirtimeFormulatedRequest{
				TransactionId: txn.Result.TransactionRefNumber,
				Amount:        req.Amount,
				Network:       req.Network,
				Destination:   destinationPhoneNumber,
				SourceSystem:  sourceSystem,
				PhoneNumber:   phoneNumber,
			}

			var response responses.BuyAirtimeAPIResponse = responses.BuyAirtimeAPIResponse{
				StatusCode:    false,
				StatusMessage: "Something went wrong",
				Result:        nil,
			}

			if accountNumber != "" {
				accountResp := apifunctions.GetCustomerAccount(&c.Controller, accountNumber)

				proceed := false
				if accountResp.StatusCode == "200" {
					accountCheckResp := helpers.LogAccountActivity(&c.Controller, accountNumber, "Airtime Purchase", req.Amount, req.ClientId, "debit")

					if !accountCheckResp.StatusCode {
						logs.Debug("Error: ", accountCheckResp.StatusMessage)
						response = responses.BuyAirtimeAPIResponse{
							StatusCode:    false,
							StatusMessage: accountCheckResp.StatusMessage,
							Result:        nil,
						}
					} else {
						logs.Info("Account activity logged successfully for account number: ", accountNumber)
						// Log payment request
						makePaymentRequest := requests.PaymentRequestApiRequestDTO{
							ClientId:        req.ClientId,
							Amount:          req.Amount,
							PaymentMethod:   "ACCOUNT",
							Service:         "AIRTIME",
							SenderAccount:   accountNumber,
							ReceiverAccount: destinationPhoneNumber,
							Network:         network,
							ServiceNetwork:  req.Network,
							MobileNumber:    phoneNumber,
							TransactionId:   txn.Result.TransactionRefNumber,
						}

						helpers.MakePaymentMain(&c.Controller, makePaymentRequest)

						proceed = true
					}
				} else {
					// Get customer by number before registering
					logs.Error("Error fetching account details for account number: ", accountNumber)
					logs.Info("Register Customer")

					req := requests.PaymentRequestApiRequestDTO{
						ClientId:        req.ClientId,
						Amount:          req.Amount,
						PaymentMethod:   "MOBILEMONEY",
						Service:         "AIRTIME",
						SenderAccount:   accountNumber,
						ReceiverAccount: destinationPhoneNumber,
						Network:         network,
						ServiceNetwork:  req.Network,
						ServicePackage:  strconv.FormatFloat(req.Amount, 'f', -1, 64),
						MobileNumber:    accountNumber,
						TransactionId:   txn.Result.TransactionRefNumber,
					}
					//

					resp, err := helpers.RequestPaymentMain(&c.Controller, req)
					if err != nil {
						logs.Error("Error requesting payment: ", err)
						isSuccess = false
						message = "Error requesting payment: " + err.Error()
					} else {
						logs.Info("Payment requested successfully: ", resp)
						if !resp.Success {
							isSuccess = false
							message = resp.StatusMessage
						} else {
							isSuccess = true
							message = "Airtime purchase is being processed"
						}
					}
				}

				if proceed {
					logs.Info("Formatted request for Buy Airtime: ", buyAirtimeRequest)
					resp := services.BuyAirtime(&c.Controller, buyAirtimeRequest)
					logs.Info("Response from Buy Airtime API: ", resp)

					if !resp.StatusCode {
						isSuccess = false
						message = resp.StatusMessage
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

						isSuccess = true
						message = "Airtime purchase is being processed"
						result = *resp.Result
					}
				}
			} else {
				logs.Error("Account number is required to process airtime purchase")
				isSuccess = false
				message = "Account number is required to process this request"
			}
		}

		c.Ctx.Output.SetStatus(200)

	} else {
		isSuccess = false
		message = "Something went wrong:: " + err.Error()
	}

	var response responses.BuyAirtimeAPIResponse = responses.BuyAirtimeAPIResponse{
		StatusCode:    isSuccess,
		StatusMessage: message,
		Result:        &result,
	}
	c.Data["json"] = response
	c.ServeJSON()
}

// AccountQuery ...
// @Title Account Query
// @Description Account Query
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	AccountNumber		header 	string true		"header for Customer's account number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	BillerCode		header 	string true		"header for network"
// @Param	body		body 	requests.DSTVAccountQueryApiRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /account-query [post]
func (c *Api_requestsController) AccountQuery() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	// accountNumber := c.Ctx.Input.Header("AccountNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	network := c.Ctx.Input.Header("Network")

	logs.Info("Network from header: %s", network)

	var req requests.BillPaymentAccountQueryRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	accountNumber_ := req.AccountNumber

	logs.Info("AccountQuery called with PhoneNumber: %s, SourceSystem: %s, Network: %s, AccountNumber: %s", phoneNumber, sourceSystem, accountNumber_)
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
			AccountNumber: accountNumber_,
			SourceSystem:  sourceSystem,
			PhoneNumber:   phoneNumber,
			BillerCode:    req.BillerCode,
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

// GhanaWaterAccountQuery ...
// @Title Ghana Water Account Query
// @Description Ghana Water Account Query
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	AccountNumber		header 	string true		"header for Customer's account number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	Network		header 	string true		"header for network"
// @Param	body		body 	requests.BillPaymentAccountQueryApiRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /ghana-water-account-query [post]
func (c *Api_requestsController) GhanaWaterAccountQuery() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	// accountNumber := c.Ctx.Input.Header("AccountNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	network := c.Ctx.Input.Header("Network")

	logs.Info("Network from header: %s", network)

	var req requests.BillPaymentAccountQueryApiRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	accountNumber_ := req.AccountNumber

	logs.Info("AccountQuery called with PhoneNumber: %s, SourceSystem: %s, Network: %s, AccountNumber: %s", phoneNumber, sourceSystem, accountNumber_)
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
			AccountNumber: accountNumber_,
			SourceSystem:  sourceSystem,
			PhoneNumber:   phoneNumber,
			BillerCode:    req.BillerCode,
		}

		logs.Info("Formatted request for account query ", accountQueryRequest)
		resp := apifunctions.GhanaWaterAccountQuery(&c.Controller, accountQueryRequest)
		logs.Info("Response from Account query API: ", resp)

		var response responses.AccountQueryResponse = responses.AccountQueryResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong",
			Result:        resp.Result,
		}

		if !resp.StatusCode {
			response = responses.AccountQueryResponse{
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

			response = responses.AccountQueryResponse{
				StatusCode:    true,
				StatusMessage: "Accounts queried successfully",
				Result:        resp.Result,
			}
		}

		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = response

	} else {
		var response responses.AccountQueryResponse = responses.AccountQueryResponse{
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
// @Param	AccountNumber		header 	string true		"header for Customer's account number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	body		body 	requests.DSTVPaymentApiRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /pay-dstv [post]
func (c *Api_requestsController) PayDSTV() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	accountNumber := c.Ctx.Input.Header("AccountNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	network := c.Ctx.Input.Header("Network")

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
	responseStatus := false
	responseMessage := "Something went wrong"
	result := responses.DSTVBillPaymentDataResponse{}
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

		requestIdStr := fmt.Sprintf("%d", v.Id)
		transactionLog := requests.LogTransactionRequest{
			RequestId:                requestIdStr,
			SourceAccountNumber:      accountNumber,
			DestinationAccountNumber: req.DestinationAccount,
			Amount:                   req.Amount,
			Charge:                   0.0,
			TransactionType:          "DSTV",
			ServiceCode:              "BILL_PAYMENT",
			TransactionReference:     "DSTV",
			StatusCode:               "PENDING",
			ExtraDetails1:            req.PackageType,
			ExtraDetails2:            strconv.FormatFloat(req.Amount, 'f', -1, 64),
			ExtraDetails3:            req.PackageType,
			Reference:                req.PackageType,
			ClientID:                 req.ClientId,
			PhoneNumber:              phoneNumber,
			TransactionPackage:       req.PackageType,
			ExternalReferenceNumber:  "",
		}

		if txn, err := helpers.LogTransaction(&c.Controller, transactionLog); err != nil {
			logs.Error("Error logging transaction: ", err)
			responseStatus = false
			responseMessage = "Error logging transaction: " + err.Error()

		} else {
			var response responses.DSTVBillPaymentResponse = responses.DSTVBillPaymentResponse{
				StatusCode:    false,
				StatusMessage: "Something went wrong",
				Result:        nil,
			}

			if accountNumber != "" {

				accountResp := apifunctions.GetCustomerAccount(&c.Controller, accountNumber)

				proceed := false
				if accountResp.StatusCode == "200" {
					accountCheckResp := helpers.LogAccountActivity(&c.Controller, accountNumber, "Pay DSTV", req.Amount, req.ClientId, "debit")

					if !accountCheckResp.StatusCode {
						response = responses.DSTVBillPaymentResponse{
							StatusCode:    false,
							StatusMessage: accountCheckResp.StatusMessage,
							Result:        nil,
						}
					} else {
						logs.Info("Account activity logged successfully for account number: ", accountNumber)
						// Log payment request
						makePaymentRequest := requests.PaymentRequestApiRequestDTO{
							ClientId:        req.ClientId,
							Amount:          req.Amount,
							PaymentMethod:   "ACCOUNT",
							Service:         "BILL PAYMENT",
							SenderAccount:   accountNumber,
							ReceiverAccount: req.DestinationAccount,
							Network:         network,
							ServiceNetwork:  "DSTV",
							ServicePackage:  req.PackageType,
							MobileNumber:    phoneNumber,
							TransactionId:   txn.Result.TransactionRefNumber,
						}

						helpers.MakePaymentMain(&c.Controller, makePaymentRequest)

						proceed = true
					}
				} else {
					logs.Error("Error fetching account details for account number: ", accountNumber)
					logs.Info("Register Customer")

					req := requests.PaymentRequestApiRequestDTO{
						ClientId:        req.ClientId,
						Amount:          req.Amount,
						PaymentMethod:   "MOBILEMONEY",
						Service:         "BILL PAYMENT",
						SenderAccount:   accountNumber,
						ReceiverAccount: req.DestinationAccount,
						Network:         network,
						ServiceNetwork:  "DSTV",
						ServicePackage:  req.PackageType,
						MobileNumber:    accountNumber,
						TransactionId:   txn.Result.TransactionRefNumber,
					}
					//

					resp, err := helpers.RequestPaymentMain(&c.Controller, req)
					if err != nil {
						logs.Error("Error requesting payment: ", err)
						responseStatus = false
						responseMessage = "Error requesting payment: " + err.Error()

					} else {
						logs.Info("Payment requested successfully: ", resp)
						if !resp.Success {
							responseStatus = false
							responseMessage = resp.StatusMessage
						} else {
							responseStatus = true
							responseMessage = "DSTV payment is being processed"
						}
					}
				}

				if proceed {
					// txnString := fmt.Sprintf("%d", txn.Result.TransactionId)
					payDSTVRequest := requests.DSTVPaymentRequest{
						TransactionId:      txn.Result.TransactionRefNumber,
						Amount:             req.Amount,
						DestinationAccount: destinationAccount,
						PackageType:        req.PackageType,
						SourceSystem:       sourceSystem,
						PhoneNumber:        phoneNumber,
					}

					logs.Info("Formatted request to pay DSTV: ", payDSTVRequest)
					resp := services.PayDstv(&c.Controller, payDSTVRequest)
					logs.Info("Response from DSTV payment: ", resp)

					if !resp.StatusCode {
						response = responses.DSTVBillPaymentResponse{
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
						responseStatus = true
						responseMessage = resp.StatusMessage
						result = *resp.Result
					}
				}
			} else {
				logs.Error("Account number is required to process DSTV payment")
				responseStatus = false
				responseMessage = "Account number is required to process this request"
			}
		}

		c.Ctx.Output.SetStatus(200)

	} else {
		responseStatus = false
		responseMessage = "Something went wrong:: " + err.Error()
	}

	var response responses.DSTVBillPaymentResponse = responses.DSTVBillPaymentResponse{
		StatusCode:    responseStatus,
		StatusMessage: responseMessage,
		Result:        &result,
	}
	c.Data["json"] = response
	c.ServeJSON()
}

// PayGOTV ...
// @Title Pay GOTV
// @Description Pay GOTV
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	AccountNumber		header 	string true		"header for Customer's account number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	body		body 	requests.GoTvPaymentRequest1	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /pay-gotv [post]
func (c *Api_requestsController) PayGOTV() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	accountNumber := c.Ctx.Input.Header("AccountNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	network := c.Ctx.Input.Header("Network")

	var req requests.GoTvPaymentRequest1
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

	responseStatus := false
	responseMessage := "Something went wrong"
	result := responses.GoTvBillPaymentDataResponse{}
	var v models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		RequestType:  "Buy GOTV",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)
		requestIdStr := fmt.Sprintf("%d", v.Id)
		transactionLog := requests.LogTransactionRequest{
			RequestId:                requestIdStr,
			SourceAccountNumber:      accountNumber,
			DestinationAccountNumber: req.DestinationAccount,
			Amount:                   req.Amount,
			Charge:                   0.0,
			TransactionType:          "ECG",
			ServiceCode:              "BILL_PAYMENT",
			TransactionReference:     "ECG",
			StatusCode:               "PENDING",
			ExtraDetails1:            req.PackageType,
			ExtraDetails2:            strconv.FormatFloat(req.Amount, 'f', -1, 64),
			ExtraDetails3:            req.PackageType,
			Reference:                req.PackageType,
			ClientID:                 req.ClientId,
			PhoneNumber:              phoneNumber,
			TransactionPackage:       req.PackageType,
			ExternalReferenceNumber:  "",
		}

		if txn, err := helpers.LogTransaction(&c.Controller, transactionLog); err != nil {
			logs.Error("Error logging transaction: ", err)
			responseStatus = false
			responseMessage = "Error logging transaction: " + err.Error()

		} else {
			logs.Info("Transaction logged successfully: ", txn)
			// txnString := fmt.Sprintf("%d", txn.Result.TransactionId)
			payGOTVRequest := requests.GoTvPaymentApiRequest{
				TransactionId:      txn.Result.TransactionRefNumber,
				Amount:             req.Amount,
				DestinationAccount: destinationAccount,
				PackageType:        req.PackageType,
				SourceSystem:       sourceSystem,
				PhoneNumber:        phoneNumber,
			}

			if accountNumber != "" {
				accountResp := apifunctions.GetCustomerAccount(&c.Controller, accountNumber)

				proceed := false
				if accountResp.StatusCode == "200" {
					accountCheckResp := helpers.LogAccountActivity(&c.Controller, accountNumber, "Pay GOTV", req.Amount, req.ClientId, "debit")

					if !accountCheckResp.StatusCode {
						responseStatus = false
						responseMessage = accountCheckResp.StatusMessage
					} else {
						logs.Info("Account activity logged successfully for account number: ", accountNumber)

						// Log payment request
						makePaymentRequest := requests.PaymentRequestApiRequestDTO{
							ClientId:        req.ClientId,
							Amount:          req.Amount,
							PaymentMethod:   "ACCOUNT",
							Service:         "BILL PAYMENT",
							SenderAccount:   accountNumber,
							ReceiverAccount: req.DestinationAccount,
							Network:         network,
							ServiceNetwork:  "GOTV",
							ServicePackage:  req.PackageType,
							MobileNumber:    phoneNumber,
							TransactionId:   txn.Result.TransactionRefNumber,
						}

						helpers.MakePaymentMain(&c.Controller, makePaymentRequest)

						proceed = true
					}
				} else {
					logs.Error("Error fetching account details for account number: ", accountNumber)
					logs.Info("Register Customer")

					req := requests.PaymentRequestApiRequestDTO{
						ClientId:        req.ClientId,
						Amount:          req.Amount,
						PaymentMethod:   "MOBILEMONEY",
						Service:         "BILL PAYMENT",
						SenderAccount:   accountNumber,
						ReceiverAccount: req.DestinationAccount,
						Network:         network,
						ServiceNetwork:  "GOTV",
						ServicePackage:  req.PackageType,
						MobileNumber:    accountNumber,
						TransactionId:   txn.Result.TransactionRefNumber,
					}
					//

					resp, err := helpers.RequestPaymentMain(&c.Controller, req)
					if err != nil {
						logs.Error("Error requesting payment: ", err)
						responseStatus = false
						responseMessage = "Error requesting payment: " + err.Error()
					} else {
						logs.Info("Payment requested successfully: ", resp)
						if !resp.Success {
							responseStatus = false
							responseMessage = resp.StatusMessage
						} else {

							responseStatus = true
							responseMessage = "GOTV payment is being processed"

						}
					}
				}

				if proceed {
					logs.Info("Formatted request to pay GOTV: ", payGOTVRequest)
					resp := services.PayGotv(&c.Controller, payGOTVRequest)
					logs.Info("Response from GOTV API: ", resp)

					var response responses.GoTvBillPaymentResponse = responses.GoTvBillPaymentResponse{
						StatusCode:    false,
						StatusMessage: "Something went wrong",
						Result:        resp.Result,
					}

					if !resp.StatusCode {
						responseStatus = false
						responseMessage = resp.StatusMessage
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

						responseStatus = true
						responseMessage = resp.StatusMessage
						result = *resp.Result
					}
				}
			} else {
				responseStatus = false
				responseMessage = "Account number is required to process this request"
				logs.Error("Account number is required to process GOTV payment")
			}
		}

		c.Ctx.Output.SetStatus(200)

	} else {
		responseStatus = false
		responseMessage = "Something went wrong"
	}

	var response responses.GoTvBillPaymentResponse = responses.GoTvBillPaymentResponse{
		StatusCode:    responseStatus,
		StatusMessage: responseMessage,
		Result:        &result,
	}
	c.Data["json"] = response
	c.ServeJSON()
}

// PayECG ...
// @Title Pay ECG
// @Description Pay ECG
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	AccountNumber		header 	string true		"header for Customer's account number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	body		body 	requests.ECGPaymentRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /pay-ecg [post]
func (c *Api_requestsController) PayECG() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	accountNumber := c.Ctx.Input.Header("AccountNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	network := c.Ctx.Input.Header("Network")

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

	responseStatus := false
	responseMessage := "Something went wrong"
	result := responses.ECGBillPaymentDataResponse{}
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
		requestIdStr := fmt.Sprintf("%d", v.Id)
		transactionLog := requests.LogTransactionRequest{
			RequestId:                requestIdStr,
			SourceAccountNumber:      accountNumber,
			DestinationAccountNumber: req.DestinationAccount,
			Amount:                   req.Amount,
			Charge:                   0.0,
			TransactionType:          "ECG",
			ServiceCode:              "BILL_PAYMENT",
			TransactionReference:     "ECG",
			StatusCode:               "PENDING",
			ExtraDetails1:            req.PackageType,
			ExtraDetails2:            strconv.FormatFloat(req.Amount, 'f', -1, 64),
			ExtraDetails3:            req.PackageType,
			Reference:                req.PackageType,
			ClientID:                 req.ClientId,
			PhoneNumber:              phoneNumber,
			TransactionPackage:       req.PackageType,
			ExternalReferenceNumber:  "",
		}

		if txn, err := helpers.LogTransaction(&c.Controller, transactionLog); err != nil {
			logs.Error("Error logging transaction: ", err)
			responseStatus = false
			responseMessage = "Error logging transaction: " + err.Error()

		} else {
			// txnString := fmt.Sprintf("%d", txn.Result.TransactionId)
			payECGRequest := requests.ECGPaymentApiRequest{
				TransactionId:      txn.Result.TransactionRefNumber,
				Amount:             req.Amount,
				DestinationAccount: destinationAccount,
				PackageType:        req.PackageType,
				SourceSystem:       sourceSystem,
				PhoneNumber:        phoneNumber,
			}

			var response responses.ECGBillPaymentResponse = responses.ECGBillPaymentResponse{
				StatusCode:    false,
				StatusMessage: "Something went wrong",
				Result:        nil,
			}

			if accountNumber != "" {
				accountResp := apifunctions.GetCustomerAccount(&c.Controller, accountNumber)

				proceed := false
				if accountResp.StatusCode == "200" {
					accountCheckResp := helpers.LogAccountActivity(&c.Controller, accountNumber, "Pay ECG", req.Amount, req.ClientId, "debit")

					if !accountCheckResp.StatusCode {
						responseStatus = false
						responseMessage = accountCheckResp.StatusMessage
					} else {
						logs.Info("Account activity logged successfully for account number: ", accountNumber)
						// Log payment request
						makePaymentRequest := requests.PaymentRequestApiRequestDTO{
							ClientId:        req.ClientId,
							Amount:          req.Amount,
							PaymentMethod:   "ACCOUNT",
							Service:         "BILL PAYMENT",
							SenderAccount:   accountNumber,
							ReceiverAccount: req.DestinationAccount,
							Network:         network,
							ServiceNetwork:  "ECG",
							ServicePackage:  req.PackageType,
							MobileNumber:    phoneNumber,
							TransactionId:   txn.Result.TransactionRefNumber,
						}

						helpers.MakePaymentMain(&c.Controller, makePaymentRequest)

						proceed = true
					}
				} else {
					logs.Error("Error fetching account details for account number: ", accountNumber)
					logs.Info("Register Customer")

					req := requests.PaymentRequestApiRequestDTO{
						ClientId:        req.ClientId,
						Amount:          req.Amount,
						PaymentMethod:   "MOBILEMONEY",
						Service:         "BILL PAYMENT",
						SenderAccount:   accountNumber,
						ReceiverAccount: req.DestinationAccount,
						Network:         network,
						ServiceNetwork:  "ECG",
						ServicePackage:  req.PackageType,
						MobileNumber:    accountNumber,
						TransactionId:   txn.Result.TransactionRefNumber,
					}
					//

					resp, err := helpers.RequestPaymentMain(&c.Controller, req)
					if err != nil {
						logs.Error("Error requesting payment: ", err)
						responseStatus = false
						responseMessage = "Error requesting payment: " + err.Error()
					} else {
						logs.Info("Payment requested successfully: ", resp)
						if !resp.Success {
							responseStatus = false
							responseMessage = resp.StatusMessage
						} else {
							responseStatus = true
							responseMessage = "Payment is being processed"
						}
					}
				}

				if proceed {
					logs.Info("Formatted request to pay ECB: ", payECGRequest)
					resp := services.PayEcg(&c.Controller, payECGRequest)
					logs.Info("Response from pay ECB: ", resp)

					if !resp.StatusCode {
						response = responses.ECGBillPaymentResponse{
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

						responseStatus = true
						responseMessage = "Payment is being processed"
						result = *resp.Result
					}
				}
			} else {
				responseStatus = false
				responseMessage = "Account number is required to process this request"
			}
		}

		c.Ctx.Output.SetStatus(200)
	} else {
		responseStatus = false
		responseMessage = "Something went wrong:: " + err.Error()
	}

	var response responses.ECGBillPaymentResponse = responses.ECGBillPaymentResponse{
		StatusCode:    responseStatus,
		StatusMessage: responseMessage,
		Result:        &result,
	}
	c.Data["json"] = response
	c.ServeJSON()
}

// PayWater ...
// @Title Pay Water bill
// @Description Pay Water Bill
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	AccountNumber		header 	string true		"header for Customer's account number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	body		body 	requests.GhanaWaterPaymentRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /pay-water [post]
func (c *Api_requestsController) PayWaterBill() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	accountNumber := c.Ctx.Input.Header("AccountNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	network := c.Ctx.Input.Header("Network")

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

	responseStatus := false
	responseMessage := "Something went wrong"

	result_ := responses.GhanaWaterBillPaymentDataResponse{}

	var response responses.GhanaWaterBillPaymentResponse = responses.GhanaWaterBillPaymentResponse{
		StatusCode:    responseStatus,
		StatusMessage: responseMessage,
		Result:        nil,
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
		logs.Info("Log transaction")

		requestIdStr := fmt.Sprintf("%d", v.Id)
		transactionLog := requests.LogTransactionRequest{
			RequestId:                requestIdStr,
			SourceAccountNumber:      accountNumber,
			DestinationAccountNumber: req.DestinationAccount,
			Amount:                   req.Amount,
			Charge:                   0.0,
			TransactionType:          "GH_WATER",
			ServiceCode:              "BILL_PAYMENT",
			TransactionReference:     "GH_WATER",
			StatusCode:               "PENDING",
			ExtraDetails1:            req.CustomerName,
			ExtraDetails2:            req.CustomerEmail,
			ExtraDetails3:            req.PackageType,
			Reference:                req.PackageType,
			ClientID:                 req.ClientId,
			PhoneNumber:              phoneNumber,
			TransactionPackage:       req.PackageType,
			ExternalReferenceNumber:  "",
		}

		if txn, err := helpers.LogTransaction(&c.Controller, transactionLog); err != nil {
			logs.Error("Error logging transaction: ", err)
			responseStatus = false
			responseMessage = "Error logging transaction: " + err.Error()

		} else {
			logs.Info("Transaction logged successfully")

			// transactionString := fmt.Sprintf("%d", txn.Result.TransactionId)
			payWaterRequest := requests.GhanaWaterPaymentFuncRequest{
				TransactionId:      txn.Result.TransactionRefNumber,
				Amount:             req.Amount,
				DestinationAccount: destinationAccount,
				PackageType:        req.PackageType,
				SourceSystem:       sourceSystem,
				PhoneNumber:        phoneNumber,
				Name:               req.CustomerName,
				Email:              req.CustomerEmail,
			}

			if accountNumber != "" {
				accountResp := apifunctions.GetCustomerAccount(&c.Controller, accountNumber)

				proceed := false
				if accountResp.StatusCode == "200" {
					logs.Info("Client ID::: ", req.ClientId)
					accountCheckResp := helpers.LogAccountActivity(&c.Controller, accountNumber, "Pay Water", req.Amount, req.ClientId, "debit")

					if !accountCheckResp.StatusCode {
						logs.Error("Error logging account activity for account number: ", accountNumber)
						responseStatus = false
						responseMessage = accountCheckResp.StatusMessage
					} else {
						logs.Info("Account activity logged successfully for account number: ", accountNumber)

						// Log payment request
						makePaymentRequest := requests.PaymentRequestApiRequestDTO{
							ClientId:        req.ClientId,
							Amount:          req.Amount,
							PaymentMethod:   "ACCOUNT",
							Service:         "BILL PAYMENT",
							SenderAccount:   accountNumber,
							ReceiverAccount: req.DestinationAccount,
							Network:         network,
							ServiceNetwork:  "WATER",
							ServicePackage:  req.PackageType,
							MobileNumber:    phoneNumber,
							TransactionId:   txn.Result.TransactionRefNumber,
						}

						helpers.MakePaymentMain(&c.Controller, makePaymentRequest)

						proceed = true
					}
				} else {
					logs.Error("Error fetching account details for account number: ", accountNumber)
					logs.Info("Register Customer")

					req := requests.PaymentRequestApiRequestDTO{
						ClientId:        req.ClientId,
						Amount:          req.Amount,
						PaymentMethod:   "MOBILEMONEY",
						Service:         "BILL PAYMENT",
						SenderAccount:   accountNumber,
						ReceiverAccount: req.DestinationAccount,
						Network:         network,
						ServiceNetwork:  "WATER",
						ServicePackage:  req.PackageType,
						MobileNumber:    accountNumber,
						TransactionId:   txn.Result.TransactionRefNumber,
					}
					//

					resp, err := helpers.RequestPaymentMain(&c.Controller, req)
					if err != nil {
						logs.Error("Error requesting payment: ", err)
						responseStatus = false
						responseMessage = "Error requesting payment: " + err.Error()
					} else {
						logs.Info("Payment requested successfully: ", resp)
						if !resp.Success {
							responseStatus = false
							responseMessage = resp.StatusMessage
						} else {
							responseStatus = true
							responseMessage = "Water bill purchase is being processed"
						}
					}
				}

				if proceed {
					logs.Info("Formatted request to pay water: ", payWaterRequest)
					resp := services.PayWater(&c.Controller, payWaterRequest)
					logs.Info("Response from pay water: ", resp)

					if !resp.StatusCode {
						responseStatus = false
						responseMessage = resp.StatusMessage
						result_ = *resp.Result
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
						responseStatus = true
						responseMessage = resp.StatusMessage
						result_ = *resp.Result
					}
				}
			} else {
				responseStatus = false
				responseMessage = "Account number is required to process this request"
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

	response = responses.GhanaWaterBillPaymentResponse{
		StatusCode:    responseStatus,
		StatusMessage: responseMessage,
		Result:        &result_,
	}
	c.Data["json"] = response
	c.ServeJSON()
}

// PayStartimes ...
// @Title Pay Startimes bill
// @Description Pay Startimes Bill
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	AccountNumber		header 	string true		"header for Customer's account number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	body		body 	requests.StartimesPaymentRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /pay-startimes [post]
func (c *Api_requestsController) PayStartimesTvBill() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	accountNumber := c.Ctx.Input.Header("AccountNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	network := c.Ctx.Input.Header("Network")

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

	responseStatus := false
	responseMessage := "Something went wrong"
	result := responses.StartimesBillPaymentDataResponse{}
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
		logs.Info("Log transaction")

		requestIdStr := fmt.Sprintf("%d", v.Id)
		transactionLog := requests.LogTransactionRequest{
			RequestId:                requestIdStr,
			SourceAccountNumber:      accountNumber,
			DestinationAccountNumber: req.DestinationAccount,
			Amount:                   req.Amount,
			Charge:                   0.0,
			TransactionType:          "STARTIMES",
			ServiceCode:              "BILL_PAYMENT",
			TransactionReference:     "STARTIMES",
			StatusCode:               "PENDING",
			ExtraDetails1:            req.PhoneNumber,
			ExtraDetails2:            req.PackageType,
			ExtraDetails3:            req.PackageType,
			Reference:                req.PackageType,
			ClientID:                 req.ClientId,
			PhoneNumber:              phoneNumber,
			TransactionPackage:       req.PackageType,
			ExternalReferenceNumber:  "",
		}

		if txn, err := helpers.LogTransaction(&c.Controller, transactionLog); err != nil {
			logs.Error("Error logging transaction: ", err)
			responseStatus = false
			responseMessage = "Error logging transaction: " + err.Error()

		} else {
			// transactionString := fmt.Sprintf("%d", txn.Result.TransactionId)
			payStartimesRequest := requests.StartimesPaymentApiRequest{
				TransactionId:      txn.Result.TransactionRefNumber,
				Amount:             req.Amount,
				DestinationAccount: destinationAccount,
				PackageType:        req.PackageType,
				SourceSystem:       sourceSystem,
				PhoneNumber:        phoneNumber,
			}

			accountResp := apifunctions.GetCustomerAccount(&c.Controller, accountNumber)

			if accountNumber != "" {

				proceed := false
				if accountResp.StatusCode == "200" {
					accountCheckResp := helpers.LogAccountActivity(&c.Controller, accountNumber, "Pay Startimes", req.Amount, req.ClientId, "debit")

					if !accountCheckResp.StatusCode {
						responseStatus = false
						responseMessage = accountCheckResp.StatusMessage
					} else {
						logs.Info("Account activity logged successfully for account number: ", accountNumber)

						// Log payment request
						makePaymentRequest := requests.PaymentRequestApiRequestDTO{
							ClientId:        req.ClientId,
							Amount:          req.Amount,
							PaymentMethod:   "ACCOUNT",
							Service:         "BILL PAYMENT",
							SenderAccount:   accountNumber,
							ReceiverAccount: req.DestinationAccount,
							Network:         network,
							ServiceNetwork:  "STARTIMES",
							ServicePackage:  req.PackageType,
							MobileNumber:    phoneNumber,
							TransactionId:   txn.Result.TransactionRefNumber,
						}

						helpers.MakePaymentMain(&c.Controller, makePaymentRequest)

						proceed = true
					}
				} else {
					logs.Error("Error fetching account details for account number: ", accountNumber)
					logs.Info("Register Customer")

					req := requests.PaymentRequestApiRequestDTO{
						ClientId:        req.ClientId,
						Amount:          req.Amount,
						PaymentMethod:   "MOBILEMONEY",
						Service:         "BILL PAYMENT",
						SenderAccount:   accountNumber,
						ReceiverAccount: req.DestinationAccount,
						Network:         network,
						ServiceNetwork:  "STARTIMES",
						ServicePackage:  req.PackageType,
						MobileNumber:    accountNumber,
						TransactionId:   txn.Result.TransactionRefNumber,
					}
					//

					resp, err := helpers.RequestPaymentMain(&c.Controller, req)
					if err != nil {
						logs.Error("Error requesting payment: ", err)
						responseStatus = false
						responseMessage = "Error requesting payment: " + err.Error()
					} else {
						logs.Info("Payment requested successfully: ", resp)
						if !resp.Success {
							responseStatus = false
							responseMessage = resp.StatusMessage
						} else {
							responseStatus = true
							responseMessage = "Startimes payment is being processed"
						}
					}
				}

				if proceed {
					logs.Info("Formatted request for pay startimes: ", payStartimesRequest)
					resp := apifunctions.PayStartimesBill(&c.Controller, payStartimesRequest)
					logs.Info("Response from pay startimes API: ", resp)

					if !resp.StatusCode {
						responseStatus = false
						responseMessage = resp.StatusMessage
						if resp.Result != nil {
							result = *resp.Result
						}
					} else {
						responseText, err := json.Marshal(resp.Result)
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

						if accountNumber != "" {
							helpers.LogAccountActivity(&c.Controller, accountNumber, "Startimes Purchase", req.Amount, req.ClientId, "debit")
						}

						responseStatus = true
						responseMessage = "Payment is being processed"
						if resp.Result != nil {
							result = *resp.Result
						}
					}
				}
			} else {
				responseStatus = false
				responseMessage = "Account number is required to process this request"
			}
		}

		c.Ctx.Output.SetStatus(200)

	} else {
		responseStatus = false
		responseMessage = "Something went wrong" + err.Error()
	}

	var response responses.StartimesBillPaymentApiResponse = responses.StartimesBillPaymentApiResponse{
		StatusCode:    responseStatus,
		StatusMessage: responseMessage,
		Result:        &result,
	}
	c.Data["json"] = response
	c.ServeJSON()
}

// ValidateCustomer ...
// @Title Validate Customer
// @Description Validate Customer
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	AccountNumber		header 	string true		"header for Customer's account number"
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

// GetCustomerCorporativeAccounts ...
// @Title Get Customer Corporative Accounts
// @Description Get customer accounts
// @Param	Authorization		header 	string true		"header for User"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	body		body 	requests.NumberExistsApiRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /get-customer-corporative-accounts [post]
// func (c *Api_requestsController) GetCustomerCorporativeAccounts() {
// 	// Extract headers
// 	// phoneNumber := c.Ctx.Input.Header("PhoneNumber")
// 	sourceSystem := c.Ctx.Input.Header("SourceSystem")

// 	var req requests.NumberExistsApiRequest
// 	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

// 	phoneNumber := req.MobileNumber

// 	logs.Info("Get customer accounts called with PhoneNumber: %s, SourceSystem: %s", phoneNumber, sourceSystem)
// 	reqBody := c.Ctx.Input.RequestBody
// 	reqHeaders := c.Ctx.Request.Header

// 	requestMap := map[string]interface{}{
// 		"headers": reqHeaders,
// 		"body":    string(reqBody),
// 	}

// 	reqText, err := json.Marshal(requestMap)
// 	if err != nil {
// 		logs.Error("Error marshalling request input: ", err)
// 		c.Data["json"] = err.Error()
// 		c.ServeJSON()
// 		return
// 	}
// 	var v models.Api_requests = models.Api_requests{
// 		Request:      string(reqText),
// 		PhoneNumber:  phoneNumber,
// 		RequestType:  "List Customer Accounts",
// 		RequestDate:  time.Now(),
// 		DateCreated:  time.Now(),
// 		DateModified: time.Now(),
// 	}
// 	if _, err := models.AddApi_requests(&v); err == nil {
// 		logs.Info("API request logged successfully: ", v)

// 		var response responses.CustomerAccountsResponse = responses.CustomerAccountsResponse{
// 			StatusCode:    false,
// 			StatusMessage: "Something went wrong",
// 			Result:        nil,
// 		}

// 		var clientId int64
// 		clientId, err = strconv.ParseInt(req.ClientId, 10, 64)
// 		if err != nil {
// 			logs.Error("Error converting ClientId to int64: ", err)
// 			response = responses.CustomerAccountsResponse{
// 				StatusCode:    false,
// 				StatusMessage: "Invalid ClientId",
// 				Result:        nil,
// 			}
// 			c.Ctx.Output.SetStatus(400)
// 			c.Data["json"] = response
// 		}

// 		if client, err := models.GetClientsById(clientId); err == nil {

// 			clientCorpId := client.ClientCorpId

// 			listAccountsRequest := requests.NumberExistsApiRequest{
// 				MobileNumber: phoneNumber,
// 				ClientId:     clientCorpId,
// 			}

// 			logs.Info("Formatted request for customer accounts: ", listAccountsRequest)
// 			resp := apifunctions.ListCustomerAccounts(&c.Controller, listAccountsRequest)
// 			logs.Info("Response from customer accounts API: ", resp)

// 			if resp.Data.StatusCode != 200 {
// 				response = responses.CustomerAccountsResponse{
// 					StatusCode:    false,
// 					StatusMessage: resp.Data.StatusMessage,
// 					Result:        nil,
// 				}
// 			} else {
// 				responseText, err := json.Marshal(response.Result)
// 				if err != nil {
// 					logs.Error("Error marshalling response result: ", err)
// 					responseText = []byte("[]")
// 				}
// 				v.RequestResponse = string(responseText)
// 				v.DateModified = time.Now()
// 				v.ResponseDate = time.Now()
// 				if err := models.UpdateApi_requestsById(&v); err != nil {
// 					logs.Error("Error updating API request with response: ", err)
// 				} else {
// 					logs.Info("API request updated with response successfully: ", v)
// 				}
// 				response = responses.CustomerAccountsResponse{
// 					StatusCode:    true,
// 					StatusMessage: "Accounts fetched successfully",
// 					Result:        resp.Data.Result,
// 				}
// 			}

// 			c.Ctx.Output.SetStatus(200)
// 			c.Data["json"] = response
// 		} else {
// 			logs.Error("Error fetching client details: ", err)
// 			response = responses.CustomerAccountsResponse{
// 				StatusCode:    false,
// 				StatusMessage: "Something went wrong:: " + err.Error(),
// 				Result:        nil,
// 			}
// 			c.Ctx.Output.SetStatus(400)
// 			c.Data["json"] = response
// 		}

// 	} else {
// 		var response responses.CustomerAccountsResponse = responses.CustomerAccountsResponse{
// 			StatusCode:    false,
// 			StatusMessage: "Something went wrong:: " + err.Error(),
// 			Result:        nil,
// 		}

// 		c.Data["json"] = response
// 	}
// 	c.ServeJSON()
// }
