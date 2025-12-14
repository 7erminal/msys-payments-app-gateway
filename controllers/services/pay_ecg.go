package services

import (
	apifunctions "msys_payment_app_gateway/controllers/api_functions"
	"msys_payment_app_gateway/structs/requests"
	"msys_payment_app_gateway/structs/responses"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Pay ECG
func PayEcg(c *beego.Controller, payEcgRequest requests.ECGPaymentApiRequest) (response responses.ECGBillPaymentApiResponse) {
	logs.Info("Formatted request for Pay ECG: ", payEcgRequest)
	status := "401"
	message := ""
	result := responses.ECGBillPaymentDataResponse{}

	resp := apifunctions.PayECGBill(c, payEcgRequest)
	logs.Info("Response from Pay ECG API: ", resp)

	if resp.StatusCode != "200" {
		status = "01"
		message = resp.StatusMessage
	} else {
		status = "200"
		logs.Info("ECG payment successful: ", resp)
		message = "ECG payment is being processed"
		result = *resp.Result
	}

	response = responses.ECGBillPaymentApiResponse{
		StatusCode:    status,
		StatusMessage: message,
		Result:        &result,
	}

	return response
}
