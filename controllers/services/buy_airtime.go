package services

import (
	apifunctions "msys_payment_app_gateway/controllers/api_functions"
	"msys_payment_app_gateway/structs/requests"
	"msys_payment_app_gateway/structs/responses"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func BuyAirtime(c *beego.Controller, buyAirtimeRequest requests.BuyAirtimeFormulatedRequest) (response responses.BuyAirtimeAPIResponse) {
	logs.Info("Formatted request for Buy Airtime: ", buyAirtimeRequest)
	status := false
	message := ""
	result := responses.AirtimeResponseResult{}

	resp := apifunctions.BuyAirtime(c, buyAirtimeRequest)
	logs.Info("Response from Buy Airtime API: ", resp)

	if !resp.StatusCode {
		status = false
		message = resp.StatusMessage
	} else {
		status = true
		logs.Info("Airtime purchase successful: ", resp)
		message = "Airtime purchase is being processed"
		result = *resp.Result
	}

	response = responses.BuyAirtimeAPIResponse{
		StatusCode:    status,
		StatusMessage: message,
		Result:        &result,
	}

	return response
}
