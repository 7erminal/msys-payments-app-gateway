package helpers

import (
	"encoding/json"
	"errors"
	apifunctions "msys_payment_app_gateway/controllers/api_functions"
	"msys_payment_app_gateway/models"
	"msys_payment_app_gateway/structs/requests"
	"msys_payment_app_gateway/structs/responses"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func AccountProcessor(c *beego.Controller, req requests.VerifyCustomerApiRequest) {
	logs.Info("Processing Account Verification for PhoneNumber: ", req.MobileNumber)
	apifunctions.VerifyCustomer(c, req)
}

func CheckProfileCompletion(c *beego.Controller, customerData *responses.Customer) {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

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
	if customerCorporatives, err := models.GetAllCustomer_corporatives(query, fields, sortby, order, offset, limit); err == nil {
		logs.Info("Customer corporatives fetched successfully: ", customerCorporatives)
		// Process each corporative
		// var customerCorporativesDTO []responses.CustomerCorporativesResponseDTO
		for _, corpc := range customerCorporatives {
			logs.Info("Processing customer corporative: ", corpc)
			// Convert to DTO
			var corpDTO models.Customer_corporatives
			corpBytes, err := json.Marshal(corpc)
			if err != nil {
				logs.Error("Error marshalling customer corporative data: ", err)
				continue
			}
			if err := json.Unmarshal(corpBytes, &corpDTO); err != nil {
				logs.Error("Error unmarshalling customer corporative data: ", err)
			}
			logs.Info("Customer corporative DTO: ", corpDTO)
			// customerCorporativesDTO = append(customerCorporativesDTO, corpDTO)

			// Get corp code
			if corpDTO.IsActive == 0 {
				if corp, err := models.GetClientsById(corpDTO.CorpId.Id); err == nil {
					corpCode := corp.ClientCode

					approvedAccountsResp := apifunctions.FetchApprovedAccounts(c, corpCode)

					logs.Info("Approved accounts response: ", approvedAccountsResp)
					// Process approved accounts response
					if approvedAccountsResp.StatusCode != 200 {
						logs.Error("Error fetching approved accounts: ", approvedAccountsResp.StatusDesc)
					} else {
						logs.Info("Approved accounts fetched successfully: ", approvedAccountsResp.Result)

						for _, acc := range *approvedAccountsResp.Result {
							if strings.EqualFold(acc.MobileNumber, customerData.PhoneNumber) {
								logs.Info("Customer account already exists in approved accounts: ", acc)
								// Account exists - Update account to be verified
								activateAccountReq := requests.ActivateVerifiedCustomerApiRequest{
									Username:     acc.Username,
									ClientId:     corpCode,
									MobileNumber: customerData.PhoneNumber,
								}
								authorizeAccountResp := apifunctions.ActivateVerifiedCustomer(c, activateAccountReq)
								logs.Info("Authorize account response: ", authorizeAccountResp)
								if authorizeAccountResp.StatusCode != 200 {
									logs.Error("Error authorizing account: ", authorizeAccountResp.StatusDesc)
								} else {
									logs.Info("Account authorized successfully: ", authorizeAccountResp.Result)
									FetchCustomerAccounts(c, customerData, corpDTO, acc.MobileNumber)
								}
								break
							}
						}

						allActive := true
						for _, corpc := range customerCorporatives {
							var corpDTO models.Customer_corporatives
							corpBytes, err := json.Marshal(corpc)
							if err != nil {
								logs.Error("Error marshalling customer corporative data: ", err)
								allActive = false
								break
							}
							if err := json.Unmarshal(corpBytes, &corpDTO); err != nil {
								logs.Error("Error unmarshalling customer corporative data: ", err)
								allActive = false
								break
							}
							if corpDTO.IsActive != 1 {
								allActive = false
								break
							}
						}

						status := customerData.Active

						if allActive {
							status = 1
						} else {
							status = 2
						}

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
						logs.Info("About to update customer status to ", status)
						logs.Info("Before update, Customer was created on ", customerData.DateCreated)
						updateCustomerResp := apifunctions.UpdateCustomer(c, updatedcustmer)
						logs.Info("Update customer response: ", updateCustomerResp)
						if updateCustomerResp.StatusCode != 200 {
							logs.Error("Error updating customer status: ", updateCustomerResp.StatusDesc)
						} else {
							logs.Info("Customer status updated successfully: ", updateCustomerResp.Customer)
							logs.Info("After update, Customer was created on ", customerData.DateCreated)
						}
					}

				} else {
					logs.Error("Error fetching client by ID: ", err)
				}
			}

		}

	}
}

func FetchCustomerAccounts(c *beego.Controller, customerData *responses.Customer, corpDTO models.Customer_corporatives, phoneNumber string) {
	logs.Info("Fetching accounts for customer: ", customerData.CustomerNumber, " and client ID: ", corpDTO.CorpId.Id)
	if corp, err := models.GetClientsById(corpDTO.CorpId.Id); err == nil {
		corpCode := corp.ClientCode

		req := requests.NumberExistsApiRequest{
			MobileNumber: phoneNumber,
			ClientId:     corpCode,
		}

		listAccountResponse := apifunctions.ListCustomerAccounts(c, req)

		if listAccountResponse.StatusCode == 200 {
			logs.Info("Accounts fetched successfully for corporative ", corp.ClientName)
			if listAccountResponse.Result != nil && len(*listAccountResponse.Result) > 0 {
				for _, account := range *listAccountResponse.Result {
					logs.Info("Processing account: ", account)

					accountAlias := account.Product
					addAccountRequest := requests.CreateCustomerAccountApiRequest{
						AccountNumber: account.AccountNumber,
						AccountAlias:  accountAlias + " - " + customerData.FullName,
						AccountType:   accountAlias,
						Reference:     strconv.FormatInt(corp.Id, 10),
						CreatedBy:     int(customerData.CustomerId),
						Active:        1,
						CustomerId:    customerData.CustomerId,
					}

					addAccountResponse := apifunctions.AddCustomerAccount(c, addAccountRequest)

					if addAccountResponse.StatusCode == "200" {
						logs.Info("Account added successfully: ", addAccountResponse.Result)
						corpDTO.IsActive = 1
						if err := models.UpdateCustomer_corporativesById(&corpDTO); err != nil {
							logs.Error("Error updating customer corporative to active: ", err)
						} else {
							logs.Info("Customer corporative updated to active successfully: ", corpDTO)
						}
					} else {
						logs.Error("Error adding account: ", addAccountResponse.StatusMessage)
					}
				}
			}
		} else {
			logs.Error("Error fetching accounts for corporative ", corp.ClientName, ": ", listAccountResponse.StatusDesc)
		}
	}
}

func CheckAccountsStatus(c *beego.Controller, customerData *responses.Customer) {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

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
	if customerCorporatives, err := models.GetAllCustomer_corporatives(query, fields, sortby, order, offset, limit); err == nil {
		logs.Info("Customer corporatives fetched successfully: ", customerCorporatives)
		// Process each corporative
		// var customerCorporativesDTO []responses.CustomerCorporativesResponseDTO
		for _, corpc := range customerCorporatives {
			logs.Info("Processing customer corporative: ", corpc)
			// Convert to DTO
			var corpDTO models.Customer_corporatives
			corpBytes, err := json.Marshal(corpc)
			if err != nil {
				logs.Error("Error marshalling customer corporative data: ", err)
				continue
			}
			if err := json.Unmarshal(corpBytes, &corpDTO); err != nil {
				logs.Error("Error unmarshalling customer corporative data: ", err)
			}
			logs.Info("Customer corporative DTO: ", corpDTO)
			// customerCorporativesDTO = append(customerCorporativesDTO, corpDTO)

			// Get corp code
			if corpDTO.IsActive == 0 {
				if corp, err := models.GetClientsById(corpDTO.CorpId.Id); err == nil {
					corpCode := corp.ClientCode

					approvedAccountsResp := apifunctions.FetchApprovedAccounts(c, corpCode)

					logs.Info("Approved accounts response: ", approvedAccountsResp)
					// Process approved accounts response
					if approvedAccountsResp.StatusCode != 200 {
						logs.Error("Error fetching approved accounts: ", approvedAccountsResp.StatusDesc)
					} else {
						logs.Info("Approved accounts fetched successfully: ", approvedAccountsResp.Result)

						for _, acc := range *approvedAccountsResp.Result {
							if strings.EqualFold(acc.MobileNumber, customerData.PhoneNumber) {
								logs.Info("Customer account already exists in approved accounts: ", acc)
								// Account exists - Update account to be verified
								activateAccountReq := requests.ActivateVerifiedCustomerApiRequest{
									Username:     acc.Username,
									ClientId:     corpCode,
									MobileNumber: customerData.PhoneNumber,
								}
								authorizeAccountResp := apifunctions.ActivateVerifiedCustomer(c, activateAccountReq)
								logs.Info("Authorize account response: ", authorizeAccountResp)
								if authorizeAccountResp.StatusCode != 200 {
									logs.Error("Error authorizing account: ", authorizeAccountResp.StatusDesc)
								} else {
									logs.Info("Account authorized successfully: ", authorizeAccountResp.Result)
									// Update corporative as active
									corpDTO.IsActive = 1
									if err := models.UpdateCustomer_corporativesById(&corpDTO); err != nil {
										logs.Error("Error updating customer corporative as active: ", err)
									} else {
										logs.Info("Customer corporative updated as active successfully")
									}

								}
								break
							}
						}

						allActive := true
						for _, corpc := range customerCorporatives {
							var corpDTO models.Customer_corporatives
							corpBytes, err := json.Marshal(corpc)
							if err != nil {
								logs.Error("Error marshalling customer corporative data: ", err)
								allActive = false
								break
							}
							if err := json.Unmarshal(corpBytes, &corpDTO); err != nil {
								logs.Error("Error unmarshalling customer corporative data: ", err)
								allActive = false
								break
							}
							if corpDTO.IsActive != 1 {
								allActive = false
								break
							}
						}

						if allActive {
							customerData.Active = 1
						} else {
							customerData.Active = 2
						}

						updatedcustmer := requests.UpdateCustomer{
							Name:        customerData.FullName,
							Email:       customerData.Email,
							PhoneNumber: customerData.PhoneNumber,
							IdNumber:    customerData.IdentificationNumber,
							Location:    customerData.Location,
							Branch:      1, // Assuming default branch
							UserId:      1, // Assuming system user
							CustomerId:  customerData.CustomerId,
						}
						updateCustomerResp := apifunctions.UpdateCustomer(c, updatedcustmer)
						logs.Info("Update customer response: ", updateCustomerResp)
						if updateCustomerResp.StatusCode != 200 {
							logs.Error("Error updating customer status: ", updateCustomerResp.StatusDesc)
						} else {
							logs.Info("Customer status updated successfully: ", updateCustomerResp.Customer)
						}
					}

				}
			}

		}

	}
}

func LogTransaction(c *beego.Controller, transactionRequest requests.LogTransactionRequest) (transactionResponse responses.LogTransactionResponse, err error) {
	logs.Info("Logging transaction with request: ", transactionRequest)

	responseCode := 400
	responseMessage := "Error logging transaction"
	transaction := responses.Bil_transactions{}

	logTransactionRequest := requests.LogTransactionApiRequest{
		RequestId:                transactionRequest.RequestId,
		ServiceCode:              transactionRequest.ServiceCode,
		Amount:                   transactionRequest.Amount,
		Reference:                transactionRequest.Reference,
		Charge:                   transactionRequest.Charge,
		StatusCode:               transactionRequest.StatusCode,
		SourceAccountNumber:      transactionRequest.SourceAccountNumber,
		DestinationAccountNumber: transactionRequest.DestinationAccountNumber,
		TransactionType:          transactionRequest.TransactionType,
		ExternalReferenceNumber:  transactionRequest.ExternalReferenceNumber,
		TransactionReference:     transactionRequest.TransactionReference,
		ClientID:                 transactionRequest.ClientID,
		PhoneNumber:              transactionRequest.PhoneNumber,
		TransactionPackage:       transactionRequest.TransactionPackage,
		ExtraDetails1:            transactionRequest.ExtraDetails1,
		ExtraDetails2:            transactionRequest.ExtraDetails2,
		ExtraDetails3:            transactionRequest.ExtraDetails3,
		CorpId:                   transactionRequest.CorpId,
	}

	resp := apifunctions.LogTransaction(c, logTransactionRequest)
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		logs.Error("Error logging transaction: ", err)
		responseCode = 501
		responseMessage = "Error logging transaction: " + err.Error()
	} else {
		transaction = *resp.Result
		responseCode = 200
		responseMessage = "Successfully logged transaction"
	}

	return responses.LogTransactionResponse{
		StatusCode: responseCode,
		Result:     &transaction,
		StatusDesc: responseMessage,
	}, nil

}

func LogTransferTransaction(c *beego.Controller, transferRequest requests.TransferApiRequest, sendCommission bool) (transferResponse responses.TransferApiResponseDTO, err error) {
	logs.Info("Logging transfer transaction with request: ", transferRequest)

	responseCode := 400
	responseMessage := "Error logging transfer transaction"
	transferTransaction := responses.Trx_transactions{}

	transaction := apifunctions.LogTransferTransaction(c, transferRequest)

	if transaction.StatusCode != 200 && transaction.StatusCode != 201 {
		logs.Error("Error logging transfer transaction: ", transaction.StatusDesc)
		responseCode = 501
		responseMessage = "Error logging transfer transaction: " + transaction.StatusDesc
	} else {
		responseCode = transaction.StatusCode
		responseMessage = "Transfer transaction logged successfully"
		transferTransaction = *transaction.Result

		if sendCommission {
			req := requests.TransferCommissionApiRequest{
				RequestId:              transferRequest.RequestId,
				TransactionId:          transaction.Result.TransactionId,
				Amount:                 transferRequest.Amount,
				Description:            transferRequest.Description,
				Charge:                 transferRequest.Charge,
				Commission:             transferRequest.Commission,
				TotalDebitAmount:       transferRequest.TotalDebitAmount,
				SenderAccountNumber:    transferRequest.SenderAccountNumber,
				RecipientAccountNumber: transferRequest.RecipientAccountNumber,
				TransferCode:           transferRequest.TransferCode,
				RecipientName:          transferRequest.RecipientName,
				Status:                 transferRequest.Status,
			}

			commissionResp := apifunctions.SendCommission(c, req)

			if commissionResp.StatusCode != 200 && commissionResp.StatusCode != 201 {
				logs.Error("Error sending commission: ", commissionResp.StatusDesc)
				responseCode = 502
				responseMessage = "Error sending commission: " + commissionResp.StatusDesc
			} else {
				logs.Info("Commission sent successfully: ", commissionResp.Result)
				responseCode = 200
				responseMessage = "Transfer transaction and commission logged successfully"
				transferTransaction = *commissionResp.Result
			}
		}
	}

	resp := responses.TransferApiResponseDTO{
		StatusCode: responseCode,
		Result:     &transferTransaction,
		StatusDesc: responseMessage,
	}

	return resp, nil
}

func LogAccountActivity(c *beego.Controller, accountNumber string, reference string, amount float64, clientid string, activityType string) (activityResponse responses.AccountActivityResponse) {
	// type DebitAccountRequestV2 struct {
	// 	AccountNumber string
	// 	Amount        string
	// 	Reference     string
	// 	Channel       string
	// }

	activityResponse = responses.AccountActivityResponse{
		StatusCode:    false,
		StatusMessage: "Error processing account activity",
		Result:        "",
	}

	channel := "GHCOOPS"

	req := requests.AccountBalanceApiRequest{
		AccountNumber: accountNumber,
		ClientId:      clientid,
	}

	logs.Info("Fetching account balance for account number: ", accountNumber)
	logs.Info("Amount is ", amount)
	resp := GetAccountBalance(c, req)

	response := responses.AccountBalanceResponse{}

	if resp.StatusCode {
		logs.Info("Account balance fetched successfully: ", resp.Result)

		balance := resp.Result.AvailableBalance
		ClearBalance := resp.Result.ClearBalance

		logs.Info("Available balance: ", balance)
		logs.Info("Clear balance: ", ClearBalance)

		accResp := responses.AccountBalanceDataResp{
			AccountNumber:    req.AccountNumber,
			AccountStatus:    resp.Result.AccountStatus,
			AvailableBalance: resp.Result.AvailableBalance,
			ClearBalance:     resp.Result.ClearBalance,
			LoanBalance:      resp.Result.LoanBalance,
			SharesBalance:    resp.Result.SharesBalance,
		}

		response = responses.AccountBalanceResponse{
			StatusCode:    true,
			StatusMessage: "Account balance fetched successfully",
			Result:        &accResp,
		}

		// Debit/Credit the account
		logs.Info("Logging account activity of type ", activityType, " for account number: ", accountNumber)

		if strings.EqualFold(activityType, "debit") {
			debitAccReq := requests.DebitAccountRequest{
				Amount:     amount,
				ModifiedBy: 1, // System user
				Reason:     reference,
				AccountId:  resp.CustomerAccount.CustomerAccountId,
			}
			debitAccResp := apifunctions.DebitAccount(c, debitAccReq)
			logs.Info("Debit account internal response: ", debitAccResp)
			if debitAccResp.StatusCode != "200" {
				logs.Error("Error debiting account internally: ", debitAccResp.StatusMessage)
				activityResponse = responses.AccountActivityResponse{
					StatusCode:    false,
					StatusMessage: "Error debiting account: " + debitAccResp.StatusMessage,
					Result:        "",
				}
			} else {
				logs.Info("Account debited internally successfully: ", debitAccResp.Result)

				debitReq := requests.DebitAccountRequestV2{
					AccountNumber: accountNumber,
					Amount:        strconv.FormatFloat(amount, 'f', 2, 64),
					Reference:     reference,
					Channel:       channel,
					ClientId:      clientid,
				}
				debitResp := apifunctions.DebitAccountPro(c, debitReq)
				logs.Info("Debit account response: ", debitResp)

				if debitResp.StatusCode != 200 {
					logs.Error("Error debiting account: ", debitResp.StatusDesc, ". Reverseing internal debit...")

					creditAccReq := requests.CreditAccountRequest{
						Amount:     amount,
						ModifiedBy: 1, // System user
						Reason:     reference,
						AccountId:  resp.CustomerAccount.CustomerAccountId,
					}
					creditAccResp := apifunctions.CreditAccount(c, creditAccReq)
					logs.Info("Credit account internal response: ", creditAccResp)
					if creditAccResp.StatusCode != "200" {
						logs.Error("Error while reversing: ", creditAccResp.StatusMessage)

						activityResponse = responses.AccountActivityResponse{
							StatusCode:    false,
							StatusMessage: "Error debiting account: " + debitResp.StatusDesc + ". Reversal failed: " + creditAccResp.StatusMessage,
							Result:        "",
						}
					} else {
						logs.Info("Successful reversal: ", creditAccResp.Result)

						activityResponse = responses.AccountActivityResponse{
							StatusCode:    false,
							StatusMessage: "Error debiting account: " + debitResp.StatusDesc + ". Reversal successful.",
							Result:        "",
						}
					}
				} else {
					logs.Info("Account debited successfully: ", debitResp.Result)
					activityResponse = responses.AccountActivityResponse{
						StatusCode:    true,
						StatusMessage: "Account debited successfully",
						Result:        debitResp.Result,
					}
				}
			}
		} else if strings.EqualFold(activityType, "credit") {
			creditReq := requests.CreditAccountRequestV2{
				AccountNumber: accountNumber,
				Amount:        strconv.FormatFloat(amount, 'f', 2, 64),
				Reference:     reference,
				Channel:       channel,
				ClientId:      clientid,
			}
			creditResp := apifunctions.CreditAccountPro(c, creditReq)
			logs.Info("Credit account response: ", creditResp)

			if creditResp.StatusCode != 200 {
				logs.Error("Error crediting account: ", creditResp.StatusDesc)
				activityResponse = responses.AccountActivityResponse{
					StatusCode:    false,
					StatusMessage: "Error crediting account: " + creditResp.StatusDesc,
					Result:        "",
				}
			} else {
				logs.Info("Account credited successfully: ", creditResp.Result)
				creditAccReq := requests.CreditAccountRequest{
					Amount:     amount,
					ModifiedBy: 1, // System user
					Reason:     reference,
					AccountId:  resp.CustomerAccount.CustomerAccountId,
				}
				creditAccResp := apifunctions.CreditAccount(c, creditAccReq)
				logs.Info("Credit account internal response: ", creditAccResp)
				if creditAccResp.StatusCode != "200" {
					logs.Error("Error crediting account internally: ", creditAccResp.StatusMessage)
					activityResponse = responses.AccountActivityResponse{
						StatusCode:    false,
						StatusMessage: "Error crediting account internally: " + creditAccResp.StatusMessage,
						Result:        "",
					}
				} else {
					logs.Info("Account credited internally successfully: ", creditAccResp.Result)
					activityResponse = responses.AccountActivityResponse{
						StatusCode:    true,
						StatusMessage: "Account credited successfully",
						Result:        creditResp.Result,
					}
				}
			}
		} else {
			logs.Error("Invalid activity type specified for account activity logging")
		}

		logs.Info("Final response for account balance: ", response)

		// Log the activity
	}
	logs.Info("Account activity logging completed for account number: ", accountNumber)

	jsonBytes, err := json.MarshalIndent(activityResponse, "", "  ")
	if err != nil {
		logs.Error("Error marshalling activityResponse to JSON: ", err)
	} else {
		logs.Info("activityResponse JSON: ", string(jsonBytes))
	}
	return activityResponse
}

func GetAccountBalance(c *beego.Controller, req requests.AccountBalanceApiRequest) (data responses.CustAccountBalanceResponse) {
	logs.Info("Fetching account balance for account number: ", req.AccountNumber)
	resp := apifunctions.GetAccountBalance(c, req)

	response := responses.CustAccountBalanceResponse{}

	if resp.StatusCode == 200 {
		logs.Info("Account balance fetched successfully: ", resp.Result)

		balance := resp.Result.AvailableBalance
		ClearBalance := resp.Result.ClearBalance

		logs.Info("Available balance: ", balance)
		logs.Info("Clear balance: ", ClearBalance)

		accountsResp := apifunctions.GetCustomerAccount(c, req.AccountNumber)

		if accountsResp.StatusCode == "200" && accountsResp.Result != nil {
			logs.Info("Account details fetched successfully: ", accountsResp.Result)
			if accountsResp.Result.Balance != *balance {
				// Log the anomaly
				logs.Info("Balance mismatch detected. Logging account anomaly.")

				amountDifference := *balance - accountsResp.Result.Balance
				amountFloat, err := strconv.ParseFloat(strconv.FormatFloat(amountDifference, 'f', 2, 64), 64)
				if err != nil {
					logs.Error("Error parsing amount to float: ", err)
					amountFloat = 0.0
				}

				req := requests.CustomerAccountAnomaliesRequest{
					AccountNumber:  req.AccountNumber,
					Amount:         amountFloat,
					Desc:           "Balance mismatch detected during transaction. System Balance: " + strconv.FormatFloat(accountsResp.Result.Balance, 'f', 2, 64) + ", Actual Balance: " + strconv.FormatFloat(*balance, 'f', 2, 64),
					Balance:        accountsResp.Result.Balance,
					CheckedBalance: *balance,
					CreatedBy:      1, // System user
					ModifiedBy:     1, // System user
					Active:         1,
				}

				addAnomalyResp := apifunctions.ReportAccountAnomaly(c, req)

				if addAnomalyResp.StatusCode == "200" {
					logs.Info("Account anomaly logged successfully: ", addAnomalyResp.Result)
				} else {
					logs.Error("Error logging account anomaly: ", addAnomalyResp.StatusMessage)
				}

			}

			response = responses.CustAccountBalanceResponse{
				StatusCode:      true,
				StatusMessage:   "Account balance fetched successfully",
				Result:          resp.Result,
				CustomerAccount: accountsResp.Result,
			}

		} else {
			logs.Error("Error fetching account balance: ", resp.StatusDesc)
			response = responses.CustAccountBalanceResponse{
				StatusCode:    false,
				StatusMessage: resp.StatusDesc,
				Result:        nil,
			}
		}

		response = responses.CustAccountBalanceResponse{
			StatusCode:      true,
			StatusMessage:   "Account balance fetched successfully",
			Result:          resp.Result,
			CustomerAccount: accountsResp.Result,
		}
	} else {
		logs.Error("Error fetching account balance: ", resp.StatusDesc)
		response = responses.CustAccountBalanceResponse{
			StatusCode:    false,
			StatusMessage: resp.StatusDesc,
			Result:        nil,
		}
	}

	logs.Info("Final response for account balance: ", response)

	return response
}

func UpdateAccountBalance(c *beego.Controller, request_ requests.UpdateAccountBalanceApiRequest) (data responses.CustAccountBalanceResponse) {
	logs.Info("Fetching account balance for account number: ", request_.AccountNumber)
	req2 := requests.AccountBalanceApiRequest{
		AccountNumber: request_.AccountNumber,
		ClientId:      request_.ClientId,
	}
	resp := apifunctions.GetAccountBalance(c, req2)

	response := responses.CustAccountBalanceResponse{}

	if resp.StatusCode == 200 {
		logs.Info("Account balance fetched successfully: ", resp.Result)

		balance := resp.Result.AvailableBalance
		ClearBalance := resp.Result.ClearBalance
		sharesBalance := resp.Result.SharesBalance
		loanBalance := resp.Result.LoanBalance

		logs.Info("Available balance: ", balance)
		logs.Info("Clear balance: ", ClearBalance)
		logs.Info("Shares balance: ", sharesBalance)
		logs.Info("Loan balance: ", loanBalance)

		accountsResp := apifunctions.GetCustomerAccount(c, request_.AccountNumber)

		if accountsResp.StatusCode == "200" && accountsResp.Result != nil {
			logs.Info("Account details fetched successfully: ", accountsResp.Result)
			if accountsResp.Result.Balance != *balance {
				// Log the anomaly
				logs.Info("Balance mismatch detected. Logging account anomaly.")
				logs.Info("Balance is different. System Balance: ", accountsResp.Result.Balance, ", Actual Balance: ", *balance)

				amountDifference := *balance - accountsResp.Result.Balance
				amountFloat, err := strconv.ParseFloat(strconv.FormatFloat(amountDifference, 'f', 2, 64), 64)
				if err != nil {
					logs.Error("Error parsing amount to float: ", err)
					amountFloat = 0.0
				}

				req := requests.CustomerAccountAnomaliesRequest{
					AccountNumber:  request_.AccountNumber,
					Amount:         amountFloat,
					Desc:           "Balance mismatch detected during transaction. System Balance: " + strconv.FormatFloat(accountsResp.Result.Balance, 'f', 2, 64) + ", Actual Balance: " + strconv.FormatFloat(*balance, 'f', 2, 64),
					Balance:        accountsResp.Result.Balance,
					CheckedBalance: *balance,
					CreatedBy:      1, // System user
					ModifiedBy:     1, // System user
					Active:         1,
				}

				addAnomalyResp := apifunctions.ReportAccountAnomaly(c, req)

				if addAnomalyResp.StatusCode == "200" {
					logs.Info("Account anomaly logged successfully: ", addAnomalyResp.Result)
				} else {
					logs.Error("Error logging account anomaly: ", addAnomalyResp.StatusMessage)
				}

				newrequest_ := requests.UpdateAccountBalanceApiRequest{
					AccountId:     request_.AccountId,
					AccountNumber: request_.AccountNumber,
					Balance:       *balance,
					Reason:        request_.Reason,
					ModifiedBy:    request_.ModifiedBy,
					ClientId:      request_.ClientId,
				}

				logs.Info("Updating account balance with new balance: ", *balance)
				updateAccountBalanceResp := apifunctions.UpdateAccountBalance(c, newrequest_)

				if updateAccountBalanceResp.StatusCode == "200" {
					logs.Info("Account balance updated successfully: ", updateAccountBalanceResp.Result)
				} else {
					logs.Error("Unable to update account balance: ", updateAccountBalanceResp.StatusMessage)
				}

			}

			response = responses.CustAccountBalanceResponse{
				StatusCode:      true,
				StatusMessage:   "Account balance fetched successfully",
				Result:          resp.Result,
				CustomerAccount: accountsResp.Result,
			}

		} else {
			logs.Error("Error fetching account balance: ", resp.StatusDesc)
			response = responses.CustAccountBalanceResponse{
				StatusCode:    false,
				StatusMessage: resp.StatusDesc,
				Result:        nil,
			}
		}

		response = responses.CustAccountBalanceResponse{
			StatusCode:      true,
			StatusMessage:   "Account balance fetched successfully",
			Result:          resp.Result,
			CustomerAccount: accountsResp.Result,
		}
	} else {
		logs.Error("Error fetching account balance: ", resp.StatusDesc)
		response = responses.CustAccountBalanceResponse{
			StatusCode:    false,
			StatusMessage: resp.StatusDesc,
			Result:        nil,
		}
	}

	logs.Info("Final response for account balance: ", response)

	return response
}

func TempRegisterCustomer(c *beego.Controller, mobileNumber string, channel string, registerIfNotFound bool) (responses.CustomerResponseDTO, error) {
	logs.Info("Registering new customer with phone number: ", mobileNumber)

	responseStatus := 402
	responseMessage := "Error processing request"
	customer := responses.Customer{}
	resp := responses.CustomerResponseDTO{
		StatusCode: responseStatus,
		StatusDesc: responseMessage,
		Customer:   nil,
	}

	custDetails := apifunctions.GetCustomerDetailsWithPhoneNumber(c, mobileNumber)

	if custDetails.StatusCode != 200 {
		logs.Error("Error fetching customer details for phone number: ", mobileNumber)

		logs.Info("Register Customer")

		// Check phone number name inquiry
		nameInquiryReq := requests.NameInquiryApiRequestDTO{
			CustomerMsisdn: mobileNumber,
			Channel:        channel,
		}
		nameInquiryResp := apifunctions.NameInquiryViaMobileMoney(c, nameInquiryReq)
		logs.Info("Response from name inquiry API: ", nameInquiryResp)

		tempName := "Unknown"

		if nameInquiryResp.Success {
			// Register customer
			tempName = nameInquiryResp.Result.Name
		} else {
			logs.Error("Unable to process name inquiry for phone number: ", mobileNumber)
			responseMessage = "Unable to process name inquiry for phone number"
			responseStatus = 600
		}

		if registerIfNotFound {
			newCustomer := requests.AddCustomer{
				PhoneNumber:  mobileNumber,
				Name:         tempName,
				Email:        "",
				Location:     "",
				IdType:       "",
				IdNumber:     "",
				ImagePath:    "",
				AddedBy:      "1",
				CustomerType: "Individual", // Assuming default customer type
				Branch:       "1",
				Dob:          "2000-01-02",
				Status:       "Pending",
			}
			regResp := apifunctions.Register(c, newCustomer)
			logs.Info("Response from register customer API: ", regResp)

			if regResp.StatusCode == 200 {
				logs.Info("Customer registered successfully with phone number: ", mobileNumber)

				customer = *regResp.Customer
				responseStatus = 200
				responseMessage = "Customer registered successfully"
			} else {
				logs.Error("We are unable to register customer for phone number: ", mobileNumber)
				responseMessage = "We are unable to register customer for phone number"
			}
		} else {
			logs.Info("Registration skipped as per flag for phone number: ", mobileNumber)
			responseMessage = "Customer does not exist"
			responseStatus = 200
		}
	} else {
		logs.Info("Customer already exists with phone number: ", mobileNumber)

		customer = *custDetails.Customer
		responseStatus = 200
		responseMessage = "Customer fetched successfully"
	}

	resp = responses.CustomerResponseDTO{
		StatusCode: responseStatus,
		StatusDesc: responseMessage,
		Customer:   &customer,
	}

	return resp, nil
}

func MakePaymentMain(c *beego.Controller, req requests.PaymentRequestApiRequestDTO) (resp responses.MakePaymentResponse, err error) {
	logs.Info("Making Payment of ", req.Amount)

	isSuccess := false
	statusMessage := "Unable to process payment at the moment"
	var destinationPhoneNumber string
	var network string

	regCustResp, err := TempRegisterCustomer(c, req.MobileNumber, req.ClientId, false)
	if err != nil {
		logs.Error("Error registering customer: ", err)
	} else {
		logs.Info("Customer registered successfully: ", regCustResp)
		network = req.Network
		destinationPhoneNumber = req.ReceiverAccount

		if client, err := models.GetClientsByCode(req.ClientId); err != nil {
			logs.Error("Error getting client by Code: ", err)

			statusMessage = "Error getting client: " + err.Error()
		} else {

			requestMoney := requests.RequestMoneyApiRequestDTO{
				InitiatedBy:     regCustResp.Customer.CustomerId,
				Amount:          req.Amount,
				Service:         req.Service,
				Sender:          regCustResp.Customer.CustomerId,
				Reciever:        1,
				PhoneNumber:     destinationPhoneNumber,
				CustomerName:    regCustResp.Customer.FullName,
				CustomerMsisdn:  regCustResp.Customer.PhoneNumber,
				CustomerEmail:   regCustResp.Customer.Email,
				Currency:        "GHS",
				SenderAccount:   req.SenderAccount,
				ReceiverAccount: destinationPhoneNumber,
				PaymentMethod:   req.PaymentMethod,
				TransactionId:   req.TransactionId,
				PaymentProofUrl: "",
				ReferenceNumber: "",
				CallThirdParty:  true,
				Operator:        "HUBTEL",
				Network:         network,
				ServiceNetwork:  req.ServiceNetwork,
				ClientId:        client.ClientCorpId,
			}
			logs.Info("Sending money for ", req.Service, " purchase: ", requestMoney)

			reqMoneyResp, err := PaymentSendMoney(c, requestMoney)

			if err != nil {
				logs.Error("Error sending money to customer: ", err)
				isSuccess = false
				statusMessage = "Error sending money to customer: " + err.Error()
			} else {
				if reqMoneyResp.StatusCode == 200 {
					logs.Info("Payment request successful: ", reqMoneyResp)
					isSuccess = true
					statusMessage = "Payment request successful"
				} else {
					logs.Error("We could not process your payment request: ", reqMoneyResp.StatusDesc)
					isSuccess = false
					statusMessage = "We could not process your payment request: " + reqMoneyResp.StatusDesc
				}
			}
		}
	}

	response := responses.MakePaymentResponse{
		Success:       isSuccess,
		StatusMessage: statusMessage,
		Result:        nil,
	}

	return response, nil
}

func RequestPaymentMain(c *beego.Controller, req requests.PaymentRequestApiRequestDTO) (resp responses.MakePaymentResponse, err error) {
	logs.Info("Making Payment of ", req.Amount)

	isSuccess := false
	statusMessage := "Unable to process payment at the moment"
	var destinationPhoneNumber string
	var network string

	regCustResp, err := TempRegisterCustomer(c, req.MobileNumber, network, true)
	if err != nil {
		logs.Error("Error registering customer: ", err)
	} else {
		logs.Info("Customer registered successfully: ", regCustResp)
		network = req.Network
		destinationPhoneNumber = req.ReceiverAccount

		if req.PaymentMethod == "MOBILEMONEY" {
			logs.Info("Determined network for mobile money payment: ", network, " and service network is ", req.ServiceNetwork)

			if req.ServiceNetwork == req.Network || req.Network == "GET NETWORK FROM PHONE NUMBER" {
				logs.Info("Service network matches payment network, do name inquiry to confirm destination name")
				logs.Info("Mobile Number: ", req.MobileNumber)

				if req.PaymentMethod == "MOBILEMONEY" && (strings.HasPrefix(req.MobileNumber, "024") || strings.HasPrefix(req.MobileNumber, "054") || strings.HasPrefix(req.MobileNumber, "055") || strings.HasPrefix(req.MobileNumber, "059") || strings.HasPrefix(req.MobileNumber, "053")) {
					network = "MTN" + "_" + req.PaymentMethod
				} else if req.PaymentMethod == "MOBILEMONEY" && (strings.HasPrefix(req.MobileNumber, "020") || strings.HasPrefix(req.MobileNumber, "050")) {
					network = "TELECEL" + "_" + req.PaymentMethod
				} else if req.PaymentMethod == "MOBILEMONEY" && (strings.HasPrefix(req.MobileNumber, "027") || strings.HasPrefix(req.MobileNumber, "057")) {
					network = "AIRTELTIGO" + "_" + req.PaymentMethod
				} else {
					network = req.Network
				}

				logs.Info("Network found: ", network)
				// Check phone number name inquiry
				nameInquiryReq := requests.NameInquiryApiRequestDTO{
					CustomerMsisdn: req.MobileNumber,
					Channel:        network,
				}
				nameInquiryResp := apifunctions.NameInquiryViaMobileMoney(c, nameInquiryReq)

				if nameInquiryResp.Success {
					logs.Info("Name inquiry successful, updating destination phone number to: ", nameInquiryResp.Result.Name, " Network: ", nameInquiryResp.Result.Profile)
					// destinationPhoneNumber = nameInquiryResp.Result.Name
				} else {
					logs.Error("Name inquiry failed, cannot proceed with payment request")
					isSuccess = false
					statusMessage = "Name inquiry failed, cannot proceed with payment request: " + nameInquiryResp.StatusDesc
				}
			}
		}

		if client, err := models.GetClientsByCode(req.ClientId); err != nil {
			logs.Error("Error getting client by Code: ", err)

			statusMessage = "Error getting client: " + err.Error()
		} else {

			requestMoney := requests.RequestMoneyApiRequestDTO{
				InitiatedBy:     regCustResp.Customer.CustomerId,
				Amount:          req.Amount,
				Service:         req.Service,
				Sender:          regCustResp.Customer.CustomerId,
				Reciever:        1,
				PhoneNumber:     destinationPhoneNumber,
				CustomerName:    regCustResp.Customer.FullName,
				CustomerMsisdn:  regCustResp.Customer.PhoneNumber,
				CustomerEmail:   regCustResp.Customer.Email,
				Currency:        "GHS",
				SenderAccount:   req.SenderAccount,
				ReceiverAccount: destinationPhoneNumber,
				PaymentMethod:   req.PaymentMethod,
				TransactionId:   req.TransactionId,
				PaymentProofUrl: "",
				ReferenceNumber: "",
				CallThirdParty:  true,
				Operator:        "HUBTEL",
				Network:         network,
				ServiceNetwork:  req.ServiceNetwork,
				ServicePackage:  req.ServicePackage,
				ClientId:        client.ClientCorpId,
			}
			logs.Info("Requesting money from customer for : ", requestMoney)

			reqMoneyResp, err := PaymentRequestMoney(c, requestMoney)

			if err != nil {
				logs.Error("Error requesting money from customer: ", err)
				isSuccess = false
				statusMessage = "Error requesting money from customer: " + err.Error()
			} else {
				if reqMoneyResp.StatusCode == 200 {
					logs.Info("Payment request successful: ", reqMoneyResp)
					isSuccess = true
					statusMessage = "Payment request successful"
				} else {
					logs.Error("Error in payment request response: ", reqMoneyResp.StatusDesc)
					isSuccess = false
					statusMessage = "Error requesting money from customer: " + reqMoneyResp.StatusDesc
				}
			}
		}
	}

	response := responses.MakePaymentResponse{
		Success:       isSuccess,
		StatusMessage: statusMessage,
		Result:        nil,
	}

	return response, nil
}

func PaymentRequestMoney(c *beego.Controller, req requests.RequestMoneyApiRequestDTO) (resp responses.PaymentApiResponseDTO, err error) {
	logs.Info("Requesting Money ", req.Amount, " from ", req.InitiatedBy)
	requestPayment := requests.MakePaymentApiRequestDTO{
		InitiatedBy:     req.InitiatedBy,
		Amount:          req.Amount,
		Service:         req.Service,
		Sender:          req.Sender,
		SenderAccount:   req.SenderAccount,
		ReceiverAccount: req.ReceiverAccount,
		Reciever:        req.Reciever,
		PaymentMethod:   req.PaymentMethod,
		TransactionId:   req.TransactionId,
		PaymentProofUrl: req.PaymentProofUrl,
		ReferenceNumber: req.ReferenceNumber,
		CallThirdParty:  req.CallThirdParty,
		Operator:        req.Operator,
		Network:         req.Network,
		ServiceNetwork:  req.ServiceNetwork,
		ServicePackage:  req.ServicePackage,
		ClientId:        req.ClientId,
	}

	resp = apifunctions.MakePayment(c, requestPayment)

	if resp.StatusCode != 200 {
		logs.Error("Error making payment request: ", resp.StatusDesc)
		return resp, errors.New(resp.StatusDesc)
	} else {
		logs.Info("Payment request made successfully: ", resp.Payment)
		logs.Info("Initiating Momo Payment for Payment ID: ", resp.Payment.PaymentId)
		logs.Info("Amount after sending for payment is ", resp.Payment.Amount)

		momoRequest := requests.MomoPaymentApiRequestDTO{
			PaymentId:          resp.Payment.PaymentId,
			Amount:             resp.Payment.Amount,
			CustomerName:       resp.Payment.Sender,
			CustomerMsisdn:     req.CustomerMsisdn,
			CustomerEmail:      req.CustomerEmail,
			Operator:           req.Operator,
			PrimaryCallbackUrl: resp.Payment.CallbackUrl,
			Description:        "Payment Request for " + strconv.FormatFloat(resp.Payment.Amount, 'f', 2, 64),
			ClientReference:    req.TransactionId,
			Channel:            req.Network,
			ClientId:           req.ClientId,
		}

		momoResp := apifunctions.RequestMoneyViaMobileMoney(c, momoRequest)

		if momoResp.StatusCode != 200 {
			logs.Error("Error initiating momo payment: ", momoResp.StatusDesc)
			return resp, errors.New(momoResp.StatusDesc)
		} else {
			logs.Info("Momo payment initiated successfully: ", momoResp.Payment)

			logs.Info("Updated payment with momo details: ", resp.Payment)
		}
	}
	return resp, nil

}

func PaymentSendMoney(c *beego.Controller, req requests.RequestMoneyApiRequestDTO) (resp responses.PaymentApiResponseDTO, err error) {
	logs.Info("Requesting Money ", req.Amount, " from ", req.InitiatedBy)
	requestPayment := requests.MakePaymentApiRequestDTO{
		InitiatedBy:     req.InitiatedBy,
		Amount:          req.Amount,
		Service:         req.Service,
		Sender:          req.Sender,
		SenderAccount:   req.SenderAccount,
		ReceiverAccount: req.ReceiverAccount,
		Reciever:        req.Reciever,
		PaymentMethod:   req.PaymentMethod,
		TransactionId:   req.TransactionId,
		PaymentProofUrl: req.PaymentProofUrl,
		ReferenceNumber: req.ReferenceNumber,
		CallThirdParty:  req.CallThirdParty,
		Operator:        req.Operator,
		Network:         req.Network,
		ServiceNetwork:  req.ServiceNetwork,
		ServicePackage:  req.ServicePackage,
		ClientId:        req.ClientId,
	}

	resp = apifunctions.MakePayment(c, requestPayment)

	if resp.StatusCode != 200 {
		logs.Error("Error making payment request: ", resp.StatusDesc)
		return resp, errors.New(resp.StatusDesc)
	} else {
		if strings.EqualFold(req.Service, "WITHDRAWAL") {
			logs.Info("Payment request made successfully: ", resp.Payment)
			logs.Info("Initiating Momo Payment for Payment ID: ", resp.Payment.PaymentId)

			momoRequest := requests.MomoPaymentApiRequestDTO{
				PaymentId:          resp.Payment.PaymentId,
				Amount:             resp.Payment.Amount,
				CustomerName:       resp.Payment.Sender,
				CustomerMsisdn:     req.CustomerMsisdn,
				CustomerEmail:      req.CustomerEmail,
				Operator:           req.Operator,
				PrimaryCallbackUrl: resp.Payment.CallbackUrl,
				Description:        "Payment Request for " + strconv.FormatFloat(resp.Payment.Amount, 'f', 2, 64),
				ClientReference:    req.TransactionId,
				Channel:            req.Network,
				ClientId:           req.ClientId,
			}

			momoResp := apifunctions.SendMoneyViaMobileMoney(c, momoRequest)

			if momoResp.StatusCode != 200 {
				logs.Error("Error initiating momo payment: ", momoResp.StatusDesc)
				return resp, errors.New(momoResp.StatusDesc)
			} else {
				logs.Info("Momo payment initiated successfully: ", momoResp.Payment)

				logs.Info("Updated payment with momo details: ", resp.Payment)
			}
		}
	}
	return resp, nil

}
