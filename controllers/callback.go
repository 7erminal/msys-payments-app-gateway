package controllers

import (
	"encoding/json"
	apifunctions "msys_payment_app_gateway/controllers/api_functions"
	"msys_payment_app_gateway/controllers/services"
	"msys_payment_app_gateway/models"
	"msys_payment_app_gateway/structs/requests"
	"msys_payment_app_gateway/structs/responses"
	"time"

	beego "github.com/beego/beego/v2/server/web"

	"github.com/beego/beego/v2/core/logs"
)

// CallbackController operations for Callback
type CallbackController struct {
	beego.Controller
}

// URLMapping ...
func (c *CallbackController) URLMapping() {
	c.Mapping("Callback", c.Callback)
	c.Mapping("CheckTransactionStatus", c.CheckTransactionStatus)
	c.Mapping("RequestMoneyCallback", c.RequestMoneyCallback)
}

// Callback ...
// @Title Callback
// @Description create Callback
// @Param	body		body 	requests.CallbackAPIRequest	true		"body for Callback content"
// @Success 201 {object} models.Callback
// @Failure 403 body is empty
// @router / [post]
func (c *CallbackController) Callback() {
	var v requests.CallbackAPIRequest

	logs.Info("Received callback request: ", string(c.Ctx.Input.RequestBody))
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid request body"}
		c.Ctx.Output.SetStatus(400)
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
	var j models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		RequestType:  "Callback",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&j); err == nil {
		callbackReq := requests.CallbackFormulateRequest(v)

		logs.Info("Sending callback request: ", callbackReq)
		resp := apifunctions.Callback(&c.Controller, callbackReq)
		logs.Info("Callback response: ", resp)

		var response responses.CallbackAPIResponse = responses.CallbackAPIResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong",
			Result:        resp.Result,
		}

		if !resp.StatusCode {
			logs.Error("Callback failed with response: ", resp)
			response = responses.CallbackAPIResponse{
				StatusCode:    false,
				StatusMessage: "Something went wrong",
				Result:        resp.Result,
			}

			c.Data["json"] = response
		} else {
			responseText, err := json.Marshal(resp.Result)
			if err != nil {
				logs.Error("Error marshalling response result: ", err)
				responseText = []byte("[]")
			}
			j.RequestResponse = string(responseText)
			j.DateModified = time.Now()
			j.ResponseDate = time.Now()
			if err := models.UpdateApi_requestsById(&j); err != nil {
				logs.Error("Error updating API request with response: ", err)
			} else {
				logs.Info("API request updated with response successfully: ", v)
			}

			response = responses.CallbackAPIResponse{
				StatusCode:    true,
				StatusMessage: "Callback processed successfully",
				Result:        resp.Result,
			}

			c.Data["json"] = response
		}
	}

	c.ServeJSON()
}

// Callback ...
// @Title Callback
// @Description create Callback
// @Param	body		body 	requests.PaymentCallbackFormulateRequest	true		"body for Callback content"
// @Success 201 {object} models.Callback
// @Failure 403 body is empty
// @router /payment-callback [post]
func (c *CallbackController) RequestMoneyCallback() {
	var v requests.PaymentCallbackFormulateRequest

	logs.Info("Received callback request: ", string(c.Ctx.Input.RequestBody))
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid request body"}
		c.Ctx.Output.SetStatus(400)
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
	var j models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		RequestType:  "Callback",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&j); err == nil {
		var response responses.CallbackResponse = responses.CallbackResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong",
			Result:        nil,
		}

		callbackReq := requests.PaymentCallbackData{}

		if v.ResponseCode == "0000" {
			logs.Info("Successful payment callback received for transaction: ", v.Data.ClientReference)
			callbackReq = requests.PaymentCallbackData{
				AmountCharged:         v.Data.Amount,
				TransactionId:         v.Data.TransactionId,
				ClientReference:       v.Data.ClientReference,
				Description:           v.Data.Description,
				ExternalTransactionId: v.Data.ExternalTransactionId,
				Amount:                v.Data.Amount,
				Charges:               v.Data.Charges,
				AmountAfterCharges:    v.Data.AmountAfterCharges,
				PaymentDate:           v.Data.PaymentDate,
				OrderId:               v.Data.OrderId,
				Status:                "SUCCESS",
			}
		} else {
			logs.Info("Failed payment callback received for transaction: ", v.Data.ClientReference)
			callbackReq = requests.PaymentCallbackData{
				AmountCharged:         v.Data.Amount,
				TransactionId:         v.Data.TransactionId,
				ClientReference:       v.Data.ClientReference,
				Description:           v.Data.Description,
				ExternalTransactionId: v.Data.ExternalTransactionId,
				Amount:                v.Data.Amount,
				Charges:               v.Data.Charges,
				AmountAfterCharges:    v.Data.AmountAfterCharges,
				PaymentDate:           v.Data.PaymentDate,
				OrderId:               v.Data.OrderId,
				Status:                "FAILED",
			}
		}

		logs.Info("Sending callback request: ", callbackReq)
		resp := apifunctions.ReceivePaymentCallback(&c.Controller, callbackReq)
		logs.Info("Callback response: ", resp)

		if !resp.StatusCode {
			logs.Error("Callback failed with response: ", resp)
			response = responses.CallbackResponse{
				StatusCode:    false,
				StatusMessage: "Something went wrong",
				Result:        nil,
			}

			c.Data["json"] = response
		} else {
			responseText, err := json.Marshal(resp.Result)
			if err != nil {
				logs.Error("Error marshalling response result: ", err)
				responseText = []byte("[]")
			}
			j.RequestResponse = string(responseText)
			j.DateModified = time.Now()
			j.ResponseDate = time.Now()
			if err := models.UpdateApi_requestsById(&j); err != nil {
				logs.Error("Error updating API request with response: ", err)
			} else {
				logs.Info("API request updated with response successfully: ", v)
			}

			if resp.Result.Service == "AIRTIME" {
				network := ""
				// Process airtime request or insert in a queue for the airtime service to pick up
				for _, phist := range *resp.Result.PaymentHistory {
					logs.Info("Processing airtime for payment history: ", phist)
					// Get network

					if phist.Service == "AIRTIME" {
						network = phist.Reference
					}
				}
				airtimeReq := requests.BuyAirtimeFormulatedRequest{
					Amount:       resp.Result.PaymentAmount,
					PhoneNumber:  resp.Result.ReceiverAccount,
					Network:      network,
					Destination:  resp.Result.ReceiverAccount,
					SourceSystem: "MSYS_PAYMENT_APP_GATEWAY",
					RequestId:    j.Id,
				}

				airtimeresp := services.BuyAirtime(&c.Controller, airtimeReq)
				logs.Info("Response from Buy Airtime API: ", airtimeresp)

				if !airtimeresp.StatusCode {
					response = responses.CallbackResponse{
						StatusCode:    false,
						StatusMessage: resp.StatusMessage,
						Result:        nil,
					}
				} else {

					response = responses.CallbackResponse{
						StatusCode:    true,
						StatusMessage: "Airtime purchase successful",
						Result:        resp.Result,
					}
				}
			}

			if resp.Result.Service == "DATA_BUNDLE" {
				network := ""
				// Process airtime request or insert in a queue for the airtime service to pick up
				for _, phist := range *resp.Result.PaymentHistory {
					logs.Info("Processing airtime for payment history: ", phist)
					// Get network

					if phist.Service == "DATA_BUNDLE" {
						network = phist.Reference
					}
				}
				dataBundleReq := requests.BuyDataBundleFormulatedRequest{
					Amount:       resp.Result.PaymentAmount,
					PhoneNumber:  resp.Result.ReceiverAccount,
					Network:      network,
					Destination:  resp.Result.ReceiverAccount,
					BundleId:     resp.Result.ReferenceNumber,
					SourceSystem: "MSYS_PAYMENT_APP_GATEWAY",
					RequestId:    j.Id,
				}

				airtimeresp := services.BuyDataBundle(&c.Controller, dataBundleReq)
				logs.Info("Response from Buy data API: ", airtimeresp)

				if !airtimeresp.StatusCode {
					response = responses.CallbackResponse{
						StatusCode:    false,
						StatusMessage: resp.StatusMessage,
						Result:        nil,
					}
				} else {

					response = responses.CallbackResponse{
						StatusCode:    true,
						StatusMessage: "Airtime purchase successful",
						Result:        resp.Result,
					}
				}
			}

			c.Data["json"] = response
		}

	}

	c.ServeJSON()
}

// CheckTransactionStatus ...
// @Title Check transaction status
// @Description Check transaction status
// @Param	body		body 	requests.TransactionStatusRequest	true		"body for Transaction status check"
// @Success 201 {object} models.Callback
// @Failure 403 body is empty
// @router /check-transaction-status [post]
func (c *CallbackController) CheckTransactionStatus() {
	var v requests.TransactionStatusRequest

	logs.Info("Transaction status request: ", string(c.Ctx.Input.RequestBody))
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid request body"}
		c.Ctx.Output.SetStatus(400)
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
	var j models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		RequestType:  "Transaction Status",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&j); err == nil {
		transactionStatusReq := requests.TransactionStatusApiRequest(v)

		logs.Info("Sending transaction status request: ", transactionStatusReq)
		resp := apifunctions.CheckTransactionStatus(&c.Controller, transactionStatusReq)
		logs.Info("Transaction status response: ", resp)

		var response responses.TransactionStatusResponse = responses.TransactionStatusResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong",
			Result:        resp.Result,
		}

		if !resp.StatusCode {
			logs.Error("Transaction status failed with response: ", resp)
			response = responses.TransactionStatusResponse{
				StatusCode:    false,
				StatusMessage: "Something went wrong",
				Result:        resp.Result,
			}

			c.Data["json"] = response
		} else {
			responseText, err := json.Marshal(resp.Result)
			if err != nil {
				logs.Error("Error marshalling response result: ", err)
				responseText = []byte("[]")
			}
			j.RequestResponse = string(responseText)
			j.DateModified = time.Now()
			j.ResponseDate = time.Now()
			if err := models.UpdateApi_requestsById(&j); err != nil {
				logs.Error("Error updating API request with response: ", err)
			} else {
				logs.Info("API request updated with response successfully: ", v)
			}

			response = responses.TransactionStatusResponse{
				StatusCode:    true,
				StatusMessage: "Transaction status processed successfully",
				Result:        resp.Result,
			}

			c.Data["json"] = response
		}
	}

	c.ServeJSON()
}
