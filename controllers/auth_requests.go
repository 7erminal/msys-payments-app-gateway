package controllers

import (
	"msys_payment_app_gateway/structs/responses"

	"encoding/json"
	apifunctions "msys_payment_app_gateway/controllers/api_functions"
	"msys_payment_app_gateway/models"
	"msys_payment_app_gateway/structs/requests"
	"time"

	beego "github.com/beego/beego/v2/server/web"

	"github.com/beego/beego/v2/core/logs"
)

// Auth_requestsController operations for Auth_requests
type Auth_requestsController struct {
	beego.Controller
}

// URLMapping ...
func (c *Auth_requestsController) URLMapping() {
	c.Mapping("Login", c.Login)
	c.Mapping("Register", c.Register)
}

// Login ...
// @Title Login
// @Description Login
// @Param	PhoneNumber		header 	string true		"header for Customer's phone number"
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	body		body 	requests.LoginRequest	true		"body for Request content"
// @Success 200 {int} responses.LoginResponse
// @Failure 403 body is empty
// @router /login [post]
func (c *Auth_requestsController) Login() {
	// Extract headers
	// phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")
	phoneNumber := c.Ctx.Input.Header("PhoneNumber")

	var req requests.LoginRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	logs.Info("Login called with PhoneNumber: %s, SourceSystem: %s", phoneNumber, sourceSystem)
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
		PhoneNumber:  phoneNumber,
		RequestType:  "Login",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)
		loginRequest := requests.LoginApiRequest{
			PhoneNumber:  req.PhoneNumber,
			SourceSystem: sourceSystem,
			Password:     req.Password,
			ClientId:     req.ClientId,
		}

		logs.Info("Formatted request for Login: ", loginRequest)
		resp := apifunctions.Login(&c.Controller, loginRequest)
		logs.Info("Response from Login API: ", resp)

		var response responses.LoginResponse = responses.LoginResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong",
			Result:        "",
		}

		if resp.StatusCode != 200 {
			response = responses.LoginResponse{
				StatusCode:    false,
				StatusMessage: resp.StatusDesc,
				Result:        "",
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
			response = responses.LoginResponse{
				StatusCode:    true,
				StatusMessage: "Login successful",
				Result:        resp.Value,
			}
		}

		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = response

	} else {
		var response responses.LoginResponse = responses.LoginResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong:: " + err.Error(),
			Result:        "",
		}

		c.Data["json"] = response
	}
	c.ServeJSON()
}

// Register ...
// @Title Register
// @Description Register customer
// @Param	SourceSystem		header 	string true		"header for Source system"
// @Param	body		body 	requests.RegisterRequest	true		"body for Request content"
// @Success 201 {int} models.Api_requests
// @Failure 403 body is empty
// @router /register [post]
func (c *Auth_requestsController) Register() {
	// Extract headers
	// phoneNumber := c.Ctx.Input.Header("PhoneNumber")
	sourceSystem := c.Ctx.Input.Header("SourceSystem")

	var req requests.RegisterRequest
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		logs.Error("Error unmarshalling request body: ", err)
		c.Data["json"] = "Invalid request body"
		c.ServeJSON()
		return
	}

	logs.Info("Register called with SourceSystem: %s", sourceSystem)
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
		PhoneNumber:  req.MobileNumber,
		RequestType:  "Register",
		RequestDate:  time.Now(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
	}
	if _, err := models.AddApi_requests(&v); err == nil {
		logs.Info("API request logged successfully: ", v)

		mobileNumberValidation := requests.MobileNumberRequest{
			MobileNumber: req.MobileNumber,
		}

		checkCustomer := apifunctions.GetCustomer(&c.Controller, mobileNumberValidation)

		if checkCustomer.StatusCode == 200 && checkCustomer.Customer != nil {
			logs.Info("Customer already exists: ", checkCustomer.Customer)
			var response responses.RegisterResponse = responses.RegisterResponse{
				StatusCode:    false,
				StatusMessage: "Customer already exists",
				Result:        nil,
			}
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = response
			c.ServeJSON()
			return
		}

		usernameReq := requests.UsernameRequest{
			Username: req.Username,
		}

		checkCustomerUsername := apifunctions.GetCustomerByUsername(&c.Controller, usernameReq)

		if checkCustomerUsername.StatusCode == 200 && checkCustomerUsername.Customer != nil {
			logs.Info("Customer already exists: ", checkCustomerUsername.Customer)
			var response responses.RegisterResponse = responses.RegisterResponse{
				StatusCode:    false,
				StatusMessage: "Username already exists",
				Result:        nil,
			}
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = response
			c.ServeJSON()
			return
		}
		logs.Info("Date of Birth: ", req.Dob)
		// dob := req.Dob.Format("2006-01-02")
		// if dob == "0001-01-01" {
		// 	dob = req.Dob.Format("2006/01/02")
		// }
		dob := req.Dob
		logs.Info("Formatted Date of Birth: ", dob)
		logs.Info("Mobile Number: ", req.MobileNumber)
		logs.Info("First Name: ", req.FirstName)
		logs.Info("Last Name: ", req.LastName)
		registerRequest := requests.AddCustomer{
			PhoneNumber:  req.MobileNumber,
			Name:         req.LastName + " " + req.FirstName,
			Email:        req.Email,
			Location:     "",
			IdType:       "",
			IdNumber:     "",
			ImagePath:    "",
			AddedBy:      "1",
			CustomerType: "Individual", // Assuming default customer type
			Branch:       "1",
			Dob:          dob,
		}

		logs.Info("Formatted request for Register: ", registerRequest)
		resp := apifunctions.Register(&c.Controller, registerRequest)
		logs.Info("Response from Register API: ", resp)

		var response responses.RegisterResponse = responses.RegisterResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong",
			Result:        nil,
		}

		if resp.StatusCode != 200 {
			response = responses.RegisterResponse{
				StatusCode:    false,
				StatusMessage: resp.StatusDesc,
				Result:        nil,
			}
		} else {
			// Save customer credentials
			saveCustomerCredential := requests.AddCustomerCredential{
				CustomerId: resp.Customer.CustomerId,
				Username:   req.Username,
				Password:   req.Password,
				Pin:        "1234", // Default pin, can be changed later
			}
			logs.Info("Formatted request for AddCustomerCredential: ", saveCustomerCredential)
			credentialResp := apifunctions.AddCustomerCredential(&c.Controller, saveCustomerCredential)
			logs.Info("Response from AddCustomerCredential API: ", credentialResp)
			if credentialResp.StatusCode != 200 {
				logs.Error("Error saving customer credentials: ", credentialResp.StatusDesc)
				response = responses.RegisterResponse{
					StatusCode:    false,
					StatusMessage: "Error saving customer credentials: " + credentialResp.StatusDesc,
					Result:        nil,
				}
				c.Ctx.Output.SetStatus(400)
				c.Data["json"] = response
				c.ServeJSON()
				return
			}
			logs.Info("Customer credentials saved successfully: ", credentialResp)
			// Add customer corporative
			if client, err := models.GetClientsById(req.ClientId); err != nil {
				logs.Error("Error getting client by ID: ", err)
				response = responses.RegisterResponse{
					StatusCode:    false,
					StatusMessage: "Error getting client by ID: " + err.Error(),
					Result:        nil,
				}
			} else {
				customerCorporative := models.Customer_corporatives{
					CustomerNumber: resp.Customer.CustomerNumber,
					CorpId:         client, // Assuming default corp ID, can be changed later
				}

				if _, err := models.AddCustomer_corporatives(&customerCorporative); err != nil {
					logs.Error("An error occurred adding customer corporative ", err.Error())
					response = responses.RegisterResponse{
						StatusCode:    false,
						StatusMessage: "An error occurred adding customer corporative. " + err.Error(),
						Result:        nil,
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
					cust := responses.CustomerGateway{
						CustomerId:           resp.Customer.CustomerId,
						FullName:             resp.Customer.FullName,
						PhoneNumber:          resp.Customer.PhoneNumber,
						Email:                resp.Customer.Email,
						Location:             resp.Customer.Location,
						IdentificationType:   nil,
						IdentificationNumber: resp.Customer.IdentificationNumber,
						ImagePath:            resp.Customer.ImagePath,
						DateCreated:          resp.Customer.DateCreated,
						Status:               resp.Customer.Active,
					}
					response = responses.RegisterResponse{
						StatusCode:    true,
						StatusMessage: "Registration successful",
						Result:        &cust,
					}
				}
			}
		}

		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = response

	} else {
		var response responses.RegisterResponse = responses.RegisterResponse{
			StatusCode:    false,
			StatusMessage: "Something went wrong:: " + err.Error(),
			Result:        nil,
		}

		c.Data["json"] = response
	}
	c.ServeJSON()
}

// Post ...
// @Title Create
// @Description create Auth_requests
// @Param	body		body 	models.Auth_requests	true		"body for Auth_requests content"
// @Success 201 {object} models.Auth_requests
// @Failure 403 body is empty
// @router / [post]
func (c *Auth_requestsController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Auth_requests by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Auth_requests
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Auth_requestsController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Auth_requests
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Auth_requests
// @Failure 403
// @router / [get]
func (c *Auth_requestsController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Auth_requests
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Auth_requests	true		"body for Auth_requests content"
// @Success 200 {object} models.Auth_requests
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Auth_requestsController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Auth_requests
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Auth_requestsController) Delete() {

}
