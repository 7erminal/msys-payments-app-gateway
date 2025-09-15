package helpers

import (
	"encoding/json"
	"errors"
	apifunctions "msys_payment_app_gateway/controllers/api_functions"
	"msys_payment_app_gateway/models"
	"msys_payment_app_gateway/structs/requests"
	"msys_payment_app_gateway/structs/responses"
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
									// Update corporative as active

									logs.Info("Customer corporative updated as active successfully")

									active := 1
									updateCustomer := requests.UpdateCustomer{
										Name:        customerData.FullName,
										Email:       customerData.Email,
										PhoneNumber: customerData.PhoneNumber,
										IdNumber:    customerData.IdentificationNumber,
										Location:    customerData.Location,
										Branch:      1, // Assuming default branch
										UserId:      1, // Assuming system user
										CustomerId:  customerData.CustomerId,
										Status:      active,
									}

									logs.Info("Updating customer status to active: ", updateCustomer)
									logs.Info("Customer status is ", active)
									updateCustomerResp := apifunctions.UpdateCustomer(c, updateCustomer)
									logs.Info("Update customer response: ", updateCustomerResp)
									if updateCustomerResp.StatusCode != 200 {
										logs.Error("Error updating customer status: ", updateCustomerResp.StatusDesc)
									} else {
										logs.Info("Customer status updated successfully: ", updateCustomerResp.Customer)
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

				} else {
					logs.Error("Error fetching client by ID: ", err)
				}
			}

		}

	}
}

func FetchCustomerAccounts(c *beego.Controller, customerData *responses.Customer, corpDTO models.Customer_corporatives, phoneNumber string) {
	logs.Info("Fetching accounts for customer: ", customerData.CustomerNumber)
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

					accountAlias := account.AccountNumber
					addAccountRequest := requests.CreateCustomerAccountApiRequest{
						AccountNumber: account.AccountNumber,
						AccountAlias:  accountAlias,
						CreatedBy:     int(customerData.CustomerId),
						Active:        1,
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
