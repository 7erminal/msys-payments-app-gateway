package services

import (
	apifunctions "msys_payment_app_gateway/controllers/api_functions"
	"msys_payment_app_gateway/structs/requests"
	"msys_payment_app_gateway/structs/responses"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Pay GoTV
func PayGotv(c *beego.Controller, payGotvRequest requests.GoTvPaymentApiRequest) (response responses.GoTvBillPaymentApiResponse) {
	logs.Info("Formatted request for Pay GoTV: ", payGotvRequest)
	status := false
	message := ""
	result := responses.GoTvBillPaymentDataResponse{}

	resp := apifunctions.PayGoTvBill(c, payGotvRequest)
	logs.Info("Response from Pay GoTV API: ", resp)

	if !resp.StatusCode {
		status = false
		message = resp.StatusMessage
	} else {
		status = true
		logs.Info("GoTV payment successful: ", resp)
		message = "GoTV payment is being processed"
		result = *resp.Result
	}

	response = responses.GoTvBillPaymentApiResponse{
		StatusCode:    status,
		StatusMessage: message,
		Result:        &result,
	}

	return response
}
