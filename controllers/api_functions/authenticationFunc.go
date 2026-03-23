package apifunctions

import (
	"bytes"
	"encoding/json"
	"io"
	"msys_payment_app_gateway/api"
	"msys_payment_app_gateway/structs/requests"
	"msys_payment_app_gateway/structs/responses"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func SignInRequest(c *beego.Controller, req requests.SignIn) (resp responses.StringOriResponseDTO) {
	host, _ := beego.AppConfig.String("authenticationBaseUrl")

	logs.Info("Sending email ", req.Email)
	logs.Info("Sending password ", req.Password)

	request := api.NewRequest(
		host,
		"/v1/auth/login/token",
		api.POST)
	request.InterfaceParams["Username"] = req.Email
	request.InterfaceParams["Password"] = req.Password
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.StringOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	return data
}

func ChangePassword(c *beego.Controller, userid string, req requests.ChangePassword) (resp responses.StringOriResponseDTO) {
	host, _ := beego.AppConfig.String("authenticationBaseUrl")

	logs.Info("Sending old password ", req.OldPassword)
	logs.Info("Sending new password ", req.NewPassword)

	request := api.NewRequest(
		host,
		"/v1/auth/change-password/"+userid,
		api.PUT)
	request.InterfaceParams["OldPassword"] = req.OldPassword
	request.InterfaceParams["NewPassword"] = req.NewPassword
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.StringOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	return data
}

func ResetPassword(c *beego.Controller, userid string, req requests.ChangePassword) (resp responses.StringOriResponseDTO) {
	host, _ := beego.AppConfig.String("authenticationBaseUrl")

	logs.Info("Sending old password ", req.OldPassword)
	logs.Info("Sending new password ", req.NewPassword)

	request := api.NewRequest(
		host,
		"/v1/auth/reset-password/"+userid,
		api.PUT)
	request.InterfaceParams["NewPassword"] = req.NewPassword
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.StringOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	return data
}

func ChangeCustomerPassword(c *beego.Controller, customerid string, req requests.ChangePassword) (resp responses.StringOriResponseDTO) {
	host, _ := beego.AppConfig.String("authenticationBaseUrl")

	logs.Info("Sending old password ", req.OldPassword)
	logs.Info("Sending new password ", req.NewPassword)

	request := api.NewRequest(
		host,
		"/v1/auth/change-customer-password/"+customerid,
		api.PUT)
	request.Params["OldPassword"] = req.OldPassword
	request.Params["NewPassword"] = req.NewPassword
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.StringOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	return data
}

func ResetCustomerPassword(c *beego.Controller, customerid string, req requests.ChangePassword) (resp responses.StringOriResponseDTO) {
	host, _ := beego.AppConfig.String("authenticationBaseUrl")

	logs.Info("Sending old password ", req.OldPassword)
	logs.Info("Sending new password ", req.NewPassword)

	request := api.NewRequest(
		host,
		"/v1/auth/reset-customer-password/"+customerid,
		api.PUT)
	request.InterfaceParams["NewPassword"] = req.NewPassword
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.StringOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	return data
}

func VerifyToken(c *beego.Controller, token string) (resp responses.UserOriResponseDTO) {
	host, _ := beego.AppConfig.String("authenticationBaseUrl")

	logs.Info("About to verify token ", token)

	request := api.NewRequest(
		host,
		"/v1/auth/token/check",
		api.POST)
	request.InterfaceParams["Value"] = token
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.UserOriResponseDTO
	json.Unmarshal(read, &data)

	return data
}

func VerifyTokenNew(token string) (resp responses.CustomerResponseDTO2) {
	host, _ := beego.AppConfig.String("authenticationBaseUrl")

	logs.Info("About to verify token ", token)

	request := api.NewRequest(
		host,
		"/v1/auth/customer-token/check",
		api.POST)
	request.InterfaceParams["Value"] = token
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		// c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		// c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CustomerResponseDTO2
	json.Unmarshal(read, &data)

	return data
}

func VerifyUserToken(token string) (resp responses.UserOriResponseDTO) {
	host, _ := beego.AppConfig.String("authenticationBaseUrl")

	logs.Info("About to verify token ", token)

	request := api.NewRequest(
		host,
		"/v1/auth/token/verify",
		api.POST)
	request.InterfaceParams["Token"] = token
	request.HeaderField["Authorization"] = token
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		logs.Error("Error reading response body: ", err)
	}

	// Pretty print the raw JSON response for easier reading
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.UserOriResponseDTO
	json.Unmarshal(read, &data)

	return data
}

func RegistrationRequest(c *beego.Controller, req requests.RegisterUserApiRequest) (resp responses.UserResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Sending email ", req.Email)

	request := api.NewRequest(
		host,
		"/v1/users/sign-up",
		api.POST)
	request.InterfaceParams["Email"] = req.Email
	request.InterfaceParams["Password"] = req.Password
	request.InterfaceParams["Name"] = req.Name
	request.InterfaceParams["PhoneNumber"] = req.PhoneNumber
	request.InterfaceParams["Role"] = req.Role
	request.InterfaceParams["RoleRequired"] = false
	request.InterfaceParams["Dob"] = req.Dob
	request.InterfaceParams["Branch"] = req.Branch
	request.InterfaceParams["Username"] = req.Username

	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.UserResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	return data
}

func DeleteUser(c *beego.Controller, req requests.DeleteUserApiRequest) (resp responses.StringResponseCodeDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Sending email ", req.Username)

	request := api.NewRequest(
		host,
		"/v1/users/"+req.Username,
		api.DELETE)

	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.StringResponseCodeDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	return data
}

func CreateBranch(c *beego.Controller, req requests.BranchApiRequestDTO) (resp responses.BranchApiResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/branches",
		api.POST)
	request.InterfaceParams["Branch"] = req.BranchName
	request.InterfaceParams["CountryCode"] = req.CountryCode
	request.InterfaceParams["PhoneNumber"] = req.PhoneNumber
	request.InterfaceParams["Location"] = req.Location
	request.InterfaceParams["AddedBy"] = req.AddedBy
	request.InterfaceParams["BranchManager"] = req.BranchManager

	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.BranchApiResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	return data
}

func VerifyInviteToken(c *beego.Controller, token string) (resp responses.InviteDecodeResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("About to verify token ", token)

	request := api.NewRequest(
		host,
		"/v1/users/verify-invite",
		api.POST)
	request.InterfaceParams["Value"] = token
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.InviteDecodeResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	return data
}

func ResetPasswordLink(c *beego.Controller, email string, subject string, message string, links []*string) (resp responses.StringOriResponseDTO) {
	host, _ := beego.AppConfig.String("authenticationBaseUrl")

	logs.Info("Sending email ", email)

	request := api.NewRequest(
		host,
		"/v1/auth/reset-password-link",
		api.POST)
	request.InterfaceParams["Email"] = email
	request.InterfaceParams["Message"] = message
	request.InterfaceParams["Subject"] = subject
	request.InterfaceParams["Links"] = links
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.StringOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}
