// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"msys_payment_app_gateway/controllers"
	"msys_payment_app_gateway/middlewares"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/api",
			beego.NSBefore(middlewares.AuthMiddleware),
			beego.NSInclude(
				&controllers.Api_requestsController{},
			),
		),
		beego.NSNamespace("/callback",
			beego.NSInclude(
				&controllers.CallbackController{},
			),
		),
		beego.NSNamespace("/auth",
			beego.NSInclude(
				&controllers.Auth_requestsController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
