package controllers

import (
	"encoding/json"
	"fmt"
	apifunctions "msys_payment_app_gateway/controllers/api_functions"
	"msys_payment_app_gateway/controllers/helpers"
	"msys_payment_app_gateway/models"
	"msys_payment_app_gateway/structs/requests"
	"msys_payment_app_gateway/structs/responses"
	"msys_payment_app_gateway/utils"
	utilManager "msys_payment_app_gateway/utils"
	"path/filepath"
	"runtime"
	"strconv"
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
	c.Mapping("GetCorporatives", c.GetCorporatives)
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
func (c *Agent_api_requestsController) GetCorporatives() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	// sourceSystem := c.Ctx.Input.Header("SourceSystem")
	// network := c.Ctx.Input.Header("Network")

	var req requests.GetCorporativesRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	// destinationPhoneNumber := req.Destination

	_, file, line, ok := runtime.Caller(0)
	if ok {
		file = utils.GetFileName(file)
	} else {
		file = "unknown"
		line = 0
	}
	utilManager.Logger(filepath.Base(file), line, req.RequestId, "INFO", fmt.Sprintf("GetCorporatives request: %s", func() string { b, _ := json.Marshal(req); return string(b) }()))

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
		_, file, line, ok := runtime.Caller(0)
		if ok {
			file = utils.GetFileName(file)
		} else {
			file = "unknown"
			line = 0
		}
		utilManager.Logger(filepath.Base(file), line, req.RequestId, "INFO", fmt.Sprintf("Response from Get corporatives API: %s", func() string { b, _ := json.Marshal(resp); return string(b) }()))

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

	var req requests.DepositAPIRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	destinationPhoneNumber := req.Destination
	clientId := req.ClientId

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

	var response responses.PaymentRequestResponse = responses.PaymentRequestResponse{
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
			TransactionReference:     "SYSTEM",
			StatusCode:               "PENDING",
			ExtraDetails1:            amountString,
			ExtraDetails2:            strconv.FormatFloat(req.Amount, 'f', -1, 64),
			ExtraDetails3:            network,
			Reference:                amountString,
			ClientID:                 req.ClientId,
			PhoneNumber:              phoneNumber,
			TransactionPackage:       amountString,
			ExternalReferenceNumber:  "",
			CreatedBy:                req.CreatedBy,
		}

		if txn, err := helpers.LogTransaction(&c.Controller, transactionLog); err != nil {
			logs.Error("Error logging transaction: ", err)
			isSuccess = false
			message = "Error logging transaction: " + err.Error()

		} else {

			req := requests.PaymentRequestApiRequestDTO{
				ClientId:        clientId,
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

			logs.Info("Amount to debit is ", req.Amount)

			resp, err := helpers.RequestPaymentMain(&c.Controller, req)

			logs.Info("Response from Deposit API: ", resp)

			if err != nil {
				message = err.Error()
			} else {
				if resp.Success {
					// accountCheckResp := helpers.LogAccountActivity(&c.Controller, accountNumber, "Deposit", req.Amount, req.ClientId, "credit")

					isSuccess = true
					message = "Deposit successful"
					responseText, err := json.Marshal(resp)

					// if !accountCheckResp.StatusCode {
					// 	logs.Error("Error logging account activity for deposit: ", accountCheckResp.StatusMessage)
					// 	message = "Deposit failed: " + accountCheckResp.StatusMessage
					// }

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

// Get Account Loans ...
// @Title GetAccountLoans
// @Description Get account loans
// @Param	body		body 	requests.AccountLoansRequest	true		"body for crediting of account"
// @Param	clientId		header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /v2/list-account-loans [post]
func (c *Agent_api_requestsController) ListAccountLoans() {
	clientId := c.Ctx.Input.Header("clientId")
	logs.Debug("Client id is ", clientId)

	var v requests.AccountLoansRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	status := false
	statusMessage := "Error retrieving account loans"

	// logs.Debug("Request::: ", c.Ctx.Input.RequestBody)
	reqBody, err := json.Marshal(v)
	if err != nil {
		logs.Error("Error marshalling request body: %v", err)
	} else {
		logs.Debug("Get account loans request: %s", string(reqBody))
	}

	resp := apifunctions.ListAccountLoans(&c.Controller, clientId, v.AccountNumber)

	logs.Debug("Response is ", resp)

	var response responses.ListLoansResponse

	var result []responses.LoanData

	if resp.StatusCode == 200 {
		logs.Info("Successfully fetched account statement")
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
		logs.Error("Error fetching account statement")
		statusMessage = resp.StatusDesc
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
// @router /v2/loan-repayment [post]
func (c *Agent_api_requestsController) LoanRepayment() {
	clientId := c.Ctx.Input.Header("clientId")
	logs.Debug("Client id is ", clientId)

	var v requests.LoanRepaymentRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	status := false
	statusMessage := "Error retrieving account loans"
	result := "Repayment failed"

	// logs.Debug("Request::: ", c.Ctx.Input.RequestBody)
	reqBody, err := json.Marshal(v)
	if err != nil {
		logs.Error("Error marshalling request body: %v", err)
	} else {
		logs.Debug("Get account loans request: %s", string(reqBody))
	}

	req := requests.LoanRepaymentApiRequest{
		AccountNumber: v.AccountNumber,
		Amount:        v.Amount,
		MobileNumber:  v.MobileNumber,
		LoanId:        v.LoanId,
		ClientId:      v.ClientId,
	}

	resp := apifunctions.LoanRepayment(&c.Controller, req)

	logs.Debug("Response is ", resp)

	var response responses.RepayLoanResponse

	if resp.StatusCode == 200 {
		logs.Info("Successfully fetched account statement")
		status = true
		statusMessage = "Successfully fetched account loans"
		result = resp.Result
	} else {
		logs.Error("Error fetching account statement")
		statusMessage = resp.StatusDesc
	}

	response = responses.RepayLoanResponse{
		StatusCode: status,
		StatusDesc: statusMessage,
		Result:     result,
	}

	c.Data["json"] = response

	c.ServeJSON()
}

// Post ...
// @Title Create
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
