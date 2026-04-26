package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	apifunctions "msys_payment_app_gateway/controllers/api_functions"
	"msys_payment_app_gateway/controllers/helpers"
	"msys_payment_app_gateway/models"
	"msys_payment_app_gateway/structs/requests"
	"msys_payment_app_gateway/structs/responses"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Agent_api_requestsController operations for Agent_api_requests
type Agent_api_requestsController struct {
	beego.Controller
}

// URLMapping ...
func (c *Agent_api_requestsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("ListAccountLoans", c.ListAccountLoans)
	c.Mapping("LoanRepayment", c.LoanRepayment)
	c.Mapping("Deposit", c.Deposit)
	c.Mapping("GetUserDetails", c.GetUserDetails)
	c.Mapping("GetAgentTransactions", c.GetAgentTransactions)
	c.Mapping("AccountBalance", c.AccountBalance)
	c.Mapping("ListAccountDetails", c.ListAccountDetails)
	c.Mapping("GetBilTransactionWithTransactionRef", c.GetBilTransactionWithTransactionRef)
}

// GetUserDetails ...
// @Title Get User Details
// @Description Get User Details
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /get-user-details [post]
func (c *Agent_api_requestsController) GetUserDetails() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	// sourceSystem := c.Ctx.Input.Header("SourceSystem")
	user := c.Ctx.Input.GetData("user")

	logs.Info("User details: %s", user)
	userData, ok := user.(*responses.UsersOri)
	if !ok {
		logs.Error("Error asserting user data")
		c.Data["json"] = "Invalid user data"
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
		RequestType:  "Get User details",
		PhoneNumber:  phoneNumber,
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)

		var response responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{
			Success:    false,
			StatusDesc: "User fetch failed",
			Result:     nil,
		}

		userStatus := "INACTIVE"
		if userData.Active == 1 {
			userStatus = "ACTIVE"
		}

		branchResp := responses.BranchesResponse{
			BranchName:  userData.UserDetails.Branch.Branch,
			Country:     userData.UserDetails.Branch.Country.Country,
			Location:    userData.UserDetails.Branch.Location,
			PhoneNumber: userData.UserDetails.Branch.PhoneNumber,
			Active:      userData.UserDetails.Branch.Active,
		}

		var fields []string
		var sortby []string
		var order []string
		var query = make(map[string]string)
		var limit int64 = 10
		var offset int64

		userIdSearch := "UserId:" + strconv.FormatInt(userData.UserId, 10)

		if v := userIdSearch; v != "" {
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

		logs.Debug("Query for user corporatives is ", query)

		var userCorpsDTO []responses.UserCorporativesResponseDTO
		if userCorps, err := models.GetAllUser_corporatives(query, fields, sortby, order, offset, limit); err == nil {
			logs.Debug("Returned user corporatives data is ", userCorps)
			for _, v := range userCorps {
				logs.Debug("Processing user corporative: ", v)
				var corpDTO responses.UserCorporativesResponseDTO
				corpBytes, err := json.Marshal(v)
				if err != nil {
					logs.Error("Error marshalling user corporative data: ", err)
					continue
				}
				if err := json.Unmarshal(corpBytes, &corpDTO); err != nil {
					logs.Error("Error unmarshalling user corporative data: ", err)
					continue
				}
				userCorpsDTO = append(userCorpsDTO, corpDTO)
			}

			// Log user corporatives data as readable JSON
			corpsJSON, err := json.MarshalIndent(userCorpsDTO, "", "  ")
			if err != nil {
				logs.Error("Error marshalling user corporatives to JSON: ", err)
			} else {
				logs.Debug("Formatted user corporatives data is: %s", string(corpsJSON))
			}
		} else {
			logs.Error("Error fetching user corporatives: ", err)
		}

		userResp := responses.UserGateway{
			UserId:         userData.UserId,
			FullName:       userData.FullName,
			Username:       userData.Username,
			ImagePath:      userData.ImagePath,
			Email:          userData.Email,
			PhoneNumber:    userData.PhoneNumber,
			Status:         userStatus,
			DateRegistered: userData.DateCreated,
			Customer:       userData.UserDetails,
			IsVerified:     userData.IsVerified,
			Role:           userData.Role,
			Branch:         &branchResp,
			Corporatives:   &userCorpsDTO,
		}

		logs.Info("Formatted request for customer: ")

		response = responses.UserGatewayResponseDTO{
			Success:    true,
			StatusDesc: "Customer fetched successfully",
			Result:     &userResp,
		}

		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = response

	} else {
		var response responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{
			Success:    false,
			StatusDesc: "Something went wrong:: " + err.Error(),
			Result:     nil,
		}

		c.Data["json"] = response
	}
	logs.Info("Final response to be sent: ", c.Data["json"])
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
func (c *Agent_api_requestsController) Deposit() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	accountNumber := c.Ctx.Input.Header("AccountNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	network := c.Ctx.Input.Header("Network")

	user := c.Ctx.Input.GetData("user")

	logs.Info("User details: %s", user)
	userData, ok := user.(*responses.UsersOri)
	if !ok {
		logs.Error("Error asserting user data")
		c.Data["json"] = "Invalid user data"
		c.ServeJSON()
		return
	}

	var req requests.AgentDepositRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	destinationPhoneNumber := req.Destination
	clientId := req.ClientId
	paymentMethod := req.PaymentMethod

	if paymentMethod == "" {
		paymentMethod = "CASH"
	}

	if paymentMethod == "MOBILEMONEY" {
		paymentMethod = "MOMO"
	}

	if clientId == "" {
		clientId, _ = beego.AppConfig.String("msysconsultCode")
	}

	logs.Info("Client ID for deposit is ", clientId)
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
	txnData := responses.DepositData{}

	var response responses.DepositResponse = responses.DepositResponse{
		Success:       isSuccess,
		StatusMessage: message,
		Result:        nil,
	}
	var v models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		PhoneNumber:  phoneNumber,
		RequestType:  "Agent Deposit",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
		CreatedBy:    int(userData.UserId),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)

		requestIdStr := fmt.Sprintf("%d", v.Id)
		extraData := requests.ExtraData{
			ExtraData1: strings.Trim(req.CustomerName, " "),
			ExtraData2: strings.Trim(req.CustomerNumber, " "),
			ExtraData3: paymentMethod,
		}
		transactionLog := requests.UserTransactionRequestDTO{
			SourceChannel:            sourceSystem,
			RequestId:                requestIdStr,
			SourceAccountNumber:      accountNumber,
			DestinationAccountNumber: req.Destination,
			Amount:                   req.Amount,
			ServiceCode:              "DEPOSIT",
			ClientReference:          "",
			ExtraData:                extraData,
			Package:                  network,
			PhoneNumber:              phoneNumber,
			CreatedBy:                strconv.FormatInt(userData.UserId, 10),
			Status:                   "PENDING",
		}

		if txn := apifunctions.LogUserTransaction(&c.Controller, transactionLog); txn.StatusCode != 200 {
			logs.Error("Error logging transaction: ", txn.StatusDesc)
			isSuccess = false
			message = "Error logging transaction: " + txn.StatusDesc

		} else {

			transferRequest := requests.TransferApiRequest{
				RequestId:              requestIdStr,
				Amount:                 req.Amount,
				Charge:                 0.0,
				Commission:             0.0,
				TotalDebitAmount:       req.Amount + 0.0,
				SenderAccountNumber:    accountNumber,
				RecipientAccountNumber: req.Destination,
				TransferCode:           req.PaymentMethod,
				Description:            "Deposit for transaction " + requestIdStr,
				RecipientName:          network,
				Status:                 "PENDING",
				ServiceCode:            "DEPOSIT",
				CreatedBy:              strconv.FormatInt(userData.UserId, 10),
			}

			sendCommission := false
			if req.PaymentMethod == "MOBILEMONEY" {
				sendCommission = true
			}

			if _, err := helpers.LogTransferTransaction(&c.Controller, transferRequest, sendCommission); err != nil {
				logs.Error("Error logging transfer transaction: ", err)
				isSuccess = false
				message = "Error logging transfer transaction: " + err.Error()
			} else {
				logs.Info("Transfer transaction logged successfully: ", txn)

				if req.PaymentMethod == "MOBILEMONEY" {
					req2 := requests.PaymentRequestApiRequestDTO{
						ClientId:            clientId,
						Amount:              req.Amount,
						PaymentMethod:       req.PaymentMethod,
						Service:             "DEPOSIT",
						SenderAccount:       accountNumber,
						ReceiverAccount:     destinationPhoneNumber,
						Network:             network,
						ServiceNetwork:      req.ClientId,
						ServicePackage:      strconv.FormatFloat(req.Amount, 'f', -1, 64),
						MobileNumber:        req.Source,
						TransactionId:       txn.Result.TransactionId,
						CallbackServiceCode: "USER_PAYMENT",
					}
					//

					logs.Info("Amount to debit is ", req.Amount)

					resp, err := helpers.RequestPaymentMain(&c.Controller, req2)

					logs.Info("Response from Deposit API: ", resp)

					if err != nil {
						message = err.Error()
					} else {
						if resp.Success {
							// accountCheckResp := helpers.LogAccountActivity(&c.Controller, accountNumber, "Deposit", req.Amount, req.ClientId, "credit")

							isSuccess = true
							message = "Deposit successful"
							responseText, err := json.Marshal(resp)

							if err != nil {
								logs.Error("Error marshalling response result: ", err)
								responseText = []byte("[]")
							}
							v.RequestResponse = string(responseText)
						} else {
							message = "Deposit failed: " + resp.StatusMessage
						}
					}
				} else {
					isSuccess = true
					message = "Deposit successful"
				}

				if isSuccess {
					logs.Info("Proceed....Client ID is ", clientId)
					sendDepositRequest := requests.SendDepositRequest{
						Amount:        req.Amount,
						AccountNumber: req.Destination,
						MobileNumber:  phoneNumber,
						PaymentMethod: paymentMethod,
						ClientId:      clientId,
					}

					if resp := apifunctions.SendDeposit(&c.Controller, sendDepositRequest); resp.Data.StatusCode != 200 {
						logs.Error("Error sending deposit to core banking: ", err)
						isSuccess = false
						message = "Error sending deposit to core banking: " + resp.Data.StatusDesc
					} else {
						logs.Info("Deposit sent to core banking successfully")
					}
				}

				txnData.Amount = txn.Result.Amount
				txnData.Currency = txn.Result.TransactingCurrency
				txnData.TransactionReference = txn.Result.TransactionId

				v.DateModified = time.Now()
				v.ResponseDate = time.Now()
				if err := models.UpdateApi_requestsById(&v); err != nil {
					logs.Error("Error updating API request with response: ", err)
				} else {
					logs.Info("API request updated with response successfully: ", v)
				}
			}
		}

		c.Ctx.Output.SetStatus(200)

	} else {
		logs.Error("Error logging API request: ", err)
	}

	response = responses.DepositResponse{
		Success:       isSuccess,
		StatusMessage: message,
		Result:        &txnData,
	}

	c.Data["json"] = response

	c.ServeJSON()
}

// Get Account Loans ...
// @Title GetAccountLoans
// @Description Get account loans
// @Param	body		body 	requests.AccountLoansRequest	true		"body for crediting of account"
// @Param	clientId		header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /list-account-loans [post]
func (c *Agent_api_requestsController) ListAccountLoans() {
	clientId := c.Ctx.Input.Header("clientId")
	logs.Debug("Client id is ", clientId)

	user := c.Ctx.Input.GetData("user")

	logs.Info("User details: %s", user)
	userData, ok := user.(*responses.UsersOri)
	if !ok {
		logs.Error("Error asserting user data")
		c.Data["json"] = "Invalid user data"
		c.ServeJSON()
		return
	}

	var v requests.AccountLoansRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	status := false
	statusMessage := "Error retrieving account loans"

	// logs.Debug("Request::: ", c.Ctx.Input.RequestBody)
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

	var response responses.ListLoansResponse

	var result []responses.LoanData
	var ap models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		PhoneNumber:  v.AccountNumber,
		RequestType:  "List Account Loans",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
		CreatedBy:    int(userData.UserId),
	}
	if _, err := models.AddApi_requests(&ap); err == nil {
		logs.Info("API request logged successfully: ", v)

		logs.Info("Formatted request for account loans: ", string(reqText))
		resp := apifunctions.ListAccountLoans(&c.Controller, clientId, v.AccountNumber)

		logs.Debug("Response is ", resp)

		if resp.StatusCode == 200 {
			logs.Info("Successfully fetched loan list")
			status = true
			statusMessage = "Successfully fetched account loans"
			if resp.Result != nil {
				result = *resp.Result
			} else {
				result = []responses.LoanData{}
				logs.Info("No loans found for the account")
				statusMessage = "No loans found for the account"
			}
		} else {
			logs.Error("Error fetching account loans")
			statusMessage = resp.StatusDesc
		}
	} else {
		logs.Error("Error logging API request: ", err)
		statusMessage = "Error logging API request: " + err.Error()
	}

	response = responses.ListLoansResponse{
		StatusCode: status,
		StatusDesc: statusMessage,
		Result:     &result,
	}

	c.Data["json"] = response

	c.ServeJSON()
}

// Loan Repayment ...
// @Title LoanRepayment
// @Description Repay loan
// @Param	body		body 	requests.LoanRepaymentRequest	true		"body for crediting of account"
// @Param	clientId		header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /loan-repayment [post]
func (c *Agent_api_requestsController) LoanRepayment() {
	clientId := c.Ctx.Input.Header("clientId")
	logs.Debug("Client id is ", clientId)
	// sourceSystem := c.Ctx.Input.Header("SourceSystem")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	network := c.Ctx.Input.Header("Network")

	user := c.Ctx.Input.GetData("user")

	logs.Info("User details: %s", user)
	userData, ok := user.(*responses.UsersOri)
	if !ok {
		logs.Error("Error asserting user data")
		c.Data["json"] = "Invalid user data"
		c.ServeJSON()
		return
	}

	var v requests.LoanRepaymentRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	status := false
	statusMessage := "Error retrieving account loans"
	var response responses.RepayLoanResponse
	txnData := responses.DepositData{}

	// logs.Debug("Request::: ", c.Ctx.Input.RequestBody)
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

	paymentMethod := v.PaymentMethod

	if paymentMethod == "" {
		paymentMethod = "CASH"
	}

	if paymentMethod == "MOBILEMONEY" {
		paymentMethod = "MOMO"
	}

	var ap models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		PhoneNumber:  userData.PhoneNumber,
		RequestType:  "Loan Repayment",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
		CreatedBy:    int(userData.UserId),
	}
	if apiReq, err := models.AddApi_requests(&ap); err == nil {
		logs.Info("API request logged successfully: ", v)

		requestIdStr := fmt.Sprintf("%d", apiReq)
		logs.Info("Amount from request is ", v.Amount)
		vAmountFloat, _ := strconv.ParseFloat(v.Amount, 64)
		logs.Info("Float amount is ", vAmountFloat)
		// transactionLog := requests.LogTransactionRequest{
		// 	RequestId:                requestIdStr,
		// 	SourceAccountNumber:      v.AccountNumber,
		// 	DestinationAccountNumber: v.AccountNumber,
		// 	Amount:                   vAmountFloat,
		// 	Charge:                   0.0,
		// 	TransactionType:          "LOAN_REPAYMENT",
		// 	ServiceCode:              "LOAN_REPAYMENT",
		// 	TransactionReference:     "SYSTEM",
		// 	StatusCode:               "PENDING",
		// 	ExtraDetails1:            amountString,
		// 	ExtraDetails2:            strconv.FormatFloat(vAmountFloat, 'f', -1, 64),
		// 	ExtraDetails3:            network,
		// 	Reference:                amountString,
		// 	ClientID:                 v.ClientId,
		// 	PhoneNumber:              v.MobileNumber,
		// 	TransactionPackage:       amountString,
		// 	ExternalReferenceNumber:  "",
		// 	CreatedBy:                strconv.FormatInt(userData.UserId, 10),
		// }

		extraData := requests.ExtraData{
			ExtraData1: strings.Trim(v.CustomerName, " "),
			ExtraData2: strings.Trim(v.CustomerNumber, " "),
			ExtraData3: paymentMethod,
		}
		transactionLog := requests.UserTransactionRequestDTO{
			SourceChannel:            sourceSystem,
			RequestId:                requestIdStr,
			SourceAccountNumber:      userData.PhoneNumber,
			DestinationAccountNumber: v.AccountNumber,
			Amount:                   vAmountFloat,
			ServiceCode:              "LOAN_REPAYMENT",
			ClientReference:          "",
			ExtraData:                extraData,
			Package:                  network,
			PhoneNumber:              userData.PhoneNumber,
			CreatedBy:                strconv.FormatInt(userData.UserId, 10),
			Status:                   "PENDING",
		}

		if txn := apifunctions.LogUserTransaction(&c.Controller, transactionLog); txn.StatusCode != 200 {
			logs.Error("Error logging transaction: ", txn.StatusDesc)
			status = false
			statusMessage = "Error logging transaction: " + txn.StatusDesc

		} else {

			logs.Info("Payment method to be sent is ", v.PaymentMethod)

			transferRequest := requests.TransferApiRequest{
				RequestId:              requestIdStr,
				Amount:                 vAmountFloat,
				Charge:                 0.0,
				Commission:             0.0,
				TotalDebitAmount:       vAmountFloat + 0.0,
				SenderAccountNumber:    v.AccountNumber,
				RecipientAccountNumber: v.AccountNumber,
				TransferCode:           v.PaymentMethod,
				Description:            "Loan repayment for transaction " + requestIdStr,
				RecipientName:          network,
				Status:                 "PENDING",
				ServiceCode:            "LOAN_REPAYMENT",
				CreatedBy:              strconv.FormatInt(userData.UserId, 10),
			}

			sendCommission := false
			if v.PaymentMethod == "MOBILEMONEY" {
				sendCommission = true
			}

			if _, err := helpers.LogTransferTransaction(&c.Controller, transferRequest, sendCommission); err != nil {
				logs.Error("Error logging transfer transaction: ", err)
				status = false
				statusMessage = "Error logging transfer transaction: " + err.Error()
			} else {

				if v.PaymentMethod == "MOBILEMONEY" {
					logs.Info("Payment method is mobile money. About to request ", vAmountFloat, " from customer phone number ", userData.PhoneNumber)
					req2 := requests.PaymentRequestApiRequestDTO{
						ClientId:            v.ClientId,
						Amount:              vAmountFloat,
						PaymentMethod:       v.PaymentMethod,
						Service:             "LOAN_REPAYMENT",
						SenderAccount:       v.AccountNumber,
						ReceiverAccount:     v.AccountNumber,
						Network:             network,
						ServiceNetwork:      v.ClientId,
						ServicePackage:      strconv.FormatFloat(vAmountFloat, 'f', -1, 64),
						MobileNumber:        userData.PhoneNumber,
						TransactionId:       txn.Result.TransactionId,
						CallbackServiceCode: "USER_PAYMENT",
					}
					//

					logs.Info("Amount to debit is ", v.Amount)

					resp, err := helpers.RequestPaymentMain(&c.Controller, req2)

					logs.Info("Response from Loan repayment API: ", resp)

					if err != nil {
						statusMessage = err.Error()
					} else {
						if resp.Success {
							// accountCheckResp := helpers.LogAccountActivity(&c.Controller, accountNumber, "Deposit", req.Amount, req.ClientId, "credit")

							status = true
							statusMessage = "Loan repayment successful"
							responseText, err := json.Marshal(resp)

							if err != nil {
								logs.Error("Error marshalling response result: ", err)
								responseText = []byte("[]")
							}
							ap.RequestResponse = string(responseText)
						} else {
							statusMessage = "Loan repayment failed: " + resp.StatusMessage
						}
					}
				} else {
					status = true
					statusMessage = "Loan repayment successful"
				}

				txnData.Amount = txn.Result.Amount
				txnData.Currency = txn.Result.TransactingCurrency
				txnData.TransactionReference = txn.Result.TransactionId

				ap.DateModified = time.Now()
				ap.ResponseDate = time.Now()
				if err := models.UpdateApi_requestsById(&ap); err != nil {
					logs.Error("Error updating API request with response: ", err)
				} else {
					logs.Info("API request updated with response successfully: ", v)
				}

				req := requests.LoanRepaymentApiRequest{
					AccountNumber: v.AccountNumber,
					Amount:        v.Amount,
					MobileNumber:  userData.PhoneNumber,
					LoanId:        v.LoanId,
					ClientId:      v.ClientId,
					PaymentMode:   paymentMethod,
				}

				logs.Info("Loan repayment request: ", func() string { b, _ := json.Marshal(req); return string(b) }())

				resp := apifunctions.LoanRepayment(&c.Controller, req)

				logs.Debug("Response is ", resp)

				if resp.StatusCode == true {
					logs.Info("Successfully fetched account statement")
					status = true
					statusMessage = "Successfully paid account loan"

				} else {
					logs.Error("Error fetching account statement")
					statusMessage = resp.StatusDesc
				}

			}
		}
	} else {
		logs.Error("Error logging API request: ", err)
		statusMessage = "Error logging API request: " + err.Error()
	}

	response = responses.RepayLoanResponse{
		Success:       status,
		StatusMessage: statusMessage,
		Result:        txnData,
	}

	c.Data["json"] = response

	c.ServeJSON()
}

// GetFloat ...
// @Title Get Float
// @Description Get float for agent
// @Param	body		body 	requests.GetFloatRequest	true		"body for crediting of account"
// @Param	clientId		header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /get-agent-transactions [post]
func (c *Agent_api_requestsController) GetAgentTransactions() {
	clientId := c.Ctx.Input.Header("clientId")
	logs.Debug("Client id is ", clientId)

	user := c.Ctx.Input.GetData("user")

	logs.Info("User details: %s", user)
	userData, ok := user.(*responses.UsersOri)
	if !ok {
		logs.Error("Error asserting user data")
		c.Data["json"] = "Invalid user data"
		c.ServeJSON()
		return
	}

	var req requests.AgentTransactionsRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	status := false
	statusMessage := "Error retrieving agent float"
	result := []responses.TxnResp{}

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
		PhoneNumber:  req.AgentCode,
		RequestType:  "Agent Transactions",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
		CreatedBy:    int(userData.UserId),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)

		var allowedDateList [6]string = [6]string{"2006-01-02", "2006/01/02", "2006-01-02 15:04:05.000", "2006/01/02 15:04:05.000", "2006-01-02T15:04:05.000Z", "2006-01-02 15:04:05.000000 -0700 MST"}

		proceed := false
		fromDate := time.Time{}
		for _, date_ := range allowedDateList {
			logs.Debug("About to convert ", req.FromDate)
			logs.Debug("About to convert ", c.Ctx.Input.Query("Dob"))
			// Convert dob string to date
			tdobm, error := time.Parse(date_, req.FromDate)

			if error != nil {
				logs.Error("Error parsing date", error)
				statusMessage = "Invalid date format for FromDate. Please use one of the following formats: YYYY-MM-DD, YYYY/MM/DD, YYYY-MM-DD HH:MM:SS.sss, YYYY/MM/DD HH:MM:SS.sss, YYYY-MM-DDTHH:MM:SS.sssZ, or YYYY-MM-DD HH:MM:SS.ssssss -0700 MST"
				proceed = false
			} else {
				logs.Info("Date converted to time successfully", tdobm)
				fromDate = tdobm
				proceed = true

				break
			}
		}

		proceed = false
		toDate := time.Time{}
		for _, date_ := range allowedDateList {
			logs.Debug("About to convert ", req.ToDate)
			logs.Debug("About to convert ", c.Ctx.Input.Query("Dob"))
			// Convert dob string to date
			tdobm, error := time.Parse(date_, req.ToDate)

			if error != nil {
				logs.Error("Error parsing date", error)
				statusMessage = "Invalid date format for ToDate. Please use one of the following formats: YYYY-MM-DD, YYYY/MM/DD, YYYY-MM-DD HH:MM:SS.sss, YYYY/MM/DD HH:MM:SS.sss, YYYY-MM-DDTHH:MM:SS.sssZ, or YYYY-MM-DD HH:MM:SS.ssssss -0700 MST"
				proceed = false
			} else {
				logs.Info("Date converted to time successfully", tdobm)
				toDate = tdobm
				proceed = true

				break
			}
		}

		logs.Debug("From date is ", fromDate)
		logs.Debug("To date is ", toDate)
		logs.Debug("Proceed? ", proceed)

		if proceed {
			fromDateStr := fromDate.Format("2006-01-02 15:04:05")
			toDateStr := toDate.Format("2006-01-02 15:04:05")
			logs.Debug("From date is ", fromDateStr)
			logs.Debug("To date is ", toDateStr)

			getUser := apifunctions.GetUserDetailsWithCode(&c.Controller, req.AgentCode)

			if getUser.StatusCode != 200 {
				logs.Error("Error fetching user details for agent code ", req.AgentCode)
				statusMessage = "Invalid agent code"
				response := responses.AgentTransactionsResponse{
					Success:    status,
					StatusDesc: statusMessage,
					Result:     &result,
				}

				c.Data["json"] = response
				c.ServeJSON()
				return
			}

			allTrxns := []responses.TxnResp{}

			for _, user := range *getUser.Users {
				logs.Debug("User fetched for agent code ", req.AgentCode, ": ", user.FullName)

				query := "CreatedBy:" + strconv.Itoa(int(user.UserId)) + ",DateCreated__gte:" + fromDateStr + ",DateCreated__lte:" + toDateStr
				resp := apifunctions.GetAgentTransactions(&c.Controller, query)

				logs.Debug("Response is ", resp)

				if resp.StatusCode == 200 {
					logs.Info("Successfully fetched agent transactions")
					status = true
					statusMessage = "Successfully fetched agent transactions"
					if resp.Result != nil {
						for _, txn := range *resp.Result {
							txnResp := responses.TxnResp{
								TransactionRefNumber:    txn.TransactionId,
								Service:                 txn.Service,
								BillerCode:              txn.Service,
								CustomerName:            txn.ExtraDetails1,
								CustomerNumber:          txn.ExtraDetails2,
								Amount:                  txn.Amount,
								TransactingCurrency:     txn.TransactingCurrency,
								SourceChannel:           txn.SourceChannel,
								SourceAccount:           txn.Source,
								DestinationAccount:      txn.Destination,
								Package:                 txn.Package,
								Charge:                  txn.Charge,
								ExternalReferenceNumber: txn.ExternalReferenceNumber,
								Status:                  txn.Status,
								CorpId:                  userData.UserDetails.Branch.Branch,
								ExtraDetails1:           txn.ExtraDetails1,
								ExtraDetails2:           txn.ExtraDetails2,
								ExtraDetails3:           txn.ExtraDetails3,
								TransactionDate:         txn.DateCreated,
								OfficerName:             user.FullName,
								OfficerNumber:           user.PhoneNumber,
								Active:                  1,
							}
							allTrxns = append(allTrxns, txnResp)
						}
					} else {
						logs.Info("No transactions found for the agent in the specified date range")
						statusMessage = "No transactions found for the agent in the specified date range"
					}
				} else {
					logs.Error("Error fetching agent transactions")
					statusMessage = resp.StatusDesc
				}
			}
			result = allTrxns
		}
	} else {
		logs.Error("Error logging API request: ", err)
		statusMessage = "Error logging API request: " + err.Error()
	}

	response := responses.AgentTransactionsResponse{
		Success:    status,
		StatusDesc: statusMessage,
		Result:     &result,
	}

	c.Data["json"] = response

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
func (c *Agent_api_requestsController) AccountBalance() {
	// Extract headers
	// phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")

	user := c.Ctx.Input.GetData("user")

	logs.Info("User details: %s", user)
	userData, ok := user.(*responses.UsersOri)
	if !ok {
		logs.Error("Error asserting user data")
		c.Data["json"] = "Invalid user data"
		c.ServeJSON()
		return
	}

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
		CreatedBy:    int(userData.UserId),
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
		} else {
			logs.Info("Formatted request for account balance: ", accountBalanceRequest)
			resp := apifunctions.GetAccountBalance(&c.Controller, accountBalanceRequest)
			logs.Info("Response from account balance API: ", resp)

			var response responses.AccountDetailsResponse = responses.AccountDetailsResponse{
				StatusCode:    false,
				StatusMessage: "Something went wrong",
				Result:        nil,
			}

			if resp.StatusCode != 200 {
				response = responses.AccountDetailsResponse{
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

				accBal := responses.AccountDetailsDataResp{
					CustomerName:     accountResp.Result.CustomerName,
					AccountAlias:     accountResp.Result.AccountAlias,
					AccountNumber:    req.AccountNumber,
					AccountStatus:    resp.Result.AccountStatus,
					AvailableBalance: *resp.Result.AvailableBalance,
					ClearBalance:     *resp.Result.ClearBalance,
					LoanBalance:      *resp.Result.LoanBalance,
					SharesBalance:    *resp.Result.SharesBalance,
				}
				response = responses.AccountDetailsResponse{
					StatusCode:    true,
					StatusMessage: "Account balance fetched succeefully",
					Result:        &accBal,
				}
			}

			c.Ctx.Output.SetStatus(200)
			c.Data["json"] = response
		}

	} else {
		var response responses.AccountDetailsResponse = responses.AccountDetailsResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong:: " + err.Error(),
			Result:        nil,
		}

		c.Data["json"] = response
	}
	c.ServeJSON()
}

// AccountDetails ...
// @Title Account Details
// @Description Account Details
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	body		body 	requests.NumberExistsApiRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /account-details [post]
func (c *Agent_api_requestsController) ListAccountDetails() {
	// Extract headers
	// phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")

	user := c.Ctx.Input.GetData("user")

	logs.Info("User details: %s", user)
	userData, ok := user.(*responses.UsersOri)
	if !ok {
		logs.Error("Error asserting user data")
		c.Data["json"] = "Invalid user data"
		c.ServeJSON()
		return
	}

	var req requests.AccountIdApiRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	logs.Info("Account Details called with AccountId: %s, SourceSystem: %s", req.AccountId, sourceSystem)
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
		RequestType:  "Account Details",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
		CreatedBy:    int(userData.UserId),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)
		accountBalanceRequest := requests.AccountIdApiRequest{
			AccountId: req.AccountId,
			ClientId:  req.ClientId,
		}

		logs.Info("Formatted request for account details: ", accountBalanceRequest)
		resp := apifunctions.ListAccounts(&c.Controller, accountBalanceRequest)
		logs.Info("Response from account details API: ", resp)

		var response responses.AccountDetailsResponse = responses.AccountDetailsResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong",
			Result:        nil,
		}

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

		if resp.Data.StatusCode != 200 {
			response = responses.AccountDetailsResponse{
				StatusCode:    false,
				StatusMessage: resp.Data.StatusDesc,
				Result:        nil,
			}
		} else if len(resp.Data.Result) > 0 {

			customerName := ""
			accountNumber := ""
			accountAlias := ""
			product := ""

			for i, acc := range resp.Data.Result {
				if i == 0 {
					customerName = acc.AccountName
					accountNumber = acc.AccountNumber
					accountAlias = acc.AccountNumber
					product = acc.Product
				} else {
					customerName += ", " + acc.AccountName
					accountNumber += ", " + acc.AccountNumber
					accountAlias += ", " + acc.AccountNumber
					product += ", " + acc.Product
				}
			}

			accBal := responses.AccountDetailsDataResp{
				CustomerName:     customerName,
				AccountAlias:     accountAlias,
				AccountNumber:    accountNumber,
				AccountStatus:    req.AccountId,
				AvailableBalance: 0.00,
				ClearBalance:     0.00,
				LoanBalance:      0.00,
				SharesBalance:    0.00,
			}
			response = responses.AccountDetailsResponse{
				StatusCode:    true,
				StatusMessage: "Account details fetched successfully",
				Result:        &accBal,
			}
		} else {
			response = responses.AccountDetailsResponse{
				StatusCode:    false,
				StatusMessage: "No account found for the provided account ID",
				Result:        nil,
			}
		}

		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = response

	} else {
		var response responses.AccountDetailsResponse = responses.AccountDetailsResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong:: " + err.Error(),
			Result:        nil,
		}

		c.Data["json"] = response
	}
	c.ServeJSON()
}

// GetBilTransactionWithTransactionRef ...
// @Title Get Biller Transaction By Reference
// @Description Get a biller transaction using transaction reference
// @Param	Authorization		header 	string true		"header for User"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	body		body 	object true		"body containing transactionReference"
// @Success 200 {object} interface{}
// @Failure 400 invalid request body
// @router /get-transaction-by-reference [post]
func (c *Agent_api_requestsController) GetBilTransactionWithTransactionRef() {
	var req struct {
		TransactionReference string `json:"transactionReference"`
	}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"status":  false,
			"message": "Invalid request body",
		}
		c.ServeJSON()
		return
	}

	responseStatus := false
	responseMessage := "Something went wrong"
	result := responses.TxnResp{}

	reference := strings.TrimSpace(req.TransactionReference)

	if reference == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{
			"status":  false,
			"message": "transactionReference is required",
		}
		c.ServeJSON()
		return
	}

	resp := apifunctions.GetUserTransactionWithTransactionRef(&c.Controller, reference)

	if resp.StatusCode != 200 {
		responseStatus = false
		responseMessage = resp.StatusDesc
	} else {
		responseStatus = true
		responseMessage = "Transaction fetched successfully"
		if resp.Result != nil {
			result = responses.TxnResp{
				TransactionRefNumber:    resp.Result.TransactionId,
				Amount:                  resp.Result.Amount,
				Charge:                  resp.Result.Charge,
				CustomerName:            resp.Result.ExtraDetails1,
				CustomerNumber:          resp.Result.ExtraDetails2,
				Status:                  resp.Result.Status,
				Service:                 resp.Result.Service,
				TransactingCurrency:     resp.Result.TransactingCurrency,
				SourceChannel:           resp.Result.SourceChannel,
				SourceAccount:           resp.Result.Source,
				DestinationAccount:      resp.Result.Destination,
				Package:                 resp.Result.Package,
				ExternalReferenceNumber: resp.Result.ExternalReferenceNumber,
				TransactionDate:         resp.Result.DateCreated,
				OfficerName:             resp.Result.CreatedBy,
				OfficerNumber:           resp.Result.CreatedBy,
			}
		}

		logs.Info("Transaction reference number: ", result.TransactionRefNumber, " Status: ", result.Status, resp.Result.Status)
		logs.Info("Customer name ", result.CustomerName, " customer number: ", result.CustomerNumber)
	}

	c.Ctx.Output.SetStatus(200)

	response := responses.Bil_TransactionResponse{
		Success:       responseStatus,
		StatusMessage: responseMessage,
		Result:        &result,
	}
	c.Data["json"] = response
	c.ServeJSON()
}

// @Description create Agent_api_requests
// @Param	body		body 	models.Agent_api_requests	true		"body for Agent_api_requests content"
// @Success 201 {object} models.Agent_api_requests
// @Failure 403 body is empty
// @router / [post]
func (c *Agent_api_requestsController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Agent_api_requests by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Agent_api_requests
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Agent_api_requestsController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Agent_api_requests
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Agent_api_requests
// @Failure 403
// @router / [get]
func (c *Agent_api_requestsController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Agent_api_requests
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Agent_api_requests	true		"body for Agent_api_requests content"
// @Success 200 {object} models.Agent_api_requests
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Agent_api_requestsController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Agent_api_requests
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Agent_api_requestsController) Delete() {

}
