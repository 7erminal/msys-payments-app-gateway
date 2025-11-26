package services

import (
	apifunctions "msys_payment_app_gateway/controllers/api_functions"
	"msys_payment_app_gateway/models"
	"msys_payment_app_gateway/structs/requests"
	"msys_payment_app_gateway/structs/responses"
	"time"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Open Account
func OpenAccount(c *beego.Controller, openAccountRequest requests.OpenAccountApiRequest, client models.Clients, customerData responses.Customer) (response responses.RegisterAccountResponse) {
	logs.Info("Formatted request for account opening: ", openAccountRequest)
	status := false
	message := ""
	result := false

	resp := apifunctions.OpenAccount(c, openAccountRequest)
	logs.Info("Response from open API: ", resp)

	// if resp.Data.StatusCode != 200 {
	// 	status = false
	// 	message = resp.Data.StatusDesc
	// } else {
	// 	status = true
	// 	logs.Info("Account opening successful: ", resp)
	// 	message = "Account opening is successful"
	// 	result = true
	// }
	if resp.Data.StatusCode != 200 {
		logs.Info("Response returned is not successful...", resp.Data.StatusDesc)

		status = false
		message = resp.Data.StatusDesc
		result = false
	} else {
		logs.Info("Account registered successfully, adding customer corporative...")
		logs.Info("Client details: ", client.Id, " - ", client.ClientCode, " - ", client.ClientName)
		customerCorporative := models.Customer_corporatives{
			CustomerNumber: customerData.CustomerNumber,
			CorpId:         &client, // Assuming default corp ID, can be changed later
			IsActive:       1,       // Set to inactive until verified
			CreatedBy:      1,
			ModifiedBy:     1,
			IsDefault:      1,
			DateCreated:    time.Now(),
			DateModified:   time.Now(),
		}

		if cl, err := models.GetCustomer_corporativesByClient(customerData.CustomerNumber, client.Id); err == nil {

			logs.Info("Customer corporative exists, adding new record...")
			logs.Info("Customer corporative already exists: ", cl)

			status = true
			message = "Customer corporative already exists"
			result = true

		} else {
			if _, err := models.AddCustomer_corporatives(&customerCorporative); err != nil {
				logs.Error("An error occurred adding customer corporative ", err.Error())

				status = false
				message = "An error occurred adding customer corporative. " + err.Error()
				result = false
			} else {
				logs.Info("Customer corporative added successfully.")
				status = true
				message = resp.Data.StatusDesc
				result = true
			}
		}

	}

	response = responses.RegisterAccountResponse{
		StatusCode:    status,
		StatusMessage: message,
		Result:        &result,
	}

	return response
}
