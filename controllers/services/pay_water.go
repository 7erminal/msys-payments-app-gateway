package services

import (
	apifunctions "msys_payment_app_gateway/controllers/api_functions"
	"msys_payment_app_gateway/structs/requests"
	"msys_payment_app_gateway/structs/responses"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Pay Water
func PayWater(c *beego.Controller, payWaterRequest requests.GhanaWaterPaymentApiRequest) (response responses.GhanaWaterBillPaymentApiResponse) {
	logs.Info("Formatted request for Pay Water: ", payWaterRequest)
	status := false
	message := ""
	result := responses.GhanaWaterBillPaymentDataResponse{}

	resp := apifunctions.PayGhanaWaterBill(c, payWaterRequest)
	logs.Info("Response from Pay Water API: ", resp)

	if !resp.StatusCode {
		status = false
		message = resp.StatusMessage
	} else {
		status = true
		logs.Info("Water payment successful: ", resp)
		message = "Water payment is being processed"
		result = *resp.Result
	}

	response = responses.GhanaWaterBillPaymentApiResponse{
		StatusCode:    status,
		StatusMessage: message,
		Result:        &result,
	}

	return response
}
