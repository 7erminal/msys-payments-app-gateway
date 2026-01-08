package services

import (
	apifunctions "msys_payment_app_gateway/controllers/api_functions"
	"msys_payment_app_gateway/structs/requests"
	"msys_payment_app_gateway/structs/responses"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Pay Water
func PayWater(c *beego.Controller, payWaterRequest requests.GhanaWaterPaymentFuncRequest) (response responses.GhanaWaterBillPaymentApiResponse) {
	logs.Info("Formatted request for Pay Water: ", payWaterRequest)
	status := "500"
	message := ""
	result := responses.GhanaWaterBillPaymentDataResponse{}
	req := requests.GhanaWaterPaymentApiRequest{
		TransactionId:      payWaterRequest.TransactionId,
		Amount:             payWaterRequest.Amount,
		DestinationAccount: payWaterRequest.DestinationAccount,
		PackageType:        payWaterRequest.PackageType,
		PhoneNumber:        payWaterRequest.PhoneNumber,
		SourceSystem:       payWaterRequest.SourceSystem,
		ExtraData: requests.GhanaWaterExtraData{
			Bundle:    payWaterRequest.DestinationAccount,
			Email:     payWaterRequest.Email,
			SessionId: payWaterRequest.PackageType,
		},
	}

	resp := apifunctions.PayGhanaWaterBill(c, req)
	logs.Info("Response from Pay Water API: ", resp)

	if resp.StatusCode != "200" {
		status = resp.StatusCode
		message = resp.StatusMessage
	} else {
		status = resp.StatusCode
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
