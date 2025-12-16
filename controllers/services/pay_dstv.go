package services

import (
	apifunctions "msys_payment_app_gateway/controllers/api_functions"
	"msys_payment_app_gateway/structs/requests"
	"msys_payment_app_gateway/structs/responses"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Pay DStv
func PayDstv(c *beego.Controller, payDstvRequest requests.DSTVPaymentRequest) (response responses.DSTVBillPaymentApiResponse) {
	logs.Info("Formatted request for Pay DStv: ", payDstvRequest)
	status := "500"
	message := ""
	result := responses.DSTVBillPaymentDataResponse{}

	resp := apifunctions.PayDSTVBill(c, payDstvRequest)
	logs.Info("Response from Pay DStv API: ", resp)

	if resp.StatusCode != "200" {
		status = "501"
		message = resp.StatusMessage
	} else {
		status = "200"
		logs.Info("DStv payment successful: ", resp)
		message = "DStv payment is being processed"
		result = *resp.Result
	}

	response = responses.DSTVBillPaymentApiResponse{
		StatusCode:    status,
		StatusMessage: message,
		Result:        &result,
	}

	return response
}
