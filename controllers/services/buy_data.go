package services

import (
	apifunctions "msys_payment_app_gateway/controllers/api_functions"
	"msys_payment_app_gateway/structs/requests"
	"msys_payment_app_gateway/structs/responses"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Buy Data Bundle
func BuyDataBundle(c *beego.Controller, buyDataBundleRequest requests.BuyDataBundleFormulatedRequest) (response responses.BuyDataBundleAPIResponse) {
	logs.Info("Formatted request for Buy Data Bundle: ", buyDataBundleRequest)
	status := false
	message := ""
	result := responses.BuyDataBundleResponseResult{}

	resp := apifunctions.BuyDataBundle(c, buyDataBundleRequest)
	logs.Info("Response from Buy Data Bundle API: ", resp)

	if !resp.StatusCode {
		status = false
		message = resp.StatusMessage
	} else {
		status = true
		logs.Info("Data Bundle purchase successful: ", resp)
		message = "Data Bundle purchase is being processed"
		result = *resp.Result
	}

	response = responses.BuyDataBundleAPIResponse{
		StatusCode:    status,
		StatusMessage: message,
		Result:        &result,
	}

	return response
}
