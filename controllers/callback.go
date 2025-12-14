package controllers

import (
	"encoding/json"
	apifunctions "msys_payment_app_gateway/controllers/api_functions"
	"msys_payment_app_gateway/controllers/helpers"
	"msys_payment_app_gateway/controllers/services"
	"msys_payment_app_gateway/models"
	"msys_payment_app_gateway/structs/requests"
	"msys_payment_app_gateway/structs/responses"
	"strings"
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

		// AmountCharged         float64
		// TransactionId         string
		// ClientReference       string
		// Description           string
		// ExternalTransactionId string
		// Amount                float64
		// Charges               float64
		// Status                string

		status := "FAILED"

		if v.ResponseCode == "0000" {
			status = "SUCCESS"
		} else if v.ResponseCode == "0001" {
			status = "PENDING"
		}

		callbackReq := requests.CallbackRequest{
			AmountCharged:         v.Data.AmountDebited,
			TransactionId:         v.Data.ClientReference,
			ClientReference:       v.Data.TransactionId,
			Description:           v.Data.Description,
			ExternalTransactionId: v.Data.ExternalTransactionId,
			Amount:                v.Data.Amount,
			Charges:               v.Data.Charges,
			Status:                status,
			Commission:            v.Data.Meta.Commission,
			ClientResponseCode:    v.ResponseCode,
			ClientResponseMessage: v.Data.Description,
		}

		logs.Info("Sending callback request: ", callbackReq)
		resp := apifunctions.Callback(&c.Controller, callbackReq)
		logs.Info("Callback response: ", resp)

		var response responses.CallbackAPIResponse = responses.CallbackAPIResponse{
			StatusCode:    401,
			StatusMessage: "Something went wrong",
			Result:        resp.Result,
		}

		if resp.StatusCode != 200 {
			logs.Error("Callback failed with response: ", resp)
			response = responses.CallbackAPIResponse{
				StatusCode:    400,
				StatusMessage: resp.StatusMessage,
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
				StatusCode:    200,
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

	responseStatus := false
	responseMessage := "Invalid request"
	result := responses.PaymentResponse{}

	var j models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		RequestType:  "Callback",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&j); err == nil {

		callbackReq := requests.PaymentCallbackData{}
		proceed := false

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

			proceed = true
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
				Commission:            v.Data.Commission,
				AmountAfterCharges:    v.Data.AmountAfterCharges,
				PaymentDate:           v.Data.PaymentDate,
				OrderId:               v.Data.OrderId,
				Status:                "FAILED",
			}
		}

		transaction := apifunctions.GetBilTransactionWithTransactionRef(&c.Controller, v.Data.ClientReference)

		if transaction.StatusCode == 200 {
			if proceed {
				logs.Info("Sending callback request: ", callbackReq)
				callbackReqJSON, _ := json.Marshal(callbackReq)
				logs.Info("Callback request JSON: ", string(callbackReqJSON))
				resp := apifunctions.ReceivePaymentCallback(&c.Controller, callbackReq)
				// logs.Info("Callback response: ", resp)
				respJSON, _ := json.Marshal(resp)
				logs.Info("Callback response JSON: ", string(respJSON))

				if resp.StatusCode != 200 {
					logs.Error("Callback failed with response: ", resp)
					responseStatus = false
					responseMessage = resp.StatusMessage
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

					logs.Info("Callback received. Service is ", resp.Result.Service)

					if resp.Result.Service == "AIRTIME" {
						// network := ""
						// // Process airtime request or insert in a queue for the airtime service to pick up
						// for _, phist := range *resp.Result.PaymentHistory {
						// 	logs.Info("Processing airtime for payment history: ", phist)
						// 	// Get network

						// 	if phist.Service == "AIRTIME" {
						// 		network = phist.Reference
						// 	}
						// }
						airtimeReq := requests.BuyAirtimeFormulatedRequest{
							Amount:        resp.Result.PaymentAmount,
							PhoneNumber:   resp.Result.SenderAccount,
							Network:       resp.Result.ServiceNetwork,
							Destination:   resp.Result.ReceiverAccount,
							SourceSystem:  "MSYS_PAYMENT_APP_GATEWAY",
							TransactionId: resp.Result.TransactionId,
						}

						airtimeresp := services.BuyAirtime(&c.Controller, airtimeReq)
						logs.Info("Response from Buy Airtime API: ", airtimeresp)

						if !airtimeresp.StatusCode {
							responseStatus = false
							responseMessage = resp.StatusMessage
						} else {
							responseStatus = true
							responseMessage = resp.StatusMessage
						}
					}

					if resp.Result.Service == "DATA_BUNDLE" {
						// network := ""
						// // Process airtime request or insert in a queue for the airtime service to pick up
						// for _, phist := range *resp.Result.PaymentHistory {
						// 	logs.Info("Processing data bundle for payment history: ", phist)
						// 	// Get network

						// 	if phist.Service == "DATA_BUNDLE" {
						// 		network = phist.Reference
						// 	}
						// }
						dataBundleReq := requests.BuyDataBundleFormulatedRequest{
							Amount:        resp.Result.PaymentAmount,
							PhoneNumber:   resp.Result.SenderAccount,
							Network:       resp.Result.ServiceNetwork,
							Destination:   resp.Result.ReceiverAccount,
							BundleId:      resp.Result.ServicePackage,
							SourceSystem:  "MSYS_PAYMENT_APP_GATEWAY",
							TransactionId: resp.Result.TransactionId,
						}

						dataresp := services.BuyDataBundle(&c.Controller, dataBundleReq)
						logs.Info("Response from Buy data API: ", dataresp)

						if !dataresp.StatusCode {
							responseStatus = false
							responseMessage = resp.StatusMessage
						} else {
							responseStatus = true
							responseMessage = resp.StatusMessage
						}
					}

					if resp.Result.Service == "BILL PAYMENT" {
						logs.Info("Service network is ", resp.Result.ServiceNetwork)
						if resp.Result.ServiceNetwork == "DSTV" {
							dstvReq := requests.DSTVPaymentRequest{
								Amount:             resp.Result.PaymentAmount,
								PhoneNumber:        resp.Result.SenderAccount,
								DestinationAccount: resp.Result.ReceiverAccount,
								PackageType:        resp.Result.ServicePackage,
								SourceSystem:       "MSYS_PAYMENT_APP_GATEWAY",
								TransactionId:      resp.Result.TransactionId,
							}

							dstvresp := services.PayDstv(&c.Controller, dstvReq)
							logs.Info("Response from DSTV payment API: ", dstvresp)

							if !dstvresp.StatusCode {
								responseStatus = false
								responseMessage = resp.StatusMessage
							} else {
								responseStatus = true
								responseMessage = resp.StatusMessage
							}
						}

						if resp.Result.ServiceNetwork == "GOTV" {
							gotvReq := requests.GoTvPaymentApiRequest{
								Amount:             resp.Result.PaymentAmount,
								PhoneNumber:        resp.Result.SenderAccount,
								DestinationAccount: resp.Result.ReceiverAccount,
								PackageType:        resp.Result.ServicePackage,
								SourceSystem:       "MSYS_PAYMENT_APP_GATEWAY",
								TransactionId:      resp.Result.TransactionId,
							}

							gotvresp := services.PayGotv(&c.Controller, gotvReq)
							logs.Info("Response from GOTV payment API: ", gotvresp)

							if !gotvresp.StatusCode {
								responseStatus = false
								responseMessage = resp.StatusMessage
							} else {
								responseStatus = true
								responseMessage = resp.StatusMessage
							}
						}

						if resp.Result.ServiceNetwork == "STARTIMES" {
							dstvReq := requests.DSTVPaymentRequest{
								Amount:             resp.Result.PaymentAmount,
								PhoneNumber:        resp.Result.SenderAccount,
								DestinationAccount: resp.Result.ReceiverAccount,
								PackageType:        resp.Result.ServicePackage,
								SourceSystem:       "MSYS_PAYMENT_APP_GATEWAY",
								TransactionId:      resp.Result.TransactionId,
							}

							dstvresp := services.PayDstv(&c.Controller, dstvReq)
							logs.Info("Response from STARTIMES payment API: ", dstvresp)

							if !dstvresp.StatusCode {
								responseStatus = false
								responseMessage = resp.StatusMessage
							} else {
								responseStatus = true
								responseMessage = resp.StatusMessage
							}
						}

						if resp.Result.ServiceNetwork == "WATER" {
							if bilTxn := apifunctions.GetBilTransactionWithTransactionRef(&c.Controller, resp.Result.TransactionId); bilTxn.StatusCode == 200 {
								// transactionString := fmt.Sprintf("%d", bilTxn.Result.TransactionId)
								waterbillReq := requests.GhanaWaterPaymentFuncRequest{
									Amount:             resp.Result.PaymentAmount,
									PhoneNumber:        resp.Result.SenderAccount,
									DestinationAccount: resp.Result.ReceiverAccount,
									PackageType:        bilTxn.Result.ExtraDetails3,
									SourceSystem:       "MSYS_PAYMENT_APP_GATEWAY",
									TransactionId:      resp.Result.TransactionId,
									Name:               bilTxn.Result.ExtraDetails1,
									Email:              bilTxn.Result.ExtraDetails2,
								}

								waterbillresp := services.PayWater(&c.Controller, waterbillReq)
								logs.Info("Response from WATER payment API: ", waterbillresp)

								if !waterbillresp.StatusCode {
									responseStatus = false
									responseMessage = resp.StatusMessage
								} else {
									responseStatus = true
									responseMessage = "Water bill purchase successful"
								}
							} else {
								logs.Error("Error retrieving BIL transaction for ID: ", bilTxn.StatusDesc)
								responseStatus = false
								responseMessage = "Error processing water bill payment: " + bilTxn.StatusDesc
							}
						}

						if resp.Result.ServiceNetwork == "ECG" {
							ecgbillReq := requests.ECGPaymentApiRequest{
								Amount:             resp.Result.PaymentAmount,
								PhoneNumber:        resp.Result.SenderAccount,
								DestinationAccount: resp.Result.ReceiverAccount,
								PackageType:        resp.Result.ServicePackage,
								SourceSystem:       "MSYS_PAYMENT_APP_GATEWAY",
								TransactionId:      resp.Result.TransactionId,
							}

							ecgbillresp := services.PayEcg(&c.Controller, ecgbillReq)
							logs.Info("Response from ECG payment API: ", ecgbillresp)

							if ecgbillresp.StatusCode != "200" {
								responseStatus = false
								responseMessage = resp.StatusMessage
							} else {
								responseStatus = true
								responseMessage = resp.StatusMessage
							}
						}
					}

					if resp.Result.Service == "DEPOSIT" {
						helpers.LogAccountActivity(&c.Controller, resp.Result.ReceiverAccount, "Deposit", resp.Result.PaymentAmount, resp.Result.ServiceNetwork, "credit")
						responseStatus = true
						responseMessage = resp.StatusMessage
					}

					if resp.Result.Service == "WITHDRAWAL" {
						helpers.LogAccountActivity(&c.Controller, resp.Result.SenderAccount, "Withdrawal", resp.Result.PaymentAmount, resp.Result.ServiceNetwork, "debit")
						responseStatus = true
						responseMessage = resp.StatusMessage
					}

					if resp.Result.Service == "ACCOUNT OPENING" {
						uReq := requests.UsernameRequest{
							Username: resp.Result.ReceiverAccount,
						}
						customerData := apifunctions.GetCustomerByUsername(&c.Controller, uReq)

						gender := "M"
						switch genderStr := strings.ToLower(customerData.Customer.Gender); genderStr {
						case "male":
							gender = "M"
						case "female":
							gender = "F"
						case "m":
							gender = "M"
						case "f":
							gender = "F"
						}

						firstName := ""
						lastName := ""

						nameParts := strings.Fields(customerData.Customer.FullName)
						if len(nameParts) > 0 {
							firstName = nameParts[0]
						}
						if len(nameParts) > 1 {
							lastName = strings.Join(nameParts[1:], " ")
						}

						logs.Info("Mobile Number: ", customerData.Customer.PhoneNumber)
						logs.Info("First Name: ", firstName)
						logs.Info("Last Name: ", lastName)
						logs.Info("Gender: ", customerData.Customer.Gender)

						if client, err := models.GetClientsByCode(resp.Result.ServiceNetwork); err != nil {

							registerAccountRequest := requests.OpenAccountApiRequest{
								FirstName:    firstName,
								LastName:     lastName,
								Gender:       gender,
								MobileNumber: customerData.Customer.PhoneNumber,
								ClientId:     resp.Result.ServiceNetwork,
								Source:       "GHCOOPS",
							}

							resp_ := services.OpenAccount(&c.Controller, registerAccountRequest, *client, *customerData.Customer)

							logs.Info("Response from Account opening API: ", resp)

							if !resp_.StatusCode {
								responseStatus = false
								responseMessage = resp_.StatusMessage
							} else {
								responseStatus = true
								responseMessage = "Account opening successful"
							}
						}
					}
				}

				response := responses.CallbackResponse{
					StatusCode:    responseStatus,
					StatusMessage: responseMessage,
					Result:        &result,
				}

				c.Data["json"] = response
			}
		} else {
			logs.Info("Transaction not found for ID: %s", v.Data.ClientReference)
			response := responses.CallbackResponse{
				StatusCode:    false,
				StatusMessage: "Transaction not found",
				Result:        nil,
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
