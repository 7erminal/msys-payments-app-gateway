package apifunctions

import (
	"bytes"
	"encoding/json"
	"io"
	"msys_payment_app_gateway/api"
	"msys_payment_app_gateway/structs/requests"
	"msys_payment_app_gateway/structs/responses"
	"strconv"
	"time"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func AddCustomer(c *beego.Controller, req requests.AddCustomer) (resp responses.CustomerResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Sending email ", req.Email)

	// branchid := strconv.FormatInt(branch, 10)

	// Get date
	now := time.Now()
	y, m, d := now.Date()
	d_str := strconv.Itoa(d)
	m_str := strconv.Itoa(int(m))
	if len(d_str) < 2 {
		d_str = "0" + d_str
	}
	if len(m_str) < 2 {
		m_str = "0" + m_str
	}
	dob := strconv.Itoa(y) + "/" + m_str + "/" + d_str

	request := api.NewRequest(
		host,
		"/v1/customers/add-customer",
		api.POST)
	request.Params["Name"] = req.Name
	request.Params["Email"] = req.Email
	request.Params["IdType"] = req.IdType
	request.Params["PhoneNumber"] = req.PhoneNumber
	request.Params["IdNumber"] = req.IdNumber
	request.Params["Dob"] = dob
	request.Params["AddedBy"] = req.AddedBy
	request.Params["Location"] = req.Location
	if req.ImagePath != "" {
		request.FileField["CustomerImage"] = req.ImagePath
	}
	request.Params["Category"] = req.CustomerType
	// request.Params["Branch"] = branchid
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
	var data responses.CustomerResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	// logs.Info("Response received ", c.Data["json"])
	// logs.Info("Access token ", data["access_token"])
	// logs.Info("Expires in ", data["expires_in"])
	// logs.Info("Scope is ", data["scope"])
	// logs.Info("Token Type is ", data["token_type"])
	// logs.Info("Response received ", c.Data["json"])
	// logs.Info("Access token ", data.Access_token)
	// logs.Info("Expires in ", data.Expires_in)
	// logs.Info("Scope is ", data.Scope)
	// logs.Info("Token Type is ", data.Token_type)

	return data
}

func AddCustomerCredential(c *beego.Controller, req requests.AddCustomerCredential) (resp responses.StringResponseCodeDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	// branchid := strconv.FormatInt(branch, 10)

	request := api.NewRequest(
		host,
		"/v1/customer-credentials/add-customer-credential",
		api.POST)
	request.InterfaceParams["CustomerId"] = req.CustomerId
	request.InterfaceParams["Username"] = req.Username
	request.InterfaceParams["Password"] = req.Password
	request.InterfaceParams["Pin"] = req.Pin
	// request.Params["Branch"] = branchid
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
	var data responses.StringResponseCodeDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	// logs.Info("Response received ", c.Data["json"])
	// logs.Info("Access token ", data["access_token"])
	// logs.Info("Expires in ", data["expires_in"])
	// logs.Info("Scope is ", data["scope"])
	// logs.Info("Token Type is ", data["token_type"])
	// logs.Info("Response received ", c.Data["json"])
	// logs.Info("Access token ", data.Access_token)
	// logs.Info("Expires in ", data.Expires_in)
	// logs.Info("Scope is ", data.Scope)
	// logs.Info("Token Type is ", data.Token_type)

	return data
}

func GetCustomerDetails(c *beego.Controller, userid int64) (resp responses.CustomerResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting user details ", userid)

	request := api.NewRequest(
		host,
		"/v1/customers/"+strconv.FormatInt(userid, 10),
		api.GET)
	// request.Params["username"] = username
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

	// logs.Info("Raw response received is ", res)
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CustomerResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func GetCustomers(c *beego.Controller, query string, fields string, sortby string, order string,
	offset string, limit string, search string) (resp responses.CustomersResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting users ")

	request := api.NewRequest(
		host,
		"/v1/customers/",
		api.GET)
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.Params["search"] = search
	request.Params["query"] = query
	request.Params["fields"] = fields
	request.Params["sortby"] = sortby
	request.Params["order"] = order
	request.Params["offset"] = offset
	request.Params["limit"] = limit
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

	// logs.Info("Raw response received is ", res)

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	var data responses.CustomersResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	logs.Info("Resp is ", data.Customers)

	return data
}

func GetCustomersByBranch(c *beego.Controller, query string, fields string, sortby string, order string,
	offset string, limit string, branch int64) (resp responses.CustomersResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting users ")

	branchId := strconv.FormatInt(branch, 10)

	request := api.NewRequest(
		host,
		"/v1/customers/branch/"+branchId,
		api.GET)
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.Params["query"] = query
	request.Params["fields"] = fields
	request.Params["sortby"] = sortby
	request.Params["order"] = order
	request.Params["offset"] = offset
	request.Params["limit"] = limit
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
	var data responses.CustomersResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	logs.Info("Resp is ", data.Customers)

	return data
}

func GetCustomerCount(c *beego.Controller, query string, search string) (resp responses.StringOriResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/customers/count/",
		api.GET)

	request.Params["query"] = query
	request.Params["search"] = search

	// request.FileField["UserImage"] = userImage
	// request.Params["UserId"] = strconv.FormatInt(userId, 10)
	// request.HeaderField["content-type"] = "multipart/form-data"
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
	}

	// client.Request.HeaderField["content-type"] = "multipart/form-data"
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

func UpdateCustomer(c *beego.Controller, id string, userid string, req requests.UpdateCustomer, branch int64) (resp responses.CustomerResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Sending first name ", req.Name)

	logs.Info("Sending email ", req.Email)

	branchid := strconv.FormatInt(branch, 10)

	// Get date
	now := time.Now()
	y, m, d := now.Date()
	d_str := strconv.Itoa(d)
	m_str := strconv.Itoa(int(m))
	if len(d_str) < 2 {
		d_str = "0" + d_str
	}
	if len(m_str) < 2 {
		m_str = "0" + m_str
	}
	dob := strconv.Itoa(y) + "/" + m_str + "/" + d_str

	request := api.NewRequest(
		host,
		"/v1/customers/"+id,
		api.PUT)
	request.Params["Name"] = req.Name
	request.Params["Email"] = req.Email
	request.Params["IdType"] = req.IdType
	request.Params["PhoneNumber"] = req.PhoneNumber
	request.Params["IdNumber"] = req.IdNumber
	request.Params["Dob"] = dob
	request.Params["ModifiedBy"] = userid
	request.Params["Location"] = req.Location
	request.FileField["CustomerImage"] = req.ImagePath
	request.Params["Branch"] = branchid
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
	// var dataOri responses.UserOriResponseDTO
	var data responses.CustomerResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User.Branch.Country.DefaultCurrency)

	return data
}

func DeleteCustomer(c *beego.Controller, id string, req requests.UpdateCustomer) (resp responses.StringOriResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Sending first name ", req.Name)

	request := api.NewRequest(
		host,
		"/v1/customers/"+id,
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

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	// var dataOri responses.UserOriResponseDTO
	var data responses.StringOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User.Branch.Country.DefaultCurrency)

	return data
}

func GetIdTypes(c *beego.Controller, query string, fields string, sortby string, order string,
	offset string, limit string) (resp responses.IDTypeResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting Id Types ")

	request := api.NewRequest(
		host,
		"/v1/id-types",
		api.GET)
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.Params["query"] = query
	request.Params["fields"] = fields
	request.Params["sortby"] = sortby
	request.Params["order"] = order
	request.Params["offset"] = offset
	request.Params["limit"] = limit
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
	var data responses.IDTypeResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func DeactivateUser(c *beego.Controller, userid string) (resp responses.StringOriResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Deactivating user ", userid)

	request := api.NewRequest(
		host,
		"/v1/users/deactivate/"+userid,
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

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.StringOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func AddCustomerEmergencyContact(c *beego.Controller, req requests.AddCustomerEmergencyContact, customerId int64) (resp responses.CustomerEmergencyContactResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("EContact Sending name ", req.Name)
	logs.Info("EContact Sending phone number ", req.PhoneNumber)
	logs.Info("EContact Sending customerId ", customerId)

	request := api.NewRequest(
		host,
		"/v1/customer-emergency-contacts",
		api.POST)
	request.InterfaceParams["Name"] = req.Name
	request.InterfaceParams["PhoneNumber"] = req.PhoneNumber
	request.InterfaceParams["CustomerId"] = customerId
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
	var data responses.CustomerEmergencyContactResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	// logs.Info("Response received ", c.Data["json"])
	// logs.Info("Access token ", data["access_token"])
	// logs.Info("Expires in ", data["expires_in"])
	// logs.Info("Scope is ", data["scope"])
	// logs.Info("Token Type is ", data["token_type"])
	// logs.Info("Response received ", c.Data["json"])
	// logs.Info("Access token ", data.Access_token)
	// logs.Info("Expires in ", data.Expires_in)
	// logs.Info("Scope is ", data.Scope)
	// logs.Info("Token Type is ", data.Token_type)

	return data
}

func AddCustomerGuarantor(c *beego.Controller, req requests.AddCustomerGuarantor, customerId int64) (resp responses.CustomerGuarantorResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Guarantor Sending name ", req.Name)
	logs.Info("Guarantor Sending phone number ", req.PhoneNumber)
	logs.Info("Guarantor Sending customerId ", customerId)

	request := api.NewRequest(
		host,
		"/v1/customer-guarantors",
		api.POST)
	request.InterfaceParams["Name"] = req.Name
	request.InterfaceParams["PhoneNumber"] = req.PhoneNumber
	request.InterfaceParams["CustomerId"] = customerId
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
	var data responses.CustomerGuarantorResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	// logs.Info("Response received ", c.Data["json"])
	// logs.Info("Access token ", data["access_token"])
	// logs.Info("Expires in ", data["expires_in"])
	// logs.Info("Scope is ", data["scope"])
	// logs.Info("Token Type is ", data["token_type"])
	// logs.Info("Response received ", c.Data["json"])
	// logs.Info("Access token ", data.Access_token)
	// logs.Info("Expires in ", data.Expires_in)
	// logs.Info("Scope is ", data.Scope)
	// logs.Info("Token Type is ", data.Token_type)

	return data
}

func EditCustomerEmergencyContact(c *beego.Controller, req requests.EditCustomerEmergencyContact) (resp responses.CustomerEmergencyContactResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("EContact Sending name ", req.Name)
	logs.Info("EContact Sending phone number ", req.PhoneNumber)

	request := api.NewRequest(
		host,
		"/v1/customer-emergency-contacts/"+strconv.FormatInt(req.CustomerEmergencyContactId, 10),
		api.PUT)
	request.Params["Name"] = req.Name
	request.Params["PhoneNumber"] = req.PhoneNumber
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
	var data responses.CustomerEmergencyContactResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	// logs.Info("Response received ", c.Data["json"])
	// logs.Info("Access token ", data["access_token"])
	// logs.Info("Expires in ", data["expires_in"])
	// logs.Info("Scope is ", data["scope"])
	// logs.Info("Token Type is ", data["token_type"])
	// logs.Info("Response received ", c.Data["json"])
	// logs.Info("Access token ", data.Access_token)
	// logs.Info("Expires in ", data.Expires_in)
	// logs.Info("Scope is ", data.Scope)
	// logs.Info("Token Type is ", data.Token_type)

	return data
}

func EditCustomerGuarantor(c *beego.Controller, req requests.EditCustomerGuarantor) (resp responses.CustomerGuarantorResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Guarantor Sending name ", req.Name)
	logs.Info("Guarantor Sending phone number ", req.PhoneNumber)

	request := api.NewRequest(
		host,
		"/v1/customer-guarantors/"+strconv.FormatInt(req.CustomerGuarantorId, 10),
		api.PUT)
	request.Params["Name"] = req.Name
	request.Params["PhoneNumber"] = req.PhoneNumber
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
	var data responses.CustomerGuarantorResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	// logs.Info("Response received ", c.Data["json"])
	// logs.Info("Access token ", data["access_token"])
	// logs.Info("Expires in ", data["expires_in"])
	// logs.Info("Scope is ", data["scope"])
	// logs.Info("Token Type is ", data["token_type"])
	// logs.Info("Response received ", c.Data["json"])
	// logs.Info("Access token ", data.Access_token)
	// logs.Info("Expires in ", data.Expires_in)
	// logs.Info("Scope is ", data.Scope)
	// logs.Info("Token Type is ", data.Token_type)

	return data
}
