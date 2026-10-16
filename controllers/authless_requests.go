package controllers

import (
	"encoding/json"
	"fmt"
	apifunctions "msys_payment_app_gateway/controllers/api_functions"
	"msys_payment_app_gateway/models"
	"msys_payment_app_gateway/structs/requests"
	"msys_payment_app_gateway/structs/responses"
	"msys_payment_app_gateway/utils"
	utilManager "msys_payment_app_gateway/utils"
	"path/filepath"
	"runtime"
	"time"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Agent_api_requestsController operations for Agent_api_requests
type Authless_requestsController struct {
	beego.Controller
}

// URLMapping ...
func (c *Authless_requestsController) URLMapping() {
	c.Mapping("GetCorporatives", c.GetCorporatives)
}

// GetCorporatives ...
// @Title Get Corporatives
// @Description Get Corporatives Available
// @Param	Authorization		header 	string true		"header for User"
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	Network		header 	string true		"header for network"
// @Param	body		body 	requests.GetBundlesAPIRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /get-corporatives [post]
func (c *Authless_requestsController) GetCorporatives() {
	// Extract headers
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	// sourceSystem := c.Ctx.Input.Header("SourceSystem")
	// network := c.Ctx.Input.Header("Network")

	var req requests.GetCorporativesRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	// destinationPhoneNumber := req.Destination

	_, file, line, ok := runtime.Caller(0)
	if ok {
		file = utils.GetFileName(file)
	} else {
		file = "unknown"
		line = 0
	}
	utilManager.Logger(filepath.Base(file), line, req.RequestId, "INFO", fmt.Sprintf("GetCorporatives request: %s", func() string { b, _ := json.Marshal(req); return string(b) }()))

	reqBody := c.Ctx.Input.RequestBody
	reqHeaders := c.Ctx.Request.Header

	requestMap := map[string]interface{}{
		"headers": reqHeaders,
		"body":    string(reqBody),
	}

	reqText, err := json.Marshal(requestMap)
	if err != nil {
		logs.Error("Error marshalling request input: ", err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	var v models.Api_requests = models.Api_requests{
		Request:      string(reqText),
		RequestType:  "Get Clients",
		PhoneNumber:  phoneNumber,
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)

		logs.Info("Formatted request for Corporatives: ")
		query := "HasPOS:1"
		resp := apifunctions.GetCorporatives(&c.Controller, query)
		_, file, line, ok := runtime.Caller(0)
		if ok {
			file = utils.GetFileName(file)
		} else {
			file = "unknown"
			line = 0
		}
		utilManager.Logger(filepath.Base(file), line, req.RequestId, "INFO", fmt.Sprintf("Response from Get corporatives API: %s", func() string { b, _ := json.Marshal(resp); return string(b) }()))

		var response responses.CorporativeResponse = responses.CorporativeResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong",
			Result:        nil,
		}

		if resp.StatusCode != 200 {
			response = responses.CorporativeResponse{
				StatusCode:    false,
				StatusMessage: resp.StatusMessage,
				Result:        resp.Result,
			}
		} else {
			responseText, err := json.Marshal(response.Result)
			if err != nil {
				logs.Error("Error marshalling response result: ", err)
				responseText = []byte("[]")
			}
			v.RequestResponse = string(responseText)
			v.DateModified = time.Now()
			v.ResponseDate = time.Now()
			if err := models.UpdateApi_requestsById(&v); err != nil {
				logs.Error("Error updating API request with response: ", err)
			} else {
				logs.Info("API request updated with response successfully: ", v)
			}
			response = responses.CorporativeResponse{
				StatusCode:    true,
				StatusMessage: "Corporatives fetched successfully",
				Result:        resp.Result,
			}
		}

		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = response

	} else {
		var response responses.CorporativeResponse = responses.CorporativeResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong:: " + err.Error(),
			Result:        nil,
		}

		c.Data["json"] = response
	}
	c.ServeJSON()
}
